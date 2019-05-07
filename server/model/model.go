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
