//go:build windows
// +build windows

package main

import (
	"strings"

	"golang.org/x/sys/windows/registry"
)

func addToWindowsPath(newPath string) error {
	k, err := registry.OpenKey(registry.CURRENT_USER, `Environment`, registry.QUERY_VALUE|registry.SET_VALUE)
	if err != nil {
		return err
	}
	defer k.Close()

	currentPath, _, err := k.GetStringValue("Path")
	if err != nil {
		currentPath = ""
	}

	paths := strings.Split(currentPath, ";")
	for _, p := range paths {
		if strings.EqualFold(strings.TrimSpace(p), newPath) {
			return nil
		}
	}

	var newPathValue string
	if currentPath == "" {
		newPathValue = newPath
	} else if strings.HasSuffix(currentPath, ";") {
		newPathValue = currentPath + newPath
	} else {
		newPathValue = currentPath + ";" + newPath
	}

	return k.SetStringValue("Path", newPathValue)
}
