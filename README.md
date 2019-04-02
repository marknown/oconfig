# oconfig

## useage

```
package jdconfig

import (
    "github.com/marknown/oconfig"

    "fmt"
    "sync"
)

// MailConfig mail 的配置信息
type MailConfig struct {
    Host      string // 邮件主机
    SSL       bool   // 是否 SSL 加密
    Port      int    // 主机端口
    FromEmail string // 邮件发送者地址
    FromName  string // 邮件发送者名称
    Username  string // 邮件发送者登录名
    Password  string // 邮件发送者登录密码
    Charset   string // 邮件字符集
    Weight    int    // 有多个邮件配置时，本邮件的权重
}

// Config 所有的配置信息
type Config struct {
    Mail  []MailConfig
}

var once = &sync.Once{}
var lock = &sync.Mutex{}
var packageConfigInstance *Config

// GetConfig 获取所有的配置信息 只初始化一次
func GetConfig() *Config {
    lock.Lock()
    defer lock.Unlock()

    if nil != packageConfigInstance {
        return packageConfigInstance
    }

    once.Do(func() {
        packageConfigInstance = &Config{}
        configPath := "./config.json"
        oconfig.GetConfig(configPath, packageConfigInstance)
    })

    return packageConfigInstance
}
```
