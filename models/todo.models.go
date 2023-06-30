package models

import (
	"database/sql"
	"net/http"
	"pius/databases"
	"time"
)

type Todo struct {
	TodoID          int       `json:"id" valid:"-"`
	ActivityGroupID int       `json:"activity_group_id" valid:"-"`
	Title           string    `json:"title" valid:"-"`
	Priority        string    `json:"priority" valid:"-"`
	Is_Active       bool      `json:"is_active" valid:"-"`
	CreatedAt       time.Time `json:"created_at" valid:"-"`
	UpdatedAt       time.Time `json:"updated_at" valid:"-"`
}

func GetTodoByID(ID int) (Response, error) {
	
	var obj Todo
	var res Response
	
	db := databases.CreateConn()

	sqlStatement := "SELECT * FROM todos WHERE todo_id = ?"

	err := db.QueryRow(sqlStatement, ID).Scan(&obj.ActivityGroupID, &obj.ActivityGroupID, &obj.Title, &obj.Priority, &obj.Is_Active, &obj.CreatedAt, &obj.UpdatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			res.Status = http.StatusNotFound
			res.Message = "Data not found"
			res.Data = ""

			return res, nil
		}

		res.Status = http.StatusInternalServerError
		res.Message = "Error retrieving item"
		return res, err
	}

	res.Status = http.StatusOK
	res.Message = "Success"
	res.Data = obj

	return res, nil
}