package utils

import (
	"fmt"
	"os/exec"
	"runtime"
)

func OpenInBrowser(u string) error {
	switch runtime.GOOS {
	case "darwin":
		return exec.Command("open", u).Start()
	case "linux":
		return exec.Command("xdg-open", u).Start()
	case "windows":
		return exec.Command("rundll32", "url.dll,FileProtocolHandler", u).Start()
	default:
		return fmt.Errorf("open browser failed: unsupported platform: %s", runtime.GOOS)
	}
}
