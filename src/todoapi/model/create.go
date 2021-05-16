package model

func CreateTodo(name string, todo string) error {

	insertQR, err := con.Query("INSERT INTO TODO VALUES(?,?)", name, todo)

	defer insertQR.Close()

	if err != nil {

		return err
	}

	return nil

}

func DeleteById(name string) error {

	deleteQR, err := con.Query("DELETE FROM TODO WHERE name=?", name)

	defer deleteQR.Close()

	if err != nil {

		return err
	}

	return nil

}
