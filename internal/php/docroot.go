package php

import (
	"fmt"
	"os"
)

var template = `<!doctype html>
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
	<?= 'This file is located in: ...' ?>
	</p>
	<?= "Replace it with your own file"; ?>
	<p>
	<hr>
	<p>
	<?= "Php version: " . phpversion() ?>
	<p>
</div>
</body>
</html>`

func Create(path string) error {
	return os.MkdirAll(path, 0755)
}

func Exists(path string) bool {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return false
	}

	return true
}

func CreateIndex(path, template string) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}

	defer file.Close()

	fmt.Fprintf(file, template)

	return nil
}
