package php_test

import (
	"testing"

	"github.com/marcomilon/ezphp/internal/php"
)

func TestServer(t *testing.T) {

	ioCom := php.IOCom{
		make(chan string),
		make(chan string),
		make(chan string),
		make(chan bool),
	}

	srv := php.Server{
		Exec:    "/bin/echo",
		Host:    ":3001",
		DocRoot: "/var/www",
	}

	go srv.Serve(ioCom)

	outmsg := <-ioCom.Stdout
	expected := "-S :3001 -t /var/www\n"

	if outmsg != expected {
		t.Errorf("expected %v; got %v", expected, outmsg)
	}

}
