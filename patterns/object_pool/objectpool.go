package object_pool

import (
"fmt"
"strconv"
)

type Object struct{

}

func (Object)Do(index int){
    fmt.Println("Object Do:"+strconv.Itoa(index))
}

type Pool chan *Object

/*
 对象池模式是一种创建型模式，根据需求来预测将要使用的对象，提前创建并保存在内存中。
 */
func NewPool(total int)*Pool{
    p := make(Pool,total)
    for i := 0;i<total;i++{
        p <- new(Object)
    }
    return &p
}

/*
  使用对象池的场景条件：
  1、当创建对象的代价比维护代价更高的时候，使用对象池模式是极好的。
  2、如果需求相对固定，那么维护对象的代价可能得不偿失
  3、提前初始化对象对性能有积极的影响
 */

