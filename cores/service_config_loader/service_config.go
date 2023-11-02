package service_config_loader

import (
	"infer-microservices/cores/service_config_loader/faiss_config_loader"
	"infer-microservices/cores/service_config_loader/model_config_loader"
	"infer-microservices/cores/service_config_loader/redis_config_loader"
)

var ServiceConfigs = make(map[string]*ServiceConfig, 0) //one server/dataid,one service conn

type ServiceConfig struct {
	serviceId        string                               `validate:"required,unique,min=4,max=10"` //dataid
	redisConfig      redis_config_loader.RedisConfig      `validate:"required"`                     //redis conn info
	faissIndexConfig faiss_config_loader.FaissIndexConfig //index conn info
	modelConfig      model_config_loader.ModelConfig      `validate:"required"` //model conn info
}

func init() {
}

// serviceId
func (s *ServiceConfig) setServiceId(dataId string) {
	s.serviceId = dataId
}

func (s *ServiceConfig) GetServiceId() string {
	return s.serviceId
}

// redisConfig
func (s *ServiceConfig) setRedisConfig(redisConfig redis_config_loader.RedisConfig) {
	s.redisConfig = redisConfig
}

func (s *ServiceConfig) GetRedisConfig() *redis_config_loader.RedisConfig {
	return &s.redisConfig
}

// faissIndexConfig
func (s *ServiceConfig) setFaissIndexConfig(faissIndexConfig faiss_config_loader.FaissIndexConfig) {
	s.faissIndexConfig = faissIndexConfig
}

func (s *ServiceConfig) GetFaissIndexConfig() *faiss_config_loader.FaissIndexConfig {
	return &s.faissIndexConfig
}

// modelConfig
func (s *ServiceConfig) setModelConfig(modelConfig model_config_loader.ModelConfig) {
	s.modelConfig = modelConfig
}

func (s *ServiceConfig) GetModelConfig() *model_config_loader.ModelConfig {
	return &s.modelConfig
}
