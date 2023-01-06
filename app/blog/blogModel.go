package blog

import(
	"log"
	"time"
	"mime/multipart"
	// "fmt"
	
	"gin-blog/config"
	"gorm.io/gorm"
)

type ModelBlog struct {
	gorm.Model
	ID            uint64 `gorm:"not null;primary_key"`
	Title		  string `gorm:"size:255;not null"`
	Slug	      string `gorm:"size:255;not null;uniqueIndex"`
	Author        string `gorm:"size:255;not null"`
	Content       string `gorm:"type:text;"`
	FeaturedImage string `gorm:"type:text;"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     gorm.DeletedAt
}

type ModelBlogCategory struct {
	gorm.Model
	ID            uint64 `gorm:"not null;primary_key"`
	Name     	  string `gorm:"size:255;not null"`
	Slug          string `gorm:"size:255;not null;uniqueIndex"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     gorm.DeletedAt
}

type ValidateBlogInput struct {
    // Title string `json:"title" binding:"required"`
    // Slug string `json:"slug" binding:"required"`
    // Content string `json:"content" binding:"required"`
    // Author string `json:"author" binding:"required"`
    
	Title string `form:"title" binding:"required,min=5"`
    // Slug string `form:"slug" binding:"required,min=5,lowercase"`
    Content string `form:"content" binding:"required,min=5"`
    Author string `form:"author" binding:"required,min=5"`
    
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

func (blogmodel *ModelBlog) CreateBlog() (*ModelBlog, error) {
	DB := config.Init()
    err := DB.Create(&blogmodel).Error

    if err != nil {
        return &ModelBlog{}, err
    }
    return blogmodel, nil
}

func FindBlogById(id uint64) (ModelBlog, error) {
	var blogmodel ModelBlog

	DB := config.Init()
	err := DB.Where("id=?", id).First(&blogmodel).Error
	if err != nil {
		return ModelBlog{}, err
	}
	return blogmodel, nil
}