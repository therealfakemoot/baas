package main

import (
	"log"
	"os"

	"github.com/therealfakemoot/baas/butts"
)

func main() {
	fonts, err := butts.FigletFonts()
	if err != nil {
		log.Fatalln(err)
	}

	p := butts.Printer{
		Butts:  butts.Butts,
		Colors: true,
		Fonts:  fonts, // tentative option
	}

	// log.Printf("# of fonts: %d", len(fonts))

	if _, err = p.WriteButt(os.Stderr); err != nil {
		log.Printf("error generating butts: %s", err)
	}
}
