package utils

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

// This file is mainly to parse the JSON body

func ParseBody (r *http.Request, x interface {}) error {
	// Body will be read and pass to the body
	if body,err := ioutil.ReadAll(r.Body); err == nil {
		// Now unmarshal will parse the byte data to our desired struct and pass to x
		if err := json.Unmarshal([]byte(body), x); err != nil {
			return err
		}
	}

	return nil
}