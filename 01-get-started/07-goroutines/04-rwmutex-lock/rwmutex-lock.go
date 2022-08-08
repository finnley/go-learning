package main

import (
	"fmt"
	"sync"
	"time"
)

/**
锁-资源竞争 能不用锁就不用锁，因为性能问题，加锁和解锁都需要耗时间的，当并发量过大，不停加锁解锁性能会受损，会导致并发下降
对于大部分web系统都是读多写少
有1w人同时读取数据库，现在加了把锁，A读的时候B可以读吗？就不能了，因为互斥锁
为什么要加锁? 为了数据的同步性，读和写都要加锁，而且加的还是同一把锁，才能保证读的时候不至于读错数据，如果出现写的时候你来读，就会出现问题的，可能会读到修改之前的数据的

比如原本 商品 100，现在涨价了，涨到了1w，在改成1w的期间读出来的是 100，然后立马下单了，但是支付的时候为了严谨性会重新检查价格却是1w
所以为了保证数据一致性要加锁，在写和读上加同一把锁

在读和读之间是没有产生影响的，B读了数据并不会对C读取数据造成影响，如果有影响，一定是写和读之间造成的

要做到读之间不会产生影响，写和读之间才会产生影响，也就是我在写的时候你不能读，其他人也不能写，这就是读写锁，使用读写锁提高并发
*/
var wg sync.WaitGroup

// 读写锁
var rwLock sync.RWMutex

func read() {
	defer wg.Done()
	rwLock.RLock()
	fmt.Println("开始读取数据")
	time.Sleep(time.Second)
	fmt.Println("成功读取")
	rwLock.RUnlock()
}

// 默认的是写锁
func write() {
	defer wg.Done()
	rwLock.Lock()
	fmt.Println("开始修改数据")
	time.Sleep(time.Second * 5)
	fmt.Println("成功修改")
	rwLock.Unlock()
}

func main() {
	wg.Add(6)
	// 只启动读并不会产生影响，不会出现一个在读，其他的就不能读
	for i := 0; i < 5; i++ {
		go read()
	}
	for i := 0; i < 1; i++ {
		go write()
	}
	wg.Wait()
}
