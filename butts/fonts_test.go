package butts

import (
	"log"
	"testing"
)

func TestFigletFonts(t *testing.T) {
	fonts, err := FigletFonts()
	if err != nil {
		t.Fatal(err)
	}

	log.Println("Figlet fonts:", fonts)
}
