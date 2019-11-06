package butts

import (
	"bufio"
	"io"
	"io/ioutil"
	"log"
	"math/rand"
	"os/exec"
	"strings"
	"time"
)

var (
	noise = rand.New(rand.NewSource(time.Now().UTC().UnixNano()))
)

func FigletFonts() ([]string, error) {
	var (
		fonts []string
		err   error
		buf   strings.Builder
	)

	figInfo := exec.Command("figlet", "-I2")
	figInfo.Stdout = &buf

	err = figInfo.Start()
	if err != nil {
		log.Printf("figInfo error: %s", err)
		return fonts, err
	}
	err = figInfo.Wait()
	if err != nil {
		log.Printf("figInfo error: %s", err)
		return fonts, err
	}

	r := strings.NewReader(buf.String())
	scanner := bufio.NewScanner(r)

	var fontDir string
	for scanner.Scan() {
		fontDir = scanner.Text()
		break
	}

	// log.Printf("fontDir: %s", fontDir)
	files, err := ioutil.ReadDir(fontDir)

	if err != nil {
		log.Printf("error reading fontDir: %s", err)
	}

	var split []string
	for _, f := range files {
		if strings.HasSuffix(f.Name(), ".flf") {
			split = strings.Split(f.Name(), ".")
			fonts = append(fonts, split[0])
		}
	}

	return fonts, err
}

// Printer uses configuration values to serve butts.
type Printer struct {
	Butts  []string
	Colors bool
	Fonts  []string
	Size   int
	Writer io.Writer
}

// Butt returns an ascii butt
func (p Printer) Butt() error {
	idx := noise.Intn(len(p.Butts))
	font := p.Fonts[noise.Intn(len(p.Fonts))]

	args := []string{
		"-f",
		font,
		"-k",
		"-w",
		"120",
		p.Butts[idx],
	}

	// "font: " + font + "\n" +

	// log.Println("starting figlet")
	figlet := exec.Command("figlet", args...)
	figlet.Stdout = p.Writer
	figlet.Stderr = p.Writer

	p.Writer.Write([]byte("font: " + font + "\n"))

	err := figlet.Start()
	if err != nil {
		return err
	}

	return figlet.Wait()
}
