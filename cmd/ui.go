package main

import (
	"fmt"
	"os"
	"strings"

	"golang.org/x/term"
)

const (
	ColorReset = "\033[0m"
	ColorRed   = "\033[91m"
	ColorGreen = "\033[92m"
	ColorWhite = "\033[97m"
	ColorGray  = "\033[90m"
	ColorCyan  = "\033[96m"
	ColorBold  = "\033[1m"
)

const contentWidth = 60

func getTermWidth() int {
	width, _, err := term.GetSize(int(os.Stdout.Fd()))
	if err != nil || width < 40 {
		return 90
	}
	return width
}

func colorize(color, text string) string {
	return color + text + ColorReset
}

func getLeftPadding() int {
	width := getTermWidth()
	padding := (width - contentWidth) / 2
	if padding < 0 {
		return 0
	}
	return padding
}

func padLeft(text string) string {
	return strings.Repeat(" ", getLeftPadding()) + text
}

func printBanner() {
	fmt.Println()

	bannerLines := []string{
		"██████╗ ██╗   ██╗███╗   ██╗██████╗  █████╗",
		"██╔════╝██║   ██║████╗  ██║██╔══██╗██╔══██╗",
		"███████╗██║   ██║██╔██╗ ██║██║  ██║███████║",
		"╚════██║██║   ██║██║╚██╗██║██║  ██║██╔══██║",
		"███████║╚██████╔╝██║ ╚████║██████╔╝██║  ██║",
		"╚══════╝ ╚═════╝ ╚═╝  ╚═══╝╚═════╝ ╚═╝  ╚═╝",
	}

	width := getTermWidth()
	bannerWidth := 44
	bannerPad := (width - bannerWidth) / 2
	if bannerPad < 0 {
		bannerPad = 0
	}

	for _, line := range bannerLines {
		fmt.Println(strings.Repeat(" ", bannerPad) + colorize(ColorRed+ColorBold, line))
	}

	fmt.Println()
	fmt.Println(padLeft(colorize(ColorWhite, "Bahasa Pemrograman Sunda Pandeglang")))
	fmt.Println(padLeft(colorize(ColorGray, "v1.0.4")))
	fmt.Println()
}

func printHeader(title string) {
	fmt.Println(padLeft(colorize(ColorCyan+ColorBold, ":: "+title)))
	fmt.Println()
}

func printSuccess(message string) {
	fmt.Println(padLeft(colorize(ColorGreen+ColorBold, "[OK] "+message)))
}

func printInfo(message string) {
	fmt.Println(padLeft(colorize(ColorWhite, "  "+message)))
}

func printWarning(message string) {
	fmt.Println(padLeft(colorize(ColorRed+ColorBold, "[!] "+message)))
}

func printError(message string) {
	fmt.Println(padLeft(colorize(ColorRed+ColorBold, "[x] "+message)))
}

func repeatString(s string, count int) string {
	if count <= 0 {
		return ""
	}
	return strings.Repeat(s, count)
}

func printMenuOption(number int, text string) {
	option := fmt.Sprintf("[%d] %s", number, text)
	fmt.Println(padLeft(colorize(ColorWhite, option)))
}

func printSeparator() {
	fmt.Println(padLeft(colorize(ColorGray, repeatString("─", contentWidth))))
}

func printFooter() {
	fmt.Println()
}

func printCentered(text string) {
	width := getTermWidth()
	padding := (width - len(text)) / 2
	if padding < 0 {
		padding = 0
	}
	fmt.Println(strings.Repeat(" ", padding) + text)
}

func printBox(lines []string, borderColor string) {
	boxWidth := contentWidth
	innerWidth := boxWidth - 4
	pad := strings.Repeat(" ", getLeftPadding())

	top := pad + "╭" + repeatString("─", boxWidth-2) + "╮"
	bottom := pad + "╰" + repeatString("─", boxWidth-2) + "╯"

	fmt.Println(colorize(borderColor, top))

	for _, line := range lines {
		textPad := (innerWidth - len(line)) / 2
		if textPad < 0 {
			textPad = 0
		}
		leftPad := strings.Repeat(" ", textPad)
		rightLen := innerWidth - textPad - len(line)
		if rightLen < 0 {
			rightLen = 0
		}
		rightPad := strings.Repeat(" ", rightLen)
		content := pad + "│ " + leftPad + line + rightPad + " │"
		fmt.Println(colorize(borderColor, content))
	}

	fmt.Println(colorize(borderColor, bottom))
}

func printSuccessBox(message string) {
	printBox([]string{"", message, ""}, ColorGreen+ColorBold)
}

func printPrompt(prompt string) {
	fmt.Print(padLeft(prompt))
}

func printLine(color, text string) {
	fmt.Println(padLeft(colorize(color, text)))
}
