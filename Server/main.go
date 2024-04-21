package main

/*import (
	"log"

	"example.com/mod/controllers"
	"example.com/mod/models"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// Initializes the database connection
func InitDB() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open("database.db"), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	err = db.AutoMigrate(&models.Sneaker{}, &models.MenSneaker{})
	if err != nil {
		return nil, err
	}

	return db, nil
}

func CloseDB(db *gorm.DB) {
	sqlDB, err := db.DB()
	if err != nil {
		log.Println("Error retrieving SQL DB:", err)
	} else {
		sqlDB.Close()
	}
}

func main() {

	db, err := InitDB()
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}
	defer CloseDB(db)

	router := gin.Default()

	// Set up CORS middleware
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept"},
		ExposeHeaders:    []string{"Content-Length", "Content-Disposition"},
		AllowCredentials: true,
	}))

	// Attach Db
	router.Use(func(c *gin.Context) {
		c.Set("db", db)
		c.Next()
	})

	// Women's sneakers routes
	router.POST("/womensneakers", controllers.CreateWomenSneaker)
	router.GET("/womensneakers", controllers.GetWomenSneakers)
	router.GET("/womensneakers/:id", controllers.GetWomenSneakerByID)
	router.PUT("/womensneakers/:id", controllers.UpdateWomenSneaker)
	router.DELETE("/womensneakers/:id", controllers.DeleteWomenSneaker)
	router.POST("/uploadimage", controllers.UploadImage)

	// Men's sneakers routes
	router.POST("/mensneakers", controllers.CreateMenSneaker)
	router.GET("/mensneakers", controllers.GetMenSneakers)
	router.GET("/mensneakers/:id", controllers.GetMenSneakerByID)
	router.PUT("/mensneakers/:id", controllers.UpdateMenSneaker)
	router.DELETE("/mensneakers/:id", controllers.DeleteMenSneaker)

	router.Run(":8080")
}*/
