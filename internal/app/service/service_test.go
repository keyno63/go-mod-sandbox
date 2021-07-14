package service

import (
	"fmt"
	"go-mod2/internal/app/model"
	"go-mod2/internal/app/repository"
	"reflect"
	"testing"
)

// 後で gomock で差し替える
type mockRepository struct{}

func (m mockRepository) GetUser(id string) (*model.UserAccount, error) {
	return &model.UserAccount{
		Id:        "1",
		FirstName: "first_name",
		LastName:  "last_name",
	}, nil
}

type errorMockRepository struct{}

func (m errorMockRepository) GetUser(id string) (*model.UserAccount, error) {
	return nil, fmt.Errorf("error")
}

func TestUserServiceImpl_GetUser(t *testing.T) {
	mock := mockRepository{}
	errorMock := errorMockRepository{}
	type fields struct {
		userRepository repository.UserRepository
	}
	type args struct {
		id string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *model.UserAccount
		wantErr bool
	}{
		{
			name: "green",
			fields: fields{
				userRepository: mock,
			},
			args: args{
				id: "1",
			},
			want: &model.UserAccount{
				Id:        "1",
				FirstName: "first_name",
				LastName:  "last_name",
			},
			wantErr: false,
		},
		{
			name: "red",
			fields: fields{
				userRepository: errorMock,
			},
			args: args{
				id: "1",
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			us := UserServiceImpl{
				userRepository: tt.fields.userRepository,
			}
			got, err := us.GetUser(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetUser() got = %v, want %v", got, tt.want)
			}
		})
	}
}
