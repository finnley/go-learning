package demo_sync

import "sync"

// ==================== 第一种用法 ====================
// 下面的这种写法缺点：
// PublicResource 和 PublicLock 都是 public

// 但是用户如果不去看源码或者没有注释，是无法只知道 PublicResource 是需要用 PublicResourceLock 保护起来的，
// 我希望是像上面那么用，但是我申明成公共变量，使用的人是不会管你的，所以不建议使用下面这种写法
var PublicResource map[string]string
var PublicLock sync.Mutex

// 如果在其他地方有个包需要用到下面的这些公开变量，原本的我的想法是想用 PublicLock 将 PublicResource 保护起来
func MyBusiness() {
	PublicLock.Lock()
	defer PublicLock.Unlock()

	PublicResource["a"] = "a"
}

// ==================== 第二种用法 ====================
// 这种使用方式比第一种要好些，但是也不是很建议
var privateResource map[string]string
var privateLock sync.Mutex

// ==================== 第三种用法（推荐） ====================

// 声明包变量
var safeResourceInstance safeResource

// 所有期望对资源的操作都只能通过定义在 safeResource 上的方法来进行
type safeResource struct {
	resource map[string]string
	lock     sync.Mutex
}

func (s *safeResource) Add(key, value string) {
	s.lock.Lock()
	defer s.lock.Unlock()

	s.resource[key] = value
}

// ==================== Example ====================
// SafeMap 可以看做是 map 的一个线程安全的封装。我们为它增加一个 LoadOrStore 的方法。
// 泛型写法，声明了两个参数，第一个参数是K，对应的是key，K满足的约束是可比较，即key必须是可比较的类型
type SafeMap[K comparable, V any] struct {
	values map[K]V
	lock   sync.RWMutex
}

// 判断map是否已经存在 key, 如果存在，则返回对应的值，然后 loaded = true
// 如果map中没有对应的 key, 则将key放进去，loaded = false
// loaded 表示返回老对象还是新对象
// goroutine1 => (key1, 1)
// goroutine2 => (key1, 2)
func (s *SafeMap[K, V]) LoadOrStoreV3(key K, newVale V) (val V, loaded bool) {
	// 取值-读操作
	s.lock.RLock()
	oldVal, ok := s.values[key]
	// 为什么这里不使用defer: 因为这里加了读锁，如果使用defer,读操作还么有结束就在放值又设置写锁，就会出现死锁
	s.lock.RUnlock()
	if ok {
		return oldVal, true
	}

	// 放值-写操作
	s.lock.Lock()
	defer s.lock.Unlock()
	// 如果goroutine1先进来，这里就会变成 key1 => 1
	// 然后goroutine2后进来，这里就会变成 key1 => 2，key1 已经存在了，就会将key1覆盖掉
	// 所以在goroutine2进来之前先检查 key1 是否存在，这就是 double check，这就不会出现覆盖的情况了
	// double check 核心就是加完写锁之后还需要再检查一遍
	oldVal, ok = s.values[key]
	if ok {
		return oldVal, true
	}
	s.values[key] = newVale
	return newVale, false
}

func (s *SafeMap[K, V]) LoadOrStoreV1(key K, newVale V) (val V, loaded bool) {
	s.lock.RLock()
	oldVal, ok := s.values[key]
	// 本身就一个goroutine,这里还没有释放掉读锁就去加写锁，就会出现死锁，
	defer s.lock.RUnlock()
	if ok {
		return oldVal, true
	}

	s.lock.Lock()
	defer s.lock.Unlock()
	oldVal, ok = s.values[key]
	if ok {
		return oldVal, true
	}
	s.values[key] = newVale
	return newVale, false
}

func (s *SafeMap[K, V]) LoadOrStoreV2(key K, newVale V) (val V, loaded bool) {
	s.lock.RLock()
	oldVal, ok := s.values[key]
	s.lock.RUnlock()
	if ok {
		return oldVal, true
	}

	s.lock.Lock()
	defer s.lock.Unlock()
	s.values[key] = newVale
	return newVale, false
}
