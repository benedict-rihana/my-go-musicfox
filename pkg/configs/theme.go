package configs

import "github.com/muesli/termenv"

var ThemeConfig *Theme

type Theme struct{
	AppBackground termenv.Color
	TitleBarFG termenv.Color
	StatusBarFG termenv.Color
	StatusBarBGRepeatAll termenv.Color
	StatusBarBGRepeatOne termenv.Color
	StatusBarBGShuffle termenv.Color
	StatusBarBGBeat termenv.Color
	UserFG termenv.Color
	MenuItemFG termenv.Color
	MenuItemArtistFG termenv.Color
	MenuItemSelectedFG termenv.Color
	MenuItemSelectedArtistFG termenv.Color
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
	var theme = &Theme{
		MenuItemFG: termenv.ANSIBlue,
		MenuItemArtistFG: termenv.ANSIMagenta,
		MenuItemSelectedFG: termenv.ANSIBrightBlue,
		MenuItemSelectedArtistFG: termenv.ANSIMagenta,
		MenuItemSelectedBG: termenv.ANSIBrightBlack,
		TimerFG: termenv.ANSIBlue,
		MenuTitleFG: termenv.ANSIRed,
		UserFG: termenv.ANSIBrightBlue,
		AppBackground: termenv.ANSIBlack,
		TitleBarFG: termenv.ANSIBlue,
		StatusBarFG: termenv.ANSIBlack,
		StatusBarBGRepeatAll: termenv.ANSIBlue,
		StatusBarBGRepeatOne: termenv.ANSIBrightBlue,
		StatusBarBGShuffle: termenv.ANSIMagenta,
		StatusBarBGBeat: termenv.ANSIRed,
		LyricFG: termenv.ANSIBrightCyan,
		LyricHilightFG: termenv.ANSIBrightYellow,
		MenuPointer: "",
 		TextFG: termenv.ColorProfile().Color("#1e1e1e"),
	}
	return theme
}
