package main

import (
	"log"
	"os"

	"github.com/therealfakemoot/baas/butts"
)

func main() {
	fonts, err := butts.FigletFonts()

	p := butts.Printer{
		Butts:  []string{"butt", "butts", "booty", "fanny", "derriere"},
		Colors: true,
		Fonts:  fonts, // tentative option
		Size:   0,     // tentative option
		Writer: os.Stdout,
	}

	// log.Printf("# of fonts: %d", len(fonts))

	err = p.Butt()
	if err != nil {
		log.Printf("error generating butts: %s", err)
	}
}
