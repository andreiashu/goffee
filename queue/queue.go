package queue

import (
	"github.com/garyburd/redigo/redis"
	"fmt"
)

var RedisAddressWithPort string

const batchSize = 10
const timeout = 5

func FetchBatch() (result []string) {
	c, err := redis.Dial("tcp", RedisAddressWithPort)
	if err != nil {
	    panic(err)
	}
	defer c.Close()

	for i := 0; i < batchSize; i++ {
		reply, err := redis.Values(c.Do("BRPOP", "jobs", timeout))
		if err != nil {
			break
		}
		item := string(reply[1].([]byte))
		result = append(result, item)
	}
	
	return
}


func WriteResult(result string) {
	fmt.Println("Writing result " + result)
	c, err := redis.Dial("tcp", RedisAddressWithPort)
	if err != nil {
	    panic(err)
	}
	defer c.Close()

	_, err = redis.String(c.Do("LPUSH", "results", result))
	if err != nil {
		// never mind
	}
	
	return
}