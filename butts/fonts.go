package butts

import (
	"io/ioutil"
	"log"
	"os/exec"
	"path/filepath"
	"strings"
)

func FigletFonts() ([]string, error) {
	figInfo := exec.Command("figlet", "-I2")

	output, err := figInfo.Output()
	if err != nil {
		log.Printf("figlet -I2 failed: %v", err)
		return nil, err
	}

	dir := strings.TrimSpace(string(output))

	// log.Printf("fontDir: %s", dir)
	files, err := ioutil.ReadDir(string(dir))
	if err != nil {
		log.Printf("error reading fontDir: %v", err)
		return nil, err
	}

	fonts := make([]string, 0, len(files))

	for _, f := range files {
		var (
			name = f.Name()
			ext  = filepath.Ext(name)
		)

		if ext == ".flf" {
			fonts = append(fonts, name[:len(name)-len(ext)])
		}
	}

	return fonts, err
}
