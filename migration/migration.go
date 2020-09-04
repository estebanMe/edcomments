package migration

import (
	"github.com/estebanMe/edcomments/configuration"
	"github.com/estebanMe/edcomments/models"
)

func Migrate() {
	db:= configuration.GetConnection()
	defer db.Close()
	
	db.CreateTable(&models.User{})
	db.CreateTable(&models.Comment{})
	db.CreateTable(&models.Vote{})

	db.Model(&models.Vote{}).AddUniqueIndex("Comment_id_user_id_unique", "comment_id", "user_id")

}

