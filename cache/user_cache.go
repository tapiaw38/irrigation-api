package cache

import (
	"encoding/json"
	"time"

	"github.com/tapiaw38/irrigation-api/models"
)

func (c *RedisCache) SetUser(key string, value *models.User) {
	client := c.GetClient()
	json, err := json.Marshal(&value)

	if err != nil {
		panic(err)
	}

	err = client.Set(key, json, c.Expires*time.Second).Err()

	if err != nil {
		panic(err)
	}
}

func (c *RedisCache) GetUser(key string) *models.User {
	client := c.GetClient()
	val, err := client.Get(key).Result()

	if err != nil {
		return nil
	}

	user := models.User{}
	err = json.Unmarshal([]byte(val), &user)

	if err != nil {
		panic(err)
	}

	return &user
}
