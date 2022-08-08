package main

import (
	"fmt"
	"sync"
)

// 锁-资源竞争 能不用锁就不用锁，因为性能问题，加锁和解锁都需要耗时间的，当并发量过大，不停加锁解锁性能会受损，会导致并发下降
// 对于大部分web系统都是读多写少
// 有1w人同时读取数据库，现在加了把锁，A读的时候B可以读吗？就不能了，因为互斥锁

/**
一个+1，一个-1，理论上最后结果是0
实际情况值不是0，每次运行结果不一样
理论上步骤是：
1. add 获取total=0;
2. add +1操作;
3. add 将结果1放到total中，total=1；

4. sub 获取 total=1;
5. sub -1操作；
6. sub 将结果放到total中，total=0;

但是实际上因为两个并发运行，可能出现
add先获取到total=0，+1操作为total=1; 然后轮到 sub函数运行，它拿到的total=0，然后-1， total重新设置成了-1
还有可能两个取出来都是0,sub先运行，total变成-1，然后轮到 add,因为取出来的0，加1后total就重置为1了

出现这种情况的原因是+1,-1并发运行，两个取数的时候和一个还没有执行完就取数了，想要解决这种问题，要做到一个取数的时候另外的一个不能取数
需要一个机制，我运行你不能运行，你运行我不能运行
执行的时候先锁上，执行完之后释放，同一时间点只能有一个上锁，另外没有上锁的就要等待
*/
var total int

var wg sync.WaitGroup

// 互斥锁
var lock sync.Mutex

func add() {
	defer wg.Done()
	for i := 0; i < 100000; i++ {
		//total += 1
		// 1. 取出 total 的值
		// 2. total + 1
		// 3. 将值放入 total
		lock.Lock()
		total = total + 1
		lock.Unlock()
	}
}

func sub() {
	defer wg.Done()
	for i := 0; i < 100000; i++ {
		//total -= 1
		lock.Lock()
		total = total - 1
		lock.Unlock()
	}
}

func main1() {
	wg.Add(1)
	go add()
	wg.Wait()
	fmt.Println(total)
}

// 按理说最后的结果应该是0，因为一个加一个减
// 但实际结果却不是0，而且每次运行结果还不一样
func main() {
	wg.Add(2)
	go add()
	go sub()
	wg.Wait()
	fmt.Println(total)
}

func main3() {
	wg.Add(1)
	go add()
	wg.Wait()
	fmt.Println(total)

}
