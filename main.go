package main

import (
	"bufio"
	"encoding/base64"
	"fmt"
	"image/jpeg"
	"os"
	"strings"

	"github.com/skip2/go-qrcode"
	"github.com/spf13/pflag"
	qrcodedecode "github.com/tuotoo/qrcode"
)

func readInput(inputPath string, binaryMode bool) (string, error) {
	if inputPath != "" {
		bytes, err := os.ReadFile(inputPath)
		if err != nil {
			return "", err
		}
		if binaryMode {
			return base64.StdEncoding.EncodeToString(bytes), nil
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

func decodeQRCodeImage(inputPath string, binaryMode bool, outputPath string) error {
	file, err := os.Open(inputPath)
	if err != nil {
		return err
	}
	defer file.Close()
	qrMatrix, err := qrcodedecode.Decode(file)
	if err != nil {
		return err
	}
	if binaryMode {
		decoded, err := base64.StdEncoding.DecodeString(qrMatrix.Content)
		if err != nil {
			return err
		}
		if outputPath == "" {
			fmt.Println("Decoded binary data (base64):", qrMatrix.Content)
		} else {
			err := os.WriteFile(outputPath, decoded, 0644)
			if err != nil {
				return err
			}
			fmt.Println("Decoded binary data written to:", outputPath)
		}
	} else {
		fmt.Println("Decoded QR data:", qrMatrix.Content)
	}
	return nil
}

func main() {
	inputPath := pflag.StringP("input", "i", "", "Path to input file (optional, reads from stdin if not provided)")
	outputPath := pflag.StringP("output", "o", "", "Path to output file (JPEG for encode, binary for decode)")
	ascii := pflag.BoolP("ascii", "a", false, "Display QR code in terminal as ASCII art")
	decode := pflag.BoolP("decode", "d", false, "Decode QR code from image file")
	binaryMode := pflag.BoolP("binary", "b", false, "Treat input as binary and encode as base64")

	pflag.Parse()

	if *decode {
		if *inputPath == "" {
			fmt.Fprintln(os.Stderr, "Input image path required for decoding")
			os.Exit(1)
		}
		if err := decodeQRCodeImage(*inputPath, *binaryMode, *outputPath); err != nil {
			fmt.Fprintf(os.Stderr, "Error decoding QR code: %v\n", err)
			os.Exit(1)
		}
		return
	}

	data, err := readInput(*inputPath, *binaryMode)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading input: %v\n", err)
		os.Exit(1)
	}

	if *ascii {
		if err := displayQRCodeTerminal(data); err != nil {
			fmt.Fprintf(os.Stderr, "Error displaying QR code: %v\n", err)
			os.Exit(1)
		}
	} else {
		if *outputPath == "" {
			fmt.Fprintln(os.Stderr, "Output path required for JPEG format")
			os.Exit(1)
		}
		if err := generateQRCodeJPEG(data, *outputPath); err != nil {
			fmt.Fprintf(os.Stderr, "Error generating QR code JPEG: %v\n", err)
			os.Exit(1)
		}
	}
}
