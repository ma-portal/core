package dao

import "testing"

func TestCreateCommitAndDeleteCommit(t *testing.T) {
	var (
		creatorID = "ASD"
		mainRefID = "XXXXX1"
		refID     = "XXXXX2"
		content   = "commit"
	)
	c := CreateCommit(creatorID, mainRefID, refID, content)
	if c.ID == "" ||
		c.CreateTime == 0 ||
		c.CreatorID != creatorID ||
		c.MainRefID != mainRefID ||
		c.RefID != refID ||
		c.Content != content {
		t.Fail()
	}
	DeleteCommit(c.ID)
}
