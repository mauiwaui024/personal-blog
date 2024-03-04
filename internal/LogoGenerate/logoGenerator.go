package logogenerate

import (
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"os"
)

func LogoGenerator() {
	// Create a new pixel RGBA image
	img := image.NewRGBA(image.Rect(0, 0, 300, 300))

	// Fill the background with a gradient color
	draw.Draw(img, img.Bounds(), &image.Uniform{color.RGBA{0, 0, 255, 255}}, image.ZP, draw.Src)

	// Draw a circle in the center
	center := image.Point{150, 150}
	radius := 100
	drawCircle(img, center, radius, color.RGBA{255, 0, 0, 255})

	// Draw a pencil on the background
	drawPencil(img, color.RGBA{255, 255, 0, 255})

	// Create a new file to save the image
	file, err := os.Create("./images/logo.png")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// Encode the image as PNG and save to the file
	err = png.Encode(file, img)
	if err != nil {
		panic(err)
	}
}

func drawCircle(img *image.RGBA, center image.Point, radius int, col color.RGBA) {
	x0, y0, r := center.X, center.Y, radius
	f := 1 - r
	ddF_x, ddF_y := 1, -2*r
	x, y := 0, r

	img.Set(x0, y0+r, col)
	img.Set(x0, y0-r, col)
	img.Set(x0+r, y0, col)
	img.Set(x0-r, y0, col)

	for x < y {
		if f >= 0 {
			y--
			ddF_y += 2
			f += ddF_y
		}
		x++
		ddF_x += 2
		f += ddF_x
		img.Set(x0+x, y0+y, col)
		img.Set(x0-x, y0+y, col)
		img.Set(x0+x, y0-y, col)
		img.Set(x0-x, y0-y, col)
		img.Set(x0+y, y0+x, col)
		img.Set(x0-y, y0+x, col)
		img.Set(x0+y, y0-x, col)
		img.Set(x0-y, y0-x, col)
	}
}

// Function to draw a pencil on the image
func drawPencil(img *image.RGBA, col color.RGBA) {
	// Draw pencil body
	draw.Draw(img, image.Rect(120, 100, 180, 250), &image.Uniform{col}, image.ZP, draw.Src)

	// Draw pencil tip
	draw.Draw(img, image.Rect(130, 90, 170, 100), &image.Uniform{color.Gray{128}}, image.ZP, draw.Src)

	// Draw pencil lead
	draw.Draw(img, image.Rect(135, 80, 165, 90), &image.Uniform{color.Black}, image.ZP, draw.Src)

	// Draw pencil eraser
	draw.Draw(img, image.Rect(120, 90, 180, 100), &image.Uniform{color.White}, image.ZP, draw.Src)
}
