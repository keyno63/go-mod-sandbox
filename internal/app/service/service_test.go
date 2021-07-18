package service

import (
	"fmt"
	"github.com/golang/mock/gomock"
	"go-mod2/internal/app/model"
	"go-mod2/internal/app/repository"
	"go-mod2/mock"
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

func TestUserServiceImpl_GetUser_(t *testing.T) {
	// mock controller の生成
	ctrl := gomock.NewController(t)
	// mock モジュールの生成
	mockRepository := mock.NewMockUserRepository(ctrl)

	// 戻り値の設定
	ret := model.UserAccount{
		Id:        "1",
		FirstName: "name",
	}
	// mock の振る舞い定義
	mockRepository.EXPECT().GetUser("1").Return(&ret, nil)

	// テスト対象の生成
	target := UserServiceImpl{mockRepository}
	// テスト対象のメソッド実行、実際の値取得
	actual, _ := target.GetUser("1")

	// DeepEqual による比較
	if reflect.DeepEqual(actual, ret) {
		fmt.Printf("success \n")
		return
	}
	// DeepEqual で等価にならなかった場合
	fmt.Errorf("failed")
}
