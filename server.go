package main

import (
	"bytes"
	"image"
	"io"
	"log"
	"net/http"

	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"

	"github.com/gin-gonic/gin"
	"github.com/hatobus/Teikyo/callapi"
	img "github.com/hatobus/Teikyo/imgprocessing"
	"github.com/hatobus/Teikyo/util"
)

func main() {
	r := gin.Default()

	err := util.Loadenv()
	if err != nil {
		panic(err)
	}

	r.POST("/detect", createTeikyohandler)

	r.Run(":8080")
}

func createTeikyohandler(c *gin.Context) {

	form, _ := c.MultipartForm()
	files := form.File["upload[]"]

	errch := make(map[string]string, len(form.File))
	b := new(bytes.Buffer)

	for _, file := range files {
		log.Println(file.Filename)

		f, err := file.Open()
		defer f.Close()

		// 一回DecodeConfigでファイルをいじるとファイルが壊れるために
		// 別のbufにコピーをして回避しておく
		io.Copy(b, f)

		_, format, err := image.DecodeConfig(b)
		if err != nil {
			errch[file.Filename] = err.Error()
			b.Reset()
			break
		} else if format != "jpeg" {
			errch[file.Filename] = "Filetype must be jpeg"
			b.Reset()
			break
		}

		b.Reset()

		landmark, err := callapi.DetectFace(f)
		if err != nil {
			errch[file.Filename] = err.Error()
			break
		}

		// spew.Dump(landmark)

		for _, L := range landmark {

			LM := L.ToLandmark()
			err := img.GenTeikyo(f, LM)
			if err != nil {
				errch[file.Filename] = err.Error()
			}

		}
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "craeted",
		"errors":  errch,
	})
}
