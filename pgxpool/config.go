package pgxpool

import (
	"context"
	"time"
)

// Config holds the configuration for the pool.
type Config struct {
	MaxConnIdleTime time.Duration
	HealthCheckPeriod time.Duration
	BeforeAcquire func(ctx context.Context, conn *Conn) bool
}

// NewConfig creates a default configuration with health checks enabled.
func NewConfig(connString string) (*Config, error) {
	return &Config{
		MaxConnIdleTime:   30 * time.Minute,
		HealthCheckPeriod: 1 * time.Minute,
		BeforeAcquire: func(ctx context.Context, conn *Conn) bool {
			// Default health check: ping if connection has been idle
			return conn.Ping(ctx) == nil
		},
	}, nil
}