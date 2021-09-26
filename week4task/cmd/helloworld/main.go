package main

import (
	"context"
	"errors"
	"fmt"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	"golang.org/x/sync/errgroup"
)

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	// Register reflection service on gRPC server.
	s := initApp()
	reflection.Register(s)

	eg, ctx := errgroup.WithContext(context.Background())
	eg.Go(func() error {
		if err := s.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
		return nil
	})

	eg.Go(func() error {
		<-ctx.Done()
		s.Stop()
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
