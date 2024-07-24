package application

import (
	"context"
	"sync"
)

type Application struct {
	// Dependencies
}

func NewApplication(ctx context.Context) (*Application, error) {
	a := &Application{}

	// Register dependencies
	err := a.registerAwesome()
	if err != nil {
		return nil, err
	}

	return a, nil
}

// Close closes any stateful dependencies while shutting down
func (a *Application) Close() error {
	return nil
}

// runAsync is a helper function to execute async services
func runAsync(ctx context.Context, wg *sync.WaitGroup, fn func(ctx context.Context)) {
	wg.Add(1)
	go func() {
		defer wg.Done()
		fn(ctx)
	}()
}

func (a *Application) registerAwesome() error {
	// Initialize the dependency
	// Set the a.Awesome

	return nil
}

func (a *Application) RunAwesome(ctx context.Context, wg *sync.WaitGroup) error {
	runAsync(ctx, wg, func(ctx context.Context) {
		// Run the service
		// Gracefully shutdown using the provided context
	})

	return nil
}
