package config

import "github.com/lyoshur/agentutils/task"

// 配置文件信息
type Config struct {
	Server  Server  `xml:"server"`
	Log     Log     `xml:"log"`
	Proxies []Proxy `xml:"proxy"`
}

// 服务器配置
type Server struct {
	// 服务器端口
	Port string `xml:"port,attr"`
}

// 日志配置
type Log struct {
	// 是否开启日志
	Open bool `xml:"open,attr"`
	// 日志文件所在目类
	Path string `xml:"path,attr"`
}

// 反向代理配置
type Proxy struct {
	// 请求host
	Host string `xml:"host,attr"`
	// 代理路径
	Path string `xml:"path,attr"`
	// 前置处理任务
	Tasks []task.Task
	// 代理列表
	Urls []string `xml:"url"`
}
