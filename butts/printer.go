package butts

import (
	"log"
	"math/rand"
	"os"
	"os/exec"
	"time"
)

var (
	noise = rand.New(rand.NewSource(time.Now().UTC().UnixNano()))
)

// Printer uses configuration values to serve butts.
type Printer struct {
	Butts  []string
	Colors bool
	Fonts  []string
	Size   int
}

// Butt returns an ascii butt
func (p Printer) Butt() error {
	idx := noise.Intn(len(p.Butts))
	font := p.Fonts[noise.Intn(len(p.Fonts))]
	// log.Printf("font arg: %s", font)
	args := []string{
		"-f",
		font,
		"-k",
		p.Butts[idx],
	}

	log.Println("starting figlet")
	figlet := exec.Command("figlet", args...)
	figlet.Stdout = os.Stdout
	figlet.Stderr = os.Stderr

	// log.Printf("cmd: %s", figlet.Path)
	// log.Printf("args: %s", figlet.Args)

	err := figlet.Start()
	if err != nil {
		return err
	}

	return figlet.Wait()
}
