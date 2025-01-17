package cpsv

import "C"

const (
	Async int = 0
	Sync  int = 1
)

type req struct {
	sectionId string
	data      []byte
	size      int
	offset    int
	reqType   int
	resend    int
}

type eventQ struct {
	queue chan req
}