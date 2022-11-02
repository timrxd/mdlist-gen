package gen

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/spf13/viper"
)

type LoginResponse struct {
	Result string `json:"result"`
	Token  Token  `json:"token"`
}

type Token struct {
	Session string `json:"session"`
	Refresh string `json:"refresh"`
}

func Login() error {

	body := fmt.Sprintf(`{
		"username": "%s",
		"password": "%s"
}`,
		viper.Get("username"), viper.Get("password"))

	r, err := http.Post("https://api.mangadex.org/auth/login", "application/json", bytes.NewBuffer([]byte(body)))
	if err != nil {
		return err
	}

	var resp LoginResponse
	err = json.NewDecoder(r.Body).Decode(&resp)

	viper.Set("session", resp.Token.Session)

	return nil
}
