package model

import "database/sql"

//KiAll is a function to ki array
func KiAll(db *sql.DB) ([]int, error) {
	rows, err := db.Query("SELECT ki FROM ki")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var kiList []int

	for rows.Next() {
		var ki int
		if err := rows.Scan(&ki); err != nil {
			return nil, err
		}
		kiList = append(kiList, ki)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}
	return kiList, nil
}

//KiInsert is a function to add new ki
func KiInsert(db *sql.DB, ki int) error {
	insert, err := db.Prepare("INSERT INTO ki(ki) VALUES(?)")
	if err != nil {
		return err
	}
	insert.Exec(ki)

	return nil
}
