package news

import (
	"net/http"
	
	"github.com/go-chi/chi"
	"github.com/unduu/ibadf/app/handler"
)

func Router() *chi.Mux {
	router := chi.NewRouter()

	router.Get("/{newsid}", func(w http.ResponseWriter, r *http.Request) {
		newsidParam := chi.URLParam(r, "newsid")
		handler.GetNews(w, r, newsidParam)
	})

	return router
}

