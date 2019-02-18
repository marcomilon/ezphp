package serve

import "github.com/marcomilon/ezphp/engine/php"

type Server struct {
	PhpExe  string
	Host    string
	DocRoot string
	php.IOChannels
}
