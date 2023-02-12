package blog

import (
	"fmt"
	// "io"
	"net/http"
	"os"
	"path/filepath"
	"strconv"

	// "sync"
	"time"

	"gin-blog/app/common"
	"gin-blog/config"

	"github.com/gabriel-vasile/mimetype"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/gosimple/slug"
)

func RenderBlogHello(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "render Blog Hello"})
}

func CreateBlog(c *gin.Context) {
	// body, _ := io.ReadAll(c.Request.Body)
	// fmt.Println(string(body))
	// fmt.Println("==================================")
	startTime := time.Now()

	var input ValidateBlogInput
	err := c.Bind(&input)
	// err := c.ShouldBind(&input);
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

	// var wg sync.WaitGroup
	// var mtx sync.Mutex

	// wg.Add(1)
	// mtx.Lock()

	mtype, err := mimetype.DetectReader(file)
	// fmt.Println(mtype.String(), mtype.Extension())
	filename := common.FileNameWithoutExt(filepath.Base(fl.Filename))
	extension := mtype.Extension()
	newFileName := uuid.New().String() + filename + extension

	newpath := filepath.Join(".", os.Getenv("SAVED_UPLOAD_CONTENT"))
	checkDir, err := os.Stat(newpath)
	if checkDir == nil {
		err := os.MkdirAll(newpath, 0755)
		if err != nil {
			// mtx.Unlock()
			// wg.Done()
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"message": "Can't create directory",
			})
			return
		}

	}

	if err := c.SaveUploadedFile(fl, newpath+"/"+newFileName); err != nil {
		// mtx.Unlock()
		// wg.Done()
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": "Unable to save the file",
		})
		return
	}

	slugTitle := slug.Make(input.Title)
	blog := ModelBlog{
		Title:         input.Title,
		Author:        input.Author,
		Content:       input.Content,
		Slug:          slugTitle,
		FeaturedImage: newpath + "/" + newFileName,
	}

	savedBlog, err := blog.CreateBlog()
	// mtx.Unlock()
	// wg.Done()
	// wg.Wait()
	finishTime := time.Since(startTime)

	fmt.Println("===========================================================")
	fmt.Println(finishTime)
	fmt.Println("===========================================================")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"user": savedBlog})
}

func UpdateBlog(c *gin.Context) {
	// var blog ModelBlog
	// DB := config.GetDB()
	// if err := DB.Where("id = ?", c.Param("id")).First(&blog).Error; err != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
	// 	return
	// }

	urlParam := c.Param("id")
	intUrlParam, err := strconv.ParseUint(urlParam, 10, 64)
	if err != nil {
		panic(err)
	}

	modelBlog, err := FindOneBlog(&ModelBlog{ID: intUrlParam})
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Record not found!"})
		return
	}

	var input ValidateBlogInput
	err = c.Bind(&input)
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
	filename := common.FileNameWithoutExt(filepath.Base(fl.Filename))
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

	if err := c.SaveUploadedFile(fl, newpath+"/"+newFileName); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": "Unable to save the file",
		})
		return
	}

	slugTitle := slug.Make(input.Title)
	var updatedInput ModelBlog
	updatedInput.Title = input.Title
	updatedInput.Slug = slugTitle
	updatedInput.Author = input.Author
	updatedInput.Content = input.Content
	updatedInput.FeaturedImage = newpath + "/" + newFileName
	err = modelBlog.UpdateBlog(updatedInput)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusAccepted, gin.H{"blog": modelBlog})
}

func DeleteBlog(c *gin.Context) {
	var input ValidateBlogDelete
	err := c.Bind(&input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error data validation": err.Error()})
		return
	}

	var blog ModelBlog
	DB := config.GetDB()
	if err := DB.Where("id = ?", input.ID).First(&blog).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	urlParam := c.Param("id")
	intUrlParam, err := strconv.Atoi(urlParam)
	if err != nil {
		panic(err)
	}

	if intUrlParam != input.ID {
		c.JSON(http.StatusNotAcceptable, gin.H{"error": "Delete Not Acceptable, invalid data! "})
		return
	}

	result := DB.Where("id = ?", input.ID).Delete(&blog)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": result.Error})
		return
	}

	c.JSON(http.StatusAccepted, gin.H{"message": "succes"})
}
