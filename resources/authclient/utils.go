package authclient

import (
"fmt"
"google.golang.org/grpc/codes"
"google.golang.org/grpc/status"
"net/http"
)

var(
	SomeServiceError = fmt.Errorf("Some service error")
	BadLogin = fmt.Errorf("Not Found")
)



func statusToHttp(stat *status.Status)(err error, code int){
	switch stat.Code() {
	case codes.NotFound:
		return BadLogin, http.StatusForbidden
	case codes.InvalidArgument:
		return fmt.Errorf(stat.Message()), http.StatusBadRequest
	default:
		return SomeServiceError, http.StatusInternalServerError
	}
}

func ErrorTransform(transError error) (err error, code int){
	err = transError
	if transError == ConnectionError{
		err = SomeServiceError
		code = http.StatusServiceUnavailable
	} else {
		if resp,ok  := status.FromError(transError);ok {
			return statusToHttp(resp)
		}
	}
	return
}
