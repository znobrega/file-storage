package entities

import (
	"fmt"
	"github.com/spf13/viper"
	"github.com/znobrega/file-storage/internal/configs"
	"github.com/znobrega/file-storage/pkg/dto"
	"time"
)

type File struct {
	FileID    string     `json:"fileId"`
	Name      string     `json:"name"`
	Path      string     `json:"path"`
	Extension string     `json:"extension"`
	Directory string     `json:"directory"`
	FileSize  string     `json:"fileSize"`
	UserID    uint64     `json:"userId"`
	IsPublic  bool       `json:"isPublic"`
	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt time.Time  `json:"updatedAt"`
	DeletedAt *time.Time `json:"deletedAt"`
}

// GetDir gets file user directory
func (f *File) GetDir() string {
	return fmt.Sprintf("%s/%d/%s", configs.Viper.GetString("files.directory"), f.UserID, f.Directory)
}

// GetFullPath gets file full path with directory, filename and extension
func (f *File) GetFullPath() string {
	return fmt.Sprintf("%s/%s%s", f.GetDir(), f.Name, f.Extension)
}

// GetFullPath gets file full public path with directory, filename and extension, path without app metadata
func (f *File) GetPublicFullPath() string {
	return fmt.Sprintf("%s/%s%s", f.Directory, f.Name, f.Extension)
}

// SetPath sets the file Path
func (f *File) SetPath() {
	f.Path = fmt.Sprintf("%s/%s%s", f.Directory, f.Name, f.Extension)
}

func AddUrlToFilesResponse(files []File) []dto.FilePublic {
	filesResponse := make([]dto.FilePublic, 0)
	apiStaticUrl := fmt.Sprintf("http://%s:%s/static/", viper.GetString("files.host"), viper.GetString("server.port"))
	for _, file := range files {
		filesResponse = append(filesResponse, dto.FilePublic{
			File: dto.File{
				FileID:    file.FileID,
				Name:      file.Name,
				Path:      file.Path,
				Directory: file.Directory,
				Extension: file.Extension,
				FileSize:  file.FileSize,
				UserID:    file.UserID,
				IsPublic:  file.IsPublic,
				CreatedAt: file.CreatedAt,
				UpdatedAt: file.UpdatedAt,
			},
			URL: fmt.Sprintf("%s%s", apiStaticUrl, file.Path),
		})
	}
	return filesResponse
}
