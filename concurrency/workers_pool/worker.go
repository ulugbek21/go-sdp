package main

import (
	"fmt"
	"strings"
)

// WorkerLauncher ...
type WorkerLauncher interface {
	LaunchWorker(in chan Request)
}

// PreffixSuffixWorker ...
type PreffixSuffixWorker struct {
	id      int
	prefixS string
	suffixS string
}

// LaunchWorker ...
func (w *PreffixSuffixWorker) LaunchWorker(in chan Request) {
	w.prefix(w.append(w.uppercase(in)))
}

func (w *PreffixSuffixWorker) uppercase(in <-chan Request) <-chan Request {
	out := make(chan Request)
	go func() {
		for msg := range in {
			s, ok := msg.Data.(string)
			if !ok {
				msg.handler(nil)
				continue
			}
			msg.Data = strings.ToUpper(s)
			out <- msg
		}
		close(out)
	}()
	return out
}

func (w *PreffixSuffixWorker) append(in <-chan Request) <-chan Request {
	out := make(chan Request)
	go func() {
		for msg := range in {
			uppercaseString, ok := msg.Data.(string)
			if !ok {
				msg.handler(nil)
				continue
			}
			msg.Data = fmt.Sprintf("%s%s", uppercaseString, w.suffixS)
			out <- msg
		}
		close(out)
	}()
	return out
}

func (w *PreffixSuffixWorker) prefix(in <-chan Request) {
	go func() {
		for msg := range in {
			uppercasedStringWithSuffix, ok := msg.Data.(string)
			if !ok {
				msg.handler(nil)
				continue
			}
			msg.handler(fmt.Sprintf("%s%s", w.prefixS,
				uppercasedStringWithSuffix))
		}
	}()
}
