package api

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"http_test/config"
	"http_test/pkg/util"
	"log"
	"net/http"
	"strconv"
	"time"
)

type TestClient struct {
	host string
	port int
}

func Get() *TestClient {
	cfg := *config.Get()
	return &TestClient{
		host: cfg.ApiHost,
		port: cfg.ApiPort,
	}
}
func (c *TestClient) GetToken(i int, reg bool) (string, error) {

	url := fmt.Sprintf("http://%s:%d/api/users", c.host, c.port)
	if reg {
		url = fmt.Sprintf("http://%s:%d/api/users", c.host, c.port)
	}
	login := fmt.Sprintf("f0%d@m.ru", i+1)
	password := fmt.Sprintf("f0%d", i+1)
	strLogin := fmt.Sprintf("{\"email\": \"%s\", \"password\": \"%s\"}", login, password)
	log.Printf("s=%s", strLogin)
	jsonBody := []byte(strLogin)
	bodyReader := bytes.NewReader(jsonBody)

	req, err := http.NewRequest(http.MethodPost, url, bodyReader)
	if err != nil {
		log.Printf("client: could not create request: %s", err.Error())
		return "", err
	}
	req.Header.Set("Content-Type", "application/json")

	client := http.Client{Timeout: 10 * time.Second}

	res, err := client.Do(req)
	if err != nil {
		log.Printf("client: error making http request: %s", err.Error())
		return "", err
	}
	defer res.Body.Close()

	if reg {
		return "", nil
	}
	var result map[string]interface{}

	err = json.NewDecoder(res.Body).Decode(&result)
	if err != nil {
		log.Printf("client: could not read response body: %s", err.Error())
		return "", err
	}
	// token := result["access_token"].(string)
	if id, ok := result["id"].(float64); ok == true {
		token := strconv.FormatFloat(id, 'f', -1, 64)
		log.Printf("res id: %s", token)
		return token, nil
	}
	return "", errors.New("no id")
}

/*
post: http://{{host}}/api/users/2/items

	  "title": "item1 of user2",
		"description": "descr 2"
*/
func (c *TestClient) createItem(userId string, item int, token string) error {
	url := fmt.Sprintf("http://%s:%d/api/users/%s/items", c.host, c.port, userId)
	//log.Printf("url: %s", url)

	title := fmt.Sprintf("title of item %d from user %s", item+1, userId)
	description := fmt.Sprintf("descritption of item %d from user %s", item+1, userId)
	strBody := fmt.Sprintf("{\"title\": \"%s\", \"description\": \"%s\"}", title, description)
	log.Printf("s=%s", strBody)
	jsonBody := []byte(strBody)
	bodyReader := bytes.NewReader(jsonBody)

	req, err := http.NewRequest(http.MethodPost, url, bodyReader)
	if err != nil {
		log.Printf("client: could not create request: %s", err.Error())
		return err
	}
	req.Header.Set("Content-Type", "application/json")

	client := http.Client{Timeout: 10 * time.Second}

	res, err := client.Do(req)
	if err != nil {
		log.Printf("client: error making http request: %s", err.Error())
		return err
	}
	defer res.Body.Close()

	for true {
		bs := make([]byte, 1014)
		n, err := res.Body.Read(bs)
		//fmt.Println(string(bs[:n]))
		if n == 0 || err != nil {
			break
		}
	}

	return nil
}

func (c *TestClient) CreateItems(i int) error {
	// token, err := c.GetToken(i, true)
	// if err != nil {
	// 	log.Printf("geth: %s", err.Error())
	// 	//return err
	// }
	token, err := c.GetToken(i, false)
	if err != nil {
		log.Printf("err get token: %s", err.Error())
		return err
	}
	//log.Printf("res: %s", token)
	for i := 0; i < 10; i++ {
		err = c.createItem(token, i, token)
		if err != nil {
			log.Printf("create: %s", err.Error())
			return err
		}
	}
	return nil
}

func (c *TestClient) sendFlat(name string, flat string) error {
	url := fmt.Sprintf("http://%s:%d/api/flats/%s/xx", c.host, c.port, name)
	//log.Printf("url: %s", url)

	strBody := fmt.Sprintf("{\"src\": \"%s\"}", flat)
	log.Printf("name=%s ", name)
	jsonBody := []byte(strBody)
	bodyReader := bytes.NewReader(jsonBody)

	req, err := http.NewRequest(http.MethodPost, url, bodyReader)
	if err != nil {
		log.Printf("client: could not create request: %s", err.Error())
		return err
	}

	client := http.Client{Timeout: 30 * time.Second}

	res, err := client.Do(req)
	if err != nil {
		log.Printf("client: error making http request: %s", err.Error())
		return err
	}
	defer res.Body.Close()

	for true {
		bs := make([]byte, 1014)
		n, err := res.Body.Read(bs)
		// fmt.Println(string(bs[:n]))
		if n == 0 || err != nil {
			break
		}
	}

	return nil
}
func (c *TestClient) UpdFlat(i int) error {
	name := fmt.Sprintf("%04d", i+1)
	flat := util.GetStrFlat()
	//fmt.Printf("\r\n name: %s flat: %s room: %s", name, flat, room)
	//return nil
	err := c.sendFlat(name, flat)
	return err
}

func (c *TestClient) CreateItemsWoLogin(userId int) error {
	token := fmt.Sprintf("%d", userId)
	//log.Printf("res: %s", token)
	for i := 0; i < 10; i++ {
		err := c.createItem(token, i, token)
		if err != nil {
			log.Printf("create: %s", err.Error())
			return err
		}
	}
	return nil
}

func (c *TestClient) Test() error {
	url := fmt.Sprintf("http://%s:%d/users/login", c.host, c.port)
	log.Printf("url: %s", url)
	return nil
}
