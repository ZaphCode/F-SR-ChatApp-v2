package app

import (
	"context"
	"html/template"
	"log"
	"net/http"
	"os"
	"os/signal"
	"path/filepath"
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

	signal.Notify(
		signalChan, syscall.SIGTERM, syscall.SIGINT,
		syscall.SIGQUIT, syscall.SIGTSTP,
	)

	go func() {
		<-signalChan
		log.Println("Shutting down...")
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

func RenderTmpl(w http.ResponseWriter, tmplName string, data any) error {
	tmplPath := filepath.Join("templates", tmplName+".html")

	tmplFiles := []string{
		tmplPath,
		// More components...
		filepath.Join("templates", "components", "navbar.html"),
	}

	t, err := template.ParseFiles(tmplFiles...)

	if err != nil {
		return err
	}

	return t.ExecuteTemplate(w, tmplName, data)
}

func (a *App) Run(ctx context.Context) {
	log.Printf("Serving on http://localhost%s", a.Addr)

	a.setGracefulShutdown(ctx)

	if err := a.Server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
