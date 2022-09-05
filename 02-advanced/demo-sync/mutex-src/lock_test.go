package mutex_src

// 锁本身就是个状态机，就是加锁到解锁再到加锁再到解锁，就状态在变来变去，对我们来说就两个状态，但在内部可能不止两个状态
// 自旋：有个计数器，有个for循环，里面有个CAS操作，这个过程就是自旋，比如CAS希望当前状态是 unlock状态，然后将它的状态改为 lock，
// 如果能一次性操作成功，就是自旋成功，如果失败了就继续尝试CAS操作,
// 不管拿到锁还没有么有拿到，最终都会退出，如果之前操作成功拿到了则return退出，没有就会加入到阻塞队列等待，等待唤醒

type Lock struct {
	state int
}

// compare and swap
func (l *Lock) CAS(oldValue int, newValue int) {
	if l.state == oldValue {
		l.state = newValue
	}
	// 简单的CAS就是这样，但是上面这不是线程安全的 ，这里的不是依赖go语言的实现而是依赖运行时实现，比如CPU架构之类的
	//atomic.CompareAndSwapInt32()
}

/**
面试：什么时候CPU会百分百，加锁自旋时，因为自旋时计算密集型应用
*/
