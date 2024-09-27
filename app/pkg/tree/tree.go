package tree

import "ScenarioPlayground/contract"

type Tree struct {
	server   contract.TreeHTTP
	producer contract.TreeProducer
}

func New(treeHTTP contract.TreeHTTP) *Tree {
	return &Tree{server: treeHTTP}
}

func (t *Tree) Serve() error {
	return t.server.Serve()
}
