package setup

import (
	"context"
	"fmt"
)

type contextHandler func(ctx context.Context) error
type componentHandler func() error

// 程序启动后做点事情。初始化执行
func SomeWork(ctx context.Context) (err error) {

	//component Handler
	var componentHandlers []componentHandler
	componentHandlers = append(componentHandlers,
		componentDemo,
	)

	for _, handler := range componentHandlers {
		if err = handler(); err != nil {
			return
		}
	}

	// context Handler
	var contextHandler []contextHandler
	contextHandler = append(contextHandler,
		contextDemo,
	)
	for _, handler := range contextHandler {
		if err = handler(ctx); err != nil {
			return
		}
	}
	return
}

func componentDemo() error {
	fmt.Println("handle componentDemo")
	return nil
}
func contextDemo(ctx context.Context) error {
	fmt.Println("handle contextDemo")
	return nil
}
