package singleton

import "sync"

type singleton map[string]string

var (
    once sync.Once //定义一个执行一次的结构体类型
    instance singleton //定义唯一实例，类型为一个mapping
)

func New() singleton {
    /*
      通过sync/atomic（最轻量级的锁）的原子操作，使得计算机中的任何CPU都不会进行其它的针对此值的读或写操作。这样的约束是受到底层硬件的支持的。
     */
    once.Do(func() {
        instance = make(singleton)
    }) //通过传递一个闭包函数，创建唯一实例
    return instance
}