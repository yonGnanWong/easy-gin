package Models

import (
	"gin/Database"
	"gin/Models/db"
)

func GetTestData() db.Test  {
	var test db.Test
	Database.Db.ID(1).Get(&test)
	// SELECT * FROM user Where id = 1
	return test
}
