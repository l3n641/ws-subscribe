package webSocket

import (
	"github.com/olahol/melody"
	"sync"
)

var WebSocketRouters *Routers

func init() {
	WebSocketRouters = newRouters()

}

func newRouters() *Routers {
	return &Routers{
		routers: make(map[string]Handler),
	}
}

type Handler func(*melody.Session, []byte) (interface{}, error)

// Routers 路由列表管理结构体(程序启动后就初始化完成,运行中不需要添加)
type Routers struct {
	lock    sync.RWMutex
	routers map[string]Handler
}

func (a *Routers) AddRouter(action string, f Handler) {
	a.lock.Lock()
	defer a.lock.Unlock()
	a.routers[action] = f
}

// GetRouter 获取路由
func (a *Routers) GetRouter(action string) (Handler, bool) {
	a.lock.RLock()
	defer a.lock.RUnlock()
	value, has := a.routers[action]
	return value, has
}
