package demo_context

import (
	"context"
	"database/sql"
	"fmt"
	"golang.org/x/sync/errgroup"
	"sync/atomic"
	"testing"
	"time"
)

func TestContext1(t *testing.T) {
	// 1.起点
	ctx := context.Background()
	// todo 表示现在这个context从哪里来，但是我认为将来会有一个
	//todoCtx := context.TODO()
	timeoutCtx, cancel := context.WithTimeout(ctx, time.Second)
	// 紧接着都会写一句下面这句，意思是说我退出的时候会将这个方法的时候取消掉
	// 当然也可以手动取消
	defer cancel()

	// 看下Err的调用效果
	err := timeoutCtx.Err()
	fmt.Println(err)
	switch err {
	case context.Canceled:
	case context.DeadlineExceeded:

	}
}

func TestContextCanceled(t *testing.T) {
	// 1.起点
	ctx := context.Background()
	timeoutCtx, cancel := context.WithTimeout(ctx, time.Second)
	time.Sleep(500 * time.Millisecond)
	cancel()
	err := timeoutCtx.Err()
	fmt.Println(err)
}

func TestContextDeadline(t *testing.T) {
	//ctx := context.Background()
	//timeoutCtx, cancel := context.WithTimeout(ctx, time.Second)
	//defer cancel()
	//
	//dl, ok := timeoutCtx.Deadline()
	//fmt.Println(dl, ok)

	ctx := context.Background()
	dl, ok := ctx.Deadline()
	fmt.Println(dl, ok)
}

func TestContextWithValue(t *testing.T) {
	ctx := context.Background()
	// 不断的进行WithValue就是不断的创建新的context出来
	// 源码里面因要求key还是要可比较的，因为如果是不可比较的也不能设置为map的key
	// 最后返回的是一个新的context
	valCtx := context.WithValue(ctx, "abc", 123)
	val := valCtx.Value("abc")
	fmt.Println(val)
}

// 总结：
// 可以通过err查看是超时还是主动取消

// 通常一般公开方法第一个参数都是context
// 比如和第三方打交道的，比如和数据库，和rpc 和http打交道等
// 即便不需要也需要往下传
func SomeBusiness() {
	ctx := context.TODO()

	// Done还没有结束，
	ctx.Err()
	Step1(ctx)
}

func Step1(ctx context.Context) {
	var db *sql.DB
	db.ExecContext(ctx, "UPDATE xxx", 1)
}

func TestParentCtx(t *testing.T) {
	ctx := context.Background()
	// 设置截止到具体的时间点结束
	dlCtx, cancel := context.WithDeadline(ctx, time.Now().Add(time.Minute))

	childCtx := context.WithValue(dlCtx, "key", 123)

	// err 会受下面的cancel影响
	// 这是因为上级，即父亲级别已经cancel了，所以儿子也会cancel掉
	// 这也是链路控制的原因，上层调用的时候穿个context下层，只要下层处理这个context，上层就会发送个信号给你，这个信号就是取消,
	// 至于取消代表什么含义，这是根据具体业务来决定的
	// cancel只是一个信号，究竟cancel什么，就需要看监听了这个信号的对象要去做什么
	// 怎么知道是哪一个去cancel的呢？答案是不知道，我们是不会也不应该知道是谁在上层哪个在控制下层
	cancel()

	err := childCtx.Err()
	fmt.Println(err) // context canceled
}

func TestParentValueCtx(t *testing.T) {
	ctx := context.Background()
	// 父亲
	childCtx := context.WithValue(ctx, "key1", 123)
	// 儿子
	ccCtx := context.WithValue(childCtx, "key2", 124)
	// 像是一个继承关系，key2是儿子的，儿子不会继承给父亲，所以父亲拿不到儿子的key
	val := childCtx.Value("key2")
	fmt.Println(val)
	// 像是一个继承关系，key1是父亲的，儿子会继承自父亲，所以儿子可以按到父亲的key1
	val = ccCtx.Value("key1")
	fmt.Println(val)
	// 继承是从上到下的顺序
	// 这里就有出现问题了？比如后辈设置了一个值，但是上层祖辈想拿到设置的值，需要怎么设置呢？需要使用一些奇怪的方法，就是放一个map进去
}

