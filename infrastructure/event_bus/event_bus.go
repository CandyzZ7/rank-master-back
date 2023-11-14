package event_bus

import (
	"sync"
)

type DataEvent struct {
	Data  interface{}
	Topic Topic
}

// Topic 主题类型
type Topic string

// ISubscriber 订阅者接口
type ISubscriber interface {
	Identity() string                   // 身份标识
	EventCallBack(data DataEvent) error // 事件回调
}

type subscriberMap map[string]ISubscriber

// EventBus 存储有关订阅者感兴趣的特定主题的信息
type EventBus struct {
	subscriberMap map[Topic]subscriberMap
	rm            sync.RWMutex
}

type IEventBus interface {
	Subscribe(topic Topic, sub ISubscriber)             // 订阅
	Unsubscribe(topic Topic, sub ISubscriber)           // 取消订阅
	Publish(topic Topic, isSync bool, data interface{}) // 发布消息
}

// NewEventBus 创建一个新的事件总线
func NewEventBus() IEventBus {
	return &EventBus{
		subscriberMap: make(map[Topic]subscriberMap),
		rm:            sync.RWMutex{},
	}
}

// Subscribe 订阅
func (eb *EventBus) Subscribe(topic Topic, sub ISubscriber) {
	eb.rm.Lock()
	defer eb.rm.Unlock()
	if subMap, found := eb.subscriberMap[topic]; found {
		subMap[sub.Identity()] = sub
	} else {
		eb.subscriberMap[topic] = subscriberMap{sub.Identity(): sub}
	}
}

// Unsubscribe 取消订阅
func (eb *EventBus) Unsubscribe(topic Topic, sub ISubscriber) {
	eb.rm.Lock()
	defer eb.rm.Unlock()
	if subMap, found := eb.subscriberMap[topic]; found {
		delete(subMap, sub.Identity())
		// 如果主题的订阅者列表为空，删除该主题
		if len(subMap) == 0 {
			delete(eb.subscriberMap, topic)
		}
	}
}

// Publish 发布消息
func (eb *EventBus) Publish(topic Topic, isSync bool, data interface{}) {
	if data == nil {
		return
	}
	eb.rm.RLock()
	defer eb.rm.RUnlock()
	subscriberMap, found := eb.subscriberMap[topic]

	if found {
		event := DataEvent{Data: data, Topic: topic}
		for _, sub := range subscriberMap {
			if isSync {
				err := sub.EventCallBack(event)
				if err != nil {
					panic(err)
				}
			} else {
				go func(subscriber ISubscriber) {
					err := subscriber.EventCallBack(event)
					if err != nil {
						panic(err)
					}
				}(sub)
			}
		}
	}
}
