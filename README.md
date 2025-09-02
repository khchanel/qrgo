# GoQR: QR Code Generator & Decoder CLI Tool

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
   go build -o goqr.exe main.go
   ```

## Usage

### Generate and display QR code in terminal
```powershell
# From file
./goqr.exe --input path/to/input.txt --format display

# From stdin
Get-Content path/to/input.txt | ./goqr.exe --format display
```

### Generate QR code and save as JPEG
```powershell
# From file
./goqr.exe --input path/to/input.txt --format jpeg --output qr_output.jpg

# From stdin
Get-Content path/to/input.txt | ./goqr.exe --format jpeg --output qr_output.jpg
```

### Decode QR code from image file
```powershell
./goqr.exe --input qr_output.jpg --decode
```

## Arguments
- `--input`   : Path to input file (for encoding) or image file (for decoding)
- `--format`  : Output format: `jpeg` or `display` (default: `display`, for encoding only)
- `--output`  : Path to output JPEG file (required if `--format=jpeg`, for encoding only)
- `--decode`  : Decode QR code from image file (prints decoded content)

## Example
```powershell
./goqr.exe --input "hello.txt" --format display
./goqr.exe --input "hello.txt" --format jpeg --output "hello.jpg"
./goqr.exe --input "hello.jpg" --decode
```

## License
MIT
