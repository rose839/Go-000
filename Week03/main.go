package main

import (
	"context"
	xerrors "errors"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/pkg/errors"
	"golang.org/x/sync/errgroup"
)

// pre-defined errors
var ErrRecvSignal = xerrors.New("Recive signal")
var ErrHttpServer = xerrors.New("Recive signal")

// http handler
type httpHandler struct{}

func (h *httpHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "hello world")
}

func main() {
	// signal
	sigChan := make(chan os.Signal)
	sigNotifyCh := make(chan struct{})
	signal.Notify(sigChan, syscall.SIGINT)

	// http server
	var addr = "0.0.0.0:8080"
	var server *http.Server
	httpNotifyCh := make(chan struct{})

	// create an error group
	group, _ := errgroup.WithContext(context.Background())

	// check whether any one goroutine has quit and notify the other one
	group.Go(func() error {
		select {
		case <-httpNotifyCh:
		case <-sigNotifyCh:
		}
		server.Shutdown(context.Background())
		signal.Stop(sigChan)
		close(sigChan)
		return nil
	})

	// start http server
	group.Go(func() error {
		// notify quit
		defer close(httpNotifyCh)

		handler := &httpHandler{}
		server = &http.Server{Addr: addr, Handler: handler}
		if err := server.ListenAndServe(); err != nil {
			return errors.WithMessage(ErrHttpServer, err.Error())
		}

		return nil
	})

	// start signal proc
	group.Go(func() error {
		for {
			select {
			case <-sigChan:
				// notify quit
				close(sigNotifyCh)
				return ErrRecvSignal
			}
		}

		return nil
	})

	if err := group.Wait(); err != nil {
		fmt.Printf("Error occured: %s", err)
	}
}
