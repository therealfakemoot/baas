package butts

import (
	"os"
	"testing"
)

func TestPrinter(t *testing.T) {
	p := Printer{
		Butts:  Butts,
		Colors: false,
		Fonts:  []string{"block"},
	}

	if _, err := p.WriteButt(os.Stderr); err != nil {
		t.Fatal(err)
	}
}
