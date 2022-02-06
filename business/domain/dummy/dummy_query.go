package dummy

const (
	GET_DUMMY_LIST_STATEMENT = `SELECT id, name, age, address FROM dummy`
	CREATE_DUMMY_STATEMENT   = `
		INSERT INTO dummy (
			id,
			name,
			address,
			age,
			type
		) VALUES (?, ?, ?, ?, ?)
	`
)
