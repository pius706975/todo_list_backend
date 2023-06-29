package models

import (
	"database/sql"
	"net/http"
	"pius/databases"
	"time"
)

type Activity struct {
	ActivityID string    `json:"id" valid:"-"`
	Title      string    `json:"title" valid:"-"`
	Email      string    `json:"email" valid:"-"`
	CreatedAt  time.Time `json:"created_at" valid:"-"`
	UpdatedAt  time.Time `json:"updated_at" valid:"-"`
}

func AddActivity(title, email string) (Response, error) {
	
	var res Response

	conn := databases.CreateConn()

	sqlStatement := "INSERT INTO activities (title, email) VALUES (?, ?)"

	preparedStatement, err := conn.Prepare(sqlStatement)
	if err != nil {
		return res, err
	}

	result, err := preparedStatement.Exec(title, email)
	if err != nil {
		return res, err
	}

	ID, err := result.LastInsertId()
	if err != nil {
		return res, err
	}

	newData, err := GetByID(int(ID))
	if err != nil {
		return res, err
	}

	res.Status = http.StatusOK
	res.Message = "Activity has been created"
	res.Data = newData.Data

	return res, nil
}

func DeleteActivity(ID int) (Response, error) {

	var res Response

	conn := databases.CreateConn()

	getData, err := GetByID(ID)
	if err != nil {
		return res, err
	}

	sqlStatement := "DELETE FROM activities WHERE activity_id = ?"

	preparedStatement, err := conn.Prepare(sqlStatement)
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
	res.Message = "Activity has been deleted"
	res.Data = getData.Data

	return res, nil
}

func UpdateActivity(ID int, title, email string) (Response, error) {
	
	var res Response

	conn := databases.CreateConn()

	getData, err := GetByID(ID)
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
	data := getData.Data.(Activity)
	if title == "" {
		title = data.Title
	}
	if email == "" {
		email = data.Email
	}


	sqlStatement := "UPDATE activities SET title = ?, email = ? WHERE activity_id = ?"

	preparedStatement, err := conn.Prepare(sqlStatement)
	if err != nil {
		return res, err
	}

	_, err = preparedStatement.Exec(title, email, ID)
	if err != nil {
		return res, err
	}

	updatedActivity, err := GetByID(ID)
	if err != nil {
		return res, err
	}

	res.Status = http.StatusOK
	res.Message = "Activity is updated"
	res.Data = updatedActivity.Data

	return res, nil
}

func GetAllActivities() (Response, error) {
	
	var obj Activity
	var arrObj []Activity
	var res Response

	conn := databases.CreateConn()

	sqlStatement := "SELECT * FROM activities ORDER BY created_at DESC"

	rows, err := conn.Query(sqlStatement)
	defer rows.Close()
	if err != nil {
		return res, err
	}

	for rows.Next() {
		err := rows.Scan(&obj.ActivityID, &obj.Title, &obj.Email, &obj.CreatedAt, &obj.UpdatedAt)
		if err != nil {
			return res, err
		}

		arrObj = append(arrObj, obj)
	}

	res.Status = http.StatusOK
	res.Message = "Success"
	res.Data = arrObj

	return res, nil
}

func GetByID(ID int) (Response, error) {
	
	var obj Activity
	var res Response

	conn := databases.CreateConn()

	sqlStatement := "SELECT * FROM activities WHERE activity_id = ?"

	err := conn.QueryRow(sqlStatement, ID).Scan(&obj.ActivityID, &obj.Title, &obj.Email, &obj.CreatedAt, &obj.UpdatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			res.Status = http.StatusNotFound
			res.Message = "Data not found"
			res.Data = ""
			return res, nil
		}

		res.Status = http.StatusInternalServerError
		res.Message = "Error retrieving activity"
		return res, err
	}

	res.Status = http.StatusOK
	res.Message = "Success"
	res.Data = obj

	return res, nil
}