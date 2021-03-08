package onsignal

import (
	"context"
	"os"
	"os/signal"
)

//Do allows to launch a function when a signal is caught.
//If context is Done, signal handling stops.
func Do(ctx context.Context, signals []os.Signal, f func()) {
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, signals...)

	go func() {
		defer signal.Stop(sigChan)
		for {
			select {
			case <-sigChan:
				f()
			case <-ctx.Done():
				return
			}
		}
	}()
}

//DoAndStop allows to call function when a signal is caught and then stop handling signals
//If context is Done, signal handling stops.
func DoAndStop(ctx context.Context, signals []os.Signal, f func()) {
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, signals...)

	go func() {
		defer signal.Stop(sigChan)
		select {
		case <-sigChan:
			f()
		case <-ctx.Done():
			return
		}
	}()
}

//DoAndExit allows to call a function when a signal is caught and then exit the main program
//If context is Done, signal handling stops.
func DoAndExit(ctx context.Context, signals []os.Signal, f func(), exitCode int) {
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, signals...)

	go func() {
		defer signal.Stop(sigChan)
		select {
		case <-sigChan:
			f()
			os.Exit(exitCode)
		case <-ctx.Done():
			return
		}
	}()
}
