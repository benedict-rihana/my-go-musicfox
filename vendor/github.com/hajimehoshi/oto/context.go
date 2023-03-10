// Copyright 2019 The Oto Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package oto

import (
	"errors"
	"io"
	"sync"
	"time"

	"github.com/hajimehoshi/oto/internal/mux"
)

// Context is the main object in Oto. It interacts with the audio drivers.
//
// To play sound with Oto, first create a context. Then use the context to create
// an arbitrary number of players. Then use the players to play sound.
//
// There can only be one context at any time. Closing a context and opening a new one is allowed.
type Context struct {
	driverWriter *driverWriter
	mux          *mux.Mux
	errCh        chan error
}

var (
	theContext *Context
	contextM   sync.Mutex
)

var errClosed = errors.New("closed")

// NewContext creates a new context, that creates and holds ready-to-use Player objects.
//
// The sampleRate argument specifies the number of samples that should be played during one second.
// Usual numbers are 44100 or 48000.
//
// The channelNum argument specifies the number of channels. One channel is mono playback. Two
// channels are stereo playback. No other values are supported.
//
// The bitDepthInBytes argument specifies the number of bytes per sample per channel. The usual value
// is 2. Only values 1 and 2 are supported.
//
// The bufferSizeInBytes argument specifies the size of the buffer of the Context. This means, how
// many bytes can Context remember before actually playing them. Bigger buffer can reduce the number
// of Player's Write calls, thus reducing CPU time. Smaller buffer enables more precise timing. The
// longest delay between when samples were written and when they started playing is equal to the size
// of the buffer.
func NewContext(sampleRate, channelNum, bitDepthInBytes, bufferSizeInBytes int) (*Context, error) {
	contextM.Lock()
	defer contextM.Unlock()

	if theContext != nil {
		panic("oto: NewContext can be called only once")
	}

	d, err := newDriver(sampleRate, channelNum, bitDepthInBytes, bufferSizeInBytes)
	if err != nil {
		return nil, err
	}
	dw := &driverWriter{
		driver:         d,
		bufferSize:     bufferSizeInBytes,
		bytesPerSecond: sampleRate * channelNum * bitDepthInBytes,
	}
	c := &Context{
		driverWriter: dw,
		mux:          mux.New(channelNum, bitDepthInBytes),
		errCh:        make(chan error, 1),
	}
	theContext = c
	go func() {
		if _, err := io.Copy(c.driverWriter, c.mux); err != nil {
			c.errCh <- err
		}
		close(c.errCh)
		c.Close()
	}()
	return c, nil
}

// NewPlayer creates a new, ready-to-use Player belonging to the Context.
func (c *Context) NewPlayer() *Player {
	return newPlayer(c)
}

// Close closes the Context and its Players and frees any resources associated with it. The Context is no longer
// usable after calling Close.
func (c *Context) Close() error {
	if theContext == nil {
		return nil
	}

	contextM.Lock()
	if theContext == nil {
		contextM.Unlock()
		return nil
	}
	theContext = nil
	contextM.Unlock()

	if err := c.driverWriter.Close(); err != nil {
		return err
	}
	for _, r := range c.mux.Sources() {
		if err := r.(io.Closer).Close(); err != nil {
			return err
		}
	}
	if err := c.mux.Close(); err != nil {
		return err
	}
	return nil
}

type tryWriteCloser interface {
	io.Closer

	TryWrite([]byte) (int, error)
	tryWriteCanReturnWithoutWaiting() bool
}

type driverWriter struct {
	driver         tryWriteCloser
	bufferSize     int
	bytesPerSecond int

	m sync.Mutex
}

func (d *driverWriter) Write(buf []byte) (int, error) {
	d.m.Lock()
	defer d.m.Unlock()

	written := 0
	for len(buf) > 0 {
		if d.driver == nil {
			return written, errClosed
		}
		n, err := d.driver.TryWrite(buf)
		written += n
		if err != nil {
			return written, err
		}
		buf = buf[n:]
		if d.driver.tryWriteCanReturnWithoutWaiting() {
			// When not all buf is written, the underlying buffer is full.
			// Mitigate the busy loop by sleeping (#10).
			if len(buf) > 0 {
				t := time.Second * time.Duration(d.bufferSize) / time.Duration(d.bytesPerSecond) / 8
				time.Sleep(t)
			}
		}
	}
	return written, nil
}

func (d *driverWriter) Close() error {
	d.m.Lock()
	defer d.m.Unlock()

	// Close should be wait until the buffer data is consumed (#36).
	// This is the simplest (but ugly) fix.
	// TODO: Implement player's Close to wait the buffer played.
	time.Sleep(time.Second * time.Duration(d.bufferSize) / time.Duration(d.bytesPerSecond))
	return d.driver.Close()
}
