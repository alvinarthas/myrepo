package dummy

const (
	GET_DUMMY_LIST_STATEMENT  = `SELECT id, name, age, address, is_active, type FROM dummy`
	GET_DUMMY_BY_ID_STATEMENT = `SELECT id, name, age, address, is_active, type FROM dummy WHERE id = ?`
	CREATE_DUMMY_STATEMENT    = `
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
	DELETE_DUMMY_STATEMENT = `DELETE FROM dummy WHERE id = ?`
)
