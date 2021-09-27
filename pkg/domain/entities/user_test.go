package entities

import (
	"testing"
	"time"
)

func TestUser_Validate(t *testing.T) {
	type fields struct {
		UserID    uint64
		Name      string
		Email     string
		Password  string
		CreatedAt time.Time
		UpdatedAt time.Time
		DeletedAt *time.Time
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "is should pass all conditions",
			fields: fields{
				UserID:    1,
				Name:      "test",
				Email:     "test@gmail.com",
				Password:  "123",
				CreatedAt: time.Time{},
				UpdatedAt: time.Time{},
				DeletedAt: nil,
			},
			wantErr: false,
		},
		{
			name: "is should pass return error of required password",
			fields: fields{
				UserID:    1,
				Name:      "test",
				Email:     "test@gmail.com",
				Password:  "",
				CreatedAt: time.Time{},
				UpdatedAt: time.Time{},
				DeletedAt: nil,
			},
			wantErr: true,
		},
		{
			name: "is should pass return error of required email",
			fields: fields{
				UserID:    1,
				Name:      "test",
				Email:     "",
				Password:  "123",
				CreatedAt: time.Time{},
				UpdatedAt: time.Time{},
				DeletedAt: nil,
			},
			wantErr: true,
		},
		{
			name: "is should pass return error of required name",
			fields: fields{
				UserID:    1,
				Name:      "",
				Email:     "test@gmail.com",
				Password:  "123",
				CreatedAt: time.Time{},
				UpdatedAt: time.Time{},
				DeletedAt: nil,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &User{
				UserID:    tt.fields.UserID,
				Name:      tt.fields.Name,
				Email:     tt.fields.Email,
				Password:  tt.fields.Password,
				CreatedAt: tt.fields.CreatedAt,
				UpdatedAt: tt.fields.UpdatedAt,
				DeletedAt: tt.fields.DeletedAt,
			}
			if err := u.Validate(); (err != nil) != tt.wantErr {
				t.Errorf("Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestUser_ValidateLogin(t *testing.T) {
	type fields struct {
		UserID    uint64
		Name      string
		Email     string
		Password  string
		CreatedAt time.Time
		UpdatedAt time.Time
		DeletedAt *time.Time
	}

	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "is should pass all conditions",
			fields: fields{
				UserID:    1,
				Name:      "test",
				Email:     "test@gmail.com",
				Password:  "123",
				CreatedAt: time.Time{},
				UpdatedAt: time.Time{},
				DeletedAt: nil,
			},
			wantErr: false,
		},
		{
			name: "is should pass return error of required password",
			fields: fields{
				UserID:    1,
				Name:      "test",
				Email:     "test@gmail.com",
				Password:  "",
				CreatedAt: time.Time{},
				UpdatedAt: time.Time{},
				DeletedAt: nil,
			},
			wantErr: true,
		},
		{
			name: "is should pass return error of required email",
			fields: fields{
				UserID:    1,
				Name:      "test",
				Email:     "",
				Password:  "123",
				CreatedAt: time.Time{},
				UpdatedAt: time.Time{},
				DeletedAt: nil,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &User{
				UserID:    tt.fields.UserID,
				Name:      tt.fields.Name,
				Email:     tt.fields.Email,
				Password:  tt.fields.Password,
				CreatedAt: tt.fields.CreatedAt,
				UpdatedAt: tt.fields.UpdatedAt,
				DeletedAt: tt.fields.DeletedAt,
			}
			if err := u.ValidateLogin(); (err != nil) != tt.wantErr {
				t.Errorf("ValidateLogin() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
