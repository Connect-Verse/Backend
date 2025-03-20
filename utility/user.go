package utility

import (
	"fmt"
	"errors"
)

type MetaUser struct{
	Id	  	    string
	AvatarId	string
    RoomId      string
	UserId      string
}



func CreateUser( metaUser MetaUser) error{

	if metaUser.Id=="" { 
		fmt.Println("kindly provide all the user values")
		return errors.New("provide data appropiately")
	}

	err:=Client.Query(`INSERT INTO metaUser (id,avatarId,roomId,userId) VALUES (?,?,?,?)`,metaUser.Id,metaUser.AvatarId,metaUser.RoomId,metaUser.UserId).Exec()
	if err!=nil{
		return errors.New("err occurrred while executing query")
	}    
    return nil
}


func RetreiveUser(id string) (*MetaUser,error){
    
	if id==""{
		return nil,errors.New("no relevant id is provided")	
	}
	var metaUser MetaUser
	err:=Client.Query(`SELECT * FROM metaUser WHERE id=?`,id).Scan(&metaUser.Id,&metaUser.AvatarId,&metaUser.RoomId,&metaUser.UserId)
	if err!=nil{
		return nil,err
	
	}
	return &metaUser,nil
}







