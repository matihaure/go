package model

import "github.com/matihaure/views"

func DBTest() ([]views.PostRequest, error) {

	rows, err := con.Query("SHOW DATABASES;")

	if err != nil {
		return nil, err
	}

	todos := []views.PostRequest{}

	for rows.Next() {

		data := views.PostRequest{}

		rows.Scan(&data.Name, &data.Todo)

		todos = append(todos, data)
	}

	return todos, nil

}
