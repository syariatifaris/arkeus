package inject

import (
	"reflect"

	"github.com/karlkfi/inject"
)

//GetAssignedDependencies Gets Multiple Dependencies by Type (Use Inheritance)
func GetAssignedDependencies(graph inject.Graph, lisPtr interface{}) []reflect.Value {
	return inject.FindAssignable(graph, lisPtr)
}

//GetAssignedDependency Gets Single Dependency by Type
func GetAssignedDependency(graph inject.Graph, v interface{}) reflect.Value {
	return inject.ExtractAssignable(graph, v)
}
