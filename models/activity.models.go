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
	res.Message = "Ok"
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
	res.Message = "Ok"
	res.Data = obj

	return res, nil
}

func AddActivity(title, email string) (Response, error) {
	
	var res Response

	conn := databases.CreateConn()

	sqlStatement := "INSERT INTO activities (title, email) VALUES (?, ?)"

	prepareStatement, err := conn.Prepare(sqlStatement)
	if err != nil {
		return res, err
	}

	result, err := prepareStatement.Exec(title, email)
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
	res.Message = "Ok"
	res.Data = newData.Data

	return res, nil
}

// func DeleteActivity()  {

// }