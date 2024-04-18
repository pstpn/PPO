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

func Test_photoServiceImpl_CreatePhoto(t *testing.T) {
	ctx := context.TODO()

	type args struct {
		ctx     context.Context
		request *dto.CreatePhotoRequest
	}

	type storages struct {
		photoKeyStorage struct {
			storageArgs struct {
				ctx     context.Context
				request *dto.CreatePhotoKeyRequest
			}
			storageReturn struct {
				photoMeta *model.PhotoMeta
				err       error
			}
		}
		photoStorage struct {
			storageArgs struct {
				ctx     context.Context
				request *dto.CreatePhotoRequest
			}
			storageReturn struct {
				photoKey *model.PhotoKey
				err      error
			}
		}
	}

	photoMockStorage := mocks.NewPhotoStorage(t)
	tests := []struct {
		name    string
		p       *photoServiceImpl
		args    args
		want    *model.PhotoMeta
		wantErr bool

		storages storages
	}{
		{
			name: "incorrect photo data",
			p: &photoServiceImpl{
				logger:       NewMockLogger(),
				photoStorage: photoMockStorage,
			},
			args: args{
				ctx: ctx,
				request: &dto.CreatePhotoRequest{
					DocumentID: 1,
					Data:       nil,
				},
			},
			want:    nil,
			wantErr: true,

			storages: storages{
				photoKeyStorage: struct {
					storageArgs struct {
						ctx     context.Context
						request *dto.CreatePhotoKeyRequest
					}
					storageReturn struct {
						photoMeta *model.PhotoMeta
						err       error
					}
				}{
					storageArgs: struct {
						ctx     context.Context
						request *dto.CreatePhotoKeyRequest
					}{
						ctx: ctx,
						request: &dto.CreatePhotoKeyRequest{
							DocumentID: model.ToDocumentID(1),
							Key:        model.ToPhotoKey("soso"),
						},
					},
					storageReturn: struct {
						photoMeta *model.PhotoMeta
						err       error
					}{
						photoMeta: nil,
						err:       nil,
					},
				},
				photoStorage: struct {
					storageArgs struct {
						ctx     context.Context
						request *dto.CreatePhotoRequest
					}
					storageReturn struct {
						photoKey *model.PhotoKey
						err      error
					}
				}{
					storageArgs: struct {
						ctx     context.Context
						request *dto.CreatePhotoRequest
					}{
						ctx: ctx,
						request: &dto.CreatePhotoRequest{
							DocumentID: 1,
							Data:       nil,
						},
					},
					storageReturn: struct {
						photoKey *model.PhotoKey
						err      error
					}{
						photoKey: nil,
						err:      fmt.Errorf("incorrect photo data"),
					},
				},
			},
		},
		{
			name: "incorrect document ID",
			p: &photoServiceImpl{
				logger:       NewMockLogger(),
				photoStorage: photoMockStorage,
			},
			args: args{
				ctx: ctx,
				request: &dto.CreatePhotoRequest{
					DocumentID: -1,
					Data:       nil,
				},
			},
			want:    nil,
			wantErr: true,

			storages: storages{
				photoKeyStorage: struct {
					storageArgs struct {
						ctx     context.Context
						request *dto.CreatePhotoKeyRequest
					}
					storageReturn struct {
						photoMeta *model.PhotoMeta
						err       error
					}
				}{
					storageArgs: struct {
						ctx     context.Context
						request *dto.CreatePhotoKeyRequest
					}{
						ctx: ctx,
						request: &dto.CreatePhotoKeyRequest{
							DocumentID: model.ToDocumentID(-1),
							Key:        model.ToPhotoKey("soso"),
						},
					},
					storageReturn: struct {
						photoMeta *model.PhotoMeta
						err       error
					}{
						photoMeta: nil,
						err:       fmt.Errorf("incorrect documentID"),
					},
				},
				photoStorage: struct {
					storageArgs struct {
						ctx     context.Context
						request *dto.CreatePhotoRequest
					}
					storageReturn struct {
						photoKey *model.PhotoKey
						err      error
					}
				}{
					storageArgs: struct {
						ctx     context.Context
						request *dto.CreatePhotoRequest
					}{
						ctx: ctx,
						request: &dto.CreatePhotoRequest{
							DocumentID: -1,
							Data:       nil,
						},
					},
					storageReturn: struct {
						photoKey *model.PhotoKey
						err      error
					}{
						photoKey: model.ToPhotoKey("soso"),
						err:      nil,
					},
				},
			},
		},
		{
			name: "success",
			p: &photoServiceImpl{
				logger:       NewMockLogger(),
				photoStorage: photoMockStorage,
			},
			args: args{
				ctx: ctx,
				request: &dto.CreatePhotoRequest{
					DocumentID: 2,
					Data:       []byte{'s'},
				},
			},
			want: &model.PhotoMeta{
				ID:         model.ToPhotoID(1),
				DocumentID: model.ToDocumentID(2),
				PhotoKey:   model.ToPhotoKey("soso"),
			},
			wantErr: false,

			storages: storages{
				photoKeyStorage: struct {
					storageArgs struct {
						ctx     context.Context
						request *dto.CreatePhotoKeyRequest
					}
					storageReturn struct {
						photoMeta *model.PhotoMeta
						err       error
					}
				}{
					storageArgs: struct {
						ctx     context.Context
						request *dto.CreatePhotoKeyRequest
					}{
						ctx: ctx,
						request: &dto.CreatePhotoKeyRequest{
							DocumentID: model.ToDocumentID(2),
							Key:        model.ToPhotoKey("soso"),
						},
					},
					storageReturn: struct {
						photoMeta *model.PhotoMeta
						err       error
					}{
						photoMeta: &model.PhotoMeta{
							ID:         model.ToPhotoID(1),
							DocumentID: model.ToDocumentID(2),
							PhotoKey:   model.ToPhotoKey("soso"),
						},
						err: nil,
					},
				},
				photoStorage: struct {
					storageArgs struct {
						ctx     context.Context
						request *dto.CreatePhotoRequest
					}
					storageReturn struct {
						photoKey *model.PhotoKey
						err      error
					}
				}{
					storageArgs: struct {
						ctx     context.Context
						request *dto.CreatePhotoRequest
					}{
						ctx: ctx,
						request: &dto.CreatePhotoRequest{
							DocumentID: 2,
							Data:       []byte{'s'},
						},
					},
					storageReturn: struct {
						photoKey *model.PhotoKey
						err      error
					}{
						photoKey: model.ToPhotoKey("soso"),
						err:      nil,
					},
				},
			},
		},
	}

	for _, tt := range tests {
		photoMockStorage.
			On("Save",
				tt.storages.photoStorage.storageArgs.ctx,
				tt.storages.photoStorage.storageArgs.request,
			).
			Return(
				tt.storages.photoStorage.storageReturn.photoKey,
				tt.storages.photoStorage.storageReturn.err,
			).
			Once()
		photoMockStorage.
			On("SaveKey",
				tt.storages.photoKeyStorage.storageArgs.ctx,
				tt.storages.photoKeyStorage.storageArgs.request,
			).
			Return(
				tt.storages.photoKeyStorage.storageReturn.photoMeta,
				tt.storages.photoKeyStorage.storageReturn.err,
			).Maybe()
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.p.CreatePhoto(tt.args.ctx, tt.args.request)
			if (err != nil) != tt.wantErr {
				t.Errorf("photoServiceImpl.CreatePhoto() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("photoServiceImpl.CreatePhoto() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_photoServiceImpl_GetPhoto(t *testing.T) {
	ctx := context.TODO()

	type args struct {
		ctx     context.Context
		request *dto.GetPhotoRequest
	}

	type storages struct {
		photoKeyStorage struct {
			storageArgs struct {
				ctx     context.Context
				request *dto.GetPhotoRequest
			}
			storageReturn struct {
				meta *model.PhotoMeta
				err  error
			}
		}
		photoStorage struct {
			storageArgs struct {
				ctx context.Context
				key *model.PhotoKey
			}
			storageReturn struct {
				data []byte
				err  error
			}
		}
	}

	photoMockStorage := mocks.NewPhotoStorage(t)
	tests := []struct {
		name    string
		p       *photoServiceImpl
		args    args
		want    *model.Photo
		wantErr bool

		storages storages
	}{
		{
			name: "incorrect document ID",
			p: &photoServiceImpl{
				logger:       NewMockLogger(),
				photoStorage: photoMockStorage,
			},
			args: args{
				ctx: ctx,
				request: &dto.GetPhotoRequest{
					DocumentID: -1,
				},
			},
			want:    nil,
			wantErr: true,

			storages: storages{
				photoKeyStorage: struct {
					storageArgs struct {
						ctx     context.Context
						request *dto.GetPhotoRequest
					}
					storageReturn struct {
						meta *model.PhotoMeta
						err  error
					}
				}{
					storageArgs: struct {
						ctx     context.Context
						request *dto.GetPhotoRequest
					}{
						ctx: ctx,
						request: &dto.GetPhotoRequest{
							DocumentID: -1,
						},
					},
					storageReturn: struct {
						meta *model.PhotoMeta
						err  error
					}{
						meta: nil,
						err:  fmt.Errorf("incorrect documentID"),
					},
				},
				photoStorage: struct {
					storageArgs struct {
						ctx context.Context
						key *model.PhotoKey
					}
					storageReturn struct {
						data []byte
						err  error
					}
				}{
					storageArgs: struct {
						ctx context.Context
						key *model.PhotoKey
					}{
						ctx: ctx,
						key: nil,
					},
					storageReturn: struct {
						data []byte
						err  error
					}{
						data: nil,
						err:  nil,
					},
				},
			},
		},
		{
			name: "incorrect photo key",
			p: &photoServiceImpl{
				logger:       NewMockLogger(),
				photoStorage: photoMockStorage,
			},
			args: args{
				ctx: ctx,
				request: &dto.GetPhotoRequest{
					DocumentID: 1,
				},
			},
			want:    nil,
			wantErr: true,

			storages: storages{
				photoKeyStorage: struct {
					storageArgs struct {
						ctx     context.Context
						request *dto.GetPhotoRequest
					}
					storageReturn struct {
						meta *model.PhotoMeta
						err  error
					}
				}{
					storageArgs: struct {
						ctx     context.Context
						request *dto.GetPhotoRequest
					}{
						ctx: ctx,
						request: &dto.GetPhotoRequest{
							DocumentID: 1,
						},
					},
					storageReturn: struct {
						meta *model.PhotoMeta
						err  error
					}{
						meta: &model.PhotoMeta{
							ID:       model.ToPhotoID(-1),
							PhotoKey: model.ToPhotoKey("gogo"),
						},
						err: nil,
					},
				},
				photoStorage: struct {
					storageArgs struct {
						ctx context.Context
						key *model.PhotoKey
					}
					storageReturn struct {
						data []byte
						err  error
					}
				}{
					storageArgs: struct {
						ctx context.Context
						key *model.PhotoKey
					}{
						ctx: ctx,
						key: model.ToPhotoKey("gogo"),
					},
					storageReturn: struct {
						data []byte
						err  error
					}{
						data: nil,
						err:  fmt.Errorf("incorrect photo key"),
					},
				},
			},
		},
		{
			name: "success",
			p: &photoServiceImpl{
				logger:       NewMockLogger(),
				photoStorage: photoMockStorage,
			},
			args: args{
				ctx: ctx,
				request: &dto.GetPhotoRequest{
					DocumentID: 1,
				},
			},
			want: &model.Photo{
				Meta: &model.PhotoMeta{
					ID:       model.ToPhotoID(1),
					PhotoKey: model.ToPhotoKey("okok"),
				},
				Data: []byte{'a'},
			},
			wantErr: false,

			storages: storages{
				photoKeyStorage: struct {
					storageArgs struct {
						ctx     context.Context
						request *dto.GetPhotoRequest
					}
					storageReturn struct {
						meta *model.PhotoMeta
						err  error
					}
				}{
					storageArgs: struct {
						ctx     context.Context
						request *dto.GetPhotoRequest
					}{
						ctx: ctx,
						request: &dto.GetPhotoRequest{
							DocumentID: 1,
						},
					},
					storageReturn: struct {
						meta *model.PhotoMeta
						err  error
					}{
						meta: &model.PhotoMeta{
							ID:       model.ToPhotoID(1),
							PhotoKey: model.ToPhotoKey("okok"),
						},
						err: nil,
					},
				},
				photoStorage: struct {
					storageArgs struct {
						ctx context.Context
						key *model.PhotoKey
					}
					storageReturn struct {
						data []byte
						err  error
					}
				}{
					storageArgs: struct {
						ctx context.Context
						key *model.PhotoKey
					}{
						ctx: ctx,
						key: model.ToPhotoKey("okok"),
					},
					storageReturn: struct {
						data []byte
						err  error
					}{
						data: []byte{'a'},
						err:  nil,
					},
				},
			},
		},
	}

	for _, tt := range tests {
		photoMockStorage.
			On("GetKey",
				tt.storages.photoKeyStorage.storageArgs.ctx,
				tt.storages.photoKeyStorage.storageArgs.request,
			).
			Return(
				tt.storages.photoKeyStorage.storageReturn.meta,
				tt.storages.photoKeyStorage.storageReturn.err,
			).
			Once()
		photoMockStorage.
			On("Get",
				tt.storages.photoStorage.storageArgs.ctx,
				tt.storages.photoStorage.storageArgs.key,
			).
			Return(
				tt.storages.photoStorage.storageReturn.data,
				tt.storages.photoStorage.storageReturn.err,
			).Maybe()
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.p.GetPhoto(tt.args.ctx, tt.args.request)
			if (err != nil) != tt.wantErr {
				t.Errorf("photoServiceImpl.GetPhoto() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("photoServiceImpl.GetPhoto() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_photoServiceImpl_DeletePhoto(t *testing.T) {
	ctx := context.TODO()

	type args struct {
		ctx     context.Context
		request *dto.DeletePhotoRequest
	}

	type storages struct {
		photoKeyStorage struct {
			getKey struct {
				storageArgs struct {
					ctx     context.Context
					request *dto.GetPhotoRequest
				}
				storageReturn struct {
					meta *model.PhotoMeta
					err  error
				}
			}
			deleteKey struct {
				storageArgs struct {
					ctx     context.Context
					request *dto.DeletePhotoRequest
				}
				storageReturn struct {
					err error
				}
			}
		}
		photoStorage struct {
			storageArgs struct {
				ctx context.Context
				key *model.PhotoKey
			}
			storageReturn struct {
				err error
			}
		}
	}

	photoMockStorage := mocks.NewPhotoStorage(t)
	tests := []struct {
		name    string
		p       *photoServiceImpl
		args    args
		wantErr bool

		storages storages
	}{
		{
			name: "incorrect document ID (GetKey())",
			p: &photoServiceImpl{
				logger:       NewMockLogger(),
				photoStorage: photoMockStorage,
			},
			args: args{
				ctx: ctx,
				request: &dto.DeletePhotoRequest{
					DocumentID: -1,
				},
			},
			wantErr: true,

			storages: storages{
				photoKeyStorage: struct {
					getKey struct {
						storageArgs struct {
							ctx     context.Context
							request *dto.GetPhotoRequest
						}
						storageReturn struct {
							meta *model.PhotoMeta
							err  error
						}
					}
					deleteKey struct {
						storageArgs struct {
							ctx     context.Context
							request *dto.DeletePhotoRequest
						}
						storageReturn struct {
							err error
						}
					}
				}{
					getKey: struct {
						storageArgs struct {
							ctx     context.Context
							request *dto.GetPhotoRequest
						}
						storageReturn struct {
							meta *model.PhotoMeta
							err  error
						}
					}{
						storageArgs: struct {
							ctx     context.Context
							request *dto.GetPhotoRequest
						}{
							ctx: ctx,
							request: &dto.GetPhotoRequest{
								DocumentID: -1,
							},
						},
						storageReturn: struct {
							meta *model.PhotoMeta
							err  error
						}{
							meta: nil,
							err:  fmt.Errorf("incorrect documentID"),
						},
					},
					deleteKey: struct {
						storageArgs struct {
							ctx     context.Context
							request *dto.DeletePhotoRequest
						}
						storageReturn struct {
							err error
						}
					}{
						storageArgs: struct {
							ctx     context.Context
							request *dto.DeletePhotoRequest
						}{
							ctx: ctx,
							request: &dto.DeletePhotoRequest{
								DocumentID: -1,
							},
						},
						storageReturn: struct {
							err error
						}{
							err: nil,
						},
					},
				},
				photoStorage: struct {
					storageArgs struct {
						ctx context.Context
						key *model.PhotoKey
					}
					storageReturn struct {
						err error
					}
				}{
					storageArgs: struct {
						ctx context.Context
						key *model.PhotoKey
					}{
						ctx: ctx,
						key: nil,
					},
					storageReturn: struct {
						err error
					}{
						err: nil,
					},
				},
			},
		},
		{
			name: "incorrect photo key",
			p: &photoServiceImpl{
				logger:       NewMockLogger(),
				photoStorage: photoMockStorage,
			},
			args: args{
				ctx: ctx,
				request: &dto.DeletePhotoRequest{
					DocumentID: 1,
				},
			},
			wantErr: true,

			storages: storages{
				photoKeyStorage: struct {
					getKey struct {
						storageArgs struct {
							ctx     context.Context
							request *dto.GetPhotoRequest
						}
						storageReturn struct {
							meta *model.PhotoMeta
							err  error
						}
					}
					deleteKey struct {
						storageArgs struct {
							ctx     context.Context
							request *dto.DeletePhotoRequest
						}
						storageReturn struct {
							err error
						}
					}
				}{
					getKey: struct {
						storageArgs struct {
							ctx     context.Context
							request *dto.GetPhotoRequest
						}
						storageReturn struct {
							meta *model.PhotoMeta
							err  error
						}
					}{
						storageArgs: struct {
							ctx     context.Context
							request *dto.GetPhotoRequest
						}{
							ctx: ctx,
							request: &dto.GetPhotoRequest{
								DocumentID: 1,
							},
						},
						storageReturn: struct {
							meta *model.PhotoMeta
							err  error
						}{
							meta: &model.PhotoMeta{
								ID:       model.ToPhotoID(1),
								PhotoKey: model.ToPhotoKey("gg"),
							},
							err: nil,
						},
					},
					deleteKey: struct {
						storageArgs struct {
							ctx     context.Context
							request *dto.DeletePhotoRequest
						}
						storageReturn struct {
							err error
						}
					}{
						storageArgs: struct {
							ctx     context.Context
							request *dto.DeletePhotoRequest
						}{
							ctx: ctx,
							request: &dto.DeletePhotoRequest{
								DocumentID: 1,
							},
						},
						storageReturn: struct {
							err error
						}{
							err: nil,
						},
					},
				},
				photoStorage: struct {
					storageArgs struct {
						ctx context.Context
						key *model.PhotoKey
					}
					storageReturn struct {
						err error
					}
				}{
					storageArgs: struct {
						ctx context.Context
						key *model.PhotoKey
					}{
						ctx: ctx,
						key: model.ToPhotoKey("gg"),
					},
					storageReturn: struct {
						err error
					}{
						err: fmt.Errorf("incorrect photo key"),
					},
				},
			},
		},
		{
			name: "incorrect document ID (DeleteKey())",
			p: &photoServiceImpl{
				logger:       NewMockLogger(),
				photoStorage: photoMockStorage,
			},
			args: args{
				ctx: ctx,
				request: &dto.DeletePhotoRequest{
					DocumentID: -2,
				},
			},
			wantErr: true,

			storages: storages{
				photoKeyStorage: struct {
					getKey struct {
						storageArgs struct {
							ctx     context.Context
							request *dto.GetPhotoRequest
						}
						storageReturn struct {
							meta *model.PhotoMeta
							err  error
						}
					}
					deleteKey struct {
						storageArgs struct {
							ctx     context.Context
							request *dto.DeletePhotoRequest
						}
						storageReturn struct {
							err error
						}
					}
				}{
					getKey: struct {
						storageArgs struct {
							ctx     context.Context
							request *dto.GetPhotoRequest
						}
						storageReturn struct {
							meta *model.PhotoMeta
							err  error
						}
					}{
						storageArgs: struct {
							ctx     context.Context
							request *dto.GetPhotoRequest
						}{
							ctx: ctx,
							request: &dto.GetPhotoRequest{
								DocumentID: -2,
							},
						},
						storageReturn: struct {
							meta *model.PhotoMeta
							err  error
						}{
							meta: &model.PhotoMeta{
								ID:       model.ToPhotoID(2),
								PhotoKey: model.ToPhotoKey("ggg"),
							},
							err: nil,
						},
					},
					deleteKey: struct {
						storageArgs struct {
							ctx     context.Context
							request *dto.DeletePhotoRequest
						}
						storageReturn struct {
							err error
						}
					}{
						storageArgs: struct {
							ctx     context.Context
							request *dto.DeletePhotoRequest
						}{
							ctx: ctx,
							request: &dto.DeletePhotoRequest{
								DocumentID: -2,
							},
						},
						storageReturn: struct {
							err error
						}{
							err: fmt.Errorf("incorrect documentID"),
						},
					},
				},
				photoStorage: struct {
					storageArgs struct {
						ctx context.Context
						key *model.PhotoKey
					}
					storageReturn struct {
						err error
					}
				}{
					storageArgs: struct {
						ctx context.Context
						key *model.PhotoKey
					}{
						ctx: ctx,
						key: model.ToPhotoKey("ggg"),
					},
					storageReturn: struct {
						err error
					}{
						err: nil,
					},
				},
			},
		},
		{
			name: "success",
			p: &photoServiceImpl{
				logger:       NewMockLogger(),
				photoStorage: photoMockStorage,
			},
			args: args{
				ctx: ctx,
				request: &dto.DeletePhotoRequest{
					DocumentID: 2,
				},
			},
			wantErr: false,

			storages: storages{
				photoKeyStorage: struct {
					getKey struct {
						storageArgs struct {
							ctx     context.Context
							request *dto.GetPhotoRequest
						}
						storageReturn struct {
							meta *model.PhotoMeta
							err  error
						}
					}
					deleteKey struct {
						storageArgs struct {
							ctx     context.Context
							request *dto.DeletePhotoRequest
						}
						storageReturn struct {
							err error
						}
					}
				}{
					getKey: struct {
						storageArgs struct {
							ctx     context.Context
							request *dto.GetPhotoRequest
						}
						storageReturn struct {
							meta *model.PhotoMeta
							err  error
						}
					}{
						storageArgs: struct {
							ctx     context.Context
							request *dto.GetPhotoRequest
						}{
							ctx: ctx,
							request: &dto.GetPhotoRequest{
								DocumentID: 2,
							},
						},
						storageReturn: struct {
							meta *model.PhotoMeta
							err  error
						}{
							meta: &model.PhotoMeta{
								ID:       model.ToPhotoID(3),
								PhotoKey: model.ToPhotoKey("okok"),
							},
							err: nil,
						},
					},
					deleteKey: struct {
						storageArgs struct {
							ctx     context.Context
							request *dto.DeletePhotoRequest
						}
						storageReturn struct {
							err error
						}
					}{
						storageArgs: struct {
							ctx     context.Context
							request *dto.DeletePhotoRequest
						}{
							ctx: ctx,
							request: &dto.DeletePhotoRequest{
								DocumentID: 2,
							},
						},
						storageReturn: struct {
							err error
						}{
							err: nil,
						},
					},
				},
				photoStorage: struct {
					storageArgs struct {
						ctx context.Context
						key *model.PhotoKey
					}
					storageReturn struct {
						err error
					}
				}{
					storageArgs: struct {
						ctx context.Context
						key *model.PhotoKey
					}{
						ctx: ctx,
						key: model.ToPhotoKey("okok"),
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
		photoMockStorage.
			On("GetKey",
				tt.storages.photoKeyStorage.getKey.storageArgs.ctx,
				tt.storages.photoKeyStorage.getKey.storageArgs.request,
			).
			Return(
				tt.storages.photoKeyStorage.getKey.storageReturn.meta,
				tt.storages.photoKeyStorage.getKey.storageReturn.err,
			).
			Once()
		photoMockStorage.
			On("Delete",
				tt.storages.photoStorage.storageArgs.ctx,
				tt.storages.photoStorage.storageArgs.key,
			).
			Return(
				tt.storages.photoStorage.storageReturn.err,
			).
			Maybe()
		photoMockStorage.
			On("DeleteKey",
				tt.storages.photoKeyStorage.deleteKey.storageArgs.ctx,
				tt.storages.photoKeyStorage.deleteKey.storageArgs.request,
			).
			Return(
				tt.storages.photoKeyStorage.deleteKey.storageReturn.err,
			).Maybe()
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.p.DeletePhoto(tt.args.ctx, tt.args.request); (err != nil) != tt.wantErr {
				t.Errorf("photoServiceImpl.DeletePhoto() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
