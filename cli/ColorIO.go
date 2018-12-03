package cli

import (
	"fmt"
	"regexp"
	"strings"
	"time"

	"github.com/fatih/color"
)

type ColorIO struct{}

const DATE_FORMAT = "Mon Jan _2 15:04:05 2006"

func (ColorIO) Write(b []byte) (int, error) {
	s := string(b[0:])
	log := strings.TrimSpace(s)

	re := regexp.MustCompile(`\[?(\S+ \S+ \S+ \d+:\d+:\d+ \d+)\] (.+) (\[\d{3}\])(.+)`)
	if re.MatchString(log) {
		submatchall := re.FindStringSubmatch(log)
		green := color.New(color.FgGreen)
		green.Printf("[%s] ", submatchall[1])
		fmt.Println(submatchall[2] + submatchall[3] + submatchall[4])
	} else {
		red := color.New(color.FgRed)
		red.Printf("[%s] ", time.Now().Format(DATE_FORMAT))
		fmt.Print(s)
	}

	return len(b), nil
}

func (ColorIO) Info(s string) {
	green := color.New(color.FgGreen)
	green.Printf("[%s] ", time.Now().Format(DATE_FORMAT))
	fmt.Print(s)
}

func (ColorIO) Error(s string) {
	red := color.New(color.FgRed)
	red.Printf("[%s] ", time.Now().Format(DATE_FORMAT))
	fmt.Print(s)
}

func (ColorIO) Custom(tag string, s string) {
	red := color.New(color.FgYellow)
	red.Printf("[%-24s] ", tag)
	fmt.Print(s)
}

func (io ColorIO) Confirm(question string) bool {

	var confirmation string

	io.Info(fmt.Sprintf("%s [y/N]? ", question))
	fmt.Scan(&confirmation)

	confirmation = strings.TrimSpace(confirmation)
	confirmation = strings.ToLower(confirmation)

	if confirmation == "y" {
		return true
	}

	return false
}
