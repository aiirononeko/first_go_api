package main
import (
	"net/http"
	"github.com/labstack/echo"
)
type User struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
func main() {
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
		return c.String(http.StatusOK, "{\"token\":\""+token+"\"}")
	})
	// サーバー起動
	e.Logger.Fatal(e.Start(":1323"))
}
func helloWorld(c echo.Context) error {
	return c.String(http.StatusOK, "hello world!!")
}