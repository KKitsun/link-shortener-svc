package mocks

import (
	"gitlab.com/distributed_lab/kit/comfig"
	"gitlab.com/distributed_lab/logan/v3"
)

type FakeLogger interface {
	FakeLog() *logan.Entry
}

type fakelogger struct {
	once comfig.Once
}

type LoggerOpts struct {
	Release string
}

func NewFakeLogger() FakeLogger {
	return &fakelogger{}
}

func (l *fakelogger) FakeLog() *logan.Entry {
	return l.once.Do(func() interface{} {

		lLevel, _ := logan.ParseLevel("debug")
		entry := logan.New().Level(lLevel)

		return entry
	}).(*logan.Entry)
}
