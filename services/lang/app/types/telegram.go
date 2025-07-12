package types

type ITelegram interface {
	Start() error
	Stop() error
}
