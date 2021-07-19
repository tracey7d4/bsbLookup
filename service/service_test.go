package service

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/tracey7d4/bsbLookup/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"testing"
)

func Test_bsbValidate(t *testing.T) {
	type args struct {
		bsb string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "success",
			args: args{"638010"},
			want: true,
		},
		{
			name: "invalid bsb number - 4 digit number",
			args: args{"6380"},
			want: false,
		},
		{
			name: "invalid bsb number - contain letters",
			args: args{"aaa638"},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := bsbValidate(tt.args.bsb); got != tt.want {
				t.Errorf("bsbValidate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLookupAPI_Validate(t *testing.T) {
	type fields struct {
		UnimplementedBsbLookupServer proto.UnimplementedBsbLookupServer
		cache                        map[string]string
	}
	tests := []struct {
		name    string
		fields  fields
		bsb    string
		want    *proto.ValidateResponse
		wantErr error
	}{
		{
			name:    "successful",
			fields:  fields{proto.UnimplementedBsbLookupServer{}, map[string]string{"638010":"HBL"}},
			bsb:	"638010",
			want:    &proto.ValidateResponse{
				Valid:    true,
				BankCode: "HBL",
			},
			wantErr: nil,
		},
		{
			name:    "unsuccessful - bsb not available",
			fields:  fields{proto.UnimplementedBsbLookupServer{}, map[string]string{"638010":"HBL"}},
			bsb:	"638011",
			want:    &proto.ValidateResponse{
				Valid:    false,
				BankCode: "",
			},
			wantErr: nil,
		},
		{
			name:    "unsuccessful - invalid bsb",
			fields:  fields{proto.UnimplementedBsbLookupServer{}, map[string]string{"638010":"HBL"}},
			bsb:	"638",
			want:    nil,
			wantErr: status.Error(codes.InvalidArgument, "Invalid bsb number"),
		},

	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			s := &LookupAPI{
				UnimplementedBsbLookupServer: tt.fields.UnimplementedBsbLookupServer,
				cache:                        tt.fields.cache,
			}
			got, err := s.Validate(context.Background(),&proto.ValidateRequest{Bsb: tt.bsb})
			if tt.wantErr == nil {
				assert.Nil(t, err)
				assert.Equal(t, got.Valid, tt.want.Valid)
				assert.Equal(t, got.BankCode, tt.want.BankCode)
			} else {
				assert.NotNil(t, err)
			}
		})
	}
}