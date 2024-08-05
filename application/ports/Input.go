package ports

type InputPort[P any, R any] interface {
	Execute(*P) (*R, error)
}
