package filesystem_repository

import (
	"fmt"
	"github.com/znobrega/file-storage/pkg/domain/entities"
	"github.com/znobrega/file-storage/pkg/domain/repositories"
	"github.com/znobrega/file-storage/pkg/domain/usecases"
	"io/ioutil"
	"mime/multipart"
	"os"
	"path/filepath"
)

type saveFilesOnFileSystem struct{}

func NewSaveFilesOnFileSystem() repositories.SaveFilesOnFileSystem {
	return saveFilesOnFileSystem{}
}

func (s saveFilesOnFileSystem) Execute(fileData usecases.SaveFilesParam) ([]entities.File, error) {
	filesSaved := make([]entities.File, 0)
	for _, file := range fileData.Files {
		fileSaved, err := s.processFile(fileData, file)
		if err != nil {
			return nil, err
		}
		filesSaved = append(filesSaved, *fileSaved)
	}

	return filesSaved, nil
}

func (s saveFilesOnFileSystem) processFile(fileData usecases.SaveFilesParam, file *multipart.FileHeader) (*entities.File, error) {
	fileOpened, err := file.Open()
	if err != nil {
		return nil, err
	}
	defer fileOpened.Close()

	fileBytes, err := ioutil.ReadAll(fileOpened)
	if err != nil {
		return nil, err
	}

	fileSaved, err := s.writeFileOnFIleSystem(fileBytes, file.Filename, file.Size, fileData)
	if err != nil {
		return nil, err
	}

	return fileSaved, nil
}

func (s saveFilesOnFileSystem) writeFileOnFIleSystem(fileBytes []byte, realFileName string, fileSize int64, fileData usecases.SaveFilesParam) (*entities.File, error) {

	fileExtension := filepath.Ext(realFileName)

	file := entities.File{
		FileID:    fileData.FileId,
		Name:      fileData.Filename,
		Extension: fileExtension,
		Directory: fileData.Directory,
		FileSize:  fmt.Sprintf("%d", fileSize),
		UserID:    fileData.UserID,
		IsPublic:  fileData.IsPublic,
	}
	file.SetPath()

	err := os.MkdirAll(file.GetDir(), os.ModePerm)
	if err != nil {
		return nil, err
	}

	err = ioutil.WriteFile(file.GetFullPath(), fileBytes, 0644)
	if err != nil {
		return nil, err
	}

	return &file, nil
}
