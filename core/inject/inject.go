package inject

import (
	"reflect"

	"github.com/syariatifaris/arkeus/core/inject/karlfki"
)

//NewDependencyInjection creates injection instances
func NewDependencyInjection() Injection {
	return karlfki.NewDependencyInjection()

}

//Injection contract
type Injection interface {
	AddDependency(ptr interface{}, constructor interface{}, argPtrs ...interface{}) interface{}
	GetAddedDependencies(listPtr interface{}) []reflect.Value
	GetAddedDependency(ptr interface{}) reflect.Value
}
