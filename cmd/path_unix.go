//go:build !windows
// +build !windows

package main

func addToWindowsPath(newPath string) error {
	return nil
}
