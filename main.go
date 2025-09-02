package main

import (
	"bufio"
	"fmt"
	"image/jpeg"
	"io/ioutil"
	"os"
	"strings"

	"github.com/skip2/go-qrcode"
	"github.com/spf13/pflag"
	qrcodedecode "github.com/tuotoo/qrcode"
)

func readInput(inputPath string) (string, error) {
	if inputPath != "" {
		bytes, err := ioutil.ReadFile(inputPath)
		if err != nil {
			return "", err
		}
		return string(bytes), nil
	}
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil && err.Error() != "EOF" {
		return "", err
	}
	return strings.TrimSpace(input), nil
}

func generateQRCodeJPEG(data, outputPath string) error {
	qr, err := qrcode.New(data, qrcode.Medium)
	if err != nil {
		return err
	}
	img := qr.Image(256)
	file, err := os.Create(outputPath)
	if err != nil {
		return err
	}
	defer file.Close()
	if err := jpeg.Encode(file, img, nil); err != nil {
		return err
	}
	fmt.Println("QR code saved as JPEG:", outputPath)
	return nil
}

func displayQRCodeTerminal(data string) error {
	qr, err := qrcode.New(data, qrcode.Low)
	if err != nil {
		return err
	}
	qr.DisableBorder = true
	fmt.Println(qr.ToString(false))
	return nil
}

func decodeQRCodeImage(inputPath string) error {
	file, err := os.Open(inputPath)
	if err != nil {
		return err
	}
	defer file.Close()
	qrMatrix, err := qrcodedecode.Decode(file)
	if err != nil {
		return err
	}
	fmt.Println("Decoded QR data:", qrMatrix.Content)
	return nil
}

func main() {
	inputPath := pflag.StringP("input", "i", "", "Path to input file (optional, reads from stdin if not provided)")
	outputPath := pflag.StringP("output", "o", "", "Path to output JPEG file (required if --format=jpeg)")
	format := pflag.StringP("format", "f", "display", "Output format: 'jpeg' or 'display'")
	decode := pflag.BoolP("decode", "d", false, "Decode QR code from image file")

	pflag.Parse()

	if *decode {
		if *inputPath == "" {
			fmt.Fprintln(os.Stderr, "Input image path required for decoding")
			os.Exit(1)
		}
		if err := decodeQRCodeImage(*inputPath); err != nil {
			fmt.Fprintf(os.Stderr, "Error decoding QR code: %v\n", err)
			os.Exit(1)
		}
		return
	}

	data, err := readInput(*inputPath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading input: %v\n", err)
		os.Exit(1)
	}

	if *format == "jpeg" {
		if *outputPath == "" {
			fmt.Fprintln(os.Stderr, "Output path required for JPEG format")
			os.Exit(1)
		}
		if err := generateQRCodeJPEG(data, *outputPath); err != nil {
			fmt.Fprintf(os.Stderr, "Error generating QR code JPEG: %v\n", err)
			os.Exit(1)
		}
	} else {
		if err := displayQRCodeTerminal(data); err != nil {
			fmt.Fprintf(os.Stderr, "Error displaying QR code: %v\n", err)
			os.Exit(1)
		}
	}
}
