package adapters

import (
	"bytes"
	"image"
	"image/jpeg"
)

func ImageToBytes(image image.Image) (*[]byte, error) {
	data, err := encodeToJPEG(image, 70)
	return &data, err
}

// encodeToJPEG принимает изображение *image.RGBA и качество кодирования от 1 до 100,
// где большее значение соответствует лучшему качеству и большему размеру файла.
func encodeToJPEG(image image.Image, quality int) ([]byte, error) {
	var buffer bytes.Buffer
	err := jpeg.Encode(&buffer, image, &jpeg.Options{Quality: quality})
	if err != nil {
		return nil, err
	}
	return buffer.Bytes(), nil
}
