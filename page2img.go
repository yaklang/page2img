package main

import (
	"bytes"
	"flag"
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"os"
	"path/filepath"
	"strings"

	"github.com/gen2brain/go-fitz"
	"github.com/yaklang/page2img/common/log"
)

// 默认目标文件大小 (300 KB)
const defaultTargetSizeKB = 300

// encodeJPEGWithTargetSize 使用二分查找来寻找最佳JPEG质量，以满足目标文件大小。
// 它返回编码后的图像数据和最终选择的质量值。
func encodeJPEGWithTargetSize(img image.Image, targetSize int) ([]byte, int, error) {
	var bestImageBytes []byte
	var bestQuality int

	// 二分查找的范围是质量 1-100
	lowQuality := 1
	highQuality := 100

	// 为了确保即使最低质量也超大的情况下有输出，先用最低质量编码一次
	var buf bytes.Buffer
	err := jpeg.Encode(&buf, img, &jpeg.Options{Quality: lowQuality})
	if err != nil {
		return nil, 0, fmt.Errorf("failed to encode with lowest quality: %w", err)
	}
	bestImageBytes = buf.Bytes()
	bestQuality = lowQuality

	// 如果最低质量已经小于目标，我们才开始寻找更高的质量
	if len(bestImageBytes) < targetSize {
		// 开始二分查找
		for lowQuality <= highQuality {
			midQuality := lowQuality + (highQuality-lowQuality)/2
			if midQuality == 0 {
				break
			}

			buf.Reset() // 重置缓冲区以供下次使用

			err := jpeg.Encode(&buf, img, &jpeg.Options{Quality: midQuality})
			if err != nil {
				// 编码失败，可以认为这个质量值不可用
				highQuality = midQuality - 1
				continue
			}

			currentSize := len(buf.Bytes())

			if currentSize <= targetSize {
				// 当前大小满足要求，这是一个可能的最佳结果
				// 保存下来，然后尝试寻找更高的质量
				bestImageBytes = buf.Bytes()
				bestQuality = midQuality
				lowQuality = midQuality + 1
			} else {
				// 当前大小超标，必须降低质量
				highQuality = midQuality - 1
			}
		}
	}

	return bestImageBytes, bestQuality, nil
}

func main() {
	var inputFile, outputPattern string
	var targetSizeKB int
	flag.StringVar(&inputFile, "i", "", "input document file")
	flag.StringVar(&outputPattern, "o", "", "output image pattern (e.g., image-%d.jpg)")
	flag.IntVar(&targetSizeKB, "s", defaultTargetSizeKB, "target file size in KB (default: 300)")
	flag.Parse()

	if inputFile == "" || outputPattern == "" {
		flag.Usage()
		os.Exit(1)
	}

	// 将KB转换为字节
	targetSizeBytes := targetSizeKB * 1024

	log.Info("Opening document file:", inputFile)
	log.Info("Target file size:", targetSizeKB, "KB")
	doc, err := fitz.New(inputFile)
	if err != nil {
		log.Fatal("Failed to open document:", err)
	}
	defer doc.Close()

	outputExt := strings.ToLower(filepath.Ext(outputPattern))
	if outputExt != ".jpeg" && outputExt != ".jpg" && outputExt != ".png" {
		log.Fatal("Unsupported output format. Please use .jpeg, .jpg, or .png")
	}
	outputPrefix := strings.TrimSuffix(outputPattern, filepath.Ext(outputPattern))

	log.Info(fmt.Sprintf("Document has %d pages.", doc.NumPage()))

	processImage := func(img *image.RGBA, pageNum int) {
		outputFile := fmt.Sprintf(strings.Replace(outputPrefix, "%d", "%d", -1), pageNum) + outputExt

		tempFile := outputFile + ".tmp"
		f, err := os.Create(tempFile)
		if err != nil {
			log.Error(fmt.Sprintf("Failed to create output file %s: %v", tempFile, err))
			return
		}
		defer f.Close()

		if outputExt == ".jpeg" || outputExt == ".jpg" {
			// 使用新函数来编码JPEG
			jpegBytes, finalQuality, err := encodeJPEGWithTargetSize(img, targetSizeBytes)
			if err != nil {
				log.Error(fmt.Sprintf("Failed to encode JPEG for page %d: %v", pageNum, err))
				return
			}

			_, err = f.Write(jpegBytes)
			if err != nil {
				log.Error(fmt.Sprintf("Failed to write image data to %s: %v", tempFile, err))
				return
			}

			log.Info(fmt.Sprintf("Saved page %d (Final Quality: %d, Size: %.2f KB)",
				pageNum, finalQuality, float64(len(jpegBytes))/1024.0))

		} else if outputExt == ".png" {
			err = png.Encode(f, img)
			if err != nil {
				log.Error(fmt.Sprintf("Failed to save PNG image %s: %v", tempFile, err))
				return
			}
		}
		err = os.Rename(tempFile, outputFile)
		if err != nil {
			log.Error(fmt.Sprintf("Failed to rename %s to %s: %v", tempFile, outputFile, err))
			return
		}
		info, _ := os.Stat(outputFile)
		log.Info(fmt.Sprintf("Saved page %d to %s (Size: %.2f KB)",
			pageNum, outputFile, float64(info.Size())/1024.0))

	}

	for n := 0; n < doc.NumPage(); n++ {
		pageNum := n + 1
		log.Info(fmt.Sprintf("Processing page %d", pageNum))

		// 这里可以增加渲染时的DPI来提高清晰度，例如 doc.Image(n, 150)
		img, err := doc.Image(n) // 默认DPI是72
		if err != nil {
			log.Error(fmt.Sprintf("Failed to get image for page %d: %v", pageNum, err))
			continue
		}
		processImage(img, pageNum)
	}

	log.Info("All pages processed.")
}
