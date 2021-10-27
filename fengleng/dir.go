package fengleng

import "runtime"

var (
	Sep    string
	CfgDir string
	LogDir string
)

func init() {
	if runtime.GOOS == "windows" {
		Sep = `\`
		CfgDir = `c:\fengleng\`
		LogDir = `c:\fengleng\log\`
	} else {
		Sep = `/`
		CfgDir = `/etc/fengleng/`
		LogDir = `/var/log/fengleng/`
	}
}