// 父辈想拿到后辈设置的值-map
// 这种方式不到逼不得已不要用，比较危险，因为已经改变了context不可变性的约束，因为map是可变的，map的引用虽然没有变，但是map的内容也就是map存储的内容是可以被改变的
// 另一个危险的就是在Map里面存储指针
// 如果是跨服务，就只能通过网络协议，比如http的header来传递context
func TestParentValueMapCtx(t *testing.T) {
	ctx := context.Background()
	// 父亲
	childCtx := context.WithValue(ctx, "map", map[string]string{})
	ccChild := context.WithValue(childCtx, "key1", "value1")
	// 儿子
	// 这边要做下类型断言
	m := ccChild.Value("map").(map[string]string)
	m["key1"] = "val1"

	val := childCtx.Value("key1")
	fmt.Println(val)

	// 像是一个继承关系，key1是父亲的，儿子会继承自父亲，所以儿子可以按到父亲的key1
	val = childCtx.Value("map")
	fmt.Println(val)
}

// WithValue 典型的装饰器模式，已经有了一个实现，再原来的基础上装饰以下，在上面加点功能，装饰器强调的在已有的基础上增加新的额外功能

// 子context不会重置父context的超时时间
func TestConttext_timeout(t *testing.T) {
	bg := context.Background()
	//timeoutCtx, cancel1 := context.WithTimeout(bg, time.Second)
	timeoutCtx, _ := context.WithTimeout(bg, time.Second)

	//subCtx, cancel2 := context.WithTimeout(timeoutCtx, 3*time.Second)
	subCtx, _ := context.WithTimeout(timeoutCtx, 3*time.Second)

	go func() {
		// 一秒之后就会过期，然后输出 timeout
		<-subCtx.Done()
		fmt.Printf("timeout\n")
	}()

	time.Sleep(2 * time.Second)
	//cancel2()
	//cancel1()
}

// Example: 控制业务超时
func TestBusinessTimeout(t *testing.T) {
	// 业务睡2s，但是这里想值1s内就返回
	ctx := context.Background()
	timeoutCtx, cancel := context.WithTimeout(ctx, 1*time.Second)
	defer cancel()

	// 控制超时
	end := make(chan struct{}, 1)
	go func() {
		MyBusiness()
		end <- struct{}{}
	}()

	ch := timeoutCtx.Done()
	//<-ch
	select {
	case <-ch:
		fmt.Println("timeout")
	case <-end:
		fmt.Println("business end")
	}
}

func MyBusiness() {
	//time.Sleep(2 * time.Second)
	time.Sleep(500 * time.Millisecond)
	fmt.Println("hello world")
}

// 不建议在结构体里面使用context，因为context是线程安全的，就不应该放到结构体里面,
// 但是还是可以看到有些地方放到结构体里面了，比如http.Request
// Example: 修改context
func TestRequest(t *testing.T) {
	//http.NewRequestWithContext()

	//req := http.Request{}
	// 返回了一个新的context
	//req = req.WithContext()
}

func TestErrGroup(t *testing.T) {
	eg := errgroup.Group{}
	var result int64 = 0
	for i := 0; i < 10; i++ {
		delta := i
		// 不断的开goroutine去创建任务去执行
		eg.Go(func() error {
			atomic.AddInt64(&result, int64(delta))
			return nil
		})
	}
	// 执行完需要等下，等所有的任务都执行完才再往后去执行
	// 停下来等所有的任务去执行，如果任务一个任务返回了err，这里都能拿到err
	if err := eg.Wait(); err != nil {
		t.Fatal(err)
	}
	fmt.Println(result)
}

func TestErrGroup2(t *testing.T) {
	eg, ctx := errgroup.WithContext(context.Background())
	var result int64 = 0
	for i := 0; i < 10; i++ {
		delta := i
		// 不断的开goroutine去创建任务去执行
		eg.Go(func() error {
			atomic.AddInt64(&result, int64(delta))
			return nil
		})
	}
	// 执行完需要等下，等所有的任务都执行完才再往后去执行
	// 停下来等所有的任务去执行，如果任务一个任务返回了err，这里都能拿到err
	if err := eg.Wait(); err != nil {
		t.Fatal(err)
	}
	ctx.Err()
	fmt.Println(result)
}
