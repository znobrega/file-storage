package controllers

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"github.com/znobrega/file-storage/internal/configs"
	"github.com/znobrega/file-storage/pkg/domain/usecases"
	"github.com/znobrega/file-storage/pkg/dto"
	"github.com/znobrega/file-storage/pkg/infra/helpers"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

type FileController interface {
	HandleFileUpload() func(w http.ResponseWriter, r *http.Request)
	HandleFileDelete() func(w http.ResponseWriter, r *http.Request)
	HandleListPublicFiles() func(w http.ResponseWriter, r *http.Request)
	HandleListPrivateFiles() func(w http.ResponseWriter, r *http.Request)
	HandleFindById() func(w http.ResponseWriter, r *http.Request)
	HandleFileDeleteById() func(w http.ResponseWriter, r *http.Request)
	HandleUpdateDirectoryById() func(w http.ResponseWriter, r *http.Request)
	HandleListDirectories() func(w http.ResponseWriter, r *http.Request)
	HandleListByDirectory() func(w http.ResponseWriter, r *http.Request)
	HandleFileReplace() func(w http.ResponseWriter, r *http.Request)
}

type fileController struct {
	saveFilesService                usecases.SaveFiles
	deleteFilesService              usecases.DeleteFiles
	listFilesService                usecases.ListFiles
	listByIdFilesService            usecases.ListFilesById
	findByIdFilesService            usecases.FindFilesById
	updateDirectoryByIdFilesService usecases.UpdateFileDirectoryById
	updateFileBlobService           usecases.UpdateFileBlob
	listDirectoriesService          usecases.ListDirectories
	listFilesByDirectoriesService   usecases.ListFilesByDirectory
}

func NewFileController(
	saveFilesService usecases.SaveFiles,
	deleteFilesService usecases.DeleteFiles,
	listFilesService usecases.ListFiles,
	listByIdFilesService usecases.ListFilesById,
	findByIdFilesService usecases.FindFilesById,
	updateDirectoryByIdFilesService usecases.UpdateFileDirectoryById,
	listDirectoriesService usecases.ListDirectories,
	listFilesByDirectoriesService usecases.ListFilesByDirectory,
	updateFileBlobService usecases.UpdateFileBlob,
) FileController {
	return fileController{
		saveFilesService:                saveFilesService,
		deleteFilesService:              deleteFilesService,
		listFilesService:                listFilesService,
		listByIdFilesService:            listByIdFilesService,
		findByIdFilesService:            findByIdFilesService,
		updateDirectoryByIdFilesService: updateDirectoryByIdFilesService,
		listDirectoriesService:          listDirectoriesService,
		listFilesByDirectoriesService:   listFilesByDirectoriesService,
		updateFileBlobService:           updateFileBlobService,
	}
}

// @Title saveFile
// @Tags Files
// @Security userIdAuthentication
// @Param files formData file true "file to be uploaded"
// @Param isPublic formData boolean true "Indicates if the file is public or private"
// @Param filename formData string true "Gives a custom name to the file"
// @Param directory formData string true "Indicates the file directory"
// @Summary Upload a file
// @Description Upload a new File
// @Success 200 {object} dto.FileResponse
// @Failure 400 "Bad request"
// @Accept json
// @Router /files [post]
func (i fileController) HandleFileUpload() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		err := r.ParseMultipartForm(1 + configs.Viper.GetInt64("files.sizeLimit"))
		if err != nil {
			helpers.ReturnHttpError(w, 400, err)
			return
		}
		userID := helpers.GetUserIdFromContext(r.Context())

		isPublicBody := r.MultipartForm.Value["isPublic"]
		if isPublicBody == nil {
			helpers.ReturnHttpError(w, 400, helpers.ErrIsPublicRequired)
			return
		}

		filename := r.MultipartForm.Value["filename"]
		if filename == nil {
			helpers.ReturnHttpError(w, 400, helpers.ErrFilenameIsRequired)
			return
		}

		directory := r.MultipartForm.Value["directory"]
		if directory == nil {
			helpers.ReturnHttpError(w, 400, helpers.ErrFilenameIsRequired)
			return
		}

		isPublic, err := strconv.ParseBool(isPublicBody[0])
		if err != nil {
			helpers.ReturnHttpError(w, 400, err)
			return
		}

		files := r.MultipartForm.File["files"]
		filesResponse, err := i.saveFilesService.Execute(r.Context(), usecases.SaveFilesParam{
			Files:     files,
			Filename:  filename[0],
			UserID:    userID,
			IsPublic:  isPublic,
			Directory: directory[0],
		})
		if err != nil {
			log.Println("failed to upload files")
			helpers.ReturnHttpError(w, 400, err)
			return
		}

		helpers.WriteResponseAsJson(w, dto.FileResponse{Files: filesResponse})
	}
}

