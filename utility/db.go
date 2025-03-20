package utility

import (
	"errors"
	"fmt"
	"github.com/gocql/gocql"
)

var Client *gocql.Session

func Startdb() (error) {
	
	cluster := gocql.NewCluster("127.0.0.1:9042")
    if cluster!=nil{
	//defining keyspaces
	cluster.Keyspace="example"
	cluster.Consistency=gocql.Quorum

	fmt.Println("server started")
	
	session, err := cluster.CreateSession()
 	
	if err != nil {
 		return err
 		}
		Client=session
	    return nil
	}
	return errors.New("not connected to db kindly restart")

}


