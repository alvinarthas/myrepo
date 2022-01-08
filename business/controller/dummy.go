package controller

import "net/http"

func GetDummy(res http.ResponseWriter, req *http.Request) {

	res.Write([]byte("Get Dummy Page"))
}
