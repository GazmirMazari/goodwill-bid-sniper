package login

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const (
	baseURL       = "https://buyerapi.shopgoodwill.com/api"
	loginURL      = baseURL + "/SignIn/Login"
	encryptionKey = "6696D2E6F042FEC4D6E3F32AD541143B" // Example, replace with actual key
	iv            = "0000000000000000"
)

type Client struct {
	httpClient  *http.Client
	token       string
	accessToken interface{}
}

func NewClient() *Client {
	return &Client{
		httpClient: &http.Client{},
	}
}

func encrypt(plaintext, key, iv string) string {
	block, _ := aes.NewCipher([]byte(key))
	plaintextBytes := []byte(plaintext)
	cfb := cipher.NewCFBEncrypter(block, []byte(iv))
	ciphertext := make([]byte, len(plaintextBytes))
	cfb.XORKeyStream(ciphertext, plaintextBytes)
	return base64.StdEncoding.EncodeToString(ciphertext)
}
func (c *Client) Login(username, password string) error {
	// Encrypt the username and password
	encryptedUsername := encrypt(username, encryptionKey, iv)
	encryptedPassword := encrypt(password, encryptionKey, iv)

	// Create the request payload
	payload := struct {
		UserName   string `json:"userName"`
		Password   string `json:"password"`
		Remember   bool   `json:"remember"`
		AppVersion string `json:"appVersion"`
		Browser    string `json:"browser"`
	}{
		UserName:   encryptedUsername,
		Password:   encryptedPassword,
		Remember:   false,
		AppVersion: "51f7af627b5d26aa",
		Browser:    "chrome",
	}

	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	// Create the HTTP request
	req, err := http.NewRequest("POST", loginURL, bytes.NewBuffer(jsonPayload))
	if err != nil {
		return err
	}

	// Set the request headers
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/58.0.3029.110 Safari/537.3")

	// Send the request
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(resp.Body)

	// Check the response status code
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("login failed with status code %d", resp.StatusCode)
	}

	return nil
}
