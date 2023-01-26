package configs

import "github.com/muesli/termenv"

var ThemeConfig *Theme

type Theme struct{
	TitleBarFG termenv.Color
	StatusBarFG termenv.Color
	StatusBarBGRepeatAll termenv.Color
	StatusBarBGRepeatOne termenv.Color
	StatusBarBGShuffle termenv.Color
	UerFG termenv.Color
	ListItemFG termenv.Color
}

//init default

//init from config file
//func initThemeConfig() *Theme{

//}
