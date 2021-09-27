package container

import (
	"github.com/spf13/viper"
	"github.com/znobrega/file-storage/internal/services"
	"github.com/znobrega/file-storage/pkg/domain/repositories"
	"github.com/znobrega/file-storage/pkg/domain/usecases"
	"github.com/znobrega/file-storage/pkg/fileid"
	"github.com/znobrega/file-storage/pkg/infra/db"
	"github.com/znobrega/file-storage/pkg/infra/file_repository"
	"github.com/znobrega/file-storage/pkg/infra/filesystem_repository"
	"github.com/znobrega/file-storage/pkg/infra/user_repository"

	"log"
)

type servicesContainer struct {
	DeleteFilesService             usecases.DeleteFiles
	SaveFilesService               usecases.SaveFiles
	ListFilesService               usecases.ListFiles
	ListFilesByIdService           usecases.ListFilesById
	FindFilesByIdService           usecases.FindFilesById
	UpdateFileDirectoryByIdService usecases.UpdateFileDirectoryById
	ListDirectoriesService         usecases.ListDirectories
	ListFilesByDirectoriesService  usecases.ListFilesByDirectory
	UpdateFileBlobService          usecases.UpdateFileBlob
	UsersService                   services.UsersService
}

type repositoriesContainer struct {
	FilesRepository                 repositories.FilesRepository
	UsersRepository                 repositories.UsersRepository
	SaveFilesOnFileSystemRepository repositories.SaveFilesOnFileSystem
}

type Dependencies struct {
	DbHelper              db.Helper
	ServicesContainer     servicesContainer
	RepositoriesContainer repositoriesContainer
}

func Injector(_ *viper.Viper) Dependencies {
	dbHelper := db.Helper{}

	err := dbHelper.InitDatabase()
	if err != nil {
		log.Fatalln("could not initialize database", err)
	}

	dbConnection := dbHelper.GetDatabase()

	usersRepository := user_repository.NewUsersRepository(dbConnection)
	filesRepository := file_repository.NewFilesRepository(dbConnection)
	saveFilesOnFileSystemRepository := filesystem_repository.NewSaveFilesOnFileSystem()

	usersService := services.NewUsersService(usersRepository)

	fileIdGenerator := fileid.Factory()

	deleteFilesService := services.NewDeleteFiles(filesRepository)
	saveFilesService := services.NewSaveFiles(filesRepository, saveFilesOnFileSystemRepository, fileIdGenerator)
	listFilesService := services.NewListFiles(filesRepository)
	listFilesByIdService := services.NewListFilesById(filesRepository)
	findFilesByIdService := services.NewFindFilesById(filesRepository)
	updateFileDirectoryByIdService := services.NewUpdateFileDirectoryById(filesRepository)
	listDirectoriesService := services.NewListDirectories(filesRepository)
	listFilesByDirectoriesService := services.NewListFilesByDirectory(filesRepository)
	updateFileBlobService := services.NewUpdateFileBlob(filesRepository, saveFilesOnFileSystemRepository)

	return Dependencies{
		DbHelper: dbHelper,
		ServicesContainer: servicesContainer{
			DeleteFilesService:             deleteFilesService,
			SaveFilesService:               saveFilesService,
			ListFilesService:               listFilesService,
			ListFilesByIdService:           listFilesByIdService,
			FindFilesByIdService:           findFilesByIdService,
			UpdateFileDirectoryByIdService: updateFileDirectoryByIdService,
			ListDirectoriesService:         listDirectoriesService,
			ListFilesByDirectoriesService:  listFilesByDirectoriesService,
			UpdateFileBlobService:          updateFileBlobService,
			UsersService:                   usersService,
		},
		RepositoriesContainer: repositoriesContainer{
			FilesRepository:                 filesRepository,
			UsersRepository:                 usersRepository,
			SaveFilesOnFileSystemRepository: saveFilesOnFileSystemRepository,
		},
	}
}
