package ezio

import (
	"fmt"
	"regexp"
	"strings"
	"time"

	"github.com/fatih/color"
)

const DATE_FORMAT = "Mon Jan _2 15:04:05 2006"

type EzOut struct {
	Prompt string
}

func (e EzOut) Write(b []byte) (int, error) {
	s := string(b[0:])

	re := regexp.MustCompile(`(\[\S+ \S+ \S+ \d+:\d+:\d+ \d+\]) (.+) (\[\d{3}\])(.+)`)
	submatchall := re.FindStringSubmatch(s)
	green := color.New(color.FgGreen)
	green.Printf("%s ", submatchall[1])
	fmt.Println(submatchall[2] + submatchall[3] + submatchall[4])
	return len(b), nil
}

func Info(s string) {
	green := color.New(color.FgGreen)
	green.Printf("[%s] ", time.Now().Format(DATE_FORMAT))
	fmt.Print(s)
}

func Error(s string) {
	red := color.New(color.FgGreen)
	red.Printf("[%s] ", time.Now().Format(DATE_FORMAT))
	fmt.Print(s)
}

func Installer(s string) {
	green := color.New(color.FgGreen)
	green.Printf("[%s] ", time.Now().Format(DATE_FORMAT))
	fmt.Print(s)
}

func Custom(level string, s string) {
	green := color.New(color.FgGreen)
	green.Printf("[%s] ", time.Now().Format(DATE_FORMAT))
	fmt.Print(s)
}

func Confirm(question string) bool {

	var confirmation string

	Info(fmt.Sprintf("%s [y/N]? ", question))
	fmt.Scan(&confirmation)

	confirmation = strings.TrimSpace(confirmation)
	confirmation = strings.ToLower(confirmation)

	if confirmation == "y" {
		return true
	}

	return false
}
