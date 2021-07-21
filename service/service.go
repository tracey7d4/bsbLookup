package service

import (
	"context"
	"encoding/csv"
	"fmt"
	"github.com/tracey7d4/bsbLookup/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"os"
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

//func (s *LookupAPI) UpdateCache() error {
//	s.cache = map[string]string{"638010": "HBL"}
//	return nil
//}

func (s *LookupAPI) UpdateCache() error {
	csvFile, err := os.Open("service/data/bsbcache.csv")
	if err != nil {
		fmt.Println("Error open csv file")
		return err
	}
	defer func() {
		_ = csvFile.Close()
	}()

	csvLines, err := csv.NewReader(csvFile).ReadAll()
	if err != nil {
		return status.Error(codes.InvalidArgument, "Error reading csv file")
	}
	s.cache = make(map[string]string)
	for _, lines := range csvLines {
		line0 := lines[0]
		bsb := line0[:3]+line0[4:]
		bankCode := lines[1]
		if _, ok := s.cache[bsb]; !ok {
			s.cache[bsb] = bankCode
		}
	}
	return nil
}