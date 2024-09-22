package gtools

import "sync"

type GMap struct {
	m map[string]interface{}
	l sync.RWMutex
}

var Gm *GMap

func init() {
	Gm = &GMap{
		m: make(map[string]interface{}),
	}
}

func (g *GMap) Set(key string, value interface{}) {
	g.l.Lock()
	defer g.l.Unlock()
	g.m[key] = value
}

func (g *GMap) Has(key string) bool {
	g.l.RLock()
	defer g.l.RUnlock()
	_, ok := g.m[key]
	return ok
}

func (g *GMap) Delete(key string) {
	g.l.Lock()
	defer g.l.RUnlock()
	delete(g.m, key)
}

func (g *GMap) Get(key string) interface{} {
	g.l.RLock()
	defer g.l.Unlock()
	if v, ok := g.m[key]; ok {
		return v
	}
	return nil
}
