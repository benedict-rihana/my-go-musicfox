package ui

import (
	"fmt"
	"strings"
	"unicode/utf8"

	"go-musicfox/pkg/configs"

	"github.com/muesli/termenv"
)

type Loading struct {
	model  *NeteaseModel
	curLen int
}

func NewLoading(m *NeteaseModel) *Loading {
	return &Loading{
		model: m,
	}
}

// 开始
func (loading *Loading) start() {
	termenv.MoveCursor(loading.model.menuTitleStartRow, 0)
	fmt.Printf("%s",SetFgBgStyle(strings.Repeat(" ",loading.model.WindowWidth), configs.ThemeConfig.AppBackground,configs.ThemeConfig.AppBackground))
	termenv.MoveCursor(loading.model.menuTitleStartRow, loading.model.menuTitleStartColumn)

	loading.curLen = utf8.RuneCountInString(loading.model.menuTitle.OriginString()) + utf8.RuneCountInString(" "+configs.ConfigRegistry.MainLoadingText)

	// var repeatSpace string
	// if loading.model.menuTitleStartColumn > 0 {
	// 	repeatSpace = strings.Repeat(" ", loading.model.menuTitleStartColumn)
	// }
	fmt.Printf("%s%s",
		//SetFgBgStyle(repeatSpace,configs.ThemeConfig.AppBackground,configs.ThemeConfig.AppBackground),
		SetFgBgStyle(loading.model.menuTitle.String(), termenv.ANSIBrightGreen,configs.ThemeConfig.AppBackground),
		SetFgBgStyle(" "+configs.ConfigRegistry.MainLoadingText, termenv.ANSIBrightRed, configs.ThemeConfig.AppBackground))

	termenv.MoveCursor(0, 0)
}

// 完成
func (loading *Loading) complete() {
	termenv.MoveCursor(loading.model.menuTitleStartRow, 0)
	fmt.Printf("%s",SetFgBgStyle(strings.Repeat(" ",loading.model.WindowWidth), configs.ThemeConfig.AppBackground,configs.ThemeConfig.AppBackground))
	termenv.MoveCursor(loading.model.menuTitleStartRow, loading.model.menuTitleStartColumn)
	fmt.Printf("%s",SetFgBgStyle(loading.model.menuTitle.String(),configs.ThemeConfig.MenuTitleFG,configs.ThemeConfig.AppBackground))
	termenv.MoveCursor(0, 0)
}
