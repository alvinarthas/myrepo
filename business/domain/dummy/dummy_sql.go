package dummy

import (
	"log"

	"github.com/alvinarthas/myrepo/business/model"
	"github.com/alvinarthas/myrepo/connection/mysql"
	"github.com/google/uuid"
)

var db = mysql.Connect()

func getDummyListSQL() ([]model.Dummy, error) {
	var dummies []model.Dummy

	rows, err := db.Query(GET_DUMMY_LIST_STATEMENT)
	if err != nil {
		return []model.Dummy{}, err
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
			return []model.Dummy{}, err
		}
		dummies = append(dummies, dummy)
	}
	if err := rows.Err(); err != nil {
		return []model.Dummy{}, err
	}

	return dummies, nil
}

func getDummySQL(id string) (model.Dummy, error) {

	var data model.Dummy
	if err := db.QueryRow(GET_DUMMY_BY_ID_STATEMENT, id).Scan(
		&data.ID,
		&data.Name,
		&data.Age,
		&data.Address,
		&data.IsActive,
		&data.Type,
	); err != nil {
		log.Println(err)
		return data, err
	}

	return data, nil
}

func createDummySQL(payload model.CreateDummyRequest) error {

	tx, err := db.Begin()
	if err != nil {
		log.Println(err)
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
		log.Println(err)
		return err
	}

	err = tx.Commit()
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func updateDummySQL(payload model.UpdateDummyRequest) error {

	tx, err := db.Begin()
	if err != nil {
		log.Println(err)
	}

	// Generate UUID
	_, err = tx.Exec(UPDATE_DUMMY_STATEMENT,
		payload.Name,
		payload.Address,
		payload.Age,
		payload.Type,
		payload.ID,
	)
	if err != nil {
		tx.Rollback()
		log.Println(err)
		return err
	}

	err = tx.Commit()
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}