func (i fileController) HandleFileDelete() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		userID := helpers.GetUserIdFromContext(r.Context())

		var responseJson dto.FileList
		body, err := ioutil.ReadAll(r.Body)
		err = json.Unmarshal(body, &responseJson)

		_, err = i.deleteFilesService.Execute(r.Context(), usecases.DeleteFilesParam{
			Files:  responseJson,
			UserID: userID,
		})
		if err != nil {
			log.Println("failed to delete files")
			helpers.ReturnHttpError(w, 400, err)
			return
		}

		helpers.WriteResponseAsJson(w, "all files was successfully deleted")
	}
}

// @Title deleteFile
// @Tags Files
// @Security userIdAuthentication
// @Param id path string  true "The identifier for the file"
// @Summary Deletes a file
// @Description Deletes a  File
// @Success 200 {object} dto.FileList
// @Failure 400 "Bad request"
// @Accept json
// @Router /files/{id} [delete]
func (i fileController) HandleFileDeleteById() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		fileId := chi.URLParam(r, "id")
		file := dto.FileList{
			Files: []string{fileId},
		}

		_, err := i.deleteFilesService.Execute(r.Context(), usecases.DeleteFilesParam{
			Files:  file,
			UserID: helpers.GetUserIdFromContext(r.Context()),
		})
		if err != nil {
			log.Println("failed to delete files")
			helpers.ReturnHttpError(w, 400, err)
			return
		}

		helpers.WriteResponseAsJson(w, file)
	}
}

// @Title listPublicImages
// @Tags Files
// @Param user_id query string false "The identifier for the file's user"
// @Param limit query string  true "Limit of page records "
// @Param page query string  true "Page number"
// @Summary List public files
// @Description List public files
// @Success 200 {object} dto.FileResponse
// @Failure 400 "Bad request"
// @Accept json
// @Router /files/list/public [get]
func (i fileController) HandleListPublicFiles() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		isPublic := true
		query := r.URL.Query()

		userID, exists := query["user_id"]
		var UserIDAsInt *int

		limit, exists := query["limit"]
		if !exists {
			http.Error(w, "query param limit is required", 400)
			return
		}

		page, exists := query["page"]
		if !exists {
			http.Error(w, "query param page is required", 400)
			return
		}
		pageAsInt, _ := strconv.Atoi(page[0])
		LimitAsInt, _ := strconv.Atoi(limit[0])

		if userID != nil {
			id, _ := strconv.Atoi(userID[0])
			UserIDAsInt = &id
		}

		files, err := i.listFilesService.Execute(r.Context(), usecases.ListFilesParam{
			UserId:   UserIDAsInt,
			IsPublic: &isPublic,
			Page:     pageAsInt,
			Limit:    LimitAsInt,
		})
		if err != nil {
			log.Println(err)
			helpers.ReturnHttpError(w, 400, err)
		}
		helpers.WriteResponseAsJson(w, files)
	}
}

// @Title listPrivateImages
// @Tags Files
// @Security userIdAuthentication
// @Param limit query string  true "Limit of page records "
// @Param page query string  true "Page number"
// @Summary List public files
// @Description List public files
// @Success 200 {object} dto.FileResponse
// @Failure 400 "Bad request"
// @Accept json
// @Router /files/list/private [get]
func (i fileController) HandleListPrivateFiles() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		isPublic := false

		query := r.URL.Query()
		limit, exists := query["limit"]
		if !exists {
			helpers.ReturnHttpError(w, 400, helpers.ErrLimitRequired)
			return
		}

		page, exists := query["page"]
		if !exists {
			helpers.ReturnHttpError(w, 400, helpers.ErrPageRequired)
			return
		}

		pageAsInt, _ := strconv.Atoi(page[0])
		LimitAsInt, _ := strconv.Atoi(limit[0])
		UserIDAsInt := int(helpers.GetUserIdFromContext(r.Context()))

		files, err := i.listFilesService.Execute(r.Context(), usecases.ListFilesParam{
			UserId:   &UserIDAsInt,
			IsPublic: &isPublic,
			Page:     pageAsInt,
			Limit:    LimitAsInt,
		})
		if err != nil {
			helpers.ReturnHttpError(w, 400, err)
			return
		}

		helpers.WriteResponseAsJson(w, files)
	}
}

