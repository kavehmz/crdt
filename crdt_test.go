package crdt

import (
	"fmt"
	"time"
)

func Example() {
	c := CRDT{RedisURL: "redis://localhost:6379/0"}
	c.Connect()
	c.Add("Item", time.Now())
	c.Add("Item2", time.Now())
	c.Remove("Item", time.Now())
	c.Remove("Item2", time.Now().Add(-1*time.Second))

	l := c.Get()
	fmt.Println(l)

	c.Add("Item", time.Now())
	c.Remove("Item2", time.Now())
	l = c.Get()
	fmt.Println(l)

	// Output:
	// [Item2]
	// [Item]
}
