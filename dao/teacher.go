package dao

type Teacher struct {
	ID       string
	Account  string
	Password string
}

func CreateTeacher(account, password string) *Teacher {
	t := &Teacher{
		ID:       genUuid(),
		Account:  account,
		Password: password,
	}
	stmt, err := db.Prepare("insert TEACHER(ID, ACCOUNT, PASSWORD) values(?,?,?)")
	checkError(err)
	_, err = stmt.Exec(t.ID, t.Account, t.Password)
	checkError(err)
	return t
}

func DeleteTeacher(id string) {
	stmt, err := db.Prepare("delete from TEACHER where ID=?")
	checkError(err)
	_, err = stmt.Exec(id)
	checkError(err)
}
