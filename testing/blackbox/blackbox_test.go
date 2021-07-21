package main

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/tracey7d4/bsbLookup/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"testing"
)

func TestBSBLookup(t *testing.T) {
	cc, err := grpc.Dial("bsblookup:8080", grpc.WithInsecure())
	if err != nil {
		return
	}
	client := proto.NewBsbLookupClient(cc)

	tests := []struct {
		name 	string
		bsb 	string
		want	*proto.ValidateResponse
		wantErr	error
	}{
		{
			name:    "success",
			bsb:     "012020",
			want:    &proto.ValidateResponse{
				Valid:    true,
				BankCode: "ANZ",
			},
			wantErr: nil,
		},
		{
			name:    "unsuccessful - bsb not available",
			bsb:     "999999",
			want:    &proto.ValidateResponse{
				Valid:    false,
				BankCode: "",
			},
			wantErr: nil,
		},
		{
			name:    "unsuccessful - invalid bsb number",
			bsb:     "6380",
			want:    nil,
			wantErr: status.Error(codes.InvalidArgument, "Invalid bsb number"),
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func (t *testing.T) {
			res, err := client.Validate(context.Background(), &proto.ValidateRequest{Bsb: tt.bsb})
			if tt.wantErr == nil {
				assert.Nil(t, err)
				assert.Equal(t, res.BankCode, tt.want.BankCode)
				assert.Equal(t, res.Valid, tt.want.Valid)
			} else {
				assert.NotNil(t, err)
			}
		})
	}


}
