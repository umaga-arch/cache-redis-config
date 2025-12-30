package helpers

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/go-redis/redis/v8"
)

// RedisConfig represents the configuration for Redis.
type RedisConfig struct {
	Host     string
	Port     int
	Password string
	DB       int
}

// GetRedisClient returns a Redis client with the given configuration.
func GetRedisClient(config RedisConfig) (*redis.Client, error) {
	return redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", config.Host, config.Port),
		Password: config.Password,
		DB:       config.DB,
	})
}

// ProcessResponse processes the response from the server.
func ProcessResponse(w http.ResponseWriter, r *http.Request, response interface{}) {
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// ProcessError processes the error from the server.
func ProcessError(w http.ResponseWriter, err error) {
	log.Println(err)
	http.Error(w, err.Error(), http.StatusInternalServerError)
}

// ProcessSuccess processes the success response from the server.
func ProcessSuccess(w http.ResponseWriter, response interface{}) {
	w.WriteHeader(http.StatusOK)
	ProcessResponse(w, nil, response)
}

// SetContext sets the context for the request.
func SetContext(ctx context.Context, key string, value interface{}) context.Context {
	return context.WithValue(ctx, key, value)
}

// GetContext gets the value from the context.
func GetContext(ctx context.Context, key string) (interface{}, bool) {
	v, ok := ctx.Value(key)
	return v, ok
}