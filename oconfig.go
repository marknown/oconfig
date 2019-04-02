// Package oconfig is own config
package oconfig

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"sync"
)

// 包内变量，存储实例相关对象
var packageOnce = map[string]*sync.Once{}
var packageInstance = map[string]interface{}{}
var packageMutex = &sync.Mutex{}

// getConfigWithError 获取所有的配置信息
func getConfigWithError(configPath string, configResult interface{}) error {
	// configPath := "config.json"
	// if flag.Lookup("test.v") != nil {
	// 	//flag不为空,则说明存在test所拥有的参数,是在 go test 模式
	// 	configPath = "../" + configPath
	// }
	bytes, err := ioutil.ReadFile(configPath)

	if nil != err {
		return err
	}

	// log.Println("Get config form file" + configPath)
	err = json.Unmarshal(bytes, configResult)
	if nil != err {
		return err
	}

	return err
}

// GetConfig 获取所有的配置信息，如果有问题直接 panic 只初始化一次
func GetConfig(configPath string, configResult interface{}) {
	packageMutex.Lock()
	defer packageMutex.Unlock()

	md5byte := md5.Sum([]byte(configPath))
	md5key := fmt.Sprintf("%x", md5byte)

	// 如果有值直接返回
	if v, ok := packageInstance[md5key]; ok {
		// fmt.Println(reflect.TypeOf(configResult))
		btyes, _ := json.Marshal(v)
		json.Unmarshal(btyes, configResult)
		// todo 不要用 unmarshal 来处理
		// configResult = v
		// fmt.Printf("cache %s %v %v\n", md5key, v, configResult)
		// fmt.Printf("cached %s\n", md5key)
		return
	}

	// 如果once 不存在
	if _, ok := packageOnce[md5key]; !ok {
		var once = &sync.Once{}
		once.Do(func() {
			err := getConfigWithError(configPath, configResult)
			if nil != err {
				panic(configPath + " " + err.Error())
			}

			packageInstance[md5key] = configResult
			packageOnce[md5key] = once

			// fmt.Printf("init %s %v\n", md5key, packageInstance[md5key])
			// fmt.Printf("init %s\n", md5key)
		})
	}
}
