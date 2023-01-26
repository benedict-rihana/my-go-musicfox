package constants

import (
	"time"
)

var (
	// AppVersion Inject by -ldflags
	AppVersion   = "v2.2.1"
	LastfmKey    = ""
	LastfmSecret = ""
)

const AppName = "musicfox"
const GroupID = "com.anhoder.musicfox"
const AppDescription = "<cyan>Musicfox - 命令行版网易云音乐</>"
const AppShowStartup = true
const AppGithubUrl = "https://github.com/anhoder/go-musicfox"
const AppLatestReleases = "https://github.com/anhoder/go-musicfox/releases/latest"
const AppCheckUpdateUrl = "https://api.github.com/repos/anhoder/go-musicfox/releases/latest"
const LastfmAuthUrl = "https://www.last.fm/api/auth/?api_key=%s&token=%s"
const ProgressFullChar = "#"
const ProgressEmptyChar = " "
const StartupProgressOutBounce = true
const StartupLoadingSeconds = 2
const StartupTickDuration = time.Millisecond * 16
const StartupSignIn = true
const StartupCheckUpdate = true

const AppLocalDataDir = ".go-musicfox"
const AppDBName = "musicfox"
const AppIniFile = "go-musicfox.ini"
const AppThemeFile = "theme.ini"
const AppPrimaryRandom = "random"
const AppPrimaryColor = "#f90022"
const AppHttpTimeout = time.Second * 5

const MainShowTitle = true
const MainLoadingText = "[加载中...]"
const MainShowLyric = true
const MainShowNotify = true
const MainPProfPort = 9876
const MainAltScreen = true
const UNMDefaultSources = "kuwo"

const BeepPlayer = "beep" // beep
const MpdPlayer = "mpd"   // mpd
const OsxPlayer = "osx"   // osx

const BeepGoMp3Decoder = "go-mp3"
const BeepMiniMp3Decoder = "minimp3"

const SearchPageSize = 100

const AppHelpTemplate = `%s

{{.Description}} (Version: <info>{{.Version}}</>)

<comment>Usage:</>
  {$binName} [Global Options...] <info>{command}</> [--option ...] [argument ...]

<comment>Global Options:</>
{{.GOpts}}
<comment>Available Commands:</>{{range $module, $cs := .Cs}}{{if $module}}
<comment> {{ $module }}</>{{end}}{{ range $cs }}
  <info>{{.Name | paddingName }}</> {{.UseFor}}{{if .Aliases}} (alias: <cyan>{{ join .Aliases ","}}</>){{end}}{{end}}{{end}}

  <info>{{ paddingName "help" }}</> Display help information

Use "<cyan>{$binName} {COMMAND} -h</>" for more information about a command
`
