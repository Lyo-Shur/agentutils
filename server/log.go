package server

import (
	"bytes"
	"github.com/lyoshur/golog"
	"io/ioutil"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

// 请求日志打印工具
type RequestLogger struct {
	logger  *golog.Logger
	request *http.Request
	id      int
}

// 获取请求日志打印工具
func GetRequestLogger(logger *golog.Logger, request *http.Request) *RequestLogger {
	rl := RequestLogger{}
	rl.logger = logger
	rl.request = request
	rl.id = randomNumber()
	return &rl
}

// 打印日志
func (r *RequestLogger) Log() {
	var buffer bytes.Buffer
	buffer.WriteString("[id:" + strconv.Itoa(r.id) + "] [" + r.request.Method + "] " + r.request.Host + r.request.RequestURI + "\n")
	buffer.WriteString("----" + "UserAgent:" + r.request.UserAgent() + "\n")
	buffer.WriteString("----" + "Referer:" + r.request.Referer() + "\n")
	bs, err := ioutil.ReadAll(r.request.Body)
	if err != nil {
		buffer.WriteString("----" + err.Error() + "\n")
	}
	if len(bs) != 0 {
		buffer.WriteString("----" + "Body:" + string(bs) + "\n")
	}
	r.logger.Info(buffer.String())
}

// 产生一个随机数
func randomNumber() int {
	rand.Seed(time.Now().Unix())
	return rand.Intn(99999999)
}
