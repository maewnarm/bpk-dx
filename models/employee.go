package models

type Employee struct {
	Id           string `db:"id"`
	Surname      string `db:"surname"`
	Lastname     string `db:"lastname"`
	Department   string `db:"department"`
	Section_name string `db:"section_name"`
	Section_code string `db:"section_code"`
}

func GetAllEmployees() ([]*Employee, error) {
	db, err := connect()
	if err != nil {
		return nil, err
	}

	var result []*Employee
	err = db.Select(&result, `SELECT * FROM employees ORDER BY id ASC`)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func CreateEmployee(e *Employee) (*Employee, error) {
	db, err := connect()
	if err != nil {
		return nil, err
	}

	err = db.Get(e, `
		INSERT INTO employees (id,surname,lastname,department,section_name,section_code)
		VALUES ($1,$2,$3,$4,$5,$6)
		RETURNING *
	`, e.Id, e.Surname, e.Lastname, e.Department, e.Section_name, e.Section_code)
	if err != nil {
		return nil, err
	}
	return e, nil
}

func DeleteEmployee(list string) ([]*Employee, error) {
	db, err := connect()
	if err != nil {
		return nil, err
	}

	var deleted []*Employee
	err = db.Select(&deleted, `
		DELETE FROM employees WHERE id = ANY(string_to_array($1,'|'))
		RETURNING *
	`, list)
	if err != nil {
		return nil, err
	}
	return deleted, nil
}
