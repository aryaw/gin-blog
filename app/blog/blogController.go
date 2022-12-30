package blog

import (
	// "fmt"
	// "io"
	"os"
	"net/http"
	"path/filepath"

	"github.com/google/uuid"
	"github.com/gabriel-vasile/mimetype"
	"github.com/gin-gonic/gin"
)

func RenderBlogHello(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "render Blog Hello"})
}

func SaveBlog(c *gin.Context) {
	// body, _ := io.ReadAll(c.Request.Body)
    // println(string(body))

	var input ValidateBlogInput
    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error validate": err.Error()})
        return
    }

	// file, err := c.FormFile("featuredimage")
	file, fl, err := c.Request.FormFile("featuredimage")
	if err != nil {
        c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
            "message": "No featured image is received",
        })
        return
    }

	mtype, err := mimetype.DetectReader(file)
	// fmt.Println(mtype.String(), mtype.Extension())
	filename := fileNameWithoutExt(filepath.Base(fl.Filename))
	extension := mtype.Extension()
	newFileName := uuid.New().String() + filename + extension

	newpath := filepath.Join(".", os.Getenv("SAVED_UPLOAD_CONTENT"))
	checkDir, err := os.Stat(newpath)
	if checkDir == nil {
		err := os.MkdirAll(newpath, 0755)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"message": "Can't create directory",
			})
			return
		}

	}
	
	if err := c.SaveUploadedFile(fl, newpath+"/" + newFileName); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": "Unable to save the file",
		})
		return
	}
    
    blog := ModelBlog {
        Title: input.Title,
		Author: input.Author,
		Content: input.Content,
		FeaturedImage: newpath+"/" + newFileName,
    }
    
    savedBlog, err := blog.Save()
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusCreated, gin.H{"user": savedBlog})
}


func fileNameWithoutExt(fileName string) string {
	return fileName[:len(fileName)-len(filepath.Ext(fileName))]
}