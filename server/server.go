package server

import (
	"github.com/Lyo-Shur/agentutils/config"
	"github.com/Lyo-Shur/agentutils/task"
	"github.com/Lyo-Shur/golog"
	"io"
	"math/rand"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"
	"time"
)

// 启动服务
func StartServer(conf config.Config, tasks []task.Task) {
	// 日志服务
	logger := golog.GetLogger()
	logger.AddHandler(golog.GetPrintHandler())
	// 判断是否启用日志
	if conf.Log.Open {
		logger.AddHandler(golog.GetFileHandler(conf.Log.Path))
	}
	// 启动服务
	addr := "0.0.0.0:" + conf.Server.Port
	logger.Info("agent starts at " + addr)
	http.HandleFunc("/", getHandler(conf, tasks, logger))
	err := http.ListenAndServe(addr, nil)
	if err != nil {
		logger.Error(err.Error())
	}
}

// 获取请求处理程序
func getHandler(conf config.Config, tasks []task.Task, logger *golog.Logger) func(w http.ResponseWriter, r *http.Request) {
	// 请求处理器
	return func(w http.ResponseWriter, r *http.Request) {
		// 优先处理前置请求
		for i := range tasks {
			if !tasks[i].Do(w, r) {
				return
			}
		}
		// 打印请求相关内容
		requestLogger := GetRequestLogger(logger, r)
		requestLogger.Log()

		// 请求路径
		prefix := r.URL.Path
		parameter := r.URL.RawQuery
		var proxy config.Proxy
		for index := range conf.Proxies {
			item := conf.Proxies[index]
			if strings.HasPrefix(prefix, item.Path) {
				proxy = item
			}
		}
		// 如果没有匹配的规则
		if proxy.Path == "" {
			_, err := io.WriteString(w, "unknown service")
			if err != nil {
				logger.Error(err.Error())
			}
			return
		}

		// 代理转向的地址
		address := randomOne(proxy.Urls)

		// 代理服务器地址
		remote, err := url.Parse(address)
		if err != nil {
			logger.Error(err.Error())
		}

		// 修改请求地址与HOST
		path := strings.Replace(prefix, proxy.Path, "", 1)
		r.RequestURI = path
		if parameter != "" {
			r.RequestURI = path + "?" + parameter
		}
		r.URL.Path = path
		r.Host = remote.Host
		r.URL.Host = r.Host

		// 打印代理后内容
		requestLogger.Log()
		// 执行代理
		reverseProxy := httputil.NewSingleHostReverseProxy(remote)
		reverseProxy.ServeHTTP(w, r)
	}
}

// 从数组中随机返回一个
func randomOne(list []string) string {
	length := len(list)
	if length == 0 {
		return ""
	}
	if length == 1 {
		return list[0]
	}
	// 设置随机数种子
	rand.Seed(time.Now().Unix())
	// 使用余数来扩大范围 30为任意数
	index := rand.Intn(length*30) % length
	return list[index]
}
