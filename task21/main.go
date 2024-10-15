package main

import "fmt"

type LegacyWorker interface {
	DoStuff(text string) string
}
type NewWorker interface {
	DoMoreStuff(text string, copies int) string
}

type LegacyWorkerStruct struct{}

type NewWorkerStruct struct{}

type LegacyWorkerAdapter struct {
	LegacyWorkerAdapter NewWorker
}

func (p *LegacyWorkerStruct) DoStuff(text string) string {
	message := fmt.Sprintf("doing some stuff with: '%s'", text)
	return message
}

func (p *NewWorkerStruct) DoMoreStuff(text string, copies int) string {
	message := fmt.Sprintf("acquired data '%s' in %d copies", text, copies)
	return message
}

func (adapter *LegacyWorkerAdapter) DoStuff(text string) string {
	return adapter.LegacyWorkerAdapter.DoMoreStuff(text, 1)
}

func main() {
	lw := &LegacyWorkerStruct{}
	fmt.Println(lw.DoStuff("{Date:2021-12-12}"))

	nws := &NewWorkerStruct{}
	adapter := &LegacyWorkerAdapter{
		LegacyWorkerAdapter: nws,
	}

	fmt.Println(adapter.DoStuff("Adapter data"))
}
