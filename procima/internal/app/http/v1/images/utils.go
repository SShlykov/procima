package images

import "regexp"

const PartsWhenImage = 2

func getImageType(dataURL string) (string, bool) {
	re := regexp.MustCompile(`^data:image/([^;]+);base64,`)
	matches := re.FindStringSubmatch(dataURL[:50])

	if len(matches) == PartsWhenImage {
		return matches[1], true
	}
	return "", false
}
