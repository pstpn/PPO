package service

import (
	"context"
	"testing"
)

func Test_authServiceImpl_RegisterEmployee(t *testing.T) {
	type args struct {
		ctx     context.Context
		request *RegisterEmployeeRequest
	}
	tests := []struct {
		name    string
		a       *authServiceImpl
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.a.RegisterEmployee(tt.args.ctx, tt.args.request); (err != nil) != tt.wantErr {
				t.Errorf("authServiceImpl.RegisterEmployee() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_authServiceImpl_LoginEmployee(t *testing.T) {
	type args struct {
		ctx     context.Context
		request *LoginEmployeeRequest
	}
	tests := []struct {
		name    string
		a       *authServiceImpl
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.a.LoginEmployee(tt.args.ctx, tt.args.request); (err != nil) != tt.wantErr {
				t.Errorf("authServiceImpl.LoginEmployee() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
