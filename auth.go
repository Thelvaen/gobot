package main

import (
	"fmt"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"github.com/thelvaen/gobot/csrf"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// User structure made exportable to be used with Gorm ORM
type User struct {
	gorm.Model
	Name     string `gorm:"not null;unique" form:"username"` // Utilisateur unique!
	Password string `gorm:"not null" form:"password"`
	Email    string `gorm:"not null"`
	Roles    []Role `gorm:"many2many:user_roles;"`
}

// Role struct allow to gives Roles to user
type Role struct {
	gorm.Model
	Name  string `gorm:"not null;unique"` // Role unique!
	Users []User `gorm:"many2many:user_roles;"`
}

func initAuth() {
	// Initialisation du cookie store
	seedRoles := true
	seedUsers := true
	if BotConfig.DataStore.HasTable("roles") {
		seedRoles = false
	}
	if BotConfig.DataStore.HasTable("users") {
		seedUsers = false
	}
	BotConfig.DataStore.AutoMigrate(&User{}, &Role{})
	if seedRoles {
		adminRole := Role{Name: "admin"}
		userRole := Role{Name: "user"}
		BotConfig.DataStore.Create(&adminRole)
		BotConfig.DataStore.Create(&userRole)
	}
	if seedUsers {
		viper.SetConfigName("init")
		viper.MergeInConfig()
		var (
			role      Role
			adminRole []Role
		)
		err = BotConfig.DataStore.Where("Name = ?", "admin").First(&role).Error
		if err != nil {
			myPanic("can't fetch default admin role :", err)
		}
		adminRole = append(adminRole, role)
		clearPassword := viper.GetString("Seed.Password")
		hashPassword, _ := bcrypt.GenerateFromPassword(
			[]byte(clearPassword),
			bcrypt.DefaultCost,
		)
		userSeed := User{
			Name:     viper.GetString("Seed.Name"),
			Password: string(hashPassword),
			Email:    viper.GetString("Seed.Email"),
			Roles:    adminRole,
		}
		BotConfig.DataStore.Create(&userSeed)
	}
}

func loginHandlerForm(c *gin.Context) {
	c.HTML(200, "login_form.html", gin.H{
		"CSRF_Token": csrf.GetToken(c),
	})
}

func logoutHandler(c *gin.Context) {
	session := sessions.Default(c)

	session.Delete("userID")
	session.Save()
	c.Redirect(302, "/")
}

func loginHandler(c *gin.Context) {
	var user User
	session := sessions.Default(c)

	if c.Bind(&user) == nil {
		clearPassword := user.Password
		if err := BotConfig.DataStore.Where("name = ?", user.Name).First(&user).Error; err == nil {
			if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(clearPassword)); err == nil {
				session.Set("userID", user.ID)
				if err := session.Save(); err != nil {
					c.HTML(http.StatusInternalServerError, "error.html", gin.H{
						"error": "Failed to save session",
					})
					return
				}
				c.Redirect(302, "/")
				return
			} else {
				fmt.Println(fmt.Errorf("erreur de verification du password", err))
			}
		}
	}

}

func getNewUserToken(c *gin.Context) {
	c.String(200, "test")
}
