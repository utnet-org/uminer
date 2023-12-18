package codec

import (
	"io/ioutil"
	"net/http"
	"strings"
	commctx "uminer/common/context"
	"uminer/common/errors"
	"uminer/common/utils"

	"github.com/go-kratos/kratos/v2/encoding"
	"github.com/go-kratos/kratos/v2/transport/http/binding"
)

const (
	requestIdHeader = "Request-Id"
)

const baseContentType = "application"

var (
	// acceptHeader      = http.CanonicalHeaderKey("Accept")
	contentTypeHeader = http.CanonicalHeaderKey("Content-Type")
)

func contentType(subtype string) string {
	return strings.Join([]string{baseContentType, subtype}, "/")
}

func contentSubtype(contentType string) string {
	if contentType == baseContentType {
		return ""
	}
	if !strings.HasPrefix(contentType, baseContentType) {
		return ""
	}
	// guaranteed since != baseContentType and has baseContentType prefix
	switch contentType[len(baseContentType)] {
	case '/', ';':
		if i := strings.Index(contentType, ";"); i != -1 {
			return contentType[len(baseContentType)+1 : i]
		}
		return contentType[len(baseContentType)+1:]
	default:
		return ""
	}
}

func DecodeRequest(req *http.Request, v interface{}) error {
	*req = *req.WithContext(commctx.RequestIdToContext(req.Context(), utils.GetUUIDWithoutSeparator()))
	subtype := contentSubtype(req.Header.Get(contentTypeHeader))
	if codec := encoding.GetCodec(subtype); codec != nil {
		data, err := ioutil.ReadAll(req.Body)
		if err != nil {
			return errors.Errorf(err, errors.ErrorHttpReadBody)
		}
		err = codec.Unmarshal(data, v)
		if err != nil {
			return errors.Errorf(err, errors.ErrorJsonUnmarshal)
		}
		return nil
	}
	err := binding.BindForm(req, v)
	if err != nil {
		return errors.Errorf(err, errors.ErrorHttpBindFormFailed)
	}
	return nil
}

func EncodeResponse(res http.ResponseWriter, req *http.Request, v interface{}, f func(v interface{}) interface{}) error {
	subtype := contentSubtype(req.Header.Get("accept"))
	codec := encoding.GetCodec(subtype)
	if codec == nil {
		codec = encoding.GetCodec("json")
	}

	data, err := codec.Marshal(f(v))
	if err != nil {
		return errors.Errorf(err, errors.ErrorJsonMarshal)
	}
	res.Header().Set("content-type", contentType(codec.Name()))
	res.Header().Set(requestIdHeader, commctx.RequestIdFromContext(req.Context()))
	_, err = res.Write(data)
	if err != nil {
		return errors.Errorf(err, errors.ErrorHttpWriteFailed)
	}
	return nil
}

func EncodeError(res http.ResponseWriter, req *http.Request, err error, f func(e *errors.SpiderError) interface{}) {
	se, ok := errors.FromError(err)
	if !ok {
		se, _ = errors.FromError(errors.Errorf(nil, errors.ErrorUnknown))
	}

	subtype := contentSubtype(req.Header.Get("accept"))
	codec := encoding.GetCodec(subtype)
	if codec == nil {
		codec = encoding.GetCodec("json")
	}
	data, err := codec.Marshal(f(se))
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		return
	}
	res.Header().Set("content-type", contentType(codec.Name()))
	res.Header().Set(requestIdHeader, commctx.RequestIdFromContext(req.Context()))
	res.WriteHeader(http.StatusOK)
	_, err = res.Write(data)
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		return
	}
}
