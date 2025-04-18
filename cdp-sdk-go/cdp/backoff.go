// Copyright 2024 Cloudera. All Rights Reserved.
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
	"crypto/rand"
	"log"
	"math"
	"math/big"
	"os"
	"time"
)

const (
	expDeltaMin = 0.75
	expDeltaMax = 1.0
)

func backoff(retries int) time.Duration {
	switch os.Getenv("CDP_TF_BACKOFF_STRATEGY") {
	case "linear":
		{
			step := intFromEnvOrDefault("CDP_TF_BACKOFF_STEP", defaultLinearBackoffStep)
			log.Default().Println("Using linear backoff strategy with step: ", step)
			return linearBackoff(retries, step)
		}
	default:
		{
			log.Default().Println("Using exponential backoff strategy")
			return exponentialBackoff(retries)
		}
	}
}

// Generates a cryptographically secure random float64 in [0, 1)
func cryptoRandFloat64() float64 {
	const precision = 1 << 53
	nBig, err := rand.Int(rand.Reader, big.NewInt(precision))
	if err != nil {
		panic(err)
	}
	return float64(nBig.Int64()) / float64(precision)
}

func exponentialBackoff(retries int) time.Duration {
	delta := expDeltaMax - expDeltaMin
	jitter := expDeltaMin + cryptoRandFloat64()*delta
	return time.Duration((math.Pow(2, float64(retries))*jitter)*float64(time.Millisecond)) * 1000
}

func linearBackoff(retries int, step int) time.Duration {
	return time.Duration((retries+1)*step) * time.Second
}
