package dao

import "time"

const (
	ArticlePosting = iota
	QuestionPosting
	NotificationPosting
)

type PostingType int

type Posting struct {
	ID         string
	CreatorID  string
	CreateTime int64
	Type       PostingType
	Content    string
}

func CreatePosting(creatorID string, pType PostingType, content string) *Posting {
	p := &Posting{
		ID:         genUuid(),
		CreatorID:  creatorID,
		CreateTime: time.Now().Unix(),
		Type:       pType,
		Content:    content,
	}
	stmt, err := db.Prepare("insert POSTING(ID,CREATORID,CREATETIME,TYPE,CONTENT) values(?,?,?,?,?)")
	checkError(err)
	_, err = stmt.Exec(p.ID, p.CreatorID, p.CreateTime, p.Type, p.Content)
	checkError(err)
	return p
}

func DeletePosting(id string) {
	stmt, err := db.Prepare("delete from POSTING where ID=?")
	checkError(err)
	_, err = stmt.Exec(id)
	checkError(err)
}
