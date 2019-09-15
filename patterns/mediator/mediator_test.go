package mediator

import (
	//"fmt"
	"testing"
)

func TestMediator(test *testing.T) {
	room_1 := NewRoom("金沙绿岛")
	room_2 := NewRoom("大汉龙城")
	consumer_1 := NewConsumer("海阳之新")
	consumer_2 := NewConsumer("凤丫")

	m := NewRoomMediator()
	m.SetRoom(room_1)
	m.SetConsumerHash(consumer_1)
	m.SetConsumerHash(consumer_2)
	m.SetRoom(room_2)
}
