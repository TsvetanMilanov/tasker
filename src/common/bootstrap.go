package common

import (
	"github.com/TsvetanMilanov/go-lambda-workflow/workflow"
	"github.com/TsvetanMilanov/go-simple-di/di"
)

func createDIContainer(deps ...*di.Dependency) workflow.Injector {
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
