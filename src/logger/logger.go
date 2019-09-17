package logger

import (
	"fmt"
	"log"

	"github.com/ttacon/chalk"
)

// logger is
type logger struct {
	VLogEnabled bool
}

// Logger is
var Logger = logger{
	VLogEnabled: false,
}

// Warn is
func (l *logger) Warn(info string) {
	yellowOnBlack := chalk.Yellow.NewStyle().WithBackground(chalk.Black)
	fmt.Println(yellowOnBlack.Style(info))
}

// Fatal is
func (l *logger) Fatal(info string) {
	redOnBlack := chalk.Red.NewStyle().WithBackground(chalk.Black).WithTextStyle(chalk.Bold)
	fmt.Println(redOnBlack.Style(info))
	log.Fatalf("CloudEvent Generator has stopped running. For troubleshooting, enable verbose logging with option --verbose.\n")
}

// Success is
func (l *logger) Success(info string) {
	greenOnBlack := chalk.Green.NewStyle().WithBackground(chalk.Black).WithTextStyle(chalk.Bold)
	fmt.Println(greenOnBlack.Style(info))
}

// Info is
func (l *logger) Info(info string) {
	if l.VLogEnabled {
		blueOnBlack := chalk.Blue.NewStyle().WithBackground(chalk.Black)
		fmt.Println(blueOnBlack.Style(info))
	}
}
