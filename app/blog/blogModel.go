package blog

import (
	"log"
	"mime/multipart"
	"time"

	// "fmt"

	"gin-blog/config"

	"gorm.io/gorm"
)

// blog
type ModelBlog struct {
	gorm.Model
	ID            uint64 `gorm:"not null;primary_key"`
	Title         string `gorm:"size:255;not null"`
	Slug          string `gorm:"size:255;not null;uniqueIndex"`
	Author        string `gorm:"size:255;not null"`
	Content       string `gorm:"type:text;"`
	FeaturedImage string `gorm:"type:text;"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     gorm.DeletedAt
}

type ValidateBlogInput struct {
	// Title string `json:"title" binding:"required"`
	// Slug string `json:"slug" binding:"required"`
	// Content string `json:"content" binding:"required"`
	// Author string `json:"author" binding:"required"`

	Title   string `form:"title" binding:"required,min=5"`
	Slug    string `form:"slug" binding:"required,min=5,lowercase"`
	Content string `form:"content" binding:"required,min=5"`
	Author  string `form:"author" binding:"required,min=5"`

	FeaturedImage *multipart.FileHeader `form:"featuredimage" binding:"required"`
}

type ValidateBlogDelete struct {
	ID int `form:"id" binding:"required"`
}

func Migrate(db *gorm.DB) {
	err := db.Debug().AutoMigrate(&ModelBlog{}, &ModelBlog{})
	if err != nil {
		log.Fatal(err)
	}

	// database.DBCon.Model(&models.User{}).AddForeignKey("address_id", "address(id)", "CASCADE", "RESTRICT")
}

func (modelblog *ModelBlog) CreateBlog() (*ModelBlog, error) {
	DB := config.GetDB()
	err := DB.Create(&modelblog).Error

	if err != nil {
		return &ModelBlog{}, err
	}
	return modelblog, nil
}

func FindOneBlog(condition interface{}) (ModelBlog, error) {
	DB := config.GetDB()
	var modelblog ModelBlog
	err := DB.Where(condition).First(&modelblog).Error
	return modelblog, err
}

func (modelblog *ModelBlog) UpdateBlog(data interface{}) error {
	DB := config.GetDB()
	err := DB.Model(&modelblog).Updates(data).Error
	return err
}

// blog category
type ModelBlogCategory struct {
	gorm.Model
	ID            uint64 `gorm:"not null;primary_key"`
	Name          string `gorm:"size:255;not null"`
	Slug          string `gorm:"size:255;not null;uniqueIndex"`
	FeaturedImage string `gorm:"type:text;"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     gorm.DeletedAt
}

type ValidateBlogCategoryInput struct {
	Name string `form:"title" binding:"required,min=5"`
	Slug string `form:"slug" binding:"required,min=5,lowercase"`

	FeaturedImage *multipart.FileHeader `form:"featuredimage" binding:"required"`
}

type ValidateBlogCategoryDelete struct {
	ID int `form:"id" binding:"required"`
}

func (modelblogcategory *ModelBlogCategory) CreateBlogCategory() (*ModelBlogCategory, error) {
	DB := config.GetDB()
	err := DB.Create(&modelblogcategory).Error

	if err != nil {
		return &ModelBlogCategory{}, err
	}
	return modelblogcategory, nil
}

func FindOneBlogCategory(condition interface{}) (ModelBlogCategory, error) {
	DB := config.GetDB()
	var modelblogcategory ModelBlogCategory
	err := DB.Where(condition).First(&modelblogcategory).Error
	return modelblogcategory, err
}

func (modelblogcategory *ModelBlogCategory) UpdateBlogCategory(data interface{}) error {
	DB := config.GetDB()
	err := DB.Model(&modelblogcategory).Updates(data).Error
	return err
}
