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
			return
		}
	}
	c.JSON(404, gin.H{
		"errorMessage": "Not found",
	})
}

func deletePostById(c *gin.Context) {
	id := c.Param("id")
	for i := 0; i < len(blog); i++ {
		if blog[i].ID == id {
			newBlog := remove(blog, i)
			blog = newBlog
			//fmt.Println(blog)
			c.JSON(204, blog)
			return
		}
	}
}

func postBlog(c *gin.Context) {
	var newBlogPost BlogPost

	// Call BindJSON to bind the received JSON to
	// newBlog.
	if err := c.BindJSON(&newBlogPost); err != nil {
		return
	}

	// Add the new album to the slice.

	blog = append(blog, newBlogPost)
	c.JSON(201, blog)
}

func main() {
	r := gin.Default()
	r.GET("/posts", getBlogPosts)
	r.POST("/posts", postBlog)
	r.GET("/posts/:id", getBlogPostById)
	r.DELETE("/posts/:id", deletePostById)
	r.Run()
}

func remove(slice []BlogPost, s int) []BlogPost {
	return append(slice[:s], slice[s+1:]...)
}
