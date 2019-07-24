package index

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"

	"github.com/npthinhdev/valexer/internal/app/types"
	"github.com/npthinhdev/valexer/internal/pkg/parse"
)

type (
	val map[string]interface{}
	// Handler is web handler
	Handler struct{}
)

var tmlp = template.Must(template.ParseGlob("web/template/*.html"))

// New return new web handler
func New() *Handler {
	return &Handler{}
}

// Index render index page
func (h *Handler) Index(w http.ResponseWriter, r *http.Request) {
	ctx := val{"Title": "Homepage"}
	apiURL := "http://localhost:8080/api/exercise"
	body := parse.Get(apiURL)
	exers := []types.Exercise{}
	_ = json.Unmarshal(body, &exers)
	ctx["Exers"] = exers
	err := tmlp.ExecuteTemplate(w, "index.html", ctx)
	if err != nil {
		log.Println(err)
	}
}
