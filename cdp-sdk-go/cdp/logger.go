// Copyright 2023 Cloudera. All Rights Reserved.
//
// This file is licensed under the Apache License Version 2.0 (the "License").
// You may not use this file except in compliance with the License.
// You may obtain a copy of the License at http://www.apache.org/licenses/LICENSE-2.0.
//
// This file is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS
// OF ANY KIND, either express or implied. Refer to the License for the specific
// permissions and limitations governing your use of the file.

package cdp

import (
	"context"
	"log"
	"os"
)

// Logger is a simple interface so that SDK users can write their own logging adaptors
type Logger interface {
	Errorf(ctx context.Context, format string, args ...any)
	Warnf(ctx context.Context, format string, args ...any)
	Infof(ctx context.Context, format string, args ...any)
	Debugf(ctx context.Context, format string, args ...any)
}

type DefaultLogger struct {
	logger *log.Logger
}

func NewDefaultLogger() *DefaultLogger {
	return &DefaultLogger{logger: log.New(os.Stdout, "", 5)}
}

func (l *DefaultLogger) Errorf(_ context.Context, format string, args ...any) {
	l.logger.Printf(format, args...)
}

func (l *DefaultLogger) Warnf(_ context.Context, format string, args ...any) {
	l.logger.Printf(format, args...)
}

func (l *DefaultLogger) Infof(_ context.Context, format string, args ...any) {
	l.logger.Printf(format, args...)
}

func (l *DefaultLogger) Debugf(_ context.Context, format string, args ...any) {
	l.logger.Printf(format, args...)
}
