package main

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"sync"

	"github.com/Artenso/calendar/internal/app/calendar"
	"github.com/Artenso/calendar/internal/service"
	"github.com/Artenso/calendar/internal/storage"
	desc "github.com/Artenso/calendar/pkg/calendar/github.com/Artenso/calendar/pkg/calendar"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	grpcPort = ":50051"
	httpPort = ":8000"
)

func main() {
	wg := sync.WaitGroup{}

	wg.Add(1)
	go func() {
		defer wg.Done()
		if err := runGRPC(); err != nil {
			return
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()

		if err := runHTTP(); err != nil {
			return
		}
	}()

	wg.Wait()
}

func runGRPC() error {
	list, err := net.Listen("tcp", grpcPort)
	if err != nil {
		return fmt.Errorf("failed to mapping port: %s", err.Error())
	}

	repository := storage.NewStorage()
	srv := service.NewService(repository)
	calendarAPI := calendar.NewCalendar(srv)

	s := grpc.NewServer()
	desc.RegisterCalendarServer(s, calendarAPI)

	reflection.Register(s)

	if err = s.Serve(list); err != nil {
		return fmt.Errorf("failed to server: %s", err.Error())
	}

	return nil
}

func runHTTP() error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()} // nolint: staticcheck

	err := desc.RegisterCalendarHandlerFromEndpoint(ctx, mux, grpcPort, opts)
	if err != nil {
		return err
	}

	return http.ListenAndServe(httpPort, mux)
}
