package app

type Listener interface {
	Listen() error
	Shutdown()
}
