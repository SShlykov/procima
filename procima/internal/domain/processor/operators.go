package processor

import (
	"github.com/SShlykov/procima/procima/internal/domain/processor/parallel"
	"image"
)

const Procs = 10

func GrayScale(img image.Image) image.Image {
	src := newScanner(img)
	dst := image.NewRGBA(image.Rect(0, 0, img.Bounds().Max.X, img.Bounds().Max.Y))

	parallel.Parallel(0, img.Bounds().Max.Y, Procs, func(c <-chan int) {
		for y := range c {
			i := y * dst.Stride
			src.scan(0, y, src.w, y+1, dst.Pix[i:i+src.w*4])
			for x := 0; x < dst.Stride; x += 4 {
				d := dst.Pix[i : i+3 : i+3]
				f := 0.299*float64(d[0]) + 0.587*float64(d[1]) + 0.114*float64(d[2])
				col := uint8(f + 0.5)
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
	src := newScanner(img)
	dst := image.NewRGBA(image.Rect(0, 0, img.Bounds().Max.X, img.Bounds().Max.Y))

	parallel.Parallel(0, img.Bounds().Max.Y, Procs, func(c <-chan int) {
		for y := range c {
			i := y * dst.Stride
			src.scan(0, y, src.w, y+1, dst.Pix[i:i+src.w*4])
			for x := 0; x < dst.Stride; x += 4 {
				d := dst.Pix[i : i+3 : i+3]
				d[0] = 255 - d[0]
				d[1] = 255 - d[1]
				d[2] = 255 - d[2]
				i += 4
			}
		}
	})

	return dst
}

func Rotate90(img image.Image) *image.NRGBA {
	src := newScanner(img)
	dstW := src.h
	dstH := src.w
	rowSize := dstW * 4
	dst := image.NewNRGBA(image.Rect(0, 0, dstW, dstH))
	parallel.Parallel(0, dstH, Procs, func(ys <-chan int) {
		for dstY := range ys {
			i := dstY * dst.Stride
			srcX := dstH - dstY - 1
			src.scan(srcX, 0, srcX+1, src.h, dst.Pix[i:i+rowSize])
		}
	})
	return dst
}
