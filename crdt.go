package crdt

import "github.com/kavehmz/lww"

//CRDT is a struct
type CRDT struct {
	RedisURL string
	lww.LWW
}
