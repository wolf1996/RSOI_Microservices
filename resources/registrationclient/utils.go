package registrationclient

import (
"fmt"
"google.golang.org/grpc/codes"
"google.golang.org/grpc/status"
"net/http"
)

var SomeServiceError = fmt.Errorf("Some service error")

func ErrorTransform(transError error) (err error){
	err = transError
	if transError == ConnectionError{
		err = SomeServiceError
	} else {
		if resp,ok  := status.FromError(transError);ok {
			if resp.Code() != codes.OK {
				err = SomeServiceError
			}
		}
	}
	return
}

func StatusCodeFromError(err error) int{
	if err == nil {
		return http.StatusOK
	}
	if err == SomeServiceError {
		return http.StatusGatewayTimeout
	}
	return http.StatusNotFound
}
