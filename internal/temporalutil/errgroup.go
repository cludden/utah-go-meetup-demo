package temporalutil

import "go.temporal.io/sdk/workflow"

type ErrGroup struct {
	g      workflow.WaitGroup
	cancel func()
	err    error
}

func NewErrGroup(ctx workflow.Context) *ErrGroup {
	ctx, cancel := workflow.WithCancel(ctx)
	return &ErrGroup{workflow.NewWaitGroup(ctx), cancel, nil}
}

func (g *ErrGroup) Go(ctx workflow.Context, fn func(workflow.Context) error) {
	g.g.Add(1)
	workflow.Go(ctx, func(ctx workflow.Context) {
		defer g.g.Done()
		if err := fn(ctx); err != nil {
			g.err = err
			if g.cancel != nil {
				g.cancel()
			}
		}
	})
}

func (g *ErrGroup) Wait(ctx workflow.Context) error {
	g.g.Wait(ctx)
	if g.cancel != nil {
		g.cancel()
	}
	return g.err
}
