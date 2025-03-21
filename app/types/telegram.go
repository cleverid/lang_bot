package types

type ITelegram interface {
	Start()
	Stop() error
}
