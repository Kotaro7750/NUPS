package model

import (
	"database/sql"
	"fmt"
	"strconv"
)

//Kouji is a structure of kouji
type Kouji struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
}

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

//KoujiAll is a function to list up all kouji of ki
func KoujiAll(db *sql.DB, ki int) ([]*Kouji, error) {
	rows, err := db.Query(fmt.Sprintf("SELECT id,title FROM `%d`", ki))
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var koujiList []*Kouji

	for rows.Next() {
		var kouji Kouji
		if err := rows.Scan(&kouji.ID, &kouji.Title); err != nil {
			return nil, err
		}
		koujiList = append(koujiList, &kouji)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return koujiList, nil
}

//KiInsert is a function to add new ki
func KiInsert(db *sql.DB, ki int) error {
	insert, err := db.Prepare("INSERT INTO ki(ki) VALUES(?)")
	if err != nil {
		return err
	}

	new, err := db.Prepare(fmt.Sprintf("CREATE TABLE `%s`(id INT NOT NULL,title TEXT NOT NULL)", strconv.Itoa(ki)))
	if err != nil {
		return err
	}
	insert.Exec(ki)
	new.Exec()

	return nil
}
