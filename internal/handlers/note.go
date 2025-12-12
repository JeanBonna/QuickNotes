package handlers

import (
	"errors"
	"fmt"
	"html/template"
	"log/slog"
	"net/http"
)

type noteHandler struct{}

func NewNoteHandler() *noteHandler {
	return &noteHandler{}
}

func (nh *noteHandler) NoteList(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	files := []string{
		"views/templates/base.html",
		"views/templates/pages/home.html",
	}
	t, err := template.ParseFiles(files...)
	if err != nil {
		http.Error(w, "Acxonteceu um erro ao execuytar essa pagina", http.StatusInternalServerError)
		return
	}

	slog.Info("Executou o hander /")
	t.ExecuteTemplate(w, "base", nil)
}

func (nh *noteHandler) NoteView(w http.ResponseWriter, r *http.Request) error {
	id := r.URL.Query().Get("id")
	if id == "" {
		return errors.New("anotação não encontrada")
	}

	files := []string{
		"views/templates/base.html",
		"views/templates/pages/note-view.html",
	}
	t, err := template.ParseFiles(files...)
	if err != nil {
		return errors.New("aconteceu um erro ao executar essa página")
	}
	return t.ExecuteTemplate(w, "base", id)

}

func (nh *noteHandler) NoteNew(w http.ResponseWriter, r *http.Request) {
	files := []string{
		"views/templates/base.html",
		"views/templates/pages/note-new.html",
	}
	t, err := template.ParseFiles(files...)
	if err != nil {
		http.Error(w, "Aconteceu um erro ao executar essa pagina", http.StatusInternalServerError)
		return
	}
	t.ExecuteTemplate(w, "base", nil)
}

func (nh *noteHandler) NoteCreate(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {

		w.Header().Set("Allow", http.MethodPost)

		http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
		return
	}
	fmt.Fprint(w, "Criando uma nova nota...")
}
