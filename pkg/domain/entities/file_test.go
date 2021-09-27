package entities

import (
	"fmt"
	"github.com/spf13/viper"
	"github.com/znobrega/file-storage/internal/configs"
	"github.com/znobrega/file-storage/pkg/dto"
	"reflect"
	"testing"
	"time"
)

func TestAddUrlToFilesResponse(t *testing.T) {
	configs.LoadConfig("../../../resources")
	apiStaticUrl := fmt.Sprintf("http://%s:%s/static/", viper.GetString("files.host"), viper.GetString("server.port"))
	userID := uint64(0)
	filesUnderTest := []File{{
		FileID:    "d173faa6-5c5e-1809-0b63-0fb27e500e19",
		Name:      "testfile",
		Extension: ".png",
		UserID:    userID,
		IsPublic:  false,
		CreatedAt: time.Time{},
		DeletedAt: nil,
	}, {
		FileID:    "d173faa6-5c5e-1809-0b63-0fb27e500e17",
		Name:      "testfile",
		Extension: ".png",
		UserID:    userID,
		IsPublic:  false,
		CreatedAt: time.Time{},
		DeletedAt: nil,
	}}

	type args struct {
		files []File
	}
	tests := []struct {
		name string
		args args
		want []dto.FilePublic
	}{
		{
			name: "it should add url to file response",
			args: args{
				files: filesUnderTest,
			},
			want: func() []dto.FilePublic {
				return []dto.FilePublic{
					{
						File: dto.File{
							FileID:    filesUnderTest[0].FileID,
							Name:      filesUnderTest[0].Name,
							Extension: filesUnderTest[0].Extension,
							UserID:    filesUnderTest[0].UserID,
							IsPublic:  filesUnderTest[0].IsPublic,
							CreatedAt: filesUnderTest[0].CreatedAt,
							UpdatedAt: filesUnderTest[0].UpdatedAt,
						},
						URL: fmt.Sprintf("%s%s", apiStaticUrl, filesUnderTest[0].Path),
					},
					{
						File: dto.File{
							FileID:    filesUnderTest[1].FileID,
							Name:      filesUnderTest[1].Name,
							Extension: filesUnderTest[1].Extension,
							UserID:    filesUnderTest[1].UserID,
							IsPublic:  filesUnderTest[1].IsPublic,
							CreatedAt: filesUnderTest[1].CreatedAt,
							UpdatedAt: filesUnderTest[1].UpdatedAt,
						},
						URL: fmt.Sprintf("%s%s", apiStaticUrl, filesUnderTest[1].Path),
					},
				}
			}(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := AddUrlToFilesResponse(tt.args.files); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AddUrlToFilesResponse() = %v, want %v", got, tt.want)
			}
		})
	}
}
