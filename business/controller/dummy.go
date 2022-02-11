package controller

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/alvinarthas/myrepo/business/domain/dummy"
	"github.com/alvinarthas/myrepo/business/model"
	"github.com/alvinarthas/myrepo/utils/common"
	"github.com/alvinarthas/myrepo/utils/logger"
	"github.com/alvinarthas/myrepo/utils/response"
	"github.com/go-chi/chi"
	"gopkg.in/go-playground/validator.v9"
)

var validate = validator.New()

// GetDummyList godoc
// @Summary      List accounts
// @Description  get accounts
// @Tags         accounts
// @Accept       json
// @Produce      json
// @Param        q    query     string  false  "name search by q"  Format(email)
// @Router       /dummies [get]

func GetDummyList(res http.ResponseWriter, req *http.Request) {
	data, err := dummy.GetDummyList()
	if err.Error != nil {
		response.Error(res, err)
		return
	}

	response.Success(res, http.StatusOK, data, nil)
}

// GetDummyID godoc
// @Summary      List accounts
// @Description  get accounts
// @Tags         accounts
// @Accept       json
// @Produce      json
// @Param        q    query     string  false  "name search by q"  Format(email)
// @Router       /dummies/{id} [get]
func GetDummy(res http.ResponseWriter, req *http.Request) {
	id := chi.URLParam(req, "id")
	if id == "" {
		message := "Failed on Get Dummy"
		err := fmt.Errorf(`%s%s`, "[GetDummy] ", errors.New("ID should not be empty"))
		logger.Error(message, err)
		response.Error(res, common.RecordError(err, http.StatusBadRequest, message))
		return
	}

	data, err := dummy.GetDummy(id)
	if err.Error != nil {
		response.Error(res, err)
		return
	}

	response.Success(res, http.StatusOK, data, nil)

}

func CreateDummy(res http.ResponseWriter, req *http.Request) {
	var payload model.CreateDummyRequest

	decoder := json.NewDecoder(req.Body)
	if err := decoder.Decode(&payload); err != nil {
		message := "Failed on Creating Dummy"
		err = fmt.Errorf(`%s%s`, "[CreateDummy] ", err.Error())
		logger.Error(message, err)
		response.Error(res, common.RecordError(err, http.StatusBadRequest, message))
		return
	}

	err := validate.Struct(payload)
	if err != nil {
		message := "Please check your data input"
		err = fmt.Errorf(`%s%s`, "[CreateDummy] ", err.Error())
		logger.Error(message, err)
		response.Error(res, common.RecordError(err, http.StatusBadRequest, message))
		return
	}

	if err := dummy.CreateDummy(payload); err.Error != nil {
		response.Error(res, err)
		return
	}

	response.Success(res, http.StatusCreated, []string{}, nil)
}

func UpdateDummy(res http.ResponseWriter, req *http.Request) {

	var payload model.UpdateDummyRequest

	id := chi.URLParam(req, "id")
	if id == "" {
		message := "Failed on Updating Dummy"
		err := fmt.Errorf(`%s%s`, "[GetDummy] ", errors.New("ID should not be empty"))
		logger.Error(message, err)
		response.Error(res, common.RecordError(err, http.StatusBadRequest, message))
		return
	}

	decoder := json.NewDecoder(req.Body)
	if err := decoder.Decode(&payload); err != nil {
		message := "Failed on Updating Dummy"
		err = fmt.Errorf(`%s%s`, "[UpdateDummy] ", err.Error())
		logger.Error(message, err)
		response.Error(res, common.RecordError(err, http.StatusBadRequest, message))
		return
	}

	payload.ID = id

	err := validate.Struct(payload)
	if err != nil {
		message := "Please check your data input"
		err = fmt.Errorf(`%s%s`, "[UpdateDummy] ", err.Error())
		logger.Error(message, err)
		response.Error(res, common.RecordError(err, http.StatusBadRequest, message))
		return
	}

	if err := dummy.UpdateDummy(payload); err.Error != nil {
		response.Error(res, err)
		return
	}

	response.Success(res, http.StatusOK, []string{}, nil)
}

func DeleteDummy(res http.ResponseWriter, req *http.Request) {
	id := chi.URLParam(req, "id")
	if id == "" {
		message := "Failed on Deleting Dummy"
		err := fmt.Errorf(`%s%s`, "[GetDummy] ", errors.New("ID should not be empty"))
		logger.Error(message, err)
		response.Error(res, common.RecordError(err, http.StatusBadRequest, message))
		return
	}

	err := dummy.DeleteDummy(id)
	if err.Error != nil {
		response.Error(res, err)
		return
	}

	response.Success(res, http.StatusOK, []string{}, nil)
}

func UpsertDummy(res http.ResponseWriter, req *http.Request) {

}

func CreateDummies(res http.ResponseWriter, req *http.Request) {

}

func UpdateDummies(res http.ResponseWriter, req *http.Request) {

}
