package main

import (
	"bufio"
	"github.com/therealfakemoot/baas/butts"
	"io/ioutil"
	"log"
	"os/exec"
	"strings"
)

func FigletFonts() ([]string, error) {
	var (
		fonts []string
		err   error
		buf   strings.Builder
	)

	figInfo := exec.Command("figlet", "-I2")
	figInfo.Stdout = &buf

	figInfo.Start()
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

	log.Printf("fontDir: %s", fontDir)
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

func main() {
	fonts, err := FigletFonts()

	p := butts.Printer{
		Butts:  []string{"butt", "butts", "booty", "fanny", "derriere"},
		Colors: true,
		Fonts:  fonts, // tentative option
		Size:   0,     // tentative option
	}

	log.Printf("# of fonts: %d", len(fonts))

	err = p.Butt()
	if err != nil {
		log.Printf("error generating butts: %s", err)
	}
}
