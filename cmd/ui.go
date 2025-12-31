package main

import "fmt"

const (
	ColorReset = "\033[0m"
	ColorRed   = "\033[91m"
	ColorWhite = "\033[97m"
	ColorGray  = "\033[90m"
	ColorBold  = "\033[1m"
)

func colorize(color, text string) string {
	return color + text + ColorReset
}

func printBanner() {
	fmt.Println()
	fmt.Println(
		colorize(ColorRed+ColorBold, "  ██████╗ ██╗   ██╗███╗   ██╗██████╗  █████╗ ") +
			colorize(ColorWhite+ColorBold, "██╗      █████╗ ███╗   ██╗ ██████╗ "))
	fmt.Println(
		colorize(ColorRed+ColorBold, "  ██╔════╝██║   ██║████╗  ██║██╔══██╗██╔══██╗") +
			colorize(ColorWhite+ColorBold, "██║     ██╔══██╗████╗  ██║██╔════╝ "))
	fmt.Println(
		colorize(ColorRed+ColorBold, "  ███████╗██║   ██║██╔██╗ ██║██║  ██║███████║") +
			colorize(ColorWhite+ColorBold, "██║     ███████║██╔██╗ ██║██║  ███╗"))
	fmt.Println(
		colorize(ColorRed+ColorBold, "  ╚════██║██║   ██║██║╚██╗██║██║  ██║██╔══██║") +
			colorize(ColorWhite+ColorBold, "██║     ██╔══██║██║╚██╗██║██║   ██║"))
	fmt.Println(
		colorize(ColorRed+ColorBold, "  ███████║╚██████╔╝██║ ╚████║██████╔╝██║  ██║") +
			colorize(ColorWhite+ColorBold, "███████╗██║  ██║██║ ╚████║╚██████╔╝"))
	fmt.Println(
		colorize(ColorRed+ColorBold, "  ╚══════╝ ╚═════╝ ╚═╝  ╚═══╝╚═════╝ ╚═╝  ╚═╝") +
			colorize(ColorWhite+ColorBold, "╚══════╝╚═╝  ╚═╝╚═╝  ╚═══╝ ╚═════╝ "))
	fmt.Println()
	fmt.Println(colorize(ColorWhite, "                   Bahasa Pemrograman Sunda Pandeglang"))
	fmt.Println(colorize(ColorGray, "                              v1.0.3"))
	fmt.Println()
}

func printHeader(title string) {
	fmt.Println(colorize(ColorRed+ColorBold, " ▸ "+title))
	fmt.Println()
}

func printSuccess(message string) {
	fmt.Println(colorize(ColorRed+ColorBold, " ✓ ") + colorize(ColorWhite, message))
}

func printInfo(message string) {
	fmt.Println(colorize(ColorWhite+ColorBold, " • ") + colorize(ColorWhite, message))
}

func printWarning(message string) {
	fmt.Println(colorize(ColorRed+ColorBold, " ⚠ ") + colorize(ColorRed, message))
}

func printError(message string) {
	fmt.Println(colorize(ColorRed+ColorBold, " ✗ ") + colorize(ColorRed, message))
}

func repeatString(s string, count int) string {
	if count <= 0 {
		return ""
	}
	result := ""
	for i := 0; i < count; i++ {
		result += s
	}
	return result
}

func printMenuOption(number int, text string) {
	fmt.Printf(colorize(ColorRed+ColorBold, "   [%d] ")+colorize(ColorWhite, "%s\n"), number, text)
}

func printSeparator() {
	fmt.Println(colorize(ColorGray, " "+repeatString("─", 80)))
}

func printFooter() {
	fmt.Println()
	fmt.Print(colorize(ColorRed+ColorBold, " ▸ "))
}
