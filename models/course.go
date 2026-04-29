package model

import (
	"fmt"
	"myapp/datastore/postgres"
)

// Creating go object that has all data needed for our api
type Course struct {
	Cid        int64  `json:"cid"`
	CourseName string `json:"coursename"`
} 

// Sql query to create data for api
const Insertquery = `INSERT INTO course(Cid,CourseName)
VALUES($1,$2)`

//Fx to exucute the sql code through go

func (c *Course) Create() error { //return error to later know if it succeeded or failed
	fmt.Println("course model")
	//Exec is used to INSERT,UPDATE,But cannot Return Data
	_, err := postgres.Db.Exec( //connecting database,we get sql.Result(what we insert news)which we do not have any use
		Insertquery,
		c.Cid,
		c.CourseName,
	)
	return err
}

//Sql Query to get Course

const queryGetCourse = `SELECT cid,coursename FROM course WHERE cid=$1;` //cid =$1 that $1 is a place holder that can be replaced later by c.Cid

func (c *Course) Read() error {
	err := postgres.Db.QueryRow( //query row returns row matching the case
		queryGetCourse,
		c.Cid, //c.Cid replaces $1 by the value of it
	).Scan(&c.Cid, &c.CourseName) //Scan does reading in that row and fills your struct with sequential value
	return err
}

//Sql query to Update course

const UpdateCourse = `UPDATE course SET cid=$1,coursename=$2 WHERE cid=$3`

func (c *Course) UpdateCourses() error {
	_, err := postgres.Db.Exec(
		UpdateCourse,
		c.Cid,
		c.CourseName,
	)
	return err
}

const Deletequery = `DELETE FROM course WHERE cid=$1 RETURNING cid ` //Returning cid confirms the deletion

func (c *Course) DeleteCourse() error {
	if err := postgres.Db.QueryRow(
		Deletequery, //perform Deletequery on that row
		c.Cid,       //we overwrite placeholder c.Cid=$1 with new c.Cid that we call later in controller
		//after that cid is returned and used in scan
	).Scan(&c.Cid); err != nil { //we delete the row and read from database and write cid to struct c..if Cid doesn't match we get error
		return err
	}
	return nil
}

// we create function because it doesn't depend upon any struct
const getallquery = `SELECT * FROM course`

func GetCoursesAll() ([]Course, error) {
	rows, err := postgres.Db.Query(getallquery) //we use Query to select all rows and it returns rows in the form set of slice and error
	if err != nil {
		return nil, err //note if error occured than return nil for course struct and error as per your function requirement
	}
	var courses []Course

	//rows is where u have data got from sql and we do Next to move to each rows

	for rows.Next() {
		var c Course

		dbErr := rows.Scan(&c.Cid, &c.CourseName) //we read through each rows and write rows data to struct c
		if dbErr != nil {
			return nil, dbErr // if we get errror during scaning return nil slice and that error
		}
		courses = append(courses, c) //append c to struct courses

	}
	rows.Close() //once finish reading from rows close it //note rows like a slice of struct during looping

	return courses, nil // error handling is done above if any error occured it will take care of it but if no error then its return nil
}
