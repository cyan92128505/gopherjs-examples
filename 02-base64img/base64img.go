package main

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"image"
	"image/color"
	"image/png"
	"strconv"
	"strings"
	"github.com/gopherjs/gopherjs/js"
)

// Convert hex color string to RGBA color
func hexToColor(hex string) (color.NRGBA, error) {
    // Remove '#' if present
    hex = strings.TrimPrefix(hex, "#")
    
    // Parse hex string to RGB values
    rgb, err := strconv.ParseUint(hex, 16, 32)
    if err != nil {
        return color.NRGBA{}, err
    }
    
    return color.NRGBA{
        R: uint8(rgb >> 16),
        G: uint8((rgb >> 8) & 0xFF),
        B: uint8(rgb & 0xFF),
        A: 255,
    }, nil
}

// Generate base64 encoded image with single color
func base64img(width, height int, hexColor string) string {
    // Create new image
    m := image.NewNRGBA(image.Rectangle{
        Min: image.Point{0, 0},
        Max: image.Point{width, height},
    })
    
    // Convert hex to color
    c, err := hexToColor(hexColor)
    if err != nil {
        // Fall back to black if invalid hex
        c = color.NRGBA{0, 0, 0, 255}
    }
    
    // Fill image with single color
    for y := 0; y < height; y++ {
        for x := 0; x < width; x++ {
            m.SetNRGBA(x, y, c)
        }
    }
    
    // Encode to PNG
    buf := bytes.NewBuffer([]byte{})
    if err := png.Encode(buf, m); err != nil {
        fmt.Println(err)
    }
    
    // Convert to base64
    return base64.StdEncoding.EncodeToString(buf.Bytes())
}

func main() {
    if js.Global != nil {
        // Running from browser
        js.Global.Set("gopkg", map[string]interface{}{
            "base64img": base64img,
        })
    } else {
        // Test run
        width := 32
        height := 4
        hexColor := "FF0000" // Red color
        b64str := base64img(width, height, hexColor)
        fmt.Println(b64str)
    }
}
