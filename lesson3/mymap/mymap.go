package main

import (
	"reflect"
	"sync"
)

var tm sync.Map

type MyMap struct {
	dirty map[interface{}]interface{}
	sync.RWMutex
}

func NewMap() *MyMap {
	return &MyMap{
		dirty:   make(map[interface{}]interface{}),
		RWMutex: sync.RWMutex{},
	}
}

func (m *MyMap) Load(key interface{}) (interface{}, bool) {
	if reflect.TypeOf(key).Comparable() {
		return nil, false
	}
	m.RLock()
	defer m.RUnlock()
	val, ok := m.dirty[key]
	return val, ok
}

func (m *MyMap) Store(key, val interface{}) {
	if reflect.TypeOf(key).Comparable() {
		return
	}

	m.Lock()
	defer m.Unlock()
	m.dirty[key] = val
}

func (m *MyMap) Delete(key interface{}) {
	if reflect.TypeOf(key).Comparable() {
		return
	}

	m.Lock()
	defer m.Unlock()
	delete(m.dirty, key)
}

func (m *MyMap) LoadAndDelete(key interface{}) (interface{}, bool) {
	if reflect.TypeOf(key).Comparable() {
		return nil, false
	}

	m.Lock()
	defer m.Unlock()

	val, ok := m.dirty[key]
	delete(m.dirty, key)
	return val, ok
}

// double check
func (m *MyMap) LoadOrStore(key, val interface{}) (interface{}, bool) {
	if reflect.TypeOf(key).Comparable() {
		return nil, false
	}

	m.RLock()
	actual, ok := m.dirty[key]
	m.RUnlock()
	if ok {
		return actual, true
	}

	m.Lock()
	defer m.Unlock()
	actual, ok = m.dirty[key]
	if ok {
		return actual, true
	}

	m.dirty[key] = val
	return val, false
}
