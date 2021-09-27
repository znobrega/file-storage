package dto

type Directory struct {
	Path string `json:"path"`
}

type Directories struct {
	Directories []Directory `json:"directories"`
}
