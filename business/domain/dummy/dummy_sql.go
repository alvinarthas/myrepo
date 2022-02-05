package dummy

import (
	"log"

	"github.com/alvinarthas/myrepo/business/model"
	"github.com/alvinarthas/myrepo/connection/mysql"
)

var db = mysql.Connect()

func getDummyListSQL() ([]model.Dummy, error) {
	var dummies []model.Dummy

	rows, err := db.Query("SELECT id, name, age, address FROM dummy")
	if err != nil {
		return []model.Dummy{}, err
		// return nil, fmt.Errorf("albumsByArtist %q: %v", name, err)
	}
	defer rows.Close()

	// Loop through rows, using Scan to assign column data to struct fields.
	for rows.Next() {
		var dummy model.Dummy
		if err := rows.Scan(
			&dummy.ID,
			&dummy.Name,
			&dummy.Age,
			&dummy.Address,
		); err != nil {
			log.Println(err)
			return []model.Dummy{}, err
			// return nil, fmt.Errorf("albumsByArtist %q: %v", name, err)
		}
		dummies = append(dummies, dummy)
	}
	if err := rows.Err(); err != nil {
		log.Println(err)
		return []model.Dummy{}, err
		// return nil, fmt.Errorf("albumsByArtist %q: %v", name, err)
	}

	return dummies, nil
}
