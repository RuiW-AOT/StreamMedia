package main

import "log"

// Bucket token algorithm
// N token in the buckets, each request will get assigned a token
// Token is released when request gets its response
// There can be at most N concurrent streaming request

type ConnectionLimiter struct {
	concurrentConn int
	bucket         chan int
}

func NewConnectionLimiter(cc int) *ConnectionLimiter {
	return &ConnectionLimiter{
		concurrentConn: cc,
		bucket:         make(chan int, cc),
	}
}

func (cl *ConnectionLimiter) GetConnection() bool {
	if len(cl.bucket) >= cl.concurrentConn {
		log.Printf("Cannot get token, reached the rate limitation")
		return false
	}
	cl.bucket <- 1
	return true
}

func (cl *ConnectionLimiter) ReleaseConnection() {
	_ = <-cl.bucket
}
