package curl

import (
	"bytes"
	"encoding/json"
	"net/url"
	"strings"
)

type HeaderHandlers func(*Curl)

func FormHeader() HeaderHandlers {
	return func(c *Curl) {
		c.request.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	}
}

func JsonHeader() HeaderHandlers {
	return func(c *Curl) {
		c.request.Header.Add("Content-Type", "application/json")
	}
}

func SetHeader(key, val string) HeaderHandlers {
	return func(c *Curl) {
		c.request.Header.Add(key, val)
	}
}

func Authorization(token string) HeaderHandlers {
	return func(c *Curl) {
		c.request.Header.Add("Authorization", token)
	}
}

func (c *Curl) JsonData(data interface{}) (*Curl, error) {
	bs, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	c.data = bytes.NewReader(bs)
	return c, nil
}

func (c *Curl) QueryData(data map[string]string) *Curl {
	c.data = strings.NewReader(HttpBuildQuery(data))
	return c
}

func (c *Curl) Data(data *bytes.Buffer) *Curl {
	c.data = data
	return c
}

func HttpBuildQuery(params map[string]string) string {
	v := make(url.Values)
	for key, val := range params {
		v.Set(key, val)
	}
	return v.Encode()
}
