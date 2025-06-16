package env

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"

	"github.com/Makovey/fakelist_utils/pkg/logger"
)

const (
	authRunAddress = "AUTH_RUN_ADDRESS"
)

type resolver struct {
	log logger.Logger
}

type envAggregator struct {
	authAddress string
}

func newEnvAggregator(log logger.Logger) (*envAggregator, error) {
	fn := "env.newEnvConfig"

	if err := godotenv.Load(); err != nil {
		return nil, fmt.Errorf("[%s]: %s", fn, err.Error())
	}

	r := &resolver{log: log}

	authAddress, err := r.resolveVariable(authRunAddress)
	if err != nil {
		return nil, fmt.Errorf("[%s]: %w", fn, err)
	}

	return &envAggregator{
		authAddress: authAddress,
	}, nil
}

func (r *resolver) resolveVariable(name string) (string, error) {
	fn := "env.resolveVariable"

	value := os.Getenv(name)
	if value != "" {
		return value, nil
	}

	return "", fmt.Errorf("[%s]: required variable - %s is empty", fn, name)
}

func (r *resolver) resolveVariableWithDefault(name, defaultValue string) (string, error) {
	fn := "env.resolveVariable"

	value := os.Getenv(name)
	if value != "" {
		return value, nil
	}

	if defaultValue != "" {
		r.log.Warnf("[%s]: %s variable is empty, but default value is %s", fn, name, defaultValue)
		return defaultValue, nil
	}

	return "", fmt.Errorf("[%s]: required variable - %s is empty", fn, name)
}
