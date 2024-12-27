package ikuai

import (
	"bytes"
	"crypto/md5"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"strings"
)

const (
	Salt                       = "salt_11"
	ActionLoginPath            = "/Action/login"
	UsernameParam              = "username"
	PasswordParam              = "passwd"
	PassParam                  = "pass"
	RememberPasswordParam      = "remember_password"
	HeaderSetCookie            = "Set-Cookie"
	ContentTypeApplicationJSON = "application/json"
)

func md5Hash(input string) string {
	hash := md5.New()
	hash.Write([]byte(input))
	return fmt.Sprintf("%x", hash.Sum(nil))
}

func base64EncodeWithSalt(password, salt string) string {
	combined := salt + password
	return base64.StdEncoding.EncodeToString([]byte(combined))
}

func (c *RpcClient) accessToken() (token string, err error) {
	config := c.config
	loginActionUrl := config.Url + ActionLoginPath
	passwd := md5Hash(config.Password)
	pass := base64EncodeWithSalt(config.Password, Salt)
	params := map[string]any{
		UsernameParam:         config.Username,
		PasswordParam:         passwd,
		PassParam:             pass,
		RememberPasswordParam: true,
	}
	jsonParams, err := json.Marshal(params)
	if err != nil {
		return token, err
	}
	resp, err := c.client.Post(loginActionUrl, ContentTypeApplicationJSON, bytes.NewBuffer(jsonParams))
	if err != nil {
		return token, err
	}
	defer resp.Body.Close()

	cookie := resp.Header.Get(HeaderSetCookie)
	// Split the cookie string by ";"
	cookies := strings.Split(cookie, ";")

	// Create a map to store key-value pairs
	cookieMap := make(map[string]string)

	// Iterate over each cookie and extract key-value pairs
	for _, cookie := range cookies {
		cookie = strings.TrimSpace(cookie) // Trim leading and trailing spaces
		if cookie != "" {
			parts := strings.Split(cookie, "=")
			if len(parts) == 2 {
				cookieMap[parts[0]] = parts[1]
			}
		}
	}

	token = cookieMap["sess_key"]
	c.Printf("Get access token: %s", token)
	return token, err
}

func (c *RpcClient) login() (err error) {
	c.token, err = c.accessToken()
	return err
}
