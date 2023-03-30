package figures

import (
	"os"
	"runtime"
)

var platform string = runtime.GOOS

func isUnicodeSupported() bool {
	if platform != "win32" {
		// TERM is 'linux' in the kernet
		val, ok := os.LookupEnv("TERM")
		if !ok {
			return false
		}
		return val != "linux" // Linux console (kernel)
	}
	_, oksess := os.LookupEnv("WT_SESSION")
	_, okterminus := os.LookupEnv("TERMINUS_SUBLIME")
	_, okci := os.LookupEnv("CI")
	if oksess || okterminus || okci {
		return true
	}

	return envKeyVal("TERM_PROGRAM", "Terminus-Sublime", "vscode") ||
		envKeyVal("TERM", "xterm-256color", "alacritty") ||
		envKeyVal("TERMINAL_EMULATOR", "JetBrains-JediTerm") ||
		envKeyVal("ConEmuTask", "{cmd::Cmder}")
}

func envKeyVal(key string, vals ...string) bool {
	v, ok := os.LookupEnv(key)

	if !ok {
		return false
	}

	for _, val := range vals {
		if v == val {
			return true
		}
	}

	return false

}
