package models

type Image struct {
	Data []byte `json:"data"`
	Name string `json:"name"`
}

type RequestImage struct {
	Image     string `json:"image"`
	Operation string `json:"operation"`
}
