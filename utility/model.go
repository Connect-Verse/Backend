package utility

import (
	"errors"
	"fmt"
)

func CreateTables() (bool,error) {

 err:= Client.Query("CREATE TABLE rooms (id text PRIMARY KEY, host text,metaUserId set<text>)").Exec()
 if err!=nil{
	fmt.Println("error occurred while executing query",err)
	return false,errors.New("query not executed")
 }
 err=Client.Query("CREATE TABLE metaUser (id text PRIMARY KEY ,avatarId text, userId text, roomId text)").Exec()
 if err!=nil{
	fmt.Println("error occurred while creating metausers")
	return false,errors.New("query not executed")
 }
 return true,nil 
}
