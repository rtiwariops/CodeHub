<!-- @format -->

# WebP to PNG and JPEG Converter

📦 This Go package is a command-line utility that converts WebP image files to PNG and JPEG formats.

Usage
👉 Pass the path to the WebP file as an input flag to the program:

```
$ go run main.go -input /path/to/image.webp
```

👉 The program will generate two output files in the same directory as the input file, with the same filename but with .png and .jpeg extensions respectively.

## Dependencies

🔗 This package relies on the golang.org/x/image/webp, image/png, and image/jpeg packages.

## License

📝 This package is licensed under the MIT License.
