package handler
import (
	"fmt"
	"net/http"
	"github.com/gin-gonic/gin"
)

func Upload(c *gin.Context) {
	form, _ := c.MultipartForm()
	files := form.File["file"]

	for _, file := range files {
		err := c.SaveUploadFile(file, "images/" + file.Filename)
		if err != nil {
			c.JSON(http.statusInternalServerError, gin.H{"message": err.Error()})
		}
	}

	c.JSON(http.StatusOK, gin.H{"message": "success!"})
}