package main

import (
	"fmt"
	"log/slog"
	"net/http"
	"os"

	"github.com/jeanbonna/quicknotes/internal/handlers"
)

func main() {
	config := loadConfig()

	// logger := newLogger(os.Stderr, slog.LevelInfo)

	slog.SetDefault(newLogger(os.Stderr, config.GetLevelLog()))

	slog.Info(fmt.Sprintf("DB_PASSWORD: %s\n", config.DBPassword))
	slog.Info(fmt.Sprintf("Servidor rodando na porta: %s\n", config.ServerPort))

	staticHandler := http.FileServer(http.Dir("views/static"))

	mux := http.NewServeMux()
	mux.Handle("/static/", http.StripPrefix("/static/", staticHandler))
	mux.HandleFunc("/", handlers.NoteList)
	mux.HandleFunc("/note/view", handlers.NoteView)
	mux.HandleFunc("/note/new", handlers.NoteNew)
	mux.HandleFunc("/note/create", handlers.NoteCreate)

	http.ListenAndServe(fmt.Sprintf(":%s", config.ServerPort), mux)

}
