package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/Mind2Screen-Dev-Team/go-skeleton/bootstrap"
	"github.com/Mind2Screen-Dev-Team/go-skeleton/config"
)

func main() {
	stopCh := make(chan os.Signal, 1)
	signal.Notify(
		stopCh,
		os.Interrupt,
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT,
	)

	deps := bootstrap.Load(
		context.Background(),
		config.NewMySQLX(config.DSN{}),
	)
	_ = deps

	addr := "127.0.0.1:8081"
	handler := http.NewServeMux()

	handler.HandleFunc("GET /health/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "OK")
	})

	httpServerOption := config.NewHttpServerOption()
	httoServer, err := config.NewHTTPServer(
		addr,
		handler,

		// # Options
		httpServerOption.WithIdleTimeout(30*time.Second),
		httpServerOption.WithReadHeaderTimeout(5*time.Second),
		httpServerOption.WithReadTimeout(15*time.Second),
		httpServerOption.WithWriteTimeout(20*time.Second),
	)
	if err != nil {
		log.Fatalf("Failed to Start HTTP API, got error %+v\n", err)
	}

	srv, _ := httoServer.Create(context.Background())

	go func() {
		log.Printf("Start Service HTTP API on address %s\n", addr)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Start Service HTTP API on address %s, listen and serve : %+v\n", addr, err)
		}
		log.Printf("Stop Service HTTP API on address %s\n", addr)
	}()

	<-stopCh

	log.Printf("Perform shutdown with a maximum timeout of 5 seconds | Service HTTP API on address %s\n", addr)
	releaseCtx, releaseFn := context.WithTimeout(context.Background(), 5*time.Second)

	defer func() {
		releaseFn()
		log.Printf("Successfully Stop Service HTTP API on address %s is exited properly\n", addr)
	}()

	if err := srv.Shutdown(releaseCtx); err != nil {
		log.Printf("Shutdown Service HTTP API on address %s, err : %+v\n", addr, err)
	}
}
