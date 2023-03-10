package ui

import (
	"fmt"
	"strconv"

	"go-musicfox/pkg/structs"
	"go-musicfox/utils"

	"github.com/anhoder/netease-music/service"
)

type DjRadioDetailMenu struct {
	DefaultMenu
	menus     []MenuItem
	songs     []structs.Song
	djRadioId int64
	limit     int
	offset    int
}

func NewDjRadioDetailMenu(djRadioId int64) *DjRadioDetailMenu {
	return &DjRadioDetailMenu{
		djRadioId: djRadioId,
		limit:     50,
		offset:    0,
	}
}

func (m *DjRadioDetailMenu) IsSearchable() bool {
	return true
}

func (m *DjRadioDetailMenu) MenuData() interface{} {
	return m.songs
}

func (m *DjRadioDetailMenu) IsPlayable() bool {
	return true
}

func (m *DjRadioDetailMenu) GetMenuKey() string {
	return fmt.Sprintf("dj_radio_detail_%d", m.djRadioId)
}

func (m *DjRadioDetailMenu) MenuViews() []MenuItem {
	return m.menus
}

func (m *DjRadioDetailMenu) BeforeEnterMenuHook() Hook {
	return func(model *NeteaseModel) bool {

		djProgramService := service.DjProgramService{
			RID:    strconv.FormatInt(m.djRadioId, 10),
			Limit:  strconv.Itoa(m.limit),
			Offset: strconv.Itoa(m.offset),
		}
		code, response := djProgramService.DjProgram()
		codeType := utils.CheckCode(code)
		if codeType != utils.Success {
			return false
		}
		m.songs = utils.GetSongsOfDjRadio(response)
		m.menus = GetViewFromSongs(m.songs)

		return true
	}
}
