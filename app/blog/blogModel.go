package blog

import(
	"log"
	"time"
	"mime/multipart"
	
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
    
	Title string `form:"title" binding:"-"`
    Slug string `form:"slug" binding:"-"`
    Content string `form:"content" binding:"-"`
    Author string `form:"author" binding:"-"`
    
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

// func (blogmodel *ModelBlog) FindBlogs() (*ModelBlog, error) {
// 	var blogs []ModelBlog
// 	DB := config.Init()
//     blogs, err := DB.Find(&blogs)
// 	if err != nil {
// 		return err
// 	}

// 	return blogs
// }

// func (blogmodel *ModelBlog) FindBlog(id int) (*ModelBlog, error) {
// 	var blog ModelBlog
// 	DB := config.Init()
// 	if err := DB.Where("id = ?", id).First(&blog).Error; err != nil {
// 		return nil
// 	}

// 	return blog
// }