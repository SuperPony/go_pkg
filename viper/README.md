# Index

- 介绍
- 安装
- 使用
  - 设置配置文件
  - 读取键值
  - 设置键值
  - 写入配置
  - 监听配置文件
  - 绑定结构体

# 介绍

viper 是一款很成熟的配置解决方案，拥有丰富的的特性，例如：

- 支持 JSON/TOML/YAML/HCL/envfile/Java properties 等多种格式的配置文件；
- 可以设置监听配置文件的修改，修改时自动加载新的配置；
- 从环境变量、命令行选项和 io.Reader 中读取配置；
- 从远程配置系统中读取和监听修改，如 etcd/Consul；
- 代码逻辑中显示设置键值。

# 安装

- `go get github.com/spf13/viper`

# 使用

## 读取文件

```
// 读取配置
func readConfig() (err error) {
	viper.SetConfigName("config") // 配置文件名,不要带类型
	viper.SetConfigType("yaml")   // 配置文件类型
	viper.AddConfigPath(".")      // 配置文件路径，可以设置多个

	if err = viper.ReadInConfig(); err != nil {
		return
	}
	return nil
}
```

## 读取键值

```
// 读取
func get() {
	// viper.IsSet(key string) bool 返回一个 bool，表示指定的键是否存在
	fmt.Println(viper.IsSet("dev.password"))

	// 获取键值
	// viper 获取键值时，如果指定的 key 不存在，则返回值类型的默认值
	// viper.Get(key string) 获取对应 key 的值
	// viper.GetStringMap(key) map[string]interface{} 返回指定的 key 下面的所有键值对
	// viper.GetStringMapString(key) map[string]string，与 GetStringMap 类似
	// viper.GetType 系列，返回对应的 Type 类型的键值
	// viper.AllSettings() map[string]interface{} // 返回所有设置
}
```

## 设置键值

```
// 设置
func set() {
	// 设置键值，通过 Set 设置的键值，优先级最高
	// viper.Set(key string, value interface{})

	// 通过命令行进行设置，优先级第二
	// pflag.String("demo", "default", "demo")
	// viper.BindPFlag("demo", pflag.Lookup("demo"))
	// pflag.Parse()

	// 通过环境设置环境变量的方式，优先级第三，此处由于还未掌握，暂时空缺

	// 配置文件方式，优先级第四。

	// viper.SetDefault("dev.user", "root") // 设置默认值，优选顺序最低
}
```

## 写入配置

```
// 保存配置
func writeConfig() {
	viper.WriteConfig()                           // 写到设置预定义路径中，如果没有预定义路径，返回错误。将会覆盖当前配置
	viper.SafeWriteConfig()                       // 与 WriteConfig 类似，但是如果配置文件存在，则不覆盖
	viper.WriteConfigAs("./config.yaml.copy")     // 保存配置到指定路径，如果文件存在，则覆盖
	viper.SafeWriteConfigAs("./config.yaml.copy") // 与 WriteConfigAs 类似，但是如果配置文件存在，则不覆盖
}
```

## 监听配置

```
// 监听文件
func watchConfig() {
	if err := readConfig(); err != nil {
		log.Fatalln(err)
	}

	// 设置监听文件，当配置文件被修改，viper 会自重新加载配置
	viper.WatchConfig()

	// 注册修改事件回调
	viper.OnConfigChange(func(in fsnotify.Event) {
		// in.Name 目标文件
		// in.Op.String() 事件类型
		fmt.Println(in.Name, in.Op.String())

		fmt.Println(viper.AllSettings())
	})

	// 此处阻塞程序，用于展示文件监听
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGKILL)
	<-c
	fmt.Println("is over")
}
```

## 绑定结构体

```
// 绑定结构体
func unmarshal() {
	type configDev struct {
		Port string `json:"port,omitempty"`
		Host string `json:"host,omitempty"`
	}
	type config struct {
		Dev configDev
	}
	d := config{}
	readConfig()
	viper.Unmarshal(&d)
	fmt.Println(d.Dev.Port)
}
```
