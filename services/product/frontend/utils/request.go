package utils

import "github.com/levigross/grequests"

func CreateJsonRequestOption(body interface{}) *grequests.RequestOptions {
	return &grequests.RequestOptions{
		JSON: body,
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
	}
}
