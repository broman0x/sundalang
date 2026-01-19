//go:build !windows
// +build !windows

package main

import "fmt"

func setupInstallerWindow() {
	fmt.Print("\033[8;30;90t")
}

func cleanupInstallerWindow() {
}
