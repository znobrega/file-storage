package dto

import (
	"time"
)

type FilePublic struct {
	File
	URL string `json:"url"`
}

type File struct {
	FileID    string    `json:"fileId"`
	Name      string    `json:"name"`
	Path      string    `json:"path"`
	Extension string    `json:"extension"`
	Directory string    `json:"directory"`
	FileSize  string    `json:"fileSize"`
	UserID    uint64    `json:"userId"`
	IsPublic  bool      `json:"isPublic"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type FileList struct {
	Files []string `json:"files"`
}

type FileResponse struct {
	Files []FilePublic `json:"files"`
}
