// Copyright (c) 2015 HPE Software Inc. All rights reserved.
// Copyright (c) 2013 ActiveState Software Inc. All rights reserved.

package util

import (
	"fmt"
	"log"
	"os"
	"runtime/debug"
)

type Logger struct {
	*log.Logger
}

var LOGGER = &Logger{log.New(os.Stderr, "", log.LstdFlags)}

func Error(format string, v ...interface{}) {
	// https://github.com/hpcloud/log/blob/master/log.go#L45
	LOGGER.Output(2, fmt.Sprintf("Error -- "+format, v...)+"\n"+string(debug.Stack()))
}

// partitionString partitions the string into chunks of given size,
// with the last chunk of variable size.
func PartitionString(s string, chunkSize int) []string {
	if chunkSize <= 0 {
		panic("invalid chunkSize")
	}
	length := len(s)
	chunks := 1 + length/chunkSize
	start := 0
	end := chunkSize
	parts := make([]string, 0, chunks)
	for {
		if end > length {
			end = length
		}
		parts = append(parts, s[start:end])
		if end == length {
			break
		}
		start, end = end, end+chunkSize
	}
	return parts
}
