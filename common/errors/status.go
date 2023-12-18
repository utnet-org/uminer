package errors

import (
	"strconv"

	"github.com/golang/protobuf/proto"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func ErrorEncode(err error) error {
	se, ok := FromError(err)
	if !ok {
		se, _ = FromError(Errorf(nil, ErrorUnknown))
	}
	gs := status.Newf(codes.Unknown, "%s", se.Message)
	details := []proto.Message{
		&errdetails.ErrorInfo{
			Metadata: map[string]string{"codeType": strconv.Itoa(int(se.CodeType)), "code": strconv.Itoa(se.Code), "message": se.Message},
		},
	}

	gs, err = gs.WithDetails(details...)
	if err != nil {
		return err
	}
	return gs.Err()
}

func ErrorDecode(err error) error {
	gs := status.Convert(err)

	var (
		codeType int
		code     int
		msg      string
	)

	if len(gs.Details()) <= 0 {
		codeType = int(Unknown)
		code = ErrorUnknown
		msg = ""
	} else {
		for _, detail := range gs.Details() {
			switch d := detail.(type) {
			case *errdetails.ErrorInfo:
				codeType, err = strconv.Atoi(d.Metadata["codeType"])
				if err != nil {
					codeType = int(Unknown)
				}
				code, err = strconv.Atoi(d.Metadata["code"])
				if err != nil {
					code = ErrorUnknown
				}
				msg = d.Metadata["message"]
			}
		}
	}

	return Errorw(nil, (CodeType(codeType)), code, msg)
}
