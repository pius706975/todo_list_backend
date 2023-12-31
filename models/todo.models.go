package models

import (
	"database/sql"
	"net/http"
	"pius/databases"
	"strconv"
	"time"
)

type Todo struct {
	TodoID          int       `json:"id" valid:"-"`
	ActivityGroupID int       `json:"activity_group_id" valid:"-"`
	Title           string    `json:"title" valid:"-"`
	Priority        string    `json:"priority" valid:"-"`
	IsActive        bool      `json:"is_active" valid:"-"`
	CreatedAt       time.Time `json:"created_at" valid:"-"`
	UpdatedAt       time.Time `json:"updated_at" valid:"-"`
}

func AddTodoItem(activityGroupID int, title, priority string) (Response, error) {

	var res Response

	db := databases.CreateConn()

	sqlStatement := "INSERT INTO todos (activity_group_id, title, priority) VALUES (?, ?, ?)"

	preparedStatement, err := db.Prepare(sqlStatement)
	if err != nil {
		return res, err
	}

	result, err := preparedStatement.Exec(activityGroupID, title, priority)
	if err != nil {
		return res, err
	}

	ID, err := result.LastInsertId()
	if err != nil {
		return res, err
	}

	newData, err := GetTodoByID(int(ID))
	if err != nil {
		return res, err
	}

	res.Status = http.StatusOK
	res.Message = "To Do item is added"
	res.Data = newData.Data

	return res, nil

}

func DeleteTodoItem(ID int) (Response, error) {

	var res Response

	db := databases.CreateConn()

	getData, err := GetTodoByID(ID)
	if err != nil {
		return res, err
	}

	sqlStatement := "DELETE FROM todos WHERE todo_id = ?"

	preparedStatement, err := db.Prepare(sqlStatement)
	if err != nil {
		return res, err
	}

	_, err = preparedStatement.Exec(ID)
	if err != nil {
		return res, err
	}

	if getData.Status == http.StatusNotFound {
		res.Status = http.StatusNotFound
		res.Message = "Data not found"
		res.Data = ""

		return res, nil
	}

	res.Status = http.StatusOK
	res.Message = "To Do item has been deleted"
	res.Data = getData.Data

	return res, nil
}

func UpdateTodoItem(ID, activitiGroupID int, title, priority string, isActive string) (Response, error) {

	var res Response

	db := databases.CreateConn()

	getData, err := GetTodoByID(ID)
	if err != nil {
		return res, err
	}

	if getData.Status == http.StatusNotFound {
		res.Status = http.StatusNotFound
		res.Message = "Data not found"
		res.Data = ""

		return res, nil
	}

	// getting old data
	data := getData.Data.(Todo)
	if activitiGroupID == 0 {
		activitiGroupID = data.ActivityGroupID
	}
	if title == "" {
		title = data.Title
	}
	if priority == "" {
		priority = data.Priority
	}
	if isActive == "" {
		isActive = strconv.FormatBool(data.IsActive)
	}

	isActiveBool, err := strconv.ParseBool(isActive)
	if err != nil {
		return res, err
	}

	sqlStatement := "UPDATE todos SET activity_group_id = ?, title = ?, priority = ?, is_active = ? WHERE todo_id = ?"

	preparedStatement, err := db.Prepare(sqlStatement)
	if err != nil {
		return res, err
	}

	_, err = preparedStatement.Exec(activitiGroupID, title, priority, isActiveBool, ID)
	if err != nil {
		return res, err
	}

	updatedItem, err := GetTodoByID(ID)
	if err != nil {
		return res, err
	}

	res.Status = http.StatusOK
	res.Message = "To Do item is updated"
	res.Data = updatedItem.Data

	return res, nil
}

func GetAllTodoItems() (Response, error) {
	
	var obj Todo
	var arrObj []Todo
	var res Response

	db := databases.CreateConn()

	sqlStatement := "SELECT * FROM todos"

	rows, err := db.Query(sqlStatement)
	defer rows.Close()
	if err != nil {
		return res, err
	}

	for rows.Next() {
		err := rows.Scan(&obj.TodoID, &obj.ActivityGroupID, &obj.Title, &obj.Priority, &obj.IsActive, &obj.CreatedAt, &obj.UpdatedAt)
		if err != nil {
			return res, nil
		}

		arrObj = append(arrObj, obj)
	}

	res.Status = http.StatusOK
	res.Message = "Success"
	res.Data = arrObj

	return res, nil
}

func GetTodoItemsByGroup(activityGroupID int) (Response, error) {

	var todos []Todo
	var res Response

	db := databases.CreateConn()

	rows, err := db.Query("SELECT * FROM todos WHERE activity_group_id = ?", activityGroupID)
	if err != nil {
		return res, err
	}
	defer rows.Close()

	for rows.Next() {
		var todo Todo
		
		err := rows.Scan(&todo.TodoID, &todo.ActivityGroupID, &todo.Title, &todo.Priority, &todo.IsActive, &todo.CreatedAt, &todo.UpdatedAt)
		if err != nil {
			return res, err
		}
		todos = append(todos, todo)
	}

	err = rows.Err()
	if err != nil {
		return res, err
	}

	if len(todos) == 0 {
		res.Status = http.StatusNotFound
		res.Message = "Data not found"
		res.Data = ""

		return res, nil
	}

	res.Status = http.StatusOK
	res.Message = "Success"
	res.Data = todos

	return res, nil
} 

func GetTodoByID(ID int) (Response, error) {

	var obj Todo
	var res Response

	db := databases.CreateConn()

	sqlStatement := "SELECT * FROM todos WHERE todo_id = ?"

	err := db.QueryRow(sqlStatement, ID).Scan(&obj.TodoID, &obj.ActivityGroupID, &obj.Title, &obj.Priority, &obj.IsActive, &obj.CreatedAt, &obj.UpdatedAt)
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
