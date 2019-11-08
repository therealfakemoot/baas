package butts

import (
	"io"
	"math/rand"
	"os/exec"
	"time"

	"github.com/pkg/errors"
)

func init() {
	rand.Seed(time.Now().UTC().UnixNano())
}

// Printer uses configuration values to serve butts.
type Printer struct {
	Butts  []string
	Colors bool
	Fonts  []string
	Size   int
}

type BigButt struct {
	Butt   string
	Font   string
	Output []byte
}

func NewPrinter() (*Printer, error) {
	fonts, err := FigletFonts()
	if err != nil {
		return nil, errors.Wrap(err, "Failed to get figlet fonts")
	}

	return &Printer{
		Butts:  Butts,
		Colors: true,
		Fonts:  fonts,
	}, nil
}

func (p Printer) Butt() (*BigButt, error) {
	var (
		butt = p.Butts[rand.Intn(len(p.Butts))]
		font = p.Fonts[rand.Intn(len(p.Fonts))]
	)

	text, err := figlet(butt, font)
	if err != nil {
		return nil, err
	}

	return &BigButt{
		Butt:   butt,
		Font:   font,
		Output: text,
	}, nil
}

// Butt writes an ASCII butt into the writer, returning the font
func (p Printer) WriteButt(w io.Writer) (*BigButt, error) {
	b, err := p.Butt()
	if err != nil {
		return nil, err
	}

	_, err = w.Write(b.Output)
	return b, err
}

func figlet(text, font string) ([]byte, error) {
	figlet := exec.Command("figlet",
		"-f", font,
		"-k",
		"-w", "120",
		text,
	)

	return figlet.Output()
}
