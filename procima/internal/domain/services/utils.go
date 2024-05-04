package services

import (
	"bytes"
	"encoding/base64"
	"image"
	"strings"
)

// base64ToImage декодирует строку в формате Base64 в изображение *image.RGBA.
func base64ToImage(encodedImage string) (image.Image, error) {
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

	return img, nil
}
