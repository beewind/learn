package geecache

import (
	"example/geecache/singleflight"
	"fmt"
	"sync"
)

type Group struct {
	name      string
	getter    Getter
	mainCache cache
	peers     PeerPicker
	loader *singleflight.Group
}
type Getter interface {
	Get(key string) ([]byte, error)
}
type GetterFunc func(key string) ([]byte, error)

func (f GetterFunc) Get(key string) ([]byte, error) {
	return f(key)
}

var (
	mu     sync.RWMutex
	groups = make(map[string]*Group)
)

func NewGroup(name string, cacheBytes int64, getter Getter) *Group {
	// 1.getter == nil
	if getter == nil {
		panic("nil Getter")
	}
	// 2.new cache
	mu.Lock()
	defer mu.Unlock()
	g := &Group{
		name:      name,
		getter:    getter,
		mainCache: cache{cacheBytes: cacheBytes},
		loader: &
	}

	// 3.add cache to groups
	groups[name] = g

	// 4.return g
	return g

}

func GetGroup(name string) *Group {
	mu.RLock()
	defer mu.RUnlock()
	g := groups[name]
	return g
}

// get value for a key from cache
func (g *Group) Get(key string) (ByteView, error) {
	//1.key == nil
	if key == "" {
		return ByteView{}, fmt.Errorf("key is required")
	}
	//2.ger kv from maincache
	if v, ok := g.mainCache.get(key); ok {
		return v, nil
	}

	//3.maincache 中没有key
	return g.load(key)
}
func (g *Group) load(key string) (value ByteView, err error) {
	return g.getLocally(key)
}

func (g *Group) RegisterPeers(peers PeerPicker) {
	if g.peers != nil {
		panic("RegisterPeerPicker called more than once")
	}
}
func (g *Group) getLocally(key string) (value ByteView, err error) {
	//1.使用getter获取kv
	bytes, err := g.getter.Get(key)

	//2.nil判断
	if err != nil {
		return ByteView{}, err
	}
	//3.存放热点值
	value = ByteView{b: bytes}
	//g.mainCache.add(key,value)-->抽象为函数
	g.populateCache(key, value)
	//4.返回
	return value, nil
}
func (g *Group) populateCache(key string, value ByteView) {
	g.mainCache.add(key, value)
}
