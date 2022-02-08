package dummy

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/alvinarthas/myrepo/business/model"
	"github.com/alvinarthas/myrepo/connection/mysql"
	"github.com/alvinarthas/myrepo/utils/common"
	"github.com/google/uuid"
)

var db = mysql.Connect()

func getDummyListSQL() ([]model.Dummy, common.Error) {
	var dummies []model.Dummy

	rows, err := db.Query(GET_DUMMY_LIST_STATEMENT)
	if err != nil {
		err = fmt.Errorf(`%s%s`, "[getDummyListSQL] ", err.Error())
		return dummies, common.RecordError(err, http.StatusInternalServerError, "Failed Get Dummy List")
	}
	defer rows.Close()

	for rows.Next() {
		var dummy model.Dummy
		if err := rows.Scan(
			&dummy.ID,
			&dummy.Name,
			&dummy.Age,
			&dummy.Address,
			&dummy.IsActive,
			&dummy.Type,
		); err != nil {
			err = fmt.Errorf(`%s%s`, "[getDummyListSQL] ", err.Error())
			return dummies, common.RecordError(err, http.StatusInternalServerError, "Failed Get Dummy List")
		}
		dummies = append(dummies, dummy)
	}
	if err := rows.Err(); err != nil {
		err = fmt.Errorf(`%s%s`, "[getDummyListSQL] ", err.Error())
		return dummies, common.RecordError(err, http.StatusInternalServerError, "Failed Get Dummy List")
	}

	return dummies, common.Error{}
}

func getDummySQL(id string) (model.Dummy, common.Error) {

	var data model.Dummy
	if err := db.QueryRow(GET_DUMMY_BY_ID_STATEMENT, id).Scan(
		&data.ID,
		&data.Name,
		&data.Age,
		&data.Address,
		&data.IsActive,
		&data.Type,
	); err != nil {
		if err == sql.ErrNoRows {
			err = fmt.Errorf(`%s%s`, "[getDummySQL] ", err.Error())
			return data, common.RecordError(err, http.StatusNotFound, "Dummy not Found")
		}
		err = fmt.Errorf(`%s%s`, "[getDummySQL] ", err.Error())
		return data, common.RecordError(err, http.StatusInternalServerError, "Failed Get Dummy")
	}

	return data, common.Error{}
}

func createDummySQL(payload model.CreateDummyRequest) common.Error {

	tx, err := db.Begin()
	if err != nil {
		err = fmt.Errorf(`%s%s`, "[createDummySQL] ", err.Error())
		return common.RecordError(err, http.StatusInternalServerError, "Failed Create Dummy")
	}

	// Generate UUID
	dummyID := uuid.Must(uuid.NewRandom())
	_, err = tx.Exec(CREATE_DUMMY_STATEMENT,
		dummyID,
		payload.Name,
		payload.Address,
		payload.Age,
		payload.Type,
		true,
	)
	if err != nil {
		tx.Rollback()
		err = fmt.Errorf(`%s%s`, "[createDummySQL] ", err.Error())
		return common.RecordError(err, http.StatusInternalServerError, "Failed Create Dummy")
	}

	err = tx.Commit()
	if err != nil {
		err = fmt.Errorf(`%s%s`, "[createDummySQL] ", err.Error())
		return common.RecordError(err, http.StatusInternalServerError, "Failed Create Dummy")
	}

	return common.Error{}
}

func updateDummySQL(payload model.UpdateDummyRequest) common.Error {

	tx, err := db.Begin()
	if err != nil {
		err = fmt.Errorf(`%s%s`, "[updateDummySQL] ", err.Error())
		return common.RecordError(err, http.StatusInternalServerError, "Failed Update Dummy")
	}

	_, err = tx.Exec(UPDATE_DUMMY_STATEMENT,
		payload.Name,
		payload.Address,
		payload.Age,
		payload.Type,
		payload.ID,
	)
	if err != nil {
		tx.Rollback()
		err = fmt.Errorf(`%s%s`, "[updateDummySQL] ", err.Error())
		return common.RecordError(err, http.StatusInternalServerError, "Failed Update Dummy")
	}

	err = tx.Commit()
	if err != nil {
		err = fmt.Errorf(`%s%s`, "[updateDummySQL] ", err.Error())
		return common.RecordError(err, http.StatusInternalServerError, "Failed Update Dummy")
	}

	return common.Error{}
}

func deleteDummySQL(id string) common.Error {

	tx, err := db.Begin()
	if err != nil {
		err = fmt.Errorf(`%s%s`, "[deleteDummySQL] ", err.Error())
		return common.RecordError(err, http.StatusInternalServerError, "Failed Update Dummy")
	}

	// Generate UUID
	_, err = tx.Exec(DELETE_DUMMY_STATEMENT, id)
	if err != nil {
		err = fmt.Errorf(`%s%s`, "[deleteDummySQL] ", err.Error())
		return common.RecordError(err, http.StatusInternalServerError, "Failed Update Dummy")
	}

	err = tx.Commit()
	if err != nil {
		err = fmt.Errorf(`%s%s`, "[deleteDummySQL] ", err.Error())
		return common.RecordError(err, http.StatusInternalServerError, "Failed Update Dummy")
	}

	return common.Error{}
}
