package blog

import (
	"fmt"
	// "io"
	"os"
	"net/http"
	"path/filepath"

	"github.com/google/uuid"
	"github.com/gabriel-vasile/mimetype"
	"gin-blog/config"
	"github.com/gin-gonic/gin"
)

func RenderBlogHello(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "render Blog Hello"})
}

func CreateBlog(c *gin.Context) {
	// body, _ := io.ReadAll(c.Request.Body)
    // fmt.Println(string(body))
	// fmt.Println("==================================")	
	
	var input ValidateBlogInput
	err := c.Bind(&input);
	// err := c.ShouldBind(&input);
	fmt.Println("==================================")
	fmt.Println(err)
	fmt.Println("==================================")
    if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error data validation": err.Error()})
        return
    }

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
		Slug: input.Slug,
		FeaturedImage: newpath+"/" + newFileName,
    }
    
    savedBlog, err := blog.CreateBlog()
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusCreated, gin.H{"user": savedBlog})
}

func UpdateBlog(c *gin.Context) {	
	var blog ModelBlog
	DB := config.Init()
	if err := DB.Where("id = ?", c.Param("id")).First(&blog).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}	
	
	var input ValidateBlogInput
	err := c.Bind(&input);
	// err := c.ShouldBind(&input);
	fmt.Println("==================================")
	fmt.Println(err)
	fmt.Println("==================================")
    if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error data validation": err.Error()})
        return
    }

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

	var updatedInput ModelBlog
    updatedInput.Title = input.Title
    updatedInput.Slug = input.Slug
    updatedInput.Author = input.Author
    updatedInput.Content = input.Content
    updatedInput.FeaturedImage = newpath+"/" + newFileName

    // DB.Model(&blog).Updates(updatedInput)
    result := DB.Model(&blog).Where("id = ?", c.Param("id")).Updates(updatedInput)
	if result.Error != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": result.Error})
        return
    }

    c.JSON(http.StatusCreated, gin.H{"blog": blog})
}


func fileNameWithoutExt(fileName string) string {
	return fileName[:len(fileName)-len(filepath.Ext(fileName))]
}