package inject

import "go.uber.org/dig"

var container *dig.Container

func Default() *dig.Container {
	if container == nil {
		container = dig.New()
	}
	return container
}
