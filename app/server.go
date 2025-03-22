package app

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

type App struct {
	http.Server
}

func New(addr string) *App {
	return &App{
		http.Server{
			Addr: addr,
		},
	}
}

func (a *App) setGracefulShutdown(ctx context.Context) {
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGTERM, syscall.SIGINT, syscall.SIGQUIT, syscall.SIGTSTP)

	go func() {
		<-signalChan
		log.Println("Shutting down...")
		a.Server.Shutdown(ctx)
	}()
}

func (a *App) RegisterHandlers(handlers ...Handler) {
	mux := http.NewServeMux()

	for _, handler := range handlers {
		handler.SetRoutes(mux)
	}

	a.Handler = mux
}

func (a *App) Run(ctx context.Context) {
	log.Printf("Serving on: http://localhost%s", a.Addr)

	a.setGracefulShutdown(ctx)

	if err := a.Server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
