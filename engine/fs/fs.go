package fs

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/marcomilon/ezphp/engine/php"
	"github.com/sirupsen/logrus"
)

func CreateDirIfNotExist(dir string) error {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err = os.MkdirAll(dir, 0755)
		if err != nil {
			return err
		}

		createDefaultIndex(dir)
	}
	return nil
}

func createDefaultIndex(basePath string) {
	logrus.Info("Creating default index.php in directory: " + basePath)

	file, err := os.Create(basePath + string(os.PathSeparator) + "index.php")
	if err != nil {
		logrus.Error("Cannot create default index.php:  " + err.Error())
		return
	}

	defer file.Close()

	absDoctRoot, _ := filepath.Abs(basePath)

	template := `<!doctype html>
<html lang="en">
  <head>
    <!-- Required meta tags -->
    <meta charset="utf-8">
    <meta name="viewport" content="width=device-width, initial-scale=1, shrink-to-fit=no">

    <!-- Bootstrap CSS -->
    <link rel="stylesheet" href="https://stackpath.bootstrapcdn.com/bootstrap/4.2.1/css/bootstrap.min.css" integrity="sha384-GJzZqFGwb1QTTN6wy59ffF1BuGJpLSa9DkKMp0DgiMDm4iYMj70gZWKYbI706tWS" crossorigin="anonymous">

    <title>EzPHP</title>
  </head>
  <body>
	<div class="container">
	<h1>Welcome to your personal web server</h1>
		<p>
		<?= "This file is located in: ` + absDoctRoot + `/index.php"; ?>
		</p>
		<?= "Replaced it with your own file"; ?>
		<p>
		<hr>
		<p>
		<?= "Php version: " . phpversion() ?>
		<p>
	</div>
    </body>
</html>`

	fmt.Fprintf(file, template)
}

func WhereIsPHP(installDir string) (string, error) {
	var phpPath string
	var err error

	phpPath, err = whereIsGlobalPHP(php.PHP_EXECUTABLE)
	if err != nil {
		phpPath, err = whereIsLocalPHP(php.PHP_EXECUTABLE, installDir)
	}

	return phpPath, err
}

func whereIsGlobalPHP(phpExe string) (string, error) {
	logrus.Info("Searching for PHP in $PATH")
	return exec.LookPath(phpExe)
}

func whereIsLocalPHP(phpExe string, target string) (string, error) {
	var err error
	absPath, _ := filepath.Abs(filepath.Dir(target))
	localPHP := absPath + string(os.PathSeparator) + "7.0.0" + string(os.PathSeparator) + phpExe

	logrus.Info("Searching for PHP in " + localPHP)

	if _, err = os.Stat(localPHP); !os.IsNotExist(err) {
		return localPHP, nil
	}

	return "", err
}
