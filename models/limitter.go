package models

import "time"

type Config struct {
	RedisAddr         string
	RedisPassword     string
	RedisDB           int
	DefaultRateLimit  int
	CacheTTL          time.Duration
	SubscriptionPlans map[string]Plan
	FeatureFlags      map[string]bool
}

type Plan struct {
	Name         string
	Features     []string
	RateLimit    int
	MonthlyPrice float64
}

type UserLimits struct {
	UserID      string
	Plan        string
	Features    map[string]bool
	RateLimit   int
	LastUpdated time.Time
}

const (
	DefaultCacheTTL = 1 * time.Minute
	DefaultRate     = 100
)
