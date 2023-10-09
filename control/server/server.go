package server

import (
	"log"

	"github.com/gin-gonic/gin"
)

func FileUpload(c *gin.Context) {
	file, _ := c.FormFile("file")

	log.Println(file.Filename)

	c.SaveUploadedFile(file, "files/" + file.Filename)

	c.String(200, "Uploaded")
}

func Listen() {
	r := gin.Default()
	r.MaxMultipartMemory = 8 << 20
	r.POST("/upload", FileUpload)
	r.Run(":8000")
}
