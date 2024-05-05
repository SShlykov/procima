package processor

import (
	"github.com/SShlykov/procima/procima/internal/domain/processor/parallel"
	"github.com/SShlykov/procima/procima/pkg/scanner"
	"image"
)

const (
	Procs    = 10
	MAXCOLOR = 255
	RGBAPC   = 4
	RWeight  = 0.299
	GWeight  = 0.587
	BWeight  = 0.114
	Light    = 0.5
)

func GrayScale(img image.Image) image.Image {
	src := scanner.NewScanner(img)
	dst := image.NewRGBA(image.Rect(0, 0, img.Bounds().Max.X, img.Bounds().Max.Y))

	parallel.Parallel(0, img.Bounds().Max.Y, Procs, func(c <-chan int) {
		for y := range c {
			i := y * dst.Stride
			src.Scan(0, y, src.W, y+1, dst.Pix[i:i+src.W*4])
			for x := 0; x < dst.Stride; x += 4 {
				d := dst.Pix[i : i+3 : i+3]
				f := RWeight*float64(d[0]) + GWeight*float64(d[1]) + BWeight*float64(d[2])
				col := uint8(f + Light)
				d[0] = col
				d[1] = col
				d[2] = col
				i += 4
			}
		}
	})

	return dst
}

func Negative(img image.Image) image.Image {
	src := scanner.NewScanner(img)
	dst := image.NewRGBA(image.Rect(0, 0, img.Bounds().Max.X, img.Bounds().Max.Y))

	parallel.Parallel(0, img.Bounds().Max.Y, Procs, func(c <-chan int) {
		for y := range c {
			i := y * dst.Stride
			src.Scan(0, y, src.W, y+1, dst.Pix[i:i+src.W*4])
			for x := 0; x < dst.Stride; x += RGBAPC {
				d := dst.Pix[i : i+3 : i+3]
				d[0] = MAXCOLOR - d[0]
				d[1] = MAXCOLOR - d[1]
				d[2] = MAXCOLOR - d[2]
				i += RGBAPC
			}
		}
	})

	return dst
}

func Rotate90(img image.Image) *image.NRGBA {
	src := scanner.NewScanner(img)
	dstW := src.H
	dstH := src.W
	rowSize := dstW * RGBAPC
	dst := image.NewNRGBA(image.Rect(0, 0, dstW, dstH))
	parallel.Parallel(0, dstH, Procs, func(ys <-chan int) {
		for dstY := range ys {
			i := dstY * dst.Stride
			srcX := dstH - dstY - 1
			src.Scan(srcX, 0, srcX+1, src.H, dst.Pix[i:i+rowSize])
		}
	})
	return dst
}
