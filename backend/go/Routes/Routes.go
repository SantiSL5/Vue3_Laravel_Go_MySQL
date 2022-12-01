package Routes

import (
	"namazu/Category"
	"namazu/Table"

	// "net/http"
	// "fmt"
	"github.com/gin-gonic/gin"
)

//SetupRouter ... Configure routes
func SetupRouter() *gin.Engine {
	r := gin.Default()

	grp1 := r.Group("/api")
	{
		grp1.GET("table", Table.GetTables)
		grp1.GET("table/:id", Table.GetTableByID)

		grp1.GET("category", Category.GetCategories)
		grp1.GET("category/:id", Category.GetCategoryByID)
		// grp1.POST("category", Category.CreateCategory)
		// grp1.PUT("table/:id", Category.UpdateCategory)
		// grp1.DELETE("table/:id", Category.DeleteCategory)
	}
	return r

	// r := gin.Default()

	// r.Use(CORS)

	// api := r.Group("/api");

	// Table.TablesRouting(api.Group("/table"))
	// courts.CourtsRouting(api.Group("/courts"))
	// reservations.ReservationRouting(api.Group("/reservations"))
	// users.UsersRouting(api.Group("/users"))

	return r

}

// func CORS(c *gin.Context) {

// 	fmt.Println("join CORS")
// 	// First, we add the headers with need to enable CORS
// 	// Make sure to adjust these headers to your needs
// 	c.Header("Access-Control-Allow-Origin", "*")
// 	c.Header("Access-Control-Allow-Methods", "*")
// 	c.Header("Access-Control-Allow-Headers", "*")
// 	c.Header("Content-Type", "application/json")

// 	// Second, we handle the OPTIONS problem
// 	if c.Request.Method != "OPTIONS" {

// 		c.Next()

// 	} else {

// 		// Everytime we receive an OPTIONS request,
// 		// we just return an HTTP 200 Status Code
// 		// Like this, Angular can now do the real
// 		// request using any other method than OPTIONS
// 		c.AbortWithStatus(http.StatusOK)
// 	}
// }
