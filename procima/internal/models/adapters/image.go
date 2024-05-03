package adapters

import (
	"bytes"
	"fmt"
	"github.com/SShlykov/procima/procima/internal/models"
	"image"
	"image/jpeg"
)

func ImageToModel(image image.Image) (*models.Image, error) {
	imgBytes, err := imageToBytes(image)
	if err != nil {
		fmt.Println("Error converting image to bytes:", err)
		return nil, err
	}
	return &models.Image{Data: imgBytes, Name: "test.jpeg"}, nil
}

func imageToBytes(img image.Image) ([]byte, error) {
	buffer := new(bytes.Buffer)
	if err := jpeg.Encode(buffer, img, nil); err != nil {
		return nil, err
	}
	return buffer.Bytes(), nil
}
