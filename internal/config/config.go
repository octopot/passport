package config

import (
	"net/url"
	"sync"
	"time"

	"go.octolab.org/toolkit/config"
)

// ApplicationConfig holds all configurations of the application.
type ApplicationConfig struct {
	Union struct {
		CPUs int `json:"cpus" yaml:"cpus"`

		DatabaseConfig   `json:"db"         yaml:"db"`
		GRPCConfig       `json:"grpc"       yaml:"grpc"`
		MigrationConfig  `json:"migration"  yaml:"migration"`
		MonitoringConfig `json:"monitoring" yaml:"monitoring"`
		ProfilingConfig  `json:"profiling"  yaml:"profiling"`
		ServerConfig     `json:"server"     yaml:"server"`
		ServiceConfig    `json:"service"    yaml:"service"`
	} `json:"config" yaml:"config"`
}

// DatabaseConfig contains configuration related to database.
type DatabaseConfig struct {
	DSN         config.Secret `json:"dsn"      yaml:"dsn"`
	MaxIdle     int           `json:"idle"     yaml:"idle"`
	MaxOpen     int           `json:"open"     yaml:"open"`
	MaxLifetime time.Duration `json:"lifetime" yaml:"lifetime"`

	once sync.Once
	dsn  *url.URL
}

// DriverName returns database driver name.
func (cnf *DatabaseConfig) DriverName() string {
	cnf.once.Do(func() {
		cnf.dsn, _ = url.Parse(string(cnf.DSN))
	})
	return cnf.dsn.Scheme
}

// GRPCConfig contains configuration related to gRPC server.
type GRPCConfig struct {
	Interface string        `json:"interface" yaml:"interface"`
	Timeout   time.Duration `json:"timeout"   yaml:"timeout"`
	Token     config.Secret `json:"token"     yaml:"token"`
}

// MigrationConfig contains configuration related to migrations.
type MigrationConfig struct {
	Table  string `json:"table"     yaml:"table"`
	Schema string `json:"schema"    yaml:"schema"`
	Limit  uint   `json:"limit"     yaml:"limit"`
	DryRun bool   `json:"dry-run"   yaml:"dry-run"`
}

// MonitoringConfig contains configuration related to monitoring.
type MonitoringConfig struct {
	Enabled   bool   `json:"enabled"   yaml:"enabled"`
	Interface string `json:"interface" yaml:"interface"`
}

// ProfilingConfig contains configuration related to profiler.
type ProfilingConfig struct {
	Enabled   bool   `json:"enabled"   yaml:"enabled"`
	Interface string `json:"interface" yaml:"interface"`
}

// ServerConfig contains configuration related to the server.
type ServerConfig struct {
	Interface         string        `json:"interface"           yaml:"interface"`
	ReadTimeout       time.Duration `json:"read-timeout"        yaml:"read-timeout"`
	ReadHeaderTimeout time.Duration `json:"read-header-timeout" yaml:"read-header-timeout"`
	WriteTimeout      time.Duration `json:"write-timeout"       yaml:"write-timeout"`
	IdleTimeout       time.Duration `json:"idle-timeout"        yaml:"idle-timeout"`
	BaseURL           string        `json:"base-url"            yaml:"base-url"`
}

// ServiceConfig contains configuration related to the service.
type ServiceConfig struct{}
