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
// @Summary      List of Dummy
// @Description  get list of dummy data
// @Tags         Dummies
// @Accept       json
// @Produce      json
// @Param X-Api-Key header string true "App Secret Key"
// @Success      200  {array}   model.GetDummyListResponse
// @Failure      400  {object}  response.ErrorResponse
// @Failure      404  {object}  response.ErrorResponse
// @Failure      500  {object}  response.ErrorResponse
// @Router       /dummies [get]
func GetDummyList(res http.ResponseWriter, req *http.Request) {
	dummies, err := dummy.GetDummyList()
	if err.Error != nil {
		response.Error(res, err)
		return
	}

	data := model.GetDummyListResponse{
		Data: dummies,
	}
	response.Success(res, http.StatusOK, data)
}

// GetDummy godoc
// @Summary      Dummy Single Data by ID
// @Description  get single dummy data by ID
// @Tags         Dummies
// @Accept       json
// @Produce      json
// @Param X-Api-Key header string true "App Secret Key"
// @Param id path string true "Dummy ID"
// @Success      200  {object}  model.GetDummyByIDResponse
// @Failure      400  {object}  response.ErrorResponse
// @Failure      404  {object}  response.ErrorResponse
// @Failure      500  {object}  response.ErrorResponse
// @Router       /dummies/{id} [get]
func GetDummyByID(res http.ResponseWriter, req *http.Request) {
	id := chi.URLParam(req, "id")
	if id == "" {
		message := "Failed on Get Dummy"
		err := fmt.Errorf(`%s%s`, "[GetDummy] ", errors.New("ID should not be empty"))
		logger.Error(message, err)
		response.Error(res, common.RecordError(err, http.StatusBadRequest, message))
		return
	}

	singleDummy, err := dummy.GetDummyByID(id)
	if err.Error != nil {
		response.Error(res, err)
		return
	}

	data := model.GetDummyByIDResponse{
		Data: singleDummy,
	}

	response.Success(res, http.StatusOK, data)

}

// CreateDummy godoc
// @Summary      Dummy New Data
// @Description  create dummy data
// @Tags         Dummies
// @Accept       json
// @Produce      json
// @Param X-Api-Key header string true "App Secret Key"
// @Param data body model.CreateDummyRequest true "Create New Dummy Data"
// @Success      200  {object}  model.CreateDummyResponse
// @Failure      400  {object}  response.ErrorResponse
// @Failure      404  {object}  response.ErrorResponse
// @Failure      500  {object}  response.ErrorResponse
// @Router       /dummies [post]
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

	response.Success(res, http.StatusCreated, model.CreateDummyResponse{Data: []string{}})
}

// UpdateDummy godoc
// @Summary      Dummy Update Data
// @Description  update dummy data - if attribute is null, it will updated as null
// @Tags         Dummies
// @Accept       json
// @Produce      json
// @Param X-Api-Key header string true "App Secret Key"
// @Param id path string true "Dummy ID"
// @Param data body model.UpdateDummyRequest true "Update Dummy Data"
// @Success      200  {object}  model.UpdateDummyResponse
// @Failure      400  {object}  response.ErrorResponse
// @Failure      404  {object}  response.ErrorResponse
// @Failure      500  {object}  response.ErrorResponse
// @Router       /dummies/{id} [put]
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

	response.Success(res, http.StatusOK, model.UpdateDummyResponse{Data: []string{}})
}

// DeleteDummy godoc
// @Summary      Dummy Delete Data
// @Description  delete dummy data
// @Tags         Dummies
// @Accept       json
// @Produce      json
// @Param X-Api-Key header string true "App Secret Key"
// @Param id path string true "Dummy ID"
// @Success      200  {object}  model.DeleteDummyResponse
// @Failure      400  {object}  response.ErrorResponse
// @Failure      404  {object}  response.ErrorResponse
// @Failure      500  {object}  response.ErrorResponse
// @Router       /dummies/{id} [delete]
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

	response.Success(res, http.StatusOK, model.DeleteDummyResponse{Data: []string{}})
}

func UpsertDummy(res http.ResponseWriter, req *http.Request) {

}

func CreateDummies(res http.ResponseWriter, req *http.Request) {

}

func UpdateDummies(res http.ResponseWriter, req *http.Request) {

}
