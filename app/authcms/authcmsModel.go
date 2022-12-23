package authcms

import(
	"log"
	"time"
	"html"
	"strings"
	
	"gin-blog/config"
	"gorm.io/gorm"
	"golang.org/x/crypto/bcrypt"
)

type UserModel struct {
	gorm.Model
	ID            uint64 `gorm:"not null;primary_key"`
	FirstName     string `gorm:"size:255;not null"`
	LastName      string `gorm:"size:255;not null"`
	Email         string `gorm:"size:255;not null;uniqueIndex"`
	Password      string `gorm:"size:255;not null"`
	RememberToken string `gorm:"size:255"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     gorm.DeletedAt
}

type RoleModel struct {
	gorm.Model
	ID            uint64 `gorm:"not null;primary_key"`
	Name     	  string `gorm:"size:255;not null"`
	Slug          string `gorm:"size:255;not null;uniqueIndex"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     gorm.DeletedAt
}

type RegisterInput struct {
    Email string `json:"email" binding:"required,email,min=3"`
    FirstName string `json:"firstname" binding:"required,min=3"`
    LastName string `json:"lastname" binding:"required,min=3"`
    Password string `json:"password" binding:"required,min=6,mustalphanum"`
}

type AuthenticationInput struct {
	Email string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

func Migrate(db *gorm.DB) {
	err := db.Debug().AutoMigrate(&UserModel{}, &RoleModel{})
	if err != nil {
		log.Fatal(err)
	}

	// database.DBCon.Model(&models.User{}).AddForeignKey("address_id", "address(id)", "CASCADE", "RESTRICT")
}

func (usermodel *UserModel) Save() (*UserModel, error) {
	DB := config.Init()
    err := DB.Create(&usermodel).Error

    if err != nil {
        return &UserModel{}, err
    }
    return usermodel, nil
}

func (usermodel *UserModel) BeforeSave(*gorm.DB) error {
    passwordHash, err := bcrypt.GenerateFromPassword([]byte(usermodel.Password), bcrypt.DefaultCost)
    if err != nil {
        return err
    }
    usermodel.Password = string(passwordHash)
    usermodel.FirstName = html.EscapeString(strings.TrimSpace(usermodel.FirstName))
    usermodel.LastName = html.EscapeString(strings.TrimSpace(usermodel.LastName))
    usermodel.Email = html.EscapeString(strings.TrimSpace(usermodel.Email))
    return nil
}

func (usermodel *UserModel) ValidatePassword(password string) error {
	return bcrypt.CompareHashAndPassword([]byte(usermodel.Password), []byte(password))
}

func FindUserByEmail(email string) (UserModel, error) {
	var usermodel UserModel

	DB := config.Init()
    err := DB.Where("email=?", email).Find(&usermodel).Error

	if err != nil {
		return UserModel{}, err
	}
	return usermodel, nil
}

func FindUserById(id uint) (UserModel, error) {
	var usermodel UserModel

	DB := config.Init()
	err := DB.Preload("Entries").Where("ID=?", id).Find(&usermodel).Error

	if err != nil {
		return UserModel{}, err
	}
	return usermodel, nil
}