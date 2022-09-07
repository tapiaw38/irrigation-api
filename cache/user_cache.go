package cache

import (
	"encoding/json"
	"time"

	"github.com/tapiaw38/irrigation-api/models/user"
)

type UserCache struct {
	Cache *RedisCache
}

func (ur *UserCache) Set(key string, value *user.User) {

	client := ur.Cache.GetClient()
	json, err := json.Marshal(&value)

	if err != nil {
		panic(err)
	}

	err = client.Set(key, json, ur.Cache.Expires*time.Hour).Err()

	if err != nil {
		panic(err)
	}
}

func (ur *UserCache) Get(key string) *user.User {

	client := ur.Cache.GetClient()
	val, err := client.Get(key).Result()

	if err != nil {
		return nil
	}

	user := user.User{}
	err = json.Unmarshal([]byte(val), &user)

	if err != nil {
		panic(err)
	}

	return &user
}
