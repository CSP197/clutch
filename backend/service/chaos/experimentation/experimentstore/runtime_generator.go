package experimentstore

import (
	"go.uber.org/zap"
)

type RuntimeGeneration struct {
	ConfigTypeUrl         string
	GetEnforcingCluster   func(experiment *Experiment, logger *zap.SugaredLogger) (string, error)
	RuntimeKeysGeneration func(experiment *Experiment, runtimePrefixes *RuntimePrefixes, logger *zap.SugaredLogger) ([]*RuntimeKeyValue, error)
}

type RuntimeGenerator struct {
	nameToGenerationMap map[string]*RuntimeGeneration
}

type RuntimePrefixes struct {
	IngressPrefix string
	EgressPrefix  string
	RedisPrefix   string
}

type RuntimeKeyValue struct {
	Key   string
	Value uint32
}

func NewRuntimeGenerator() RuntimeGenerator {
	runtimeGenerator := RuntimeGenerator{map[string]*RuntimeGeneration{}}
	runtimeGenerator.nameToGenerationMap = map[string]*RuntimeGeneration{}
	return runtimeGenerator
}

func (rg *RuntimeGenerator) Register(runtimeGeneration RuntimeGeneration) error {
	rg.nameToGenerationMap[runtimeGeneration.ConfigTypeUrl] = &runtimeGeneration
	return nil
}