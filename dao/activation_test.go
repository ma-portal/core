package dao

import "testing"

func TestCreateActivationAndDeleteActivation(t *testing.T) {
	a := CreateActivation("ASDASD")
	if a.ID == "" ||
		a.RefID != "ASDASD" ||
		a.CreateTime == 0 {
		t.Fail()
	}
	DeleteActivation(a.ID)
}
