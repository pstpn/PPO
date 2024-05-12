package service

import (
	"context"
	"fmt"
	"reflect"
	"testing"

	"course/internal/model"
	"course/internal/service/dto"
	"course/internal/storage/mocks"
)

func Test_infoCardServiceImpl_CreateInfoCard(t *testing.T) {
	ctx := context.TODO()

	type args struct {
		ctx     context.Context
		request *dto.CreateInfoCardRequest
	}

	type storages struct {
		infoCardStorage struct {
			storageArgs   args
			storageReturn struct {
				infoCard *model.InfoCard
				err      error
			}
		}
	}

	infoCardMockStorage := mocks.NewInfoCardStorage(t)
	tests := []struct {
		name    string
		i       *infoCardServiceImpl
		args    args
		want    *model.InfoCard
		wantErr bool

		storages storages
	}{
		{
			name: "incorrect employee ID",
			i: &infoCardServiceImpl{
				logger:          NewMockLogger(),
				infoCardStorage: infoCardMockStorage,
			},
			args: args{
				ctx: ctx,
				request: &dto.CreateInfoCardRequest{
					EmployeeID:  -1,
					IsConfirmed: false,
					CreatedDate: nil,
				},
			},
			want:    nil,
			wantErr: true,

			storages: storages{
				infoCardStorage: struct {
					storageArgs   args
					storageReturn struct {
						infoCard *model.InfoCard
						err      error
					}
				}{
					storageArgs: args{
						ctx: ctx,
						request: &dto.CreateInfoCardRequest{
							EmployeeID:  -1,
							IsConfirmed: false,
							CreatedDate: nil,
						},
					},
					storageReturn: struct {
						infoCard *model.InfoCard
						err      error
					}{
						infoCard: nil,
						err:      fmt.Errorf("incorrect employeeID"),
					},
				},
			},
		},
		{
			name: "success",
			i: &infoCardServiceImpl{
				logger:          NewMockLogger(),
				infoCardStorage: infoCardMockStorage,
			},
			args: args{
				ctx: ctx,
				request: &dto.CreateInfoCardRequest{
					EmployeeID:  1,
					IsConfirmed: false,
					CreatedDate: nil,
				},
			},
			want: &model.InfoCard{
				ID:                model.ToInfoCardID(1),
				CreatedEmployeeID: model.ToEmployeeID(1),
				IsConfirmed:       false,
				CreatedDate:       nil,
			},
			wantErr: false,

			storages: storages{
				infoCardStorage: struct {
					storageArgs   args
					storageReturn struct {
						infoCard *model.InfoCard
						err      error
					}
				}{
					storageArgs: args{
						ctx: ctx,
						request: &dto.CreateInfoCardRequest{
							EmployeeID:  1,
							IsConfirmed: false,
							CreatedDate: nil,
						},
					},
					storageReturn: struct {
						infoCard *model.InfoCard
						err      error
					}{
						infoCard: &model.InfoCard{
							ID:                model.ToInfoCardID(1),
							CreatedEmployeeID: model.ToEmployeeID(1),
							IsConfirmed:       false,
							CreatedDate:       nil,
						},
						err: nil,
					},
				},
			},
		},
	}

	for _, tt := range tests {
		infoCardMockStorage.
			On("Create",
				tt.storages.infoCardStorage.storageArgs.ctx,
				tt.storages.infoCardStorage.storageArgs.request,
			).
			Return(
				tt.storages.infoCardStorage.storageReturn.infoCard,
				tt.storages.infoCardStorage.storageReturn.err,
			).
			Once()
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.i.CreateInfoCard(tt.args.ctx, tt.args.request)
			if (err != nil) != tt.wantErr {
				t.Errorf("infoCardServiceImpl.CreateInfoCard() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("infoCardServiceImpl.CreateInfoCard() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_infoCardServiceImpl_ValidateInfoCard(t *testing.T) {
	ctx := context.TODO()

	type args struct {
		ctx     context.Context
		request *dto.ValidateInfoCardRequest
	}

	type storages struct {
		infoCardStorage struct {
			storageArgs   args
			storageReturn struct {
				err error
			}
		}
	}

	infoCardMockStorage := mocks.NewInfoCardStorage(t)
	tests := []struct {
		name    string
		i       *infoCardServiceImpl
		args    args
		wantErr bool

		storages storages
	}{
		{
			name: "incorrect info card ID",
			i: &infoCardServiceImpl{
				logger:          NewMockLogger(),
				infoCardStorage: infoCardMockStorage,
			},
			args: args{
				ctx: ctx,
				request: &dto.ValidateInfoCardRequest{
					InfoCardID:  -1,
					IsConfirmed: true,
				},
			},
			wantErr: true,

			storages: storages{
				infoCardStorage: struct {
					storageArgs   args
					storageReturn struct {
						err error
					}
				}{
					storageArgs: args{
						ctx: ctx,
						request: &dto.ValidateInfoCardRequest{
							InfoCardID:  -1,
							IsConfirmed: true,
						},
					},
					storageReturn: struct {
						err error
					}{
						err: fmt.Errorf("incorrect infocardID"),
					},
				},
			},
		},
		{
			name: "success",
			i: &infoCardServiceImpl{
				logger:          NewMockLogger(),
				infoCardStorage: infoCardMockStorage,
			},
			args: args{
				ctx: ctx,
				request: &dto.ValidateInfoCardRequest{
					InfoCardID:  1,
					IsConfirmed: true,
				},
			},
			wantErr: false,

			storages: storages{
				infoCardStorage: struct {
					storageArgs   args
					storageReturn struct {
						err error
					}
				}{
					storageArgs: args{
						ctx: ctx,
						request: &dto.ValidateInfoCardRequest{
							InfoCardID:  1,
							IsConfirmed: true,
						},
					},
					storageReturn: struct {
						err error
					}{
						err: nil,
					},
				},
			},
		},
	}

	for _, tt := range tests {
		infoCardMockStorage.
			On("Validate",
				tt.storages.infoCardStorage.storageArgs.ctx,
				tt.storages.infoCardStorage.storageArgs.request,
			).
			Return(
				tt.storages.infoCardStorage.storageReturn.err,
			).
			Once()
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.i.ValidateInfoCard(tt.args.ctx, tt.args.request); (err != nil) != tt.wantErr {
				t.Errorf("infoCardServiceImpl.ValidateInfoCard() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_infoCardServiceImpl_GetInfoCard(t *testing.T) {
	ctx := context.TODO()

	type args struct {
		ctx     context.Context
		request *dto.GetInfoCardByIDRequest
	}

	type storages struct {
		infoCardStorage struct {
			storageArgs   args
			storageReturn struct {
				infoCard *model.InfoCard
				err      error
			}
		}
	}

	infoCardMockStorage := mocks.NewInfoCardStorage(t)
	tests := []struct {
		name    string
		i       *infoCardServiceImpl
		args    args
		want    *model.InfoCard
		wantErr bool

		storages storages
	}{
		{
			name: "incorrect info card ID",
			i: &infoCardServiceImpl{
				logger:          NewMockLogger(),
				infoCardStorage: infoCardMockStorage,
			},
			args: args{
				ctx:     ctx,
				request: &dto.GetInfoCardByIDRequest{InfoCardID: -1},
			},
			want:    nil,
			wantErr: true,

			storages: storages{
				infoCardStorage: struct {
					storageArgs   args
					storageReturn struct {
						infoCard *model.InfoCard
						err      error
					}
				}{
					storageArgs: args{
						ctx:     ctx,
						request: &dto.GetInfoCardByIDRequest{InfoCardID: -1},
					},
					storageReturn: struct {
						infoCard *model.InfoCard
						err      error
					}{
						infoCard: nil,
						err:      fmt.Errorf("incorrect infocardID"),
					},
				},
			},
		},
		{
			name: "success",
			i: &infoCardServiceImpl{
				logger:          NewMockLogger(),
				infoCardStorage: infoCardMockStorage,
			},
			args: args{
				ctx:     ctx,
				request: &dto.GetInfoCardByIDRequest{InfoCardID: 1},
			},
			want: &model.InfoCard{
				ID:                model.ToInfoCardID(1),
				CreatedEmployeeID: model.ToEmployeeID(1),
				IsConfirmed:       true,
				CreatedDate:       nil,
			},
			wantErr: false,

			storages: storages{
				infoCardStorage: struct {
					storageArgs   args
					storageReturn struct {
						infoCard *model.InfoCard
						err      error
					}
				}{
					storageArgs: args{
						ctx:     ctx,
						request: &dto.GetInfoCardByIDRequest{InfoCardID: 1},
					},
					storageReturn: struct {
						infoCard *model.InfoCard
						err      error
					}{
						infoCard: &model.InfoCard{
							ID:                model.ToInfoCardID(1),
							CreatedEmployeeID: model.ToEmployeeID(1),
							IsConfirmed:       true,
							CreatedDate:       nil,
						},
						err: nil,
					},
				},
			},
		},
	}

	for _, tt := range tests {
		infoCardMockStorage.
			On("GetByID",
				tt.storages.infoCardStorage.storageArgs.ctx,
				tt.storages.infoCardStorage.storageArgs.request,
			).
			Return(
				tt.storages.infoCardStorage.storageReturn.infoCard,
				tt.storages.infoCardStorage.storageReturn.err,
			).
			Once()
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.i.GetInfoCard(tt.args.ctx, tt.args.request)
			if (err != nil) != tt.wantErr {
				t.Errorf("infoCardServiceImpl.GetInfoCard() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("infoCardServiceImpl.GetInfoCard() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_infoCardServiceImpl_ListInfoCards(t *testing.T) {
	ctx := context.TODO()

	type args struct {
		ctx     context.Context
		request *dto.ListInfoCardsRequest
	}

	type storages struct {
		infoCardStorage struct {
			storageArgs   args
			storageReturn struct {
				infoCards []*model.InfoCard
				err       error
			}
		}
	}

	infoCardMockStorage := mocks.NewInfoCardStorage(t)
	tests := []struct {
		name    string
		i       *infoCardServiceImpl
		args    args
		want    []*model.InfoCard
		wantErr bool

		storages storages
	}{
		{
			name: "incorrect request",
			i: &infoCardServiceImpl{
				logger:          NewMockLogger(),
				infoCardStorage: infoCardMockStorage,
			},
			args: args{
				ctx:     ctx,
				request: &dto.ListInfoCardsRequest{},
			},
			want:    nil,
			wantErr: true,

			storages: storages{
				infoCardStorage: struct {
					storageArgs   args
					storageReturn struct {
						infoCards []*model.InfoCard
						err       error
					}
				}{
					storageArgs: args{
						ctx:     ctx,
						request: &dto.ListInfoCardsRequest{},
					},
					storageReturn: struct {
						infoCards []*model.InfoCard
						err       error
					}{
						infoCards: nil,
						err:       fmt.Errorf("incorrect request"),
					},
				},
			},
		},
		{
			name: "success",
			i: &infoCardServiceImpl{
				logger:          NewMockLogger(),
				infoCardStorage: infoCardMockStorage,
			},
			args: args{
				ctx:     ctx,
				request: &dto.ListInfoCardsRequest{},
			},
			want: []*model.InfoCard{
				{
					ID:                model.ToInfoCardID(1),
					CreatedEmployeeID: model.ToEmployeeID(1),
					IsConfirmed:       true,
					CreatedDate:       nil,
				},
			},
			wantErr: false,

			storages: storages{
				infoCardStorage: struct {
					storageArgs   args
					storageReturn struct {
						infoCards []*model.InfoCard
						err       error
					}
				}{
					storageArgs: args{
						ctx:     ctx,
						request: &dto.ListInfoCardsRequest{},
					},
					storageReturn: struct {
						infoCards []*model.InfoCard
						err       error
					}{
						infoCards: []*model.InfoCard{
							{
								ID:                model.ToInfoCardID(1),
								CreatedEmployeeID: model.ToEmployeeID(1),
								IsConfirmed:       true,
								CreatedDate:       nil,
							},
						},
						err: nil,
					},
				},
			},
		},
	}

	for _, tt := range tests {
		infoCardMockStorage.
			On("List",
				tt.storages.infoCardStorage.storageArgs.ctx,
				tt.storages.infoCardStorage.storageArgs.request,
			).
			Return(
				tt.storages.infoCardStorage.storageReturn.infoCards,
				tt.storages.infoCardStorage.storageReturn.err,
			).
			Once()
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.i.ListInfoCards(tt.args.ctx, tt.args.request)
			if (err != nil) != tt.wantErr {
				t.Errorf("infoCardServiceImpl.ListInfoCards() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("infoCardServiceImpl.ListInfoCards() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_infoCardServiceImpl_DeleteInfoCard(t *testing.T) {
	ctx := context.TODO()

	type args struct {
		ctx     context.Context
		request *dto.DeleteInfoCardRequest
	}

	type storages struct {
		infoCardStorage struct {
			storageArgs   args
			storageReturn struct {
				err error
			}
		}
	}

	infoCardMockStorage := mocks.NewInfoCardStorage(t)
	tests := []struct {
		name    string
		i       *infoCardServiceImpl
		args    args
		wantErr bool

		storages storages
	}{
		{
			name: "incorrect info card ID",
			i: &infoCardServiceImpl{
				logger:          NewMockLogger(),
				infoCardStorage: infoCardMockStorage,
			},
			args: args{
				ctx:     ctx,
				request: &dto.DeleteInfoCardRequest{InfoCardID: -1},
			},
			wantErr: true,

			storages: storages{
				infoCardStorage: struct {
					storageArgs   args
					storageReturn struct {
						err error
					}
				}{
					storageArgs: args{
						ctx:     ctx,
						request: &dto.DeleteInfoCardRequest{InfoCardID: -1},
					},
					storageReturn: struct {
						err error
					}{
						err: fmt.Errorf("incorrect infocardID"),
					},
				},
			},
		},
		{
			name: "success",
			i: &infoCardServiceImpl{
				logger:          NewMockLogger(),
				infoCardStorage: infoCardMockStorage,
			},
			args: args{
				ctx:     ctx,
				request: &dto.DeleteInfoCardRequest{InfoCardID: 1},
			},
			wantErr: false,

			storages: storages{
				infoCardStorage: struct {
					storageArgs   args
					storageReturn struct {
						err error
					}
				}{
					storageArgs: args{
						ctx:     ctx,
						request: &dto.DeleteInfoCardRequest{InfoCardID: 1},
					},
					storageReturn: struct {
						err error
					}{
						err: nil,
					},
				},
			},
		},
	}

	for _, tt := range tests {
		infoCardMockStorage.
			On("Delete",
				tt.storages.infoCardStorage.storageArgs.ctx,
				tt.storages.infoCardStorage.storageArgs.request,
			).
			Return(
				tt.storages.infoCardStorage.storageReturn.err,
			).
			Once()
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.i.DeleteInfoCard(tt.args.ctx, tt.args.request); (err != nil) != tt.wantErr {
				t.Errorf("infoCardServiceImpl.DeleteInfoCard() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
