package main

import (
    "net/http"
    //"strconv"
    "go_api/models"
    "github.com/gin-gonic/gin"
)

//default Api Routes
func main() {
    router := gin.Default()
    router.GET("/api/users", getUsers)
    router.GET("/api/user/:code", getUser)
    router.POST("/api/create_user", createUser)

    router.Run("localhost:8080")
}

//Fetch all users
func getUsers(c *gin.Context){
    users :=models.GetUsers()

    if users == nil || len(users) == 0{
        c.AbortWithStatus(http.StatusNotFound)
    }else{
        c.IndentedJSON(http.StatusOK, users)
    }
}

//Fetch User via ID
func getUser(c *gin.Context){
    code :=c.Param("code")
    //code := uint
    //cdd := uint(code)
    user := models.GetUser(code)

    if user == nil {
        c.AbortWithStatus(http.StatusNotFound)
    }else{
        c.IndentedJSON(http.StatusOK, user)
    }
}


//Insert User
func createUser(c *gin.Context){
    var users models.Users

    if err := c.BindJSON(&users); err != nil{
        c.AbortWithStatus(http.StatusBadRequest)
    }else{
        models.CreateUser(users)
        c.IndentedJSON(http.StatusCreated, users)
    }
}
