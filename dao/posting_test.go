package dao

import "testing"

func TestCreatePostingAndDeletePosting(t *testing.T) {
	p := CreatePosting("ASDASD", QuestionPosting, "a question")
	if p.ID == "" || p.CreateTime == 0 ||
		p.CreatorID != "ASDASD" ||
		p.Type != QuestionPosting ||
		p.Content != "a question" {
		t.Fail()
	}
	DeletePosting(p.ID)
}
