package model

import (
	"fmt"
	"myapp/datastore/postgres"
)

//Creating go object that has all data needed for our api
type Course struct{
	Cid int64 `json:"cid"`
	CourseName string `json:"coursename"`

}
//Sql query to create data for api
const Insertquery=`INSERT INTO course(Cid,CourseName)
VALUES($1,$2)`

//Fx to exucute the sql code through go

func (c *Course)Create()error{//return error to later know if it succeeded or failed
	fmt.Println("course model")
	//Exec is used to INSERT,UPDATE,But cannot Return Data
	_,err:=postgres.Db.Exec( //connecting database,we get sql.Result(what we insert news)which we do not have any use
		Insertquery,
		c.Cid,
		c.CourseName,
	
	)
	return err
}

//Sql Query to get Course

const queryGetCourse=`SELECT cid,coursename FROM course WHERE cid=$1;`//cid =$1 that $1 is a place holder that can be replaced later by c.Cid

func (c *Course)Read()error{
	err:=postgres.Db.QueryRow(//query row returns row matching the case
		queryGetCourse,
		c.Cid,//c.Cid replaces $1 by the value of it
	).Scan(&c.Cid,&c.CourseName)//Scan does reading in that row and fills your struct with sequential value
	return err
}

