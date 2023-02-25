package blog

import (
	"log"
	"mime/multipart"
	"time"

	// "fmt"

	"gin-blog/config"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// blog
type ModelBlog struct {
	gorm.Model
	ID            uuid.UUID `gorm:"not null;primary_key;type:string"`
	Title         string    `gorm:"size:255;not null"`
	Slug          string    `gorm:"size:255;not null;uniqueIndex"`
	Author        string    `gorm:"size:255;not null"`
	Content       string    `gorm:"type:text;"`
	FeaturedImage string    `gorm:"type:text;"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     gorm.DeletedAt
	Blogs         []ModelBlogToCategory `gorm:"foreignKey:Blog"`
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

	Categories []uint64 `json:"categories"`
}

type ValidateBlogDelete struct {
	ID uuid.UUID `form:"id" binding:"required"`
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
	ID            uuid.UUID `gorm:"not null;primary_key;type:string"`
	Name          string    `gorm:"size:255;not null"`
	Slug          string    `gorm:"size:255;not null;uniqueIndex"`
	FeaturedImage string    `gorm:"type:text;"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     gorm.DeletedAt
	Categories    []ModelBlogToCategory `gorm:"foreignKey:Category"`
}

type ValidateBlogCategoryInput struct {
	Name string `form:"title" binding:"required,min=5"`
	// Slug string `form:"slug" binding:"required,min=5,lowercase"`

	FeaturedImage *multipart.FileHeader `form:"featuredimage" binding:"required"`
}

type ValidateBlogCategoryDelete struct {
	ID uuid.UUID `form:"id" binding:"required"`
}

func (modelblogcategory *ModelBlogCategory) SaveBlogCategory() (*ModelBlogCategory, error) {
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

// Blog to Category
type ModelBlogToCategory struct {
	gorm.Model
	Blog     int64
	Category int64
}

type ValidateBlogToCategoryInput struct {
	Blog     uint64 `form:"title" binding:""`
	Category uint64 `form:"slug" binding:""`
}

func (modelblogcategory *ModelBlogToCategory) CreateBlogToCategory() (*ModelBlogToCategory, error) {
	DB := config.GetDB()
	err := DB.Create(&modelblogcategory).Error

	if err != nil {
		return &ModelBlogToCategory{}, err
	}
	return modelblogcategory, nil
}

func (modelblogcategory *ModelBlogToCategory) UpdateBlogToCategory(data interface{}) error {
	DB := config.GetDB()
	err := DB.Model(&modelblogcategory).Updates(data).Error
	return err
}

// migrate db
func Migrate(db *gorm.DB) {
	err := db.Debug().AutoMigrate(&ModelBlog{}, &ModelBlogCategory{}, &ModelBlogToCategory{})
	if err != nil {
		log.Fatal(err)
	}

	// database.DBCon.Model(&models.User{}).AddForeignKey("address_id", "address(id)", "CASCADE", "RESTRICT")
}
