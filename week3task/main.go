package main

import (
	"context"
	"errors"
	"fmt"
	"golang.org/x/sync/errgroup"
	"io"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	eg, ctx := errgroup.WithContext(context.Background())

	srv := &http.Server{Addr: ":8080"}
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "hello world\n")
	})
	eg.Go(func() error {
		if err := srv.ListenAndServe(); err != nil {
			return err
		}
		return nil
	})

	eg.Go(func() error {
		<-ctx.Done()
		if err := srv.Shutdown(nil); err != nil {
			return err
		}
		return ctx.Err()
	})

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	eg.Go(func() error {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-sig:
			return errors.New("kill process")
		}

	})
	if err := eg.Wait(); err != nil {
		fmt.Println(err)
	}
}
