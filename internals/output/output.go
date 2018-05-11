package output

import (
  "fmt"
  "github.com/fatih/color"
)
  

func Info(s string) {
  green := color.New(color.FgGreen, color.Bold)
  green.Printf("[%-12s] ", "Information")
  fmt.Print(s)
}

func Error(s string) {
  red := color.New(color.FgRed, color.Bold)
  red.Printf("[%-12s] ", "Error")
  fmt.Print(s)
}

func Installer(s string) {
  green := color.New(color.FgGreen, color.Bold)
  green.Printf("[%-12s] ", "Installer")
  fmt.Print(s)
}

func Custom(level string, s string) {
  green := color.New(color.FgGreen, color.Bold)
  green.Printf("[%-12s] ", level)
  fmt.Print(s)
}