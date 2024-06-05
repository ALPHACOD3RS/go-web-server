package auth

import (
	"errors"
	"net/http"
	"strings"
)

func GetApiKey(headers http.Header) (string, error){

	head := headers.Get("Authorization")

	if head == ""{
		return "", errors.New("no authentication found")
	}

	vals := strings.Split(head, " ")

	if len(vals) != 2{
		return "", errors.New("malformed auth header")
	}

	if vals[0] != "ApiKey"{
		return "", errors.New("malformed the first part of the header")
	}

	return vals[1],nil

}