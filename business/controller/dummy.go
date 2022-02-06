package controller

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/alvinarthas/myrepo/business/domain/dummy"
	"github.com/alvinarthas/myrepo/business/model"
	"github.com/alvinarthas/myrepo/utils/response"
)

func GetDummyList(res http.ResponseWriter, req *http.Request) {
	data, err := dummy.GetDummyList()
	if err != nil {
		log.Println(err)
	}

	response.Success(res, http.StatusOK, data, nil)
}

func GetDummy(res http.ResponseWriter, req *http.Request) {

}

func CreateDummy(res http.ResponseWriter, req *http.Request) {

	var payload model.CreateDummyRequest

	decoder := json.NewDecoder(req.Body)
	if err := decoder.Decode(&payload); err != nil {
		log.Println(err)
	}

	if err := dummy.CreateDummy(payload); err != nil {
		log.Println(err)
	}

	response.Success(res, http.StatusOK, []string{}, nil)
}

func UpdateDummy(res http.ResponseWriter, req *http.Request) {

}

func DeleteDummy(res http.ResponseWriter, req *http.Request) {

}

func UpsertDummy(res http.ResponseWriter, req *http.Request) {

}

func CreateDummies(res http.ResponseWriter, req *http.Request) {

}

func UpdateDummies(res http.ResponseWriter, req *http.Request) {

}
