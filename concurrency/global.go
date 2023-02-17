package concurrency

import (
	"sync"
)

var (
	WG      = sync.WaitGroup{}
	JobList = make(chan job)
)
