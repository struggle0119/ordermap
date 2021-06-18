package ordermap

import (
	"container/list"
	"sync"
)

type OrderMap struct {
	mu sync.Mutex
	kv map[interface{}]*list.Element
	l  *list.List
}

func New() *OrderMap {
	return &OrderMap{
		kv: map[interface{}]*list.Element{},
		l:  list.New(),
	}
}

// add value into map and list, support cover
func (om *OrderMap) Add(key, value interface{}) (exist bool) {
	defer om.mu.Unlock()
	om.mu.Lock()
	ele := newElement(key, value)
	if val, ok := om.kv[key]; ok {
		val.Value.(*Element).value = value
		exist = true
	} else {
		// key not exist then put it in
		element := om.l.PushBack(ele)
		om.kv[key] = element
	}
	return
}

// get value by key, return value and exist or not
func (om *OrderMap) Get(key interface{}) (value interface{}, exist bool) {
	if val, ok := om.kv[key]; ok {
		value = val.Value.(*Element).value
		exist = true
	}
	return
}

// delete value from map and list
func (om *OrderMap) Del(key interface{}) {
	defer om.mu.Unlock()
	om.mu.Lock()
	if ele, ok := om.kv[key]; ok {
		delete(om.kv, key)
		om.l.Remove(ele)
	}
}

// ordered keys
func (om *OrderMap) Keys() (keys []interface{}) {
	ele := om.l.Front()
	for i := 0; ele != nil; i++ {
		keys = append(keys, ele.Value.(*Element).key)
		ele = ele.Next()
	}
	return
}
