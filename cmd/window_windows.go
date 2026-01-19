//go:build windows
// +build windows

package main

import (
	"os"
	"syscall"
	"unsafe"
)

var (
	kernel32         = syscall.NewLazyDLL("kernel32.dll")
	user32           = syscall.NewLazyDLL("user32.dll")
	getConsoleWindow = kernel32.NewProc("GetConsoleWindow")
	getWindowLong    = user32.NewProc("GetWindowLongW")
	setWindowLong    = user32.NewProc("SetWindowLongW")
	setWindowPos     = user32.NewProc("SetWindowPos")
	getSystemMetrics = user32.NewProc("GetSystemMetrics")
)

const (
	WS_SIZEBOX       = 0x00040000
	WS_MAXIMIZEBOX   = 0x00010000
	SWP_NOMOVE       = 0x0002
	SWP_NOSIZE       = 0x0001
	SWP_NOZORDER     = 0x0004
	SWP_FRAMECHANGED = 0x0020
	SM_CXSCREEN      = 0
	SM_CYSCREEN      = 1
)

var GWL_STYLE = ^uintptr(15) 

func setConsoleSize(cols, rows int) {
	handle, err := syscall.Open("CONOUT$", syscall.O_RDWR, 0)
	if err != nil {
		return
	}
	defer syscall.Close(handle)

	coord := struct {
		X, Y int16
	}{int16(cols), int16(rows)}

	kernel32.NewProc("SetConsoleScreenBufferSize").Call(uintptr(handle), *(*uintptr)(unsafe.Pointer(&coord)))

	rect := struct {
		Left, Top, Right, Bottom int16
	}{0, 0, int16(cols - 1), int16(rows - 1)}

	kernel32.NewProc("SetConsoleWindowInfo").Call(uintptr(handle), 1, uintptr(unsafe.Pointer(&rect)))
}

func lockWindowSize() {
	hwnd, _, _ := getConsoleWindow.Call()
	if hwnd == 0 {
		return
	}

	style, _, _ := getWindowLong.Call(hwnd, uintptr(GWL_STYLE))
	newStyle := style &^ WS_SIZEBOX &^ WS_MAXIMIZEBOX
	setWindowLong.Call(hwnd, uintptr(GWL_STYLE), newStyle)

	setWindowPos.Call(hwnd, 0, 0, 0, 0, 0, SWP_NOMOVE|SWP_NOSIZE|SWP_NOZORDER|SWP_FRAMECHANGED)
}

func unlockWindowSize() {
	hwnd, _, _ := getConsoleWindow.Call()
	if hwnd == 0 {
		return
	}

	style, _, _ := getWindowLong.Call(hwnd, uintptr(GWL_STYLE))
	newStyle := style | WS_SIZEBOX | WS_MAXIMIZEBOX
	setWindowLong.Call(hwnd, uintptr(GWL_STYLE), newStyle)

	setWindowPos.Call(hwnd, 0, 0, 0, 0, 0, SWP_NOMOVE|SWP_NOSIZE|SWP_NOZORDER|SWP_FRAMECHANGED)
}

func centerWindow() {
	hwnd, _, _ := getConsoleWindow.Call()
	if hwnd == 0 {
		return
	}

	screenWidth, _, _ := getSystemMetrics.Call(SM_CXSCREEN)
	screenHeight, _, _ := getSystemMetrics.Call(SM_CYSCREEN)

	var rect struct {
		Left, Top, Right, Bottom int32
	}
	user32.NewProc("GetWindowRect").Call(hwnd, uintptr(unsafe.Pointer(&rect)))

	winWidth := rect.Right - rect.Left
	winHeight := rect.Bottom - rect.Top

	x := (int32(screenWidth) - winWidth) / 2
	y := (int32(screenHeight) - winHeight) / 2

	setWindowPos.Call(hwnd, 0, uintptr(x), uintptr(y), 0, 0, SWP_NOSIZE|SWP_NOZORDER)
}

func setupInstallerWindow() {
	os.Stdout.WriteString("\033[8;30;90t")

	setConsoleSize(90, 30)
	lockWindowSize()
	centerWindow()
}

func cleanupInstallerWindow() {
	unlockWindowSize()
}
