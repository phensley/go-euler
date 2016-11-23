package euler

import (
	"image"
	"image/color"
	"image/png"
	"math"
	"os"
)

// Quickie drawing routines to render visualizations of problems as needed.

// DrawRect draws a rectangle
func DrawRect(img *image.RGBA, c color.Color, x1, y1, x2, y2 int) {
	DrawHLine(img, c, x1, y1, x2)
	DrawHLine(img, c, x1, y2, x2)
	DrawVLine(img, c, x1, y1, y2)
	DrawVLine(img, c, x2, y1, y2)
}

// DrawHLine draws a horizontal line
func DrawHLine(img *image.RGBA, c color.Color, x1, y, x2 int) {
	for ; x1 <= x2; x1++ {
		img.Set(x1, y, c)
	}
}

// DrawVLine draws a veritcal line
func DrawVLine(img *image.RGBA, c color.Color, x, y1, y2 int) {
	for ; y1 <= y2; y1++ {
		img.Set(x, y1, c)
	}
}

// DrawLine draws a line from (x1,y1) to (x2,y2) using Bresenham's algorithm
func DrawLine(img *image.RGBA, c color.Color, x1, y1, x2, y2 int) {
	dx := AbsInt(x1 - x2)
	dy := AbsInt(y1 - y2)
	sx, sy := 1, 1
	if x1 > x2 {
		sx = -1
	}
	if y1 > y2 {
		sy = -1
	}
	e := dx - dy
	for {
		img.Set(x1, y1, c)
		if x1 == x2 && y1 == y2 {
			return
		}
		e2 := e * 2
		if e2 > -dy {
			e -= dy
			x1 += sx
		}
		if e2 < dx {
			e += dx
			y1 += sy
		}
	}
}

// DrawRotateEuclidOrigin rotates an image so the origin is at lower-left
// instead of Go's default upper left.
func DrawRotateEuclidOrigin(img *image.RGBA) *image.RGBA {
	bounds := img.Bounds()
	copy := image.NewRGBA(bounds)
	for y := 0; y < bounds.Size().Y; y++ {
		for x := 0; x < bounds.Size().X; x++ {
			nx := y
			ny := bounds.Size().Y - x - 1
			copy.Set(nx, ny, img.At(x, y))
		}
	}
	return copy
}

// DrawTileImages tiles images into a large image, assuming all images have the
// same width and height.
func DrawTileImages(limit int, images []*image.RGBA) *image.RGBA {
	space := 5
	size := len(images)
	cols := int(math.Ceil(math.Sqrt(float64(size))))
	rows := (size / cols) + 1

	b := images[0].Bounds()
	width := (space * (cols - 1)) + (b.Size().X * cols)
	height := (space * (rows - 1)) + (b.Size().Y * rows)
	res := image.NewRGBA(image.Rect(0, 0, width, height))

	for y := 0; y < rows; y++ {
		for x := 0; x < cols; x++ {
			i := (y * cols) + x
			if i >= len(images) {
				break
			}

			img := images[i]
			ox := (x * b.Size().X) + (x * space)
			oy := (y * b.Size().Y) + (y * space)
			for px := 0; px < img.Bounds().Size().X; px++ {
				for py := 0; py < img.Bounds().Size().Y; py++ {
					c := img.At(px, py)
					res.Set(ox+px, oy+py, c)
				}
			}
		}
	}

	return res
}

// SaveImagePNG saves the RGBA image as a PNG format
func SaveImagePNG(name string, img *image.RGBA) error {
	f, err := os.Create(name)
	if err != nil {
		return err
	}
	defer f.Close()
	return png.Encode(f, img)
}
