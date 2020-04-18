// package main

// import (
// 	"net/http"
// 	"github.com/labstack/echo"
// 	"os"
// 	"github.com/jinzhu/gorm"
//   _ "github.com/jinzhu/gorm/dialects/sqlite"
// )

// type User struct {
// 	Email    string
// 	Password string
// }

// func main() {
// 	// サーバー用のインスタンスの取得
// 	e := echo.New()

// 	// DBConnection
// 	db, err := gorm.Open("sqlite3", "test.db")
//   if err != nil {
//     panic("データベースへの接続に失敗しました")
//   }
// 	defer db.Close()

// 	// migration
// 	db.AutoMigrate(&User{})

// 	// ルーティング設定
// 	e.GET("/helloworld", helloWorld)
// 	e.POST("/create", func(c echo.Context) error {
// 		return db.Create(&User{Email: "me@example.com", Password: "password"})
// 	})

// 	// サーバー起動
// 	e.Logger.Fatal(e.Start(":" + os.Getenv("PORT")))
// }

// func createUser(c echo.Context) error {
// 	return db.Create(&User{Email: "me@example.com", Password: "password"})
// }

// func helloWorld(c echo.Context) error {
// 	return c.String(http.StatusOK, "hello world!!")
// }

package main

import (
	"net/http"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/labstack/echo"
)

type User struct {
	gorm.Model
	Email    string `json:"email"`
	Password string `json:"password"`
}

var (
	db *gorm.DB
)

func main() {
	// gormでmysqlかsqliteに接続
	var err error
	db, err = gorm.Open("sqlite3", "/tmp/gorm.db")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	// スキーマのマイグレーション
	db.AutoMigrate(&User{})

	// サーバー用のインスタンスの取得
	e := echo.New()

	// ユーザー
	u := User{
		Email:    "me@example.com",
		Password: "password",
	}

	// ルーティング設定
	e.GET("/helloworld", helloWorld)
	e.POST("/login", func(c echo.Context) error {
		r := new(User)
		if err := c.Bind(r); err != nil {
			return c.String(http.StatusBadRequest, err.Error())
		}
		if r.Email != u.Email || r.Password != u.Password {
			return c.String(http.StatusUnauthorized, "login fail")
		}
		// 暫定
		token := "sampletoken"
		// gormでメアドとパスワードをDBに保存
		db.Create(&User{Email: "me@example.com", Password: "password"})

		return c.String(http.StatusOK, "{\"token\":\""+token+"\"}")
	})
	// 過去ログインしてきた履歴を返す
	e.GET("/loginhistory", loginhistory)

	// サーバー起動
	e.Logger.Fatal(e.Start(":" + os.Getenv("PORT")))
}

func helloWorld(c echo.Context) error {
	return c.String(http.StatusOK, "hello world!!")
}

func loginhistory(c echo.Context) error {
	// gormでメアドとパスワードを保存したテーブルから、データ取得
	var user User
	db.Find(&user)
	return c.JSON(http.StatusOK, user)
}
