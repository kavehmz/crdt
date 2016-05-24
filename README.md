# CRDT

In distributed computing, a __conflict-free replicated data type__ (CRDT) is a type of specially-designed data structure used to achieve strong eventual consistency (SEC) and monotonicity (absence of rollbacks).

One way of implementing CRDT is by using LWW-element-set.

LWW-element-set is set with timestamped adds and removes

LWW-element-set elements are stored with timestamp. Add and remove will save the operation timestamp along with data in two different sets. For each set new operation of add/remove will update the timestamp for that element.

Queries like Get/List and Len over LWW-set will check both add and remove timestamps to decide if latest state of each element is "exists" or "removed".

# crdt package

crdt package uses two libraries, lww and qset to create a conflict-free replicated storage.

## lww https://github.com/kavehmz/lww

lww implements the logic of an LWW-element-set. It is design in a modular way to use different types of underlying sets. Each set can have a different characteristic like using Go internal maps to be fast or using Redis to share state will other processes and staying persistence.

## qset https://github.com/kavehmz/qset

Is one implementation of what lww can use as underlying set. It mixes Go internal maps and Redis storage to provide lww package with a both fast and persistent underlying set.

## Installation

```bash
$ go get github.com/kavehmz/crdt
```

# Usage

```go
package main

import (
	"fmt"
	"github.com/kavehmz/crdt"
)

func main() {
	c := crdt.CRDT{redisURL:"redis://localhost:6379/"}
    c.Init()
    c.Add("Item")
    c.AddTs("Item2", Time.Now())
    c.RemoveTs("Item", Time.Now())
    c.RemoveTs("Item2", time.Now().Add(-1 * time.Second))

    l:=c.List()
    # l includes Item2

    c.Add("Item")
    l:=c.List()
    # l includes Item and Item2
}
```
