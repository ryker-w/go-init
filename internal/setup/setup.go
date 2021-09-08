package setup

import (
	"context"
	"github.com/ryker-w/go-init/internal/process"
)

type componentHandler func(ctx context.Context) (err error)

func SomeWork(ctx context.Context) (err error) {

	var handlers []componentHandler
	handlers = append(handlers,
		// ctx handlers
		process.AmqpStart,
	)
	// 执行
	for _, h := range handlers {
		if err = h(ctx); err != nil {
			return
		}
	}

	components := []func() error{
		// func handlers
	}
	// 执行
	for _, component := range components {
		if err = component(); err != nil {
			return
		}
	}
	return
}
