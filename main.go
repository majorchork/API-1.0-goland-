package main
import (
	"fmt"
	"github.com/gin-gonic/gin"
	"os"
)

type User struct {
	Name string `json:"name"`
	Age float32 `json:"age"`
	Email string `json:"email"`
	BloodType string `json:"blood_type"`
}

		//create an empty array of users
		var Users []User

func main() {
	// create a new gin router
	router := gin.Default()

	// define a single endpoint
	router.GET("/", helloWorldhandler)
	// crude endpoints for data

		//create
	_ = router.POST("/createUser", createUserhandler)
		//retrive
	_ = router.GET("/getUser", getsinglehandler)
	_ = router.GET("/getUsers", getmultipleUserhandler)
		//update
	_ = router.PATCH("/updateUser", updateUserhandler)
		//delete
	_ = router.DELETE("/deleteUser", deleteuserUserhandler)


	// run the server on the port 3000
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	_ = router.Run(":" + port)
}

func helloWorldhandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "hello world",
	})
}
func createUserhandler(c *gin.Context) {

		//create user
		user := User{}

			// gets the user data that was sent from the client
			// fills up our empty user object with the sent data
		err := c.ShouldBindJSON(&user)
		if err != nil {
			c.JSON(400, gin.H{
				"error": "invalid request data",
			 })
			return
		}
		// add single user to the list of users
		// save user somewhere ...
	Users = append(Users, user)

	c.JSON(200, gin.H{
		"message": "succesfully created user",
		"data": user,
	})
}

func getsinglehandler(c *gin.Context) {

	name := c.Param("name")

	fmt.Println("name", name)

	var user User

	for _, value := range Users {
		// check the current iteration of users
		// check if the name matches the client request
		if value.Name == name {
			//if it matches assigns the value to the empty user object we created
			user = value
		}
	}
	// if no match was found
	//the empty use we created would still be empty
	// check if user is empty, if so return a not found error

	if &user == nil {
		c.JSON(404, gin.H{
			"error": "no user with name found:" + name,
		})
		return
	}
	c.JSON(200, gin.H{
			"message" : "success",
				 "data" : user,
	})

}

func getmultipleUserhandler(c *gin.Context) {

	c.JSON(200, gin.H{
		"message": "successfully created user",
		"data": Users,
	})
}
func updateUserhandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "code successfully updated!",
	})
}
func deleteuserUserhandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "i don comot the fucking user",
	})
}

