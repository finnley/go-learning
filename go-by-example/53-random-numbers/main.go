package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	fmt.Print(rand.Intn(100), ",")
	fmt.Print(rand.Intn(100))
	fmt.Println() // 81,87

	fmt.Println(rand.Float64()) // 0.6645600532184904

	fmt.Print((rand.Float64()*5)+5, ",")
	fmt.Print((rand.Float64() * 5) + 5)
	fmt.Println() // 7.1885709359349015,7.123187485356329

	// 默认情况下，给定的种子是确定的，每次都会产生相同的随机数数字序列。 要产生不同的数字序列，需要给定一个不同的种子。 注意，对于想要加密的随机数，使用此方法并不安全， 应该使用 crypto/rand。
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)

	fmt.Print(r1.Intn(100), ",")
	fmt.Print(r1.Intn(100))
	fmt.Println() // 56,10 每次运行值不一样

	s2 := rand.NewSource(42)
	r2 := rand.New(s2)
	fmt.Print(r2.Intn(100), ",")
	fmt.Print(r2.Intn(100))
	fmt.Println()
	s3 := rand.NewSource(42)
	r3 := rand.New(s3)
	fmt.Print(r3.Intn(100), ",") // 5,87
	fmt.Print(r3.Intn(100)) // 5,87
}
