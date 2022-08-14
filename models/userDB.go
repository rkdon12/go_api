// init DB for Users
package models

 import (
 	"database/sql"
 	_ "github.com/go-sql-driver/mysql"
 	"fmt"
 )

 func GetUsers() []Users {
 	db, err := sql.Open("mysql", "root:mag123ato@tcp(127.0.0.1:3366)/api")

 	if err != nil{
 		fmt.Println("Err", err.Error())
 		return nil
 	}

    // defer the close till after this function has finished
    // executing
 	defer db.Close()

 	results, err := db.Query("SELECT * FROM users")

 	if err != nil{
 		fmt.Println("Err", err.Error())
 		return nil
 	}

 	users := []Users{}
 	for results.Next(){
 		var usr Users

 		err = results.Scan(&usr.Code, &usr.Username, &usr.Password, &usr.Email, &usr.Role, &usr.LastUpdated)

 		if err != nil {
 			panic(err.Error())
 		}

 		users = append(users, usr)
 	}

 	return users
 }


//
 func GetUser(code string) *Users {

 	db, err := sql.Open("mysql", "root:mag123ato@tcp(127.0.0.1:3366)/api")
    usr := &Users{}
    if err != nil{
        fmt.Println("Err", err.Error())
        return nil
    }

    // defer the close till after this function has finished
    // executing
    defer db.Close()

    results, err := db.Query("SELECT * FROM users where code=?", code)

    if err != nil{
        fmt.Println("Err", err.Error())
        return nil
    }

    if results.Next(){
        err = results.Scan(&usr.Code, &usr.Username, &usr.Password, &usr.Email, &usr.Role, &usr.LastUpdated)
        if err != nil{
            return nil
        }
    }else {
        return nil
    }

    return usr

 }


///insert user details
func CreateUser(users Users){
    db, err := sql.Open("mysql", "root:mag123ato@tcp(127.0.0.1:3366)/api")

    if err != nil{
        panic(err.Error())
    }

    // defer the close till after this function has finished
    // executing
    defer db.Close()

    insert, err :=db.Query("INSERT INTO users(username,password,email,role,last_updated) VALUES(?,?,?,?, now())", users.Username, users.Password, users.Email, users.Role)

    //if there is an error inserting, handle it
    if err != nil{
        panic(err.Error())
    }

    defer insert.Close()

}