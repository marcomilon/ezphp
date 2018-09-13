package output

import (
	"fmt"
	"time"

	"github.com/fatih/color"
)

func Info(s string) {
	green := color.New(color.FgGreen)
	green.Printf("| %s | ", time.Now().Format(time.Stamp))
	fmt.Print(s)
}

func Error(s string) {
	red := color.New(color.FgGreen)
	red.Printf("| %s | ", time.Now().Format(time.Stamp))
	fmt.Print(s)
}

func Installer(s string) {
	green := color.New(color.FgGreen)
	green.Printf("| %s | ", time.Now().Format(time.Stamp))
	fmt.Print(s)
}

func Custom(level string, s string) {
	green := color.New(color.FgGreen)
	green.Printf("[%s] ", time.Now().Format(time.Stamp))
	fmt.Print(s)
}
