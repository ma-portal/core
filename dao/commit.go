package dao

import "time"

type Commit struct {
	ID         string
	CreatorID  string
	CreateTime int64
	MainRefID  string
	RefID      string
	Content    string
}

func CreateCommit(creatorID string, mainRefID string, refID string, content string) *Commit {
	c := &Commit{
		ID:         genUuid(),
		CreatorID:  creatorID,
		CreateTime: time.Now().Unix(),
		MainRefID:  mainRefID,
		RefID:      refID,
		Content:    content,
	}
	stmt, err := db.Prepare("insert into COMMIT(ID, CREATORID, CREATETIME, MAINREFID, REFID, CONTENT)" +
		" values(?,?,?,?,?,?)")
	checkError(err)
	_, err = stmt.Exec(c.ID, c.CreatorID, c.CreateTime, c.MainRefID, c.RefID, c.Content)
	checkError(err)
	return c
}

func DeleteCommit(id string) {
	stmt, err := db.Prepare("delete from COMMIT where ID=?")
	checkError(err)
	_, err = stmt.Exec(id)
	checkError(err)
}
