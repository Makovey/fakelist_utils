package env

import (
	"fmt"

	"github.com/Makovey/fakelist_utils/pkg/logger"
)

type AuthConfig struct {
	authRunAddress string
}

func NewAuthConfig(log logger.Logger) (*AuthConfig, error) {
	fn := "env.NewAuthConfig"

	aggregator, err := newEnvAggregator(log)
	if err != nil {
		return nil, fmt.Errorf("[%s]: %w", fn, err)
	}

	log.Debug("AuthDatabaseDSN: " + aggregator.authAddress)

	return &AuthConfig{
		authRunAddress: aggregator.authAddress,
	}, nil
}

func (c *AuthConfig) RunAddress() string {
	return c.authRunAddress
}
