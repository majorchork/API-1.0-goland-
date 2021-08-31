package API_1_0_goland_
import (
	"github.com/gin-gonic/gin"
)

func main() {
	// create a new gin router
	router := gin.Default()

	// define a single endpoint
	router.GET("/", helloWorldhandler)

	// run the server on the port 3000
	_ = router.Run(":3000")
}

func helloWorldhandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "hello world",
	})
}
