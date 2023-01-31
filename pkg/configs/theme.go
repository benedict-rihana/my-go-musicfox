package configs

import (
	"github.com/gookit/ini/v2"
	"github.com/muesli/termenv"
)

var ThemeConfig *Theme

type Theme struct{
	AppBackground termenv.Color
	TitleBarFG termenv.Color
	StatusBarFG termenv.Color
	StatusBarBGRepeatAll termenv.Color
	StatusBarBGList termenv.Color
	StatusBarBGRepeatOne termenv.Color
	StatusBarBGShuffle termenv.Color
	StatusBarBGBeat termenv.Color
	UserFG termenv.Color
	MenuItemFG termenv.Color
	MenuItemArtistFG termenv.Color
	MenuItemSelectedFG termenv.Color
	MenuItemArtistSelectedFG termenv.Color
	MenuItemSelectedBG termenv.Color
	MenuTitleFG termenv.Color
	TimerFG termenv.Color
	TextFG termenv.Color
	LyricFG termenv.Color
	LyricHilightFG termenv.Color
	MenuPointer string
}

//init default
//init from config file
func InitThemeConfig(filePath string) *Theme{
	var theme = initDefaultConfig()
	if err := ini.LoadExists(filePath); err != nil {
		return theme
	}
	initForeground(theme)
	initBackground(theme)
	initSymbols(theme)
	return theme
}


func initBackground(theme *Theme){
	theme.AppBackground = initColor("background.app", theme.AppBackground)
	theme.StatusBarBGBeat = initColor("background.beat", theme.StatusBarBGBeat)
	theme.StatusBarBGRepeatAll = initColor("background.repeatall", theme.StatusBarBGRepeatAll)
	theme.StatusBarBGRepeatOne = initColor("background.repeatone", theme.StatusBarBGRepeatOne)
	theme.StatusBarBGList = initColor("background.list", theme.StatusBarBGList)
	theme.StatusBarBGShuffle = initColor("background.shuffle", theme.StatusBarBGShuffle)
	theme.MenuItemSelectedBG = initColor("background.itemselected",theme.MenuItemSelectedBG)
}

func initColor(config string,defVal termenv.Color) termenv.Color{
	colorStr := ini.String(config,"")
	if colorStr != "" {
		return termenv.ColorProfile().Color(colorStr)
	}
	return defVal
}

func initForeground(theme *Theme){
	theme.TitleBarFG = initColor("foreground.titlebar",theme.TitleBarFG)
	theme.StatusBarFG = initColor("foreground.statusbar",theme.StatusBarFG)
	theme.MenuTitleFG = initColor("foreground.menutitle",theme.MenuTitleFG)
	theme.UserFG = initColor("foreground.user",theme.UserFG)
	theme.MenuItemFG = initColor("foreground.menuitem",theme.MenuItemFG)
	theme.MenuItemSelectedFG = initColor("foreground.menuitemselected",theme.MenuItemSelectedFG)
	theme.MenuItemArtistFG = initColor("foreground.menuitemartist",theme.MenuItemArtistFG)
	theme.MenuItemArtistSelectedFG = initColor("foreground.menuitemartistselected",theme.MenuItemArtistSelectedFG)
	theme.LyricFG = initColor("foreground.lyric",theme.LyricFG)
	theme.LyricHilightFG = initColor("foreground.lyrichighlight",theme.LyricHilightFG)
	theme.TimerFG = initColor("foreground.timer",theme.TimerFG)
	theme.TextFG = initColor("foreground.text",theme.TextFG)
}

func initSymbols(them *Theme){
	them.MenuPointer = ini.String("symbol.pointer",them.MenuPointer)
}

func initDefaultConfig() *Theme{
	var theme = &Theme{
		MenuItemFG: termenv.ANSIBlue,
		MenuItemArtistFG: termenv.ANSIMagenta,
		MenuItemSelectedFG: termenv.ANSIBrightBlue,
		MenuItemArtistSelectedFG: termenv.ANSIMagenta,
		MenuItemSelectedBG: termenv.ANSIBrightBlack,
		TimerFG: termenv.ANSIBlue,
		MenuTitleFG: termenv.ANSIRed,
		UserFG: termenv.ANSIBrightBlue,
		AppBackground: termenv.ANSIBlack,
		TitleBarFG: termenv.ANSIBlue,
		StatusBarFG: termenv.ANSIBlack,
		StatusBarBGRepeatAll: termenv.ANSIBlue,
		StatusBarBGRepeatOne: termenv.ANSIBrightBlue,
		StatusBarBGList: termenv.ANSICyan,
		StatusBarBGShuffle: termenv.ANSIMagenta,
		StatusBarBGBeat: termenv.ANSIRed,
		LyricFG: termenv.ANSIBrightCyan,
		LyricHilightFG: termenv.ANSIBrightYellow,
		MenuPointer: "=>",
 		TextFG: termenv.ColorProfile().Color("#1e1e1e"),
	}
	return theme
}
