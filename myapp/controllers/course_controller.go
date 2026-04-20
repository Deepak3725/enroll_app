package controller

import (
	"database/sql"
	"encoding/json"
	"fmt"
	model "myapp/models"
	"myapp/utils/httpResp"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func AddCourse(w http.ResponseWriter, r *http.Request) {
	var course model.Course
	//data sent by client is in json form but go deals with struct so, we need to read json data and convert to struct
	decoder := json.NewDecoder(r.Body)              //Creates reading environment tells where to read in(request body)
	if err := decoder.Decode(&course); err != nil { //Decode does reading and sent data to course orginal not copy
		fmt.Println(err)
		return
	}
	r.Body.Close() //Once done reading no need anymore so,close it to free memory

	err := course.Create() //we calling Create method for course object that returns error

	if err != nil {
		httpResp.RespondWithError(w, http.StatusBadRequest, err.Error()) //httpResp is a package that makes our work easy if not we have to explicitly do writtin of what to do in browser
		//err: is in the form of object .Error() is used to conver object err to browser in readable string
		return
	}

}

func GetCourse(w http.ResponseWriter, r *http.Request) {
	sid := mux.Vars(r)["sid"] //mux.Vars(r) reads all parameters from URL and Get sid http://localhost:8080/course/123 this 123 is sid its in string sid="123"
	cid, err := getCourseId(sid)

	if err != nil {
		httpResp.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	c := model.Course{Cid: cid} //create empty course object with Cid $1 replaced by cid

	getErr := c.Read() //This error has 2 types No rows and other error we have to handle both ways

	if getErr != nil {
	switch {
	case getErr == sql.ErrNoRows:
		httpResp.RespondWithError(w, http.StatusNotFound, getErr.Error()) // no rows found

	default:
		httpResp.RespondWithError(w, http.StatusInternalServerError, getErr.Error()) // real server error
	}
	return
}

httpResp.RespondWithJSON(w, http.StatusOK, c)
}
// This fx converts sid to int and gives error if occured
func getCourseId(courseId string) (int64, error) {
	Cid, err := strconv.ParseInt(courseId, 10, 64) //strconv.ParseInt is used to covert string to int
	if err != nil {
		return 0, err //if we get error then Cid is 0 by default
	}
	return Cid, err
}
