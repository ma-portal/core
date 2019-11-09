package dao

import "time"

type Activation struct {
	ID         string
	RefID      string
	CreateTime int64
	RefType    PostingType
}

func CreateActivation(refID string, refType PostingType) *Activation {
	a := &Activation{
		ID:         genUuid(),
		RefID:      refID,
		CreateTime: time.Now().Unix(),
		RefType:    refType,
	}
	stmt, err := db.Prepare("insert into ACTIVATION(ID, REFID, CREATETIME, REFTYPE) values(?,?,?,?)")
	checkError(err)
	_, err = stmt.Exec(a.ID, a.RefID, a.CreateTime, a.RefType)
	checkError(err)
	return a
}

func DeleteActivation(id string) {
	stmt, err := db.Prepare("delete from ACTIVATION where ID=?")
	checkError(err)
	_, err = stmt.Exec(id)
	checkError(err)
}
