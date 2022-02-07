package controller

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/alvinarthas/myrepo/business/domain/dummy"
	"github.com/alvinarthas/myrepo/business/model"
	"github.com/alvinarthas/myrepo/utils/response"
	"github.com/go-chi/chi"
)

func GetDummyList(res http.ResponseWriter, req *http.Request) {
	data, err := dummy.GetDummyList()
	if err != nil {
		log.Println(err)
	}

	response.Success(res, http.StatusOK, data, nil)
}

func GetDummy(res http.ResponseWriter, req *http.Request) {
	id := chi.URLParam(req, "id")
	data, err := dummy.GetDummy(id)
	if err != nil {
		log.Println(err)
	}

	response.Success(res, http.StatusOK, data, nil)

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

	response.Success(res, http.StatusCreated, []string{}, nil)
}

func UpdateDummy(res http.ResponseWriter, req *http.Request) {

	var payload model.UpdateDummyRequest

	id := chi.URLParam(req, "id")

	decoder := json.NewDecoder(req.Body)
	if err := decoder.Decode(&payload); err != nil {
		log.Println(err)
	}

	payload.ID = id
	if err := dummy.UpdateDummy(payload); err != nil {
		log.Println(err)
	}

	response.Success(res, http.StatusOK, []string{}, nil)
}

func DeleteDummy(res http.ResponseWriter, req *http.Request) {
	id := chi.URLParam(req, "id")
	err := dummy.DeleteDummy(id)
	if err != nil {
		log.Println(err)
	}

	response.Success(res, http.StatusOK, []string{}, nil)
}

func UpsertDummy(res http.ResponseWriter, req *http.Request) {

}

func CreateDummies(res http.ResponseWriter, req *http.Request) {

}

func UpdateDummies(res http.ResponseWriter, req *http.Request) {

}
