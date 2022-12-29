package blog

import(
	"log"
	"time"
	
	"gorm.io/gorm"
)

type BlogModel struct {
	gorm.Model
	ID            uint64 `gorm:"not null;primary_key"`
	Title		  string `gorm:"size:255;not null"`
	Slug	      string `gorm:"size:255;not null;uniqueIndex"`
	Author        string `gorm:"size:255;not null;uniqueIndex"`
	Content       string `gorm:"type:text;"`
	FeaturedImage       string `gorm:"type:text;"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     gorm.DeletedAt
}

type BlogCategoryModel struct {
	gorm.Model
	ID            uint64 `gorm:"not null;primary_key"`
	Name     	  string `gorm:"size:255;not null"`
	Slug          string `gorm:"size:255;not null;uniqueIndex"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     gorm.DeletedAt
}

func Migrate(db *gorm.DB) {
	err := db.Debug().AutoMigrate(&BlogModel{}, &BlogModel{})
	if err != nil {
		log.Fatal(err)
	}

	// database.DBCon.Model(&models.User{}).AddForeignKey("address_id", "address(id)", "CASCADE", "RESTRICT")
}