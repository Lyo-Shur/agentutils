package agentutils

import (
	"github.com/lyoshur/agentutils/config"
	"github.com/lyoshur/agentutils/server"
	"github.com/lyoshur/agentutils/task"
)

// 配置相关
type AgentConfig = config.Config
type AgentServer = config.Server
type AgentLog = config.Log
type AgentProxy = config.Proxy

// 前置任务
type Task = task.Task

// 启动服务
func StartServer(conf AgentConfig, tasks []Task) {
	server.StartServer(conf, tasks)
}
