package authcms

import(
	"log"
	"time"
	// "gin-blog/config"
	"gorm.io/gorm"
)

type UserModel struct {
	gorm.Model
	ID            string `gorm:"size:36;not null;uniqueIndex;primary_key"`
	FirstName     string `gorm:"size:255;not null"`
	LastName      string `gorm:"size:255;not null"`
	Email         string `gorm:"size:255;not null;uniqueIndex"`
	Password      string `gorm:"size:255;not null"`
	RememberToken string `gorm:"size:255;not null"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     gorm.DeletedAt
}

type RoleModel struct {
	gorm.Model
	ID            string `gorm:"size:36;not null;uniqueIndex;primary_key"`
	Name     	  string `gorm:"size:255;not null"`
	Slug          string `gorm:"size:255;not null;uniqueIndex"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     gorm.DeletedAt
}

func Migrate(db *gorm.DB) {
	err := db.Debug().AutoMigrate(&UserModel{}, &RoleModel{})
	if err != nil {
		log.Fatal(err)
	}

	// database.DBCon.Model(&models.User{}).AddForeignKey("address_id", "address(id)", "CASCADE", "RESTRICT")
}
