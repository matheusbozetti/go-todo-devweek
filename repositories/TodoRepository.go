package repositories

import (
	"devweek/types"

	"github.com/rs/xid"
)

var db = []types.TodoType{}

func GetAll() []types.TodoType {
	return db
}

func Create(todo *types.TodoType) error {
	todo.ID = xid.New().String()

	db = append(db, *todo)

	return nil
}
