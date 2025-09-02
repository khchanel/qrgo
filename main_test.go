package main

import (
	"bytes"
	"image/png"
	"testing"

	"github.com/skip2/go-qrcode"
	qrcodedecode "github.com/tuotoo/qrcode"
)

func TestQRCodeTextRoundTrip(t *testing.T) {
	input := "Hello123"
	imgBuf := new(bytes.Buffer)
	qr, err := qrcode.New(input, qrcode.Medium)
	if err != nil {
		t.Fatalf("Failed to generate QR: %v", err)
	}
	if err := png.Encode(imgBuf, qr.Image(256)); err != nil {
		t.Fatalf("Failed to encode PNG: %v", err)
	}
	decoded, err := decodeQRCodeImageFromReader(imgBuf)
	if err != nil {
		t.Fatalf("Failed to decode QR: %v", err)
	}
	if decoded != input {
		t.Errorf("Decoded text does not match original. Got: %q, Want: %q", decoded, input)
	}
}

func TestQRCodeEmptyInput(t *testing.T) {
	_, err := qrcode.New("", qrcode.Medium)
	if err == nil {
		t.Errorf("Expected error for empty input, got nil")
	}
}

func TestQRCodeEdgeCases(t *testing.T) {
	input := "!@#$%^&*()_+-=[]{}|;:'\",.<>/?"
	imgBuf := new(bytes.Buffer)
	qr, err := qrcode.New(input, qrcode.Medium)
	if err != nil {
		t.Fatalf("Failed to generate QR: %v", err)
	}
	if err := png.Encode(imgBuf, qr.Image(256)); err != nil {
		t.Fatalf("Failed to encode PNG: %v", err)
	}
	decoded, err := decodeQRCodeImageFromReader(imgBuf)
	if err != nil {
		t.Fatalf("Failed to decode QR: %v", err)
	}
	if decoded != input {
		t.Errorf("Decoded text does not match original. Got: %q, Want: %q", decoded, input)
	}
}

func decodeQRCodeImageFromReader(r *bytes.Buffer) (string, error) {
	qrMatrix, err := qrcodedecode.Decode(r)
	if err != nil {
		return "", err
	}
	return qrMatrix.Content, nil
}
