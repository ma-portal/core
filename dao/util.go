package dao

import (
	log "github.com/Luncert/slog"
	uuid "github.com/satori/go.uuid"
)

func genUuid() string {
	id, _ := uuid.NewV4()
	return id.String()
}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
