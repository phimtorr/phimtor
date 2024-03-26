//go:build darwin && !ios

package vlc

import (
	"os/exec"
)

func openURL(url string) error {
	cmd := exec.Command("open", url, "-a", "VLC")
	return cmd.Run()
}
