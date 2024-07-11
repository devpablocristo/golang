package usecase_test

import (
	"errors"
	"fmt"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/osalomon89/go-testing/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/teamcubation/neocamp-meli/testing/repository"
	"github.com/teamcubation/neocamp-meli/testing/usecase"
)

//SaveItem(name string, stock int) error
//GetItemByID(itemID uint) error

type itemRepositoryMock struct {
	err error
}

func (repo itemRepositoryMock) SaveItem(name string, stock int) error {
	return repo.err
}

func (repo itemRepositoryMock) GetItemByID(itemID uint) error {
	return repo.err
}

func Test_itemUsecase_CreateItem_Manual(t *testing.T) {
	type args struct {
		name  string
		stock int
	}

	tests := []struct {
		name      string
		repo      repository.ItemRepository
		args      args
		wantedErr error
	}{
		{
			name: "Should works correctly",
			repo: itemRepositoryMock{
				err: nil,
			},
			args: args{
				name:  "tablet",
				stock: 12,
			},
			wantedErr: nil,
		},
		{
			name: "Should return an error when repository returns an error",
			repo: itemRepositoryMock{
				err: errors.New("the repository error"),
			},
			args: args{
				name:  "tablet",
				stock: 12,
			},
			wantedErr: fmt.Errorf("error in repository: %w", errors.New("the repository error")),
		},
		{
			name: "Should return an error when the name is empty",
			repo: itemRepositoryMock{
				err: nil,
			},
			args: args{
				name:  "",
				stock: 12,
			},
			wantedErr: fmt.Errorf("item name could not be empty"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			svc := usecase.NewItemUsecase(tt.repo)

			err := svc.CreateItem(tt.args.name, tt.args.stock)
			assert.Equal(t, err, tt.wantedErr, "itemUsecase.CreateItem() error = %v, wantErr %v", err, tt.wantedErr)
		})
	}
}

func Test_itemService_CreateItem(t *testing.T) {
	assert := assert.New(t)

	type args struct {
		name  string
		stock int
	}

	tests := []struct {
		name      string
		args      args
		repoError error
		repoTimes int
		wantedErr error
	}{
		{
			name:      "Should work correctly",
			wantedErr: nil,
			args: args{
				name:  "tablet",
				stock: 10,
			},
			repoError: nil,
			repoTimes: 1,
		},
		{
			name:      "Should return error when item name is empty",
			wantedErr: fmt.Errorf("item name could not be empty"),
			args: args{
				name:  "",
				stock: 10,
			},
			repoError: nil,
			repoTimes: 0,
		},
		{
			name:      "Should return error when item stock is zero",
			wantedErr: fmt.Errorf("stock could not be zero"),
			args: args{
				name:  "tablet",
				stock: 0,
			},
			repoError: nil,
			repoTimes: 0,
		},
		{
			name:      "Should return error when repository returns an error",
			wantedErr: fmt.Errorf("error in repository: %w", errors.New("the repository error")),
			args: args{
				name:  "tablet",
				stock: 10,
			},
			repoError: errors.New("the repository error"),
			repoTimes: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			repositoryMock := mocks.NewMockItemRepository(ctrl)

			repositoryMock.EXPECT().
				SaveItem(tt.args.name, tt.args.stock).
				Return(tt.repoError).
				Times(tt.repoTimes)

			svc := usecase.NewItemUsecase(repositoryMock)

			err := svc.CreateItem(tt.args.name, tt.args.stock)
			if tt.wantedErr != nil {
				assert.NotNil(err)
				assert.Equal(tt.wantedErr.Error(), err.Error(), "Error message is not the expected")
				return
			}

			assert.Nil(err)
		})
	}
}

func Test_itemUsecase_GetItemByID(t *testing.T) {
	type args struct {
		id uint
	}
	tests := []struct {
		name      string
		repo      repository.ItemRepository
		args      args
		wantErr   bool
		repoError error
		repoTimes int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			repositoryMock := mocks.NewMockItemRepository(ctrl)

			repositoryMock.EXPECT().
				GetItemByID(tt.args.id).
				Return(tt.repoError).
				Times(tt.repoTimes)

			svc := usecase.NewItemUsecase(repositoryMock)
			if err := svc.GetItemByID(tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("itemUsecase.GetItemByID() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
