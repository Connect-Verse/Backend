package utility

import (
	"errors"
    "fmt"

)



type Room struct{
	Id		 	string      
 	Host	 	string      
	MetaUsers 	[]string 
} 


func CreateRoom(room Room) error{
  if room.Id==""{
	return errors.New("data is not provided appropiately")
}
  err := Client.Query(`INSERT INTO rooms (id,host,metaUserId) VALUES (?,?,?)`,room.Id,room.Host,room.MetaUsers).Exec()

  if err!=nil{
	fmt.Println(err,"error occurred")
	return err
  }
  return nil
}


func AddUser(metaUserId string,roomId string) error{
 if metaUserId==""{
	return errors.New("id is not provided")
 }
 err:=Client.Query(`UPDATE rooms SET metaUserId=metaUserId+? WHERE id=?`,[]string{metaUserId},roomId).Exec()

 if err!=nil{
	return err
 }

 return nil
}

func DeleteUser(metaUserId string,roomId string) error{
	if metaUserId!=""{
		return errors.New("id is not provided")
	 }
	 err:=Client.Query(`UPDATE rooms SET metaUsers=metaUsers-? WHERE id=?`,metaUserId,roomId).Exec()
	
	 if err!=nil{
		return err
	 }
	
	 return nil
}

func DeleteRoom(id string) error{
	if id==""{
		return errors.New("provide the requierd id")
	}
	err:=Client.Query(`DELETE * FROM rooms WHERE id=?`,id).Exec()
	if err!=nil{
		fmt.Println(err)
		return err
	}
return nil
}
