package install

import "github.com/marcomilon/ezphp/engine/php"

type Installer struct {
	DownloadUrl string
	Filename    string
	InstallDir  string
	php.IOChannels
}
