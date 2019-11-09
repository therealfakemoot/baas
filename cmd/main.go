package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"

	"github.com/therealfakemoot/baas/butts"
)

func main() {
	fonts, err := butts.FigletFonts()

	p := butts.Printer{
		Butts:  []string{"butt", "butts", "booty", "derriere", "ass"},
		Colors: true,
		Fonts:  fonts, // tentative option
		Size:   0,     // tentative option
	}

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Get("/fonts", func(w http.ResponseWriter, r *http.Request) {
		for _, f := range fonts {
			w.Write([]byte(f + "\n"))
		}
	})

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		p.Writer = w
		err = p.Butt()
		if err != nil {
			log.Printf("error generating butts: %s", err)
		}
	})

	http.ListenAndServe(":8008", r)
	// log.Printf("# of fonts: %d", len(fonts))

}
