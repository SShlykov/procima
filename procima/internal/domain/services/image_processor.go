package services

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"strings"
	"sync"
)

// RGBWeights веса для RGB (чтобы изображение было светлее)
var RGBWeights = []float64{0.299, 0.587, 0.114}

// colorToBAW конвертирует цвет в черно-белый
func colorToBAW(c color.Color) color.Color {
	r, g, b, _ := c.RGBA()
	gray := uint16(float64(r)*RGBWeights[0] + float64(g)*RGBWeights[1] + float64(b)*RGBWeights[2])
	grayColor := uint8(gray >> 8)
	return color.RGBA{
		R: grayColor,
		G: grayColor,
		B: grayColor,
		A: 255,
	}
}

// colorToNegative конвертирует цвет в негатив
func colorToNegative(c color.Color) color.Color {
	r, g, b, a := c.RGBA()
	return color.RGBA{
		R: 255 - uint8(r>>8),
		G: 255 - uint8(g>>8),
		B: 255 - uint8(b>>8),
		A: uint8(a >> 8),
	}
}

// recolor перекрашивает изображение (HOF)
func recolor(img image.Image, recolorFunc func(color.Color) color.Color) image.Image {
	bounds := img.Bounds()
	maxX, maxY := bounds.Max.X, bounds.Max.Y
	newImage := image.NewRGBA(image.Rect(0, 0, maxX, maxY))

	for y := bounds.Min.Y; y < maxY; y++ {
		for x := bounds.Min.X; x < maxX; x++ {
			newImage.Set(x, y, recolorFunc(img.At(x, y)))
		}
	}

	return newImage
}

// Rotate90deg поворачивает сегмент изображения на 90 градусов по часовой стрелке, принимая во внимание смещение по x.
func Rotate90deg(img *image.RGBA, startX, width int) *image.RGBA {
	bounds := img.Bounds()
	newImage := image.NewRGBA(image.Rect(0, 0, bounds.Dy(), width))

	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := startX; x < startX+width; x++ {
			newX := bounds.Max.Y - y - 1
			newY := x - startX
			newImage.Set(newX, newY, img.At(x, y))
		}
	}

	return newImage
}

// Parallel поворачивает изображение *image.RGBA на 90 градусов с использованием n горутин
func Parallel(img *image.RGBA, n int, f func(img *image.RGBA, startX, width int) *image.RGBA) *image.RGBA {
	bounds := img.Bounds()
	maxX, maxY := bounds.Max.X, bounds.Max.Y
	fmt.Println("maxX: ", maxX, "maxY: ", maxY)

	newImage := image.NewRGBA(image.Rect(0, 0, maxY, maxX))

	var wg sync.WaitGroup
	partWidth := maxX / n

	for i := 0; i < n; i++ {
		wg.Add(1)
		startX := i * partWidth
		endX := (i + 1) * partWidth
		if i == n-1 {
			endX = maxX // На случай, если деление порежет делителем.
		}

		go func(startX, endX, index int) {
			defer wg.Done()
			subImg := img.SubImage(image.Rect(startX, 0, endX, maxY)).(*image.RGBA)
			rotatedSegment := f(subImg, startX, partWidth)

			// Расчет начальной позиции Y для каждого сегмента
			startY := index * partWidth

			// Координаты должны быть скорректированы для новой ориентации изображения
			for y := 0; y < rotatedSegment.Bounds().Max.Y; y++ {
				for x := 0; x < rotatedSegment.Bounds().Max.X; x++ {
					newX := x
					newY := startY + y
					newImage.Set(newX, newY, rotatedSegment.At(x, y))
				}
			}
		}(startX, endX, i)
	}

	wg.Wait()
	return newImage
}

// base64ToImage декодирует строку в формате Base64 в изображение *image.RGBA.
func base64ToImage(encodedImage string) (*image.RGBA, error) {
	data := encodedImage
	if i := strings.Index(data, ","); i != -1 {
		data = data[i+1:]
	}

	decodedData, err := base64.StdEncoding.DecodeString(data)
	if err != nil {
		return nil, err
	}

	img, _, err := image.Decode(bytes.NewReader(decodedData))
	if err != nil {
		return nil, err
	}

	b := img.Bounds()
	rgba := image.NewRGBA(image.Rect(0, 0, b.Dx(), b.Dy()))
	draw.Draw(rgba, rgba.Bounds(), img, b.Min, draw.Src)

	return rgba, nil
}
