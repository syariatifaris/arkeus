package karlfki

import (
	"reflect"

	"github.com/karlkfi/inject"
)

func NewDependencyInjection() *DIImpl {
	return &DIImpl{
		graph: inject.NewGraph(),
	}
}

type DIImpl struct {
	graph inject.Graph
}

func (d *DIImpl) AddDependency(ptr interface{}, constructor interface{}, argPtrs ...interface{}) interface{} {
	return d.graph.Define(ptr, inject.NewProvider(constructor, argPtrs...))
}

func (d *DIImpl) GetAddedDependencies(listPtr interface{}) []reflect.Value {
	return inject.FindAssignable(d.graph, listPtr)
}

func (d *DIImpl) GetAddedDependency(ptr interface{}) reflect.Value {
	return inject.ExtractAssignable(d.graph, ptr)
}
