package main

import (
	"cloudnativego/workerpool/pool"
	"fmt"
)

func main() {
	cancel, in, out := pool.Dispatch(10)
	defer cancel()
	for i := 0; i < 10; i++ {
		in <- pool.WorkRequest{OP: pool.Hash, Text: []byte(fmt.Sprintf("message %d", i))}
	}
	for i := 0; i < 10; i++ {
		res := <-out
		if res.Err != nil {
			panic(res.Err)
		}
		in <- pool.WorkRequest{OP: pool.Cmpare, Text: res.Wr.Text, Compare: res.Result}
	}
	for i := 0; i < 10; i++ {
		res := <-out
		if res.Err != nil {
			panic(res.Err)
		}
		fmt.Printf("string:\"%s\";match:%v\n", string(res.Wr.Text), res.Matched)
	}

}