// @Title findFileById
// @Tags Files
// @Security userIdAuthentication
// @Param id path string  true "The identifier for the file"
// @Summary Find a file
// @Description Find a  File
// @Success 200 {object} dto.FilePublic
// @Failure 400 "Bad request"
// @Accept json
// @Router /files/{id} [get]
func (i fileController) HandleFindById() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")
		userId := helpers.GetUserIdFromContext(r.Context())
		file, err := i.findByIdFilesService.Execute(r.Context(), usecases.FindFileByIdParam{
			FileID: id,
			UserID: userId,
		})
		if err != nil {
			log.Println("failed to get file", err)
			helpers.ReturnHttpError(w, 400, err)
			return
		}
		helpers.WriteResponseAsJson(w, file)
	}
}

// @Title updateDirectoryById
// @Tags Files
// @Security userIdAuthentication
// @Param id path string  true "The identifier for the file"
// @Param content body usecases.UpdateFileDirectoryByIdParam  true "The identifier for the file"
// @Summary Update a file directory
// @Description Update a file directory
// @Success 200 {object} dto.FilePublic
// @Failure 400 "Bad request"
// @Accept json
// @Router /files/{id} [patch]
func (i fileController) HandleUpdateDirectoryById() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")
		userId := helpers.GetUserIdFromContext(r.Context())

		var requestJson usecases.UpdateFileDirectoryByIdParam
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Println("failed to read body", err)
			helpers.ReturnHttpError(w, 400, err)
			return
		}
		err = json.Unmarshal(body, &requestJson)
		if err != nil {
			log.Println("failed to unmarshal request", err)
			helpers.ReturnHttpError(w, 400, err)
			return
		}
		requestJson.FileID = id
		requestJson.UserID = userId

		file, err := i.updateDirectoryByIdFilesService.Execute(r.Context(), requestJson)
		if err != nil {
			log.Println("failed to update directory", err)
			helpers.ReturnHttpError(w, 400, err)
			return
		}
		helpers.WriteResponseAsJson(w, file)
	}
}

// @Title listFilesByDirectory
// @Tags Files
// @Security userIdAuthentication
// @Summary List all directories
// @Description List all directories
// @Success 200 {object} dto.Directories
// @Failure 400 "Bad request"
// @Accept json
// @Router /files/grep [get]
func (i fileController) HandleListDirectories() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		directories, err := i.listDirectoriesService.Execute(r.Context())
		if err != nil {
			log.Println("failed to list directories", err)
			helpers.ReturnHttpError(w, 400, err)
			return
		}
		helpers.WriteResponseAsJson(w, directories)
	}
}

// @Title listFilesByDirectory
// @Tags Files
// @Security userIdAuthentication
// @Summary List files by passing a directory
// @Description List files by passing a directory
// @Param   dir 				query    string   false "Attribute to get files under a directory"
// @Success 200 {object} dto.FileResponse
// @Accept json
// @Router /files/grep [get]
func (i fileController) HandleListByDirectory() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

		query := r.URL.Query()
		dir, exists := query["dir"]
		if !exists {
			helpers.ReturnHttpError(w, 400, helpers.ErrLimitRequired)
			return
		}
		files, err := i.listFilesByDirectoriesService.Execute(r.Context(), usecases.ListFilesByDirectoryParam{
			Directory: dir[0],
		})
		if err != nil {
			log.Println("failed to list directories", err)
			helpers.ReturnHttpError(w, 400, err)
			return
		}
		helpers.WriteResponseAsJson(w, files)
	}
}

// @Title saveFile
// @Tags Files
// @Security userIdAuthentication
// @Param files formData file true "file to be uploaded"
// @Param id path string true "file to be replaced"
// @Summary Replace a blob from a existent file
// @Description Replace a blob from a existent file
// @Success 200 {object} dto.FileResponse
// @Failure 400 "Bad request"
// @Accept json
// @Router /files/{id}/replace [patch]
func (i fileController) HandleFileReplace() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		err := r.ParseMultipartForm(1 + configs.Viper.GetInt64("files.sizeLimit"))
		if err != nil {
			helpers.ReturnHttpError(w, 400, err)
			return
		}
		files := r.MultipartForm.File["files"]
		filesResponse, err := i.updateFileBlobService.Execute(r.Context(), usecases.UpdateFileBlobParam{
			FileID: chi.URLParam(r, "id"),
			Files:  files,
		})
		if err != nil {
			log.Println("failed to upload files")
			helpers.ReturnHttpError(w, 400, err)
			return
		}

		helpers.WriteResponseAsJson(w, filesResponse)
	}
}
