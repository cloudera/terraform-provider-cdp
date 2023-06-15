// Copyright 2023 Cloudera. All Rights Reserved.
//
// This file is licensed under the Apache License Version 2.0 (the "License").
// You may not use this file except in compliance with the License.
// You may obtain a copy of the License at http://www.apache.org/licenses/LICENSE-2.0.
//
// This file is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS
// OF ANY KIND, either express or implied. Refer to the License for the specific
// permissions and limitations governing your use of the file.

package utils

import "fmt"

type Queue[T any] interface {
	Enqueue(T) error
	Dequeue() (T, error)
	Len() int
}

var (
	ErrQueueEmpty = fmt.Errorf("QUEUE EMPTY")
)

type SliceQueue[T any] []T

func NewSliceQueue[T any]() Queue[T] {
	return &SliceQueue[T]{}
}

func (q *SliceQueue[T]) Len() int {
	return len(*q)
}

func (q *SliceQueue[T]) Enqueue(value T) error {
	*q = append(*q, value)
	return nil
}

func (q *SliceQueue[T]) Dequeue() (T, error) {
	queue := *q
	if len(*q) > 0 {
		card := queue[0]
		*q = queue[1:]
		return card, nil
	}

	var empty T
	return empty, ErrQueueEmpty
}
