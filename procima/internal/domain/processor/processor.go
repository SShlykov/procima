package processor

import (
	"context"
	loggerPkg "github.com/SShlykov/procima/go_pkg/logger"
	"github.com/SShlykov/procima/procima/internal/models/adapters"
	"github.com/nfnt/resize"
	"image"
)

type ImageProcessorItem struct {
	Channel   chan<- ImageResult
	Operation string
	Img       image.Image
}

type ImageResult struct {
	Err error
	Res *[]byte
}

func Run(ctx context.Context, logger loggerPkg.Logger, largestSideLimit int, procChan <-chan ImageProcessorItem) {
	for {
		select {
		case <-ctx.Done():
			return
		case item := <-procChan:
			var img image.Image

			if largestSideLimit > 0 {
				initX, initY := img.Bounds().Max.X, img.Bounds().Max.Y
				var koef float64
				if initX > initY {
					koef = 1000 / float64(initX)
				} else {
					koef = 1000 / float64(initY)
				}
				img = resize.Resize(uint(koef*float64(initX)), uint(koef*float64(initY)), img, resize.Lanczos3)
			}

			switch item.Operation {
			case "rotate:90deg":
				img = Rotate90(item.Img)
			case "recolor:grayscale":
				img = GrayScale(item.Img)
			case "recolor:negative":
				img = Negative(item.Img)
			default:
				logger.Warn("unknown operation", loggerPkg.Any("operation", item.Operation))
			}

			res, err := adapters.ImageToBytes(img)

			item.Channel <- ImageResult{Err: err, Res: res}
		}

	}
}
