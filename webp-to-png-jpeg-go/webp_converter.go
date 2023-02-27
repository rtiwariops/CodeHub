package main

import (
	"flag"
	"fmt"
	"image/jpeg"
	"image/png"
	"os"
	"path/filepath"

	"golang.org/x/image/webp"
)

func main() {
    var input string
    flag.StringVar(&input, "input", "", "path to input WebP file")
    flag.Parse()

    if input == "" {
        fmt.Println("Error: No input file specified")
        os.Exit(1)
    }

    output := filepath.Base(input)
    output = output[0:len(output)-len(filepath.Ext(output))] // Remove file extension

    f, err := os.Open(input)
    if err != nil {
        fmt.Println("Error:", err)
        os.Exit(1)
    }
    defer f.Close()

    img, err := webp.Decode(f)
    if err != nil {
        fmt.Println("Error:", err)
        os.Exit(1)
    }

    // Convert to PNG
    pngFile, err := os.Create(output + ".png")
    if err != nil {
        fmt.Println("Error:", err)
        os.Exit(1)
    }
    defer pngFile.Close()

    err = png.Encode(pngFile, img)
    if err != nil {
        fmt.Println("Error:", err)
        os.Exit(1)
    }

    // Convert to JPEG
    jpegFile, err := os.Create(output + ".jpeg")
    if err != nil {
        fmt.Println("Error:", err)
        os.Exit(1)
    }
    defer jpegFile.Close()

    err = jpeg.Encode(jpegFile, img, &jpeg.Options{Quality: 90})
    if err != nil {
        fmt.Println("Error:", err)
        os.Exit(1)
    }

    fmt.Println("Conversion completed successfully")
}

