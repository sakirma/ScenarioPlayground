package contract

type TreeProducer interface {
	ProduceLeaf() error
	Healthy() bool
	Ready() bool
	ShutDown()
}
