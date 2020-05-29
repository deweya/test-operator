package controller

import (
	"github.com/deweya/test-operator/pkg/controller/helloworld"
)

func init() {
	// AddToManagerFuncs is a list of functions to create controllers and add them to a manager.
	AddToManagerFuncs = append(AddToManagerFuncs, helloworld.Add)
}
