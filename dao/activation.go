package dao

import "time"

type Activation struct {
	ID         string
	RefID      string
	CreateTime int64
}

func CreateActivation(refID string) *Activation {
	a := &Activation{
		ID:         genUuid(),
		RefID:      refID,
		CreateTime: time.Now().Unix(),
	}
	stmt, err := db.Prepare("insert into ACTIVATION(ID, REFID, CREATETIME) values(?,?,?)")
	checkError(err)
	_, err = stmt.Exec(a.ID, a.RefID, a.CreateTime)
	checkError(err)
	return a
}

func DeleteActivation(id string) {
	stmt, err := db.Prepare("delete from ACTIVATION where ID=?")
	checkError(err)
	_, err = stmt.Exec(id)
	checkError(err)
}
