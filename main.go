package main
import (
	"github.com/gin-gonic/gin"
	"os"
)

type User struct {
	Name string `json:"name"`
	Age float64 `json:"age"`
	Email string `json:"email"`
	BloodType string `json:"blood_type"`
}

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

		err := c.ShouldBindJSON(&user)
		if err != nil {
			c.JSON(400, gin.H{
				"error": "invalid request data",
			 })
			return
		}
		// save user somewhere ...

	c.JSON(200, gin.H{
		"message": "succesfully created user",
		"data": user,
	})
}
func getsinglehandler(c *gin.Context) {
	user := User{
		Name: "Sochuks",
		Age: 04.20,
		Email: "we@email.com",
	}
	c.JSON(200, gin.H{
		"message": "succesfully created user",
		"data": user,
	})
}

func getmultipleUserhandler(c *gin.Context) {
	user := User{
		Name: "Sochuks",
		Age: 04.20,
		Email: "we@email.com",
	}
	c.JSON(200, gin.H{
		"message": "succesfully created user",
		"data": user,
	})
}
func updateUserhandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "code succesfully updated!",
	})
}
func deleteuserUserhandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "i don comot the fucking user",
	})
}

