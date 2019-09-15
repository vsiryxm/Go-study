package object_pool

import (
	"fmt"
	"strconv"
	"sync"
	"testing"
)

func Test1(test *testing.T) {
	p := NewPool(5)
	wait := sync.WaitGroup{}
	for i := 0; i < 100; i++ {
		index := i
		wait.Add(1)
		//这里使用goroutines来并发地读取pool中的对象.
		go func(pool Pool, ind int) {
			select {
			case Obj := <-pool:
				Obj.Do(ind)
				pool <- Obj
			default:
				fmt.Println("No Object:" + strconv.Itoa(ind))
			}
			wait.Done()
		}(*p, index)
	}
	wait.Wait()
}
