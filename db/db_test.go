package db

import (
	"context"
	"testing"
)

func TestRedisClient(t *testing.T) {
	dbAddress := "redis:6379"
	_, err := NewDatabase(dbAddress)
	if err != nil {
		t.Fatalf("Client for Redis failed to be created: %v", err)
	}
}

func TestRedisPing(t *testing.T) {
	dbAddress := "redis:6379"
	ctx := context.TODO()
	db, _ := NewDatabase(dbAddress)
	_, err := db.Client.Ping(ctx).Result()
	if err != nil {
		t.Fatalf("Connection to Redis could not be established: %v", err)
	}
}
