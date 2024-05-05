package adapters

import (
	"bytes"
	"image"
	"image/jpeg"
)

const ExpectedQuality = 70

func ImageToBytes(img image.Image) (*[]byte, error) {
	data, err := encodeToJPEG(img, ExpectedQuality)
	return &data, err
}

// encodeToJPEG принимает изображение *image.RGBA и качество кодирования от 1 до 100,
// где большее значение соответствует лучшему качеству и большему размеру файла.
func encodeToJPEG(img image.Image, quality int) ([]byte, error) {
	var buffer bytes.Buffer
	err := jpeg.Encode(&buffer, img, &jpeg.Options{Quality: quality})
	if err != nil {
		return nil, err
	}
	return buffer.Bytes(), nil
}
