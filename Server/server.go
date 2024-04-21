package main

import (
	"fmt"
	"log"

	"example.com/mod/graph"
	"example.com/mod/graph/model"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

const defaultPort = ":8080"

func graphqlHandler() gin.HandlerFunc {

	h := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{}}))

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

func playgroundHandler() gin.HandlerFunc {
	h := playground.Handler("GraphQL", "/query")

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

func RunMigrations(db *gorm.DB) {

}

func main() {
	db, err := gorm.Open(sqlite.Open("graph/database.db"), &gorm.Config{})
	if err != nil {
		log.Println("Unable to connect to SQLite database:", err)
	}
	fmt.Println("Running migrations...")
	db.AutoMigrate(&model.WomenSneakers{}, &model.MenSneakers{})
	fmt.Println("Migrations completed successfully.")

	r := gin.Default()
	r.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(200)
			return
		}
		c.Next()
	})
	r.POST("/query", graphqlHandler())
	r.GET("/", playgroundHandler())
	r.Run(defaultPort)

}
