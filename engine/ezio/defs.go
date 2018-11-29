package ezio

type EzIO interface {
	Info(s string)
	Error(s string)
	Custom(tag string, s string)
	Confirm(s string) bool
	Write(p []byte) (n int, err error)
}
