package crdt

import (
	"github.com/garyburd/redigo/redis"
	"github.com/kavehmz/lww"
	"github.com/kavehmz/qset"
)

//CRDT is a struct
type CRDT struct {
	RedisURL string
	Key      string
	lww.LWW

	AddWrite    redis.Conn
	AddSub      redis.Conn
	RemoveWrite redis.Conn
	RemoveSub   redis.Conn
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}

}

// Connect is a function
func (c *CRDT) Connect() {

	var err error
	if c.AddWrite == nil {
		c.AddWrite, err = redis.DialURL(c.RedisURL)
		checkErr(err)
	}
	if c.AddSub == nil {
		c.AddSub, err = redis.DialURL(c.RedisURL)
		checkErr(err)
	}
	if c.RemoveWrite == nil {
		c.RemoveWrite, err = redis.DialURL(c.RedisURL)
		checkErr(err)
	}
	if c.RemoveSub == nil {
		c.RemoveSub, err = redis.DialURL(c.RedisURL)
		checkErr(err)
	}

	c.AddSet = &qset.QSet{ConnWrite: c.AddWrite, ConnSub: c.AddSub, Marshal: func(e interface{}) string { return e.(string) }, UnMarshal: func(e string) interface{} { return e }, SetKey: "ADDSET"}
	c.RemoveSet = &qset.QSet{ConnWrite: c.RemoveWrite, ConnSub: c.RemoveSub, Marshal: func(e interface{}) string { return e.(string) }, UnMarshal: func(e string) interface{} { return e }, SetKey: "ADDSET"}
	c.Init()
}
