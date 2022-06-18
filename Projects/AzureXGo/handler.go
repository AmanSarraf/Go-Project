package main

import (
	"encoding/json"
	"net/http"
)

func compvs(w http.ResponseWriter, r *http.Request) {

	newurl := myurl{}
	_ = json.NewDecoder(r.Body).Decode(&newurl)

	value := newurl.Url
	BatchReadFileRemoteImage(value)
	return

}
