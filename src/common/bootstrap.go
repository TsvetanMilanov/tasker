package common

import (
	"github.com/TsvetanMilanov/go-lambda-workflow/workflow"
	"github.com/TsvetanMilanov/go-simple-di/di"
)

// CreateDIContainer registers the common dependencies
// and the provided ones and returns the DI container.
// This function will panic if there is a problem when registering
// the dependencies.
func CreateDIContainer(deps ...*di.Dependency) workflow.Injector {
	c := di.NewContainer()
	commonDeps := []*di.Dependency{
		&di.Dependency{Value: &Config{}},
		&di.Dependency{Value: &HTTPClient{}},
	}
	err := c.Register(append(commonDeps, deps...)...)
	if err != nil {
		panic(err)
	}

	return c
}
