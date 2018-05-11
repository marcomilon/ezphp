package output

import "fmt"

func Info(s string) {
  out("Information", s)
}

func Error(s string) {
  out("Error", s)
}

func Installer(s string) {
  out("Installer", s)
}

func Custom(level string, s string) {
  out(level, s)
}

func out(level string, s string) {
  fmt.Printf("[%-12s] %s", level, s)
}