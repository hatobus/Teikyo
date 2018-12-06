package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/hatobus/Teikyo/util"
)

func main() {
	r := gin.Default()

	util.Loadenv()

	r.POST("/detect", createTeikyohandler)

	r.Run(":8080")
}

func createTeikyohandler(c *gin.Context) {

	form, _ := c.MultipartForm()
	files := form.File["upload[]"]

	exe, _ := os.Getwd()
	fmt.Println(filepath.Join(filepath.Dir(exe), "picture", "output"))

	for _, file := range files {
		log.Println(file.Filename)
		err := c.SaveUploadedFile(file, filepath.Join(filepath.Dir(exe), "Teikyo", "picture", "output", file.Filename))
		if err != nil {
			log.Println(err)
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "craeted",
	})
}
