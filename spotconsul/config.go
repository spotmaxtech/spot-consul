package spotconsul

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

type Config struct {
	Global *GlobalConfig  `json:"global"`
	Logic  []*LogicConfig `json:"logic"`
}

type GlobalConfig struct {
	LoopingTimeS     int64 `json:"loopingTimeS"`
	LingerTimeS      int64 `json:"lingerTimeS"`
	FreshLearningStart bool  `json:"freshLearningStart"`
}

type LogicConfig struct {
	ConsulAddr        string `json:"consulAddr"`
	InstanceLoadKey   string `json:"instanceLoadKey"`
	LearningFactorKey string `json:"learningFactorKey"`
	OnlineLabKey      string `json:"onlineLabKey"`
	ServiceName       string `json:"serviceName"`
	ZoneCPUKey        string `json:"zoneCPUKey"`
}

func NewConfig(filepath string) *Config {
	jsonFile, err := os.Open(filepath)
	if err != nil {
		panic("load config file " + filepath + " error, " + err.Error())
	}
	bytes, _ := ioutil.ReadAll(jsonFile)
	var config Config
	if err := json.Unmarshal(bytes, &config); err != nil {
		panic("load config file " + filepath + " error, " + err.Error())
	}

	return &config
}
