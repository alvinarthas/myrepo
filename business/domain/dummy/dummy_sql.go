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
	)
	if err != nil {
		tx.Rollback()
		log.Println(err)
	}

	err = tx.Commit()
	if err != nil {
		log.Println(err)
	}

	return nil
}
