package signal

import (
	"context"
	"os"
	"os/signal"
	"syscall"
)

var (
	signalChan      = make(chan struct{})
	shutdownSignals = []os.Signal{syscall.SIGTERM, syscall.SIGINT}
)

type Context struct {
	funcs []func()
}

func NewContext() *Context {
	return &Context{
		funcs: make([]func(), 0),
	}
}

func NewContextWithFuncs(funcs ...func()) *Context {
	return &Context{
		funcs: funcs,
	}
}

func (c *Context) WithFunc(funcs ...func()) {
	c.funcs = append(c.funcs, funcs...)
}

func (c *Context) SignalHandler() context.Context {
	close(signalChan)

	ctx, cancel := context.WithCancel(context.Background())
	osChan := make(chan os.Signal, 1)
	signal.Notify(osChan, shutdownSignals...)

	go func() {
		<-osChan
		cancel()

		for _, f := range c.funcs {
			f()
		}
		os.Exit(0)
	}()
	return ctx
}

func (c *Context) Exit() {
	for _, f := range c.funcs {
		f()
	}
	os.Exit(0)
}
