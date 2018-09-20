package fs

import (
	"os"
	"os/exec"
	"path/filepath"
)

const local = "phpbin"

func WhereIsGlobalPHP(phpExe string) (string, error) {
	return exec.LookPath(phpExe)
}

func WhereIsLocalPHP(phpExe string) (string, error) {

	var err error

	if _, err = os.Stat(local); err == nil {
		absPath, _ := filepath.Abs(filepath.Dir(local))
		defaultExecPath := absPath + string(os.PathSeparator) + local + string(os.PathSeparator) + phpExe
		return defaultExecPath, nil
	}

	return "", err
}

func CreateDirIfNotExist(dir string) error {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err = os.MkdirAll(dir, 0755)
		if err != nil {
			return err
		}
	}
	return nil
}

// func searchForPhp(phpExe string) (string, error) {
//
// 	var defaultExecPath string
// 	var path string
// 	var err error
// 	var absPath string
//
// 	output.Info("Looking for php in default directory: " + install.PhpDir + "\n")
// 	if _, err = os.Stat(install.PhpDir); err == nil {
// 		output.Info("Local php installation founded\n")
// 		absPath, _ = filepath.Abs(filepath.Dir(install.PhpDir))
// 		defaultExecPath = absPath + string(os.PathSeparator) + install.PhpDir + string(os.PathSeparator) + phpExe
// 		return defaultExecPath, nil
// 	}
//
// 	defaultExecPath, err = exec.LookPath(phpExe)
// 	if err != nil {
// 		output.Error("php executable not found in path\n")
//
// 		if !askToInstallPhp() {
// 			return "", errors.New("php won't be installed. bye bye.")
// 		}
//
// 		output.Info("Please wait...\n")
// 		path, err = install.Installer(install.Version, install.PhpDir)
// 		if err != nil {
// 			return "", errors.New(err.Error())
// 		}
//
// 		defaultExecPath = path + string(os.PathSeparator) + phpExe
// 	}
//
// 	return defaultExecPath, nil
// }
//
// func askToInstallPhp() bool {
// 	var confirmation string
//
// 	output.Installer("Would you like to install php locally [y/N]? ")
// 	fmt.Scan(&confirmation)
//
// 	confirmation = strings.TrimSpace(confirmation)
// 	confirmation = strings.ToLower(confirmation)
//
// 	if confirmation == "y" || confirmation == "yes" {
// 		return true
// 	}
//
// 	return false
// }
