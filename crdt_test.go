package crdt

import (
	"fmt"
	"time"
)

func ExampleCRDT() {
	c := CRDT{RedisURL: "redis://localhost:6379/"}
	c.Init()
	c.Add("Item", time.Now())
	c.Add("Item2", time.Now())
	c.Remove("Item", time.Now())
	c.Remove("Item2", time.Now().Add(-1*time.Second))

	l := c.Get()
	fmt.Println(l)
	// l includes Item2

	c.Add("Item", time.Now())
	l = c.Get()
	fmt.Println(l)
	// l includes Item and Item2
}
