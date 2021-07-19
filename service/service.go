package service

import (
	"context"
	"github.com/tracey7d4/bsbLookup/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"regexp"
)

type LookupAPI struct {
	proto.UnimplementedBsbLookupServer
	cache map[string]string
}

func (s *LookupAPI) Validate(ctx context.Context, request *proto.ValidateRequest) (*proto.ValidateResponse, error){
	// check to see a bsb is valid
	if !bsbValidate(request.Bsb) {
		return nil, status.Error(codes.InvalidArgument, "Invalid BSB number")
	}
	// check to see it's a real bsb number
	if bankCode, ok := s.cache[request.GetBsb()]; ok {
		return &proto.ValidateResponse{
			Valid:    true,
			BankCode: bankCode,
		}, nil
	} else {
		return &proto.ValidateResponse{
			Valid:    false,
			BankCode: "",
		}, nil
	}
}

func bsbValidate(bsb string) bool {
	// bsb should be a 6 digit number
	var validBSB = regexp.MustCompile(`^[0-9]{6}$`)
	return validBSB.MatchString(bsb)
}