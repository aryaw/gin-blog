package blog

import (
	"fmt"
	// "io"
	"os"
	"strconv"
	"net/http"
	"path/filepath"
	// "sync"
	"time"

	"github.com/google/uuid"
	"github.com/gabriel-vasile/mimetype"
	"gin-blog/config"
	"github.com/gosimple/slug"
	"github.com/gin-gonic/gin"
)

var err error
func CreateBlogCategory(c *gin.Context) {
	startTime := time.Now()
	
	var input ValidateBlogCategoryInput
	err := c.Bind(&input);
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
	filename := fileNameWithoutExt(filepath.Base(fl.Filename))
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
	
	if err := c.SaveUploadedFile(fl, newpath+"/" + newFileName); err != nil {
		// mtx.Unlock()
		// wg.Done()
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": "Unable to save the file",
		})
		return
	}

    slugTitle := slug.Make(input.Title)
    blogcategory := ModelBlogCategory {
        Name: input.Name,
		Slug: slugTitle,
		FeaturedImage: newpath+"/" + newFileName,
    }
    
    savedBlogCategory, err := blogcategory.CreateBlogCategory()
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
	intUrlParam, err := strconv.ParseUint(urlParam, 10, 64)
    if err != nil {
        panic(err)
    }

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

	slugTitle := slug.Make(input.Title)
	var updatedInput ModelBlogCategory
    updatedInput.Name = input.Title
    updatedInput.Slug = slugTitle
    updatedInput.FeaturedImage = newpath+"/" + newFileName
	err = modelBlogCategory.UpdateBlogCategory(updatedInput)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err})
		return
	}

    c.JSON(http.StatusAccepted, gin.H{"blogcategory": modelBlogCategory})
}

func DeleteBlogCategory(c *gin.Context) {
	var input ValidateBlogCategoryDelete
	err := c.Bind(&input);
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
	intUrlParam, err := strconv.Atoi(urlParam)
    if err != nil {
        panic(err)
    }

	if (intUrlParam != input.ID) {
		c.JSON(http.StatusNotAcceptable , gin.H{"error": "Delete Not Acceptable, invalid data! "})
		return
	}

    result := DB.Where("id = ?", input.ID).Delete(&blogcategory)
	if result.Error != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": result.Error})
        return
    }

    c.JSON(http.StatusAccepted, gin.H{"message": "succes"})
}


func fileNameWithoutExt(fileName string) string {
	return fileName[:len(fileName)-len(filepath.Ext(fileName))]
}