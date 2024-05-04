package adapters

import (
	"bytes"
	"github.com/SShlykov/procima/procima/internal/models"
	"image"
	"image/jpeg"
)

func ImageToModel(image *image.RGBA) (*models.Image, error) {
	data, err := encodeToJPEG(image, 100)
	return &models.Image{Data: data, Name: "test.jpeg"}, err
}

// encodeToJPEG принимает изображение *image.RGBA и качество кодирования от 1 до 100,
// где большее значение соответствует лучшему качеству и большему размеру файла.
func encodeToJPEG(rgba *image.RGBA, quality int) ([]byte, error) {
	var buffer bytes.Buffer
	err := jpeg.Encode(&buffer, rgba, &jpeg.Options{Quality: quality})
	if err != nil {
		return nil, err
	}
	return buffer.Bytes(), nil
}
