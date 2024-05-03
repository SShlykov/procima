package models

type Image struct {
	Data []byte `json:"data"`
	Name string `json:"name"`
}
type Operation string

type RequestImage struct {
	Image      string      `json:"image"`
	Operations []Operation `json:"operations"`
}
