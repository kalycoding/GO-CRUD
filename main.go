package main

import (
	"github.com/gin-gonic/gin"
)

type BlogPost struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Author      string `json:"author"`
	Description string `json:"description"`
}

var blog = []BlogPost{
	{ID: "1", Title: "Learning Go", Author: "Kalycoding", Description: "My Path to Learning Go"},
}

func getBlogPosts(c *gin.Context) {
	c.JSON(200, blog)
}

func getBlogPostById(c *gin.Context) {
	for i := 0; i < len(blog); i++ {
		if blog[i].ID == c.Param("id") {
			c.JSON(200, blog[i])
		}
	}
	c.JSON(404, gin.H{
		"errorMessage": "Not found",
	})
}

func main() {
	r := gin.Default()
	r.GET("/posts", getBlogPosts)
	r.GET("/posts/:id", getBlogPostById)
	r.Run()
}
