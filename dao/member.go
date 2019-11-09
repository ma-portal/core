package dao

const (
	AdminRole = iota
	NormalRole
)

type MemberRole int

type Member struct {
	ID       string
	Account  string
	Password string
	Role     MemberRole
}

func CreateMember(account, password string, role MemberRole) *Member {
	m := &Member{
		ID:       genUuid(),
		Account:  account,
		Password: password,
		Role:     role,
	}
	stmt, err := db.Prepare("insert MEMBER(ID, ACCOUNT, PASSWORD, ROLE) values(?,?,?,?)")
	checkError(err)
	_, err = stmt.Exec(m.ID, m.Account, m.Password, m.Role)
	checkError(err)
	return m
}

func DeleteMember(id string) {
	stmt, err := db.Prepare("delete from MEMBER where ID=?")
	checkError(err)
	_, err = stmt.Exec(id)
	checkError(err)
}
