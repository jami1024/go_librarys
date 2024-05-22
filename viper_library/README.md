## Go项目开发中如何读取应用配置？
> 本文引用于[孔令飞的云原生实战营](https://www.yuque.com/konglingfei-vzag4/onex/bm6ms4fn18va6ii3)

Viper 是 Go 语言中一个强大的配置文件管理库，它能够处理多种配置文件格式（如 JSON、YAML、TOML 等），并且支持从环境变量、命令行参数、配置文件等不同来源读取配置信息。以下是 Viper 的一些基本用法：

### 1. 安装 Viper
首先，你需要使用 `go get` 命令安装 Viper：
```bash
go get github.com/spf13/viper
```

### 2. 初始化 Viper 实例
创建一个新的 Viper 实例，通常在应用程序的初始化阶段进行：
```go
import "github.com/spf13/viper"

func main() {
    v := viper.New()
}
```

### 3. 设置配置文件
指定配置文件的路径和名称：
```go
v.SetConfigFile("config.yaml")
```

### 4. 读取配置文件
Viper 支持多种格式的配置文件，如 JSON、YAML、TOML 等。你可以指定配置文件的格式：
```go
v.SetConfigType("yaml")
```

### 5. 自动加载配置
使用 `ReadConfig` 方法自动加载配置文件：
```go
err := v.ReadConfig()
if err != nil {
    log.Fatal(err)
}
```

### 6. 读取配置项
读取配置文件中的值，Viper 允许你使用点号（`.`）来访问嵌套的配置项：
```go
dbHost := v.GetString("database.host")
dbPort := v.GetInt("database.port")
```

### 7. 设置默认值
如果配置文件中没有找到某个配置项，可以使用 `SetDefault` 方法设置默认值：
```go
v.SetDefault("database.host", "localhost")
v.SetDefault("database.port", 5432)
```

### 8. 从环境变量读取配置
Viper 可以自动从环境变量中读取配置，只需将配置项的键设置为环境变量的键：
```go
v.AutomaticEnv()
dbHost := v.GetString("DB_HOST")
```

### 9. 从命令行参数读取配置
Viper 也支持从命令行参数读取配置，使用 `BindPFlag` 方法将配置项与命令行参数绑定：
```go
import "github.com/spf13/pflag"

func main() {
    ...
    var dbPort int
    pflag.IntVar(&dbPort, "dbPort", 5432, "database port")
    ...
    v.BindPFlag("database.port", pflag.Lookup("dbPort"))
    ...
}
```

### 10. 监听配置文件变化
Viper 可以监听配置文件的变化，并重新加载配置：
```go
v.WatchConfig()
v.OnConfigChange(func(e fsnotify.Event) {
    log.Println("Config file changed:", e.Name)
})
```

### 11. 封装配置结构体
通常，我们会定义一个结构体来封装配置项，然后使用 Viper 将其自动映射到结构体中：
```go
type Config struct {
    Database struct {
        Host     string
        Port     int
        Username string
        Password string
    } `mapstructure:"database"`
}

var cfg Config

err := v.Unmarshal(&cfg)
if err != nil {
    log.Fatal(err)
}
```

### 12. 使用第三方库解析配置文件
Viper 支持使用第三方库来解析配置文件，如 `mapstructure`：
```go
import "github.com/spf13/viper"
import "github.com/spf13/cast"

err := v.UnmarshalExact(&cfg)
if err != nil {
    log.Fatal(err)
}
```

以上是 Viper 的一些基本用法，通过这些方法，你可以方便地管理和使用你的应用程序配置。