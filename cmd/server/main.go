package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"sync"

	"github.com/Artenso/calendar/internal/app/calendar"
	"github.com/Artenso/calendar/internal/service"
	storage "github.com/Artenso/calendar/internal/storage/calendar_pg"
	desc "github.com/Artenso/calendar/pkg/calendar/github.com/Artenso/calendar/pkg/calendar"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/jackc/pgx/v5"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	grpcPort = ":50051"
	httpPort = ":8000"
)

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	wg := sync.WaitGroup{}

	wg.Add(1)
	go func() {
		defer wg.Done()
		if err := runGRPC(ctx); err != nil {
			return
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()

		if err := runHTTP(ctx); err != nil {
			return
		}
	}()

	wg.Wait()
}

func runGRPC(ctx context.Context) error {
	list, err := net.Listen("tcp", grpcPort)
	if err != nil {
		return fmt.Errorf("failed to mapping port: %s", err.Error())
	}

	dbDSN := "postgres://postgres:postgres@localhost:5432/calendar"
	conn, err := pgx.Connect(ctx, dbDSN)
	if err != nil {
		log.Fatalf("failed to init db connection: %s", err.Error())
	}

	repository := storage.New(conn)
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

func runHTTP(ctx context.Context) error {

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()} // nolint: staticcheck

	err := desc.RegisterCalendarHandlerFromEndpoint(ctx, mux, grpcPort, opts)
	if err != nil {
		return err
	}

	return http.ListenAndServe(httpPort, mux)
}
