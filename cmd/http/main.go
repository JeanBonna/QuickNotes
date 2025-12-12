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

	slog.SetDefault(newLogger(os.Stderr, config.GetLevelLog()))

	slog.Info(fmt.Sprintf("Servidor rodando na porta: %s\n", config.ServerPort))

	staticHandler := http.FileServer(http.Dir("views/static"))

	noteHandler := handlers.NewNoteHandler()
	mux := http.NewServeMux()
	mux.Handle("/static/", http.StripPrefix("/static/", staticHandler))
	mux.HandleFunc("/", noteHandler.NoteList)
	mux.Handle("/note/view", handlers.HandlerWithError(noteHandler.NoteView))
	mux.HandleFunc("/note/new", noteHandler.NoteNew)
	mux.HandleFunc("/note/create", noteHandler.NoteCreate)

	if err := http.ListenAndServe(fmt.Sprintf(":%s", config.ServerPort), mux); err != nil {
		panic(err)
	}

}
