package main

import (
	"flag"
	"fmt"
	"image/jpeg"
	"image/png"
	"os"
	"path/filepath"
	"strings"

	"github.com/yaklang/pdf2img/common/log"

	"github.com/gen2brain/go-fitz"
)

func main() {
	var inputFile, outputPattern string
	flag.StringVar(&inputFile, "i", "", "input PDF file")
	flag.StringVar(&outputPattern, "o", "", "output image pattern (e.g., image-%d.jpg)")
	flag.Parse()

	if inputFile == "" || outputPattern == "" {
		flag.Usage()
		os.Exit(1)
	}

	log.Info("Opening PDF file:", inputFile)
	doc, err := fitz.New(inputFile)
	if err != nil {
		log.Fatal("Failed to open PDF:", err)
	}
	defer doc.Close()

	outputExt := strings.ToLower(filepath.Ext(outputPattern))
	if outputExt != ".jpeg" && outputExt != ".jpg" && outputExt != ".png" {
		log.Fatal("Unsupported output format. Please use .jpeg, .jpg, or .png")
	}
	outputPrefix := strings.TrimSuffix(outputPattern, filepath.Ext(outputPattern))

	log.Info(fmt.Sprintf("PDF has %d pages.", doc.NumPage()))

	for n := 0; n < doc.NumPage(); n++ {
		pageNum := n + 1
		log.Info(fmt.Sprintf("Processing page %d", pageNum))

		img, err := doc.Image(n)
		if err != nil {
			log.Error(fmt.Sprintf("Failed to get image for page %d: %v", pageNum, err))
			continue
		}

		outputFile := fmt.Sprintf(strings.Replace(outputPrefix, "%d", "%d", -1), pageNum) + outputExt
		f, err := os.Create(outputFile)
		if err != nil {
			log.Error(fmt.Sprintf("Failed to create output file %s: %v", outputFile, err))
			continue
		}

		if outputExt == ".jpeg" || outputExt == ".jpg" {
			err = jpeg.Encode(f, img, &jpeg.Options{Quality: 95})
		} else if outputExt == ".png" {
			err = png.Encode(f, img)
		}

		if err != nil {
			log.Error(fmt.Sprintf("Failed to save image %s: %v", outputFile, err))
			f.Close()
			continue
		}

		f.Close()
		log.Info("Saved page", pageNum, "to", outputFile)
	}

	log.Info("All pages processed.")
}
