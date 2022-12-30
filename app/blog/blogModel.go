package blog

import(
	"log"
	"time"
	// "mime/multipart"
	
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
    Title string `json:"title" binding:"required,min=5"`
    Slug string `json:"slug" binding:"required,min=5"`
    Content string `json:"content" binding:"required,min=5"`
    Author string `json:"author" binding:"required"`
    // FeaturedImage multipart.File `json:"featuredimage" binding:"required"`
}

func Migrate(db *gorm.DB) {
	err := db.Debug().AutoMigrate(&ModelBlog{}, &ModelBlog{})
	if err != nil {
		log.Fatal(err)
	}

	// database.DBCon.Model(&models.User{}).AddForeignKey("address_id", "address(id)", "CASCADE", "RESTRICT")
}

func (blogmodel *ModelBlog) Save() (*ModelBlog, error) {
	DB := config.Init()
    err := DB.Create(&blogmodel).Error

    if err != nil {
        return &ModelBlog{}, err
    }
    return blogmodel, nil
}

