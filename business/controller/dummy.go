package controller

import (
	"net/http"

	"github.com/alvinarthas/myrepo/business/model"
	"github.com/alvinarthas/myrepo/utils/response"
)

func GetDummyList(res http.ResponseWriter, req *http.Request) {

	data := model.Dummy{
		Name:     "Arthas King",
		Adress:   "Wakanda",
		Age:      12,
		Type:     "Undead",
		IsActive: true,
	}

	response.Success(res, http.StatusOK, data, nil)
}

func GetDummy(res http.ResponseWriter, req *http.Request) {

}

func CreateDummy(res http.ResponseWriter, req *http.Request) {

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
