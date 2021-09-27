package services

import (
	"errors"
	"github.com/znobrega/file-storage/mocks"
	"github.com/znobrega/file-storage/pkg/domain/entities"
	"github.com/znobrega/file-storage/pkg/domain/repositories"
	"github.com/znobrega/file-storage/pkg/dto"
	"github.com/znobrega/file-storage/pkg/infra/helpers"
	"reflect"
	"testing"
	"time"
)

func TestNewUsersService(t *testing.T) {
	mockRepository := mocks.UsersRepository{}
	type args struct {
		usersRepository repositories.UsersRepository
	}
	tests := []struct {
		name string
		args args
		want UsersService
	}{
		{
			name: "it should return a user service",
			args: args{
				usersRepository: &mockRepository,
			},
			want: NewUsersService(&mockRepository),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewUsersService(tt.args.usersRepository); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewUsersService() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_usersService_Create(t *testing.T) {
	userWanted := entities.User{
		UserID:    1,
		Name:      "Test",
		Email:     "test@gmail.com",
		Password:  "123",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		DeletedAt: nil,
	}

	userWantedWithoutPassword := dto.User{
		UserID:    userWanted.UserID,
		Name:      userWanted.Name,
		Email:     userWanted.Email,
		CreatedAt: userWanted.CreatedAt,
		UpdatedAt: userWanted.UpdatedAt,
		DeletedAt: userWanted.DeletedAt,
	}

	userWithoutPassword := entities.User{Name: "Test", Password: ""}
	userWithoutName := entities.User{Name: ""}
	userWithoutEmail := entities.User{Name: "Test", Password: "123123", Email: ""}
	type fields struct {
		usersRepository repositories.UsersRepository
	}
	type args struct {
		user entities.User
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *dto.User
		wantErr bool
	}{
		{
			name: "it should create a new user",
			fields: fields{
				usersRepository: func() repositories.UsersRepository {
					repoMock := mocks.UsersRepository{}
					repoMock.On("Store", &userWanted).Return(nil)
					repoMock.On("FindByEmail", "test@gmail.com").Return(nil, nil)
					return &repoMock
				}(),
			},
			args: args{
				user: userWanted,
			},
			want:    &userWantedWithoutPassword,
			wantErr: false,
		},
		{
			name: "it should return name required",
			fields: fields{
				usersRepository: func() repositories.UsersRepository {
					repoMock := mocks.UsersRepository{}
					repoMock.On("Store", &userWithoutName).Return(nil)
					repoMock.On("FindByEmail", "test@gmail.com").Return(nil, nil)
					return &repoMock
				}(),
			},
			args: args{
				user: userWithoutName,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "it should return password required",
			fields: fields{
				usersRepository: func() repositories.UsersRepository {
					repoMock := mocks.UsersRepository{}
					repoMock.On("Store", &userWithoutPassword).Return(nil)
					repoMock.On("FindByEmail", "test@gmail.com").Return(nil, nil)
					return &repoMock
				}(),
			},
			args: args{
				user: userWithoutPassword,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "it should return a error: user already exists",
			fields: fields{
				usersRepository: func() repositories.UsersRepository {
					repoMock := mocks.UsersRepository{}
					repoMock.On("Store", &userWanted).Return(errors.New("user already exists"))
					repoMock.On("FindByEmail", "test@gmail.com").Return(&userWanted, errors.New("user already exists"))

					return &repoMock
				}(),
			},
			args: args{
				user: userWanted,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "it should return a error: user already exists",
			fields: fields{
				usersRepository: func() repositories.UsersRepository {
					repoMock := mocks.UsersRepository{}
					repoMock.On("Store", &userWithoutEmail).Return(errors.New("user already exists"))
					repoMock.On("FindByEmail", "test@gmail.com").Return(&userWanted, errors.New("user already exists"))

					return &repoMock
				}(),
			},
			args: args{
				user: userWithoutEmail,
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := usersService{
				usersRepository: tt.fields.usersRepository,
			}
			got, err := u.Create(tt.args.user)
			if (err != nil) != tt.wantErr {
				t.Errorf("Create() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Create() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_usersService_FindByEmail(t *testing.T) {
	userWanted := &entities.User{
		UserID:    0,
		Name:      "Test",
		Email:     "test@gmail.com",
		Password:  "123",
		CreatedAt: time.Now(),
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	}
	emailUnderTest := "test@gmail.com"
	type fields struct {
		usersRepository repositories.UsersRepository
	}
	type args struct {
		email string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *entities.User
		wantErr bool
	}{
		{
			name: "it should return a user",
			fields: fields{
				usersRepository: func() repositories.UsersRepository {
					repoMock := mocks.UsersRepository{}
					repoMock.On("FindByEmail", emailUnderTest).Return(userWanted, nil)
					return &repoMock
				}(),
			},
			args: args{
				email: emailUnderTest,
			},
			want:    userWanted,
			wantErr: false,
		},
		{
			name: "it should return a error",
			fields: fields{
				usersRepository: func() repositories.UsersRepository {
					repoMock := mocks.UsersRepository{}
					repoMock.On("FindByEmail", emailUnderTest).Return(nil, errors.New(""))
					return &repoMock
				}(),
			},
			args: args{
				email: emailUnderTest,
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := usersService{
				usersRepository: tt.fields.usersRepository,
			}
			got, err := u.FindByEmail(tt.args.email)
			if (err != nil) != tt.wantErr {
				t.Errorf("FindByEmail() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FindByEmail() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_usersService_FindById(t *testing.T) {

	userWanted := &dto.User{
		UserID:    0,
		Name:      "Test",
		Email:     "test@gmail.com",
		CreatedAt: time.Now(),
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	}
	type fields struct {
		usersRepository repositories.UsersRepository
	}
	type args struct {
		userID int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *dto.User
		wantErr bool
	}{
		{
			name: "it should return a user",
			fields: fields{
				usersRepository: func() repositories.UsersRepository {
					repoMock := mocks.UsersRepository{}
					repoMock.On("FindById", 0).Return(&entities.User{
						UserID:    userWanted.UserID,
						Name:      userWanted.Name,
						Email:     userWanted.Email,
						CreatedAt: userWanted.CreatedAt,
						Password:  "123",
					}, nil)
					return &repoMock
				}(),
			},
			args: args{
				userID: 0,
			},
			want:    userWanted,
			wantErr: false,
		},
		{
			name: "it should return a error",
			fields: fields{
				usersRepository: func() repositories.UsersRepository {
					repoMock := mocks.UsersRepository{}
					repoMock.On("FindById", 0).Return(nil, errors.New(""))
					return &repoMock
				}(),
			},
			args: args{
				userID: 0,
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := usersService{
				usersRepository: tt.fields.usersRepository,
			}
			got, err := u.FindById(tt.args.userID)
			if (err != nil) != tt.wantErr {
				t.Errorf("FindById() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FindById() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_usersService_ListAll(t *testing.T) {
	type fields struct {
		usersRepository repositories.UsersRepository
	}
	tests := []struct {
		name    string
		fields  fields
		want    []dto.User
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := usersService{
				usersRepository: tt.fields.usersRepository,
			}
			got, err := u.ListAll()
			if (err != nil) != tt.wantErr {
				t.Errorf("ListAll() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ListAll() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_usersService_Login(t *testing.T) {
	type fields struct {
		usersRepository repositories.UsersRepository
	}
	type args struct {
		user entities.User
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *helpers.TokenResponse
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := usersService{
				usersRepository: tt.fields.usersRepository,
			}
			got, err := u.Login(tt.args.user)
			if (err != nil) != tt.wantErr {
				t.Errorf("Login() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Login() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_usersService_Update(t *testing.T) {
	type fields struct {
		usersRepository repositories.UsersRepository
	}
	type args struct {
		user entities.User
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *entities.User
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := usersService{
				usersRepository: tt.fields.usersRepository,
			}
			got, err := u.Update(tt.args.user)
			if (err != nil) != tt.wantErr {
				t.Errorf("Update() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Update() got = %v, want %v", got, tt.want)
			}
		})
	}
}
