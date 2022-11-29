package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"
)

type Configuration interface {
	LoadConfigurationFromFile(fp string)
	GetValue(key string) string
}

type configuration struct {
	conf     map[string]interface{}
	filepath string
}

func NewConfiguration() *configuration {
	return &configuration{}
}

func (c *configuration) LoadConfigurationFromFile(fp string) {
	c.filepath = fp
	plan, _ := ioutil.ReadFile(c.filepath)
	jsonMap := make(map[string]interface{})
	err := json.Unmarshal([]byte(plan), &jsonMap)
	if err != nil {
		panic(err)
	}
	c.conf = jsonMap
}

func (c *configuration) GetValue(key string) string {
	return getValue(key, c.conf)
}
func getValue(key string, m map[string]interface{}) string {
	keys := strings.Split(key, ".")
	for k, v := range m {
		if keys[0] == k {
			if mv, ok := v.(map[string]interface{}); ok {
				return getValue(strings.Join(keys[1:], "."), mv)
			} else {
				if len(keys) == 1 {
					return fmt.Sprintf("%v", v)
				}
			}
		}
	}
	return ""
}
