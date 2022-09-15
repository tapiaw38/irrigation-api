package cache

import (
	"encoding/json"
	"time"

	"github.com/tapiaw38/irrigation-api/models"
)

func (c *RedisCache) SetProducer(key string, value *models.Producer) {
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

func (c *RedisCache) GetProducer(key string) *models.Producer {
	client := c.GetClient()
	val, err := client.Get(key).Result()
	if err != nil {
		return nil
	}
	producer := models.Producer{}
	err = json.Unmarshal([]byte(val), &producer)
	if err != nil {
		panic(err)
	}
	return &producer
}

func (c *RedisCache) SetProducers(key string, value *[]models.Producer) {
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

func (c *RedisCache) GetProducers(key string) *[]models.Producer {
	client := c.GetClient()
	val, err := client.Get(key).Result()
	if err != nil {
		return nil
	}
	producers := []models.Producer{}
	err = json.Unmarshal([]byte(val), &producers)

	if err != nil {
		panic(err)
	}
	return &producers
}
