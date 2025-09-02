
# QRGo: QR Code Generator & Decoder CLI Tool

A simple command-line tool written in Go to generate QR codes from file or stdin input, and decode QR codes from image files. Output as JPEG or display in terminal.

## Prerequisites
- Go 1.18 or newer

## Installation
1. Clone or download this repository.
2. Install dependencies:
   ```powershell
   go mod tidy
   ```
3. Build the executable:
   ```powershell
   go build -o qrgo.exe main.go
   ```

## Usage

### Display QR code in terminal (ASCII art)
```powershell
# From file
./qrgo.exe --input path/to/input.txt --ascii

# From stdin
Get-Content path/to/input.txt | ./qrgo.exe --ascii
```

### Generate QR code and save as JPEG
```powershell
# From file
./qrgo.exe --input path/to/input.txt --output qr_output.jpg

# From stdin
Get-Content path/to/input.txt | ./qrgo.exe --output qr_output.jpg
```

### Decode QR code from image file
```powershell
./qrgo.exe --input qr_output.jpg --decode
```

## Arguments
- `--input`, `-i`   : Path to input file (for encoding) or image file (for decoding)
- `--output`, `-o`  : Path to output JPEG file (required for JPEG output)
- `--ascii`, `-a`   : Display QR code in terminal as ASCII art
- `--decode`, `-d`  : Decode QR code from image file (prints decoded content)
- `--binary`, `-b`  : Treat input as binary and encode as base64

## Example
```powershell
./qrgo.exe --input "hello.txt" --ascii
./qrgo.exe --input "hello.txt" --output "hello.jpg"
./qrgo.exe --input "hello.jpg" --decode
```

## License
MIT
