package scheduler

import "crawler/engine"

type SimpleScheduler struct {
	workChan chan engine.Request
}

func (s *SimpleScheduler) ConfigureMasterWorkChan(c chan engine.Request) {
	s.workChan = c
}

func (s *SimpleScheduler) Submit(r engine.Request) {
	//send to request to work channel
	go func () {s. workChan <- r}()
}

