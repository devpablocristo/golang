package application

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"

	mocks "github.com/devpablocristo/growuphr/mocks"
	domain "github.com/devpablocristo/growuphr/number-manager/domain"
	utils "github.com/devpablocristo/growuphr/utils/test-utils"
)

func TestNumberManager_AddReserveNumber(t *testing.T) {
	type args struct {
		ctx       context.Context
		newResNum *domain.ReservedNumber
	}
	tests := []struct {
		name                     string
		args                     args
		wantErr                  bool
		expectedErr              error
		calledTimes              int
		expectedCheckZeroNumber  bool
		expectedCheckForUsername bool
		expectedCheckForNumber   bool
		expectedRN               *domain.ReservedNumber
	}{
		{
			name: "new username, new number (happy path)",
			args: args{
				ctx:       context.Background(),
				newResNum: utils.ResNum1,
			},
			wantErr:                  false,
			expectedErr:              nil,
			expectedRN:               nil,
			expectedCheckZeroNumber:  false,
			expectedCheckForUsername: false,
			expectedCheckForNumber:   false,
			calledTimes:              1,
		},
		{
			// no ocurre
			// para crear un nuevo usuario se debe asignar un nuevo valido
			// o sea, no existen nuevos usuarios sin numeros asignados
			name: "existing username, new number",
			args: args{
				ctx:       context.Background(),
				newResNum: utils.ResNum1,
			},
			wantErr:                  true,
			expectedErr:              utils.ErrorExistingUser,
			expectedRN:               utils.ResNum1,
			expectedCheckZeroNumber:  false,
			expectedCheckForNumber:   false,
			expectedCheckForUsername: true,
			calledTimes:              1,
		},
		{
			name: "new username, taken number",
			args: args{
				ctx:       context.Background(),
				newResNum: utils.ResNum1,
			},
			wantErr:                  true,
			expectedErr:              utils.ErrorTakenNumber,
			expectedRN:               utils.ResNum1,
			expectedCheckZeroNumber:  false,
			expectedCheckForUsername: false,
			expectedCheckForNumber:   true,
			calledTimes:              1,
		},
		{
			name: "existing username, taken number",
			args: args{
				ctx:       context.Background(),
				newResNum: utils.ResNum1,
			},
			wantErr:                  true,
			expectedErr:              utils.ErrorTakenNumber,
			expectedRN:               utils.ResNum1,
			expectedCheckZeroNumber:  false,
			expectedCheckForUsername: true,
			expectedCheckForNumber:   true,
			calledTimes:              1,
		},
		{
			name: "taken 0 as chosen number",
			args: args{
				ctx:       context.Background(),
				newResNum: utils.ResNum0,
			},
			wantErr:                  true,
			expectedErr:              utils.ErrorZeroNumber,
			expectedRN:               utils.ResNum1,
			expectedCheckZeroNumber:  true,
			expectedCheckForUsername: true,
			expectedCheckForNumber:   true,
			calledTimes:              1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			storageMock := mocks.NewMockStorage(ctrl)

			if !tt.expectedCheckZeroNumber {
				storageMock.EXPECT().CheckForUsername(tt.args.ctx, tt.args.newResNum.User.Username).
					Return(tt.expectedRN, tt.expectedCheckForUsername).
					Times(tt.calledTimes)

				if !tt.expectedCheckForUsername {
					storageMock.EXPECT().CheckForNumber(tt.args.ctx, tt.args.newResNum.Number.Number).
						Return(tt.expectedRN, tt.expectedCheckForNumber).
						Times(tt.calledTimes)

					if !tt.expectedCheckForNumber {
						storageMock.EXPECT().Create(tt.args.ctx, tt.args.newResNum).
							Return(tt.expectedErr).
							Times(tt.calledTimes)
					}
				}
			}

			nm := NewNumberManager(storageMock)
			err := nm.AddReserveNumber(tt.args.ctx, tt.args.newResNum)
			if (err != nil) != tt.wantErr {
				t.Errorf("NumberManager.AddReserveNumber() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
