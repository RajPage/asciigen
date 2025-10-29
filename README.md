# AsciiGen ✨

AsciiGen is a command-line tool to convert images into ASCII art. It supports various options for customization, including output width, interpolation, and color inversion for light mode.

## 🚀 Usage

```bash
go run main.go -input <image_file> [options]
```

**Options:**

- `-input`, `-i`: Input image file path (required).
- `-output`, `-o`: Output file path (default: stdout).
- `-width`, `-w`: Output width in characters (default: 100).
- `-interpolation`: Use bilinear interpolation for resizing (default: true).
- `-invert`, `-light`: Invert colors for light mode (default: false).
- `-verbose`, `-v`: Enable verbose output (default: false).

**Examples:**

```bash
# Convert an image and print to console with default width
go run main.go -input photo.jpg

# Convert an image with a specific width and save to a file
go run main.go -i photo.jpg -w 120 -o output.txt

# Convert an image for light mode display
go run main.go -input photo.jpg -invert

# Convert an image with verbose output
go run main.go -input photo.jpg -v
```

## Example Output

![Pikachu](./static/ascii_pikachu.png)
![Raj](./static/ascii_raj.png)

Output without compression:
![RajWithoutCompression](./static/ascii_raj_without_interpolation.png)

## Caveats

- Terminal Line spacing could cause skewed output. Current code assumes a line spacing of 1:2.