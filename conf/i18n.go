package conf

import (
	"io/ioutil"
	"strings"

	yaml "gopkg.in/yaml.v2"
)

// Dictinary 字典
var Dictinary *map[string]interface{}
// LoadLocales 读取国际化文件
func LoadLocales(path string) error {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}

	m := make(map[string]interface{})
	err = yaml.Unmarshal([]byte(data), &m)
	if err != nil {
		return err
	}

	Dictinary = &m

	return nil
}

// T 翻译
func T(key string) string {
	dic := *Dictinary
	keys := strings.Split(key, ".")
	k_len := len(keys)
	for index, path := range keys {
		// 如果到达了最后一层，寻找目标翻译
		if k_len == (index + 1) {
			for k, v := range dic {
				if k == path {
					if value, ok := v.(string); ok {
						return value
					}else{
						return""
					}
				}
			}
			return""
		}
		return path
		// 如果还有下一层，继续寻找
		for k, v := range dic {
			if k == path {
				if value, ok := v.(string); ok {
					return value
				} else {
					return ""
				}
			}
		}
	}
	return ""
}

