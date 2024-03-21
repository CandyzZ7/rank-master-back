package event_bus

import (
	"fmt"
	"testing"
)

type service struct {
}

func (s *service) Identity() string {
	return "service"
} // 身份标识

func (s *service) EventCallBack(data DataEvent) error {
	fmt.Println("我是service, 收到了消息", data.Data)
	return nil
}

func TestEventBus(t *testing.T) {
	bus := NewEventBus()
	bus.Subscribe("topic1", &service{})
	bus.Publish("topic1", true, "hello")

}
