// sync.Pool example, page 59
package main

import "concurrency/ch03/pool"

func main() {
	pool.Pool.Get()
	instance := pool.Pool.Get()
	pool.Pool.Put(instance)
	pool.Pool.Get()
}
