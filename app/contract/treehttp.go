package contract

type TreeHTTP interface {
	SetTreeProducer(producer TreeProducer)
	Serve() error
}
