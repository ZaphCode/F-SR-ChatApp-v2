package app

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

type App struct {
	http.Server
	onShutdown func()
}

func New(addr uint) *App {
	return &App{
		http.Server{
			Addr: fmt.Sprintf(":%d", addr),
		}, nil,
	}
}

func (a *App) setGracefulShutdown(ctx context.Context) {
	signalChan := make(chan os.Signal, 1)

	signal.Notify(
		signalChan, syscall.SIGTERM, syscall.SIGINT,
		syscall.SIGQUIT, syscall.SIGTSTP,
	)

	go func() {
		<-signalChan

		if a.onShutdown != nil {
			a.onShutdown()
		}

		a.Server.Shutdown(ctx)
	}()
}

func (a *App) RegisterHandlers(handlers ...Handler) {
	mux := http.NewServeMux()

	mux.Handle("/public",
		http.StripPrefix("/public/", http.FileServer(http.Dir("public"))),
	)

	for _, handler := range handlers {
		handler.SetRoutes(mux)
	}

	a.Handler = mux
}

func (a *App) Run(ctx context.Context) {
	InitSessionStore()

	a.setGracefulShutdown(ctx)

	log.Printf("Serving on http://localhost%s", a.Addr)

	if err := a.Server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}

func (a *App) OnShutdown(fn func()) {
	a.onShutdown = fn
}
