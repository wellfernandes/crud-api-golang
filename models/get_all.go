package models

import "crud-api-golang/db"

func GetAll() (todos []Todo, err error) {

	conn, err := db.OpenConnection()
	if err != nil {
		return
	}
	defer conn.Close()

	rows, err := conn.Query(`SELECT * FROM tb_todos`)
	if err != nil {
		return
	}

	for rows.Next() {
		var todo Todo
		err = rows.Scan(&todo.ID, &todo.Title, &todo.Description, &todo.Done)
		if err != nil {
			continue //if there is an error, it will ignore and go to the next one *just for learning
		}
		todos = append(todos, todo)
	}
	return
}
