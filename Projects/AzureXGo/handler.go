package main

import (
	"encoding/json"
	"net/http"
)

func compvs(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		{
			newurl := myurl{}
			_ = json.NewDecoder(r.Body).Decode(&newurl)

			value := newurl.Url
			BatchReadFileRemoteImage(value)

		}
	case http.MethodGet:
		{
			url := "https:" + r.FormValue("url")
			BatchReadFileRemoteImage(url)
		}
	}

}
