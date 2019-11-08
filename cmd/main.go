package main

import (
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"

	"github.com/therealfakemoot/baas/butts"
)

func main() {
	p, err := butts.NewPrinter()
	if err != nil {
		log.Fatalln(err)
	}

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Get("/fonts", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(strings.Join(p.Fonts, "\n")))
	})

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		b, err := p.WriteButt(w)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			log.Printf("error generating butts: %s", err)
		}

		w.Header().Add("X-figlet-font", b.Font)
		w.Header().Add("X-figlet-butt", b.Butt)
	})

	var addr = ":8008"
	if len(os.Args) > 1 {
		addr = os.Args[1]
	}

	if err := http.ListenAndServe(addr, r); err != nil {
		log.Fatalln("Failed listening to address", addr+":", err)
	}

	// log.Printf("# of fonts: %d", len(fonts))
}
