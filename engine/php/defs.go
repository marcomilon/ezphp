package php

import "io"

type EzInstaller interface {
	Download(w io.Writer) error
	Install(w io.Writer) error
	WhereIs(w io.Writer) (string, error)
}

type EzServe interface {
	Serve(w io.Writer) error
}
