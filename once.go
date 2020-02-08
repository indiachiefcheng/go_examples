package main

import (
	"fmt"
	"sync"
)

func main() {

	var count int
	increment := func() {
		count++
	}
	decrement := func() {
		count--
	}

	var once sync.Once

	var increments sync.WaitGroup
	increments.Add(200)
	for i := 0; i < 100; i++ {
		go func() {
			defer increments.Done()
			once.Do(increment)
		}()
	}
	for j := 0; j < 100; j++ {
		go func() {
			defer increments.Done()
			once.Do(decrement)
		}()
	}
	increments.Wait()
	fmt.Printf("Count is %d\n", count)
}

//once.Do()只会执行一次，所以结果为1




pool:


// 实现 io.Closer 接口
type Pool struct{
	m sync.Mutex
	resources chan io.Closer
	factory func()(io.Closer,error)
	closed bool
}

//工厂函数
func New(fn func()(io.Closer,error),size uint)(*Pool,error){
	if size <= 0{
		return nil,errors.New("Size value too small")
	}
	return &Pool{
		factory: fn
		resources: make(chan io.Closer,size)
	},nil
}

//Acquire从池中获取一个资源
func (p *Pool) Acquire()(io.Closer,error){
	select {
		//检查是否有空闲的资源
	case r,ok := <p.resources:
		log.Println("Acquire:", "Shared Resource")
		if !ok{
			return nil, ErrPoolClosed
		}
		return r,nil
	}
	// 因为没有空闲资源可用，所以提供一个新资源
	default:
		log.Println("Acquire:","New Resource")
		return p.factory()	
}

//Close 会让资源池停止工作，并关闭所有现有的资源
func (p *Pool) Close() {
	// 保证本操作与Release 操作的安全
	p.m.Lock()
	defer p.m.Unlock()
	// 如果 pool 已经被关闭，什么也不做
	if p.closed{
		return
	}
	// 将池关闭
	p.closed = true
	// 在清空通道里的资源之前，将通道关闭
	// 如果不这样做，会发生死锁
	close(p.resources)
	// 关闭资源
	for r := range p.resources{
		r.Close()
	}
}

//Release 将一个使用后的资源放回池里
func(p *Pool)Release(r io.Closer){
	// 保证本操作和Close 操作的安全
	p.m.Lock()
	defer p.m.Unlock()

	//如果池子关闭，销毁这个资源,关闭后就不接收新资源了
	if p.closed{
		r.Close()
		return
	}
	select{
		// 试图将这个资源放入队列
	    case p.resources <- r:
	    	log.Println("Release:", "In Queue")
	    //如果队列已满，则关闭这个资源
	    default:
	    	log.Println("Release:", "Closing")
	    	r.Close()
	}
}






