package serve

import (
	"os"
	"os/exec"
)

func Start(phpBin string, host string, docRoot string) error {

	cmd := exec.Command(phpBin, "-S", host, "-t", docRoot)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		return err
	}

	return nil
	
}
