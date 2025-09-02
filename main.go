package main

import (
	"bufio"
	"flag"
	"fmt"
	"image/jpeg"
	"io/ioutil"
	"os"
	"strings"

	"github.com/skip2/go-qrcode"
	qrcodedecode "github.com/tuotoo/qrcode"
)

func main() {
	inputPath := flag.String("input", "", "Path to input file (optional, reads from stdin if not provided)")
	outputPath := flag.String("output", "", "Path to output JPEG file (required if --format=jpeg)")
	format := flag.String("format", "display", "Output format: 'jpeg' or 'display'")
	decode := flag.Bool("decode", false, "Decode QR code from image file")
	flag.Parse()

	if *decode {
		if *inputPath == "" {
			fmt.Fprintln(os.Stderr, "Input image path required for decoding")
			os.Exit(1)
		}
		file, err := os.Open(*inputPath)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error opening image file: %v\n", err)
			os.Exit(1)
		}
		defer file.Close()
		qrMatrix, err := qrcodedecode.Decode(file)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error decoding QR code: %v\n", err)
			os.Exit(1)
		}
		fmt.Println("Decoded QR data:", qrMatrix.Content)
		return
	}

	var data string
	if *inputPath != "" {
		bytes, err := ioutil.ReadFile(*inputPath)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error reading file: %v\n", err)
			os.Exit(1)
		}
		data = string(bytes)
	} else {
		reader := bufio.NewReader(os.Stdin)
		input, err := reader.ReadString('\n')
		if err != nil && err.Error() != "EOF" {
			fmt.Fprintf(os.Stderr, "Error reading stdin: %v\n", err)
			os.Exit(1)
		}
		data = strings.TrimSpace(input)
	}

	if *format == "jpeg" {
		if *outputPath == "" {
			fmt.Fprintln(os.Stderr, "Output path required for JPEG format")
			os.Exit(1)
		}
		// Generate QR code and save as JPEG
		qr, err := qrcode.New(data, qrcode.Medium)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error generating QR code: %v\n", err)
			os.Exit(1)
		}
		img := qr.Image(256)
		file, err := os.Create(*outputPath)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error creating output file: %v\n", err)
			os.Exit(1)
		}
		defer file.Close()
		if err := jpeg.Encode(file, img, nil); err != nil {
			fmt.Fprintf(os.Stderr, "Error encoding JPEG: %v\n", err)
			os.Exit(1)
		}
		fmt.Println("QR code saved as JPEG:", *outputPath)
	} else {
		// Display QR code in terminal (ASCII)
		qr, err := qrcode.New(data, qrcode.Medium)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error generating QR code: %v\n", err)
			os.Exit(1)
		}
		fmt.Println(qr.ToString(false))
	}
}
