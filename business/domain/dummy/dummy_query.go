package dummy

const (
	GET_DUMMY_LIST_STATEMENT = `SELECT id, name, age, address FROM dummy`
	CREATE_DUMMY_STATEMENT   = `
		INSERT INTO dummy (
			id,
			name,
			address,
			age,
			type,
			is_active
		) VALUES (?, ?, ?, ?, ?, ?)
	`
	UPDATE_DUMMY_STATEMENT = `
		UPDATE dummy SET name = ?, address = ?, age = ?, type = ? WHERE id = ?
	`
)
