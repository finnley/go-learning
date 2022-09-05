package array_list

import "sync"

// 装饰器模式，在原来的功能不断叠加
type SafeList[T any] struct {
	List[T]
	lock sync.RWMutex
}

// 装饰器模式：相当于将方法手动重复一遍
func (s *SafeList[T]) Get(index int) (T, error) {
	// 为什么这里要添加读锁呢？
	// 因为其他地方有写操作呀，为了互斥，我读的时候别人不能写，我写的时候别人不能读
	s.lock.RLock()
	defer s.lock.RUnlock()
	return s.List.Get(index)
}

func (s *SafeList[T]) Append(t T) error {
	s.lock.Lock()
	defer s.lock.Unlock()
	return s.List.Append(t)
}

/**
太多读锁是否会造成写效率降低呢？
读写锁要看具体实现细节，有些读写锁是读优先，有些读写锁是写优先
读优先：如果有人先加了读锁，后面有人使用写锁会阻塞住，再之后加读锁还能加
写优先：如果有人先加了写锁，写锁卡在那边，后面有人加读锁是不会成功的
go是第二种，加了写锁在那边，就加不了读锁，会等待写锁结束，读锁才能添加进去


任何非线程安全的类型、接口都可以利用读写锁+装饰器默认无侵入式改造为线程安全的类型、接口
*/
