package blog

import (
	"fmt"
	// "io"
	"net/http"
	"os"
	"path/filepath"

	// "sync"
	"time"

	"gin-blog/app/common"
	"gin-blog/config"

	"github.com/gabriel-vasile/mimetype"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/gosimple/slug"
)

var err error

func CreateBlogCategory(c *gin.Context) {
	startTime := time.Now()

	var input ValidateBlogCategoryInput
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

	slugName := slug.Make(input.Name)
	id := uuid.New()
	// strID := id.String()
	blogcategory := ModelBlogCategory{
		ID:            id,
		Name:          input.Name,
		Slug:          slugName,
		FeaturedImage: newpath + "/" + newFileName,
	}

	savedBlogCategory, err := blogcategory.SaveBlogCategory()
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

	c.JSON(http.StatusOK, gin.H{"user": savedBlogCategory})
}

func UpdateBlogCategory(c *gin.Context) {
	urlParam := c.Param("id")
	// intUrlParam, err := strconv.ParseUint(urlParam, 10, 64)
	intUrlParam := uuid.Must(uuid.Parse(urlParam))

	modelBlogCategory, err := FindOneBlogCategory(&ModelBlogCategory{ID: intUrlParam})
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Record not found!"})
		return
	}

	var input ValidateBlogCategoryInput
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

	slugName := slug.Make(input.Name)
	var updatedInput ModelBlogCategory
	updatedInput.Name = input.Name
	updatedInput.Slug = slugName
	updatedInput.FeaturedImage = newpath + "/" + newFileName
	err = modelBlogCategory.UpdateBlogCategory(updatedInput)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusAccepted, gin.H{"blogcategory": modelBlogCategory})
}

func DeleteBlogCategory(c *gin.Context) {
	var input ValidateBlogCategoryDelete
	err := c.Bind(&input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error data validation": err.Error()})
		return
	}

	var blogcategory ModelBlogCategory
	DB := config.GetDB()
	if err := DB.Where("id = ?", input.ID).First(&blogcategory).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	urlParam := c.Param("id")
	// intUrlParam, err := strconv.Atoi(urlParam)
	intUrlParam := uuid.Must(uuid.Parse(urlParam))

	if intUrlParam != input.ID {
		c.JSON(http.StatusNotAcceptable, gin.H{"error": "Delete Not Acceptable, invalid data! "})
		return
	}

	result := DB.Where("id = ?", input.ID).Delete(&blogcategory)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": result.Error})
		return
	}

	c.JSON(http.StatusAccepted, gin.H{"message": "succes"})
}
