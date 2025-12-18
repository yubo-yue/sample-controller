/*
Copyright 2017 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package signals

import (
	"os"
	"syscall"
	"testing"
	"time"
)

func TestSetupSignalHandler(t *testing.T) {
	// Reset the signal handler for this test
	onlyOneSignalHandler = make(chan struct{})

	ctx := SetupSignalHandler()

	if ctx == nil {
		t.Fatal("SetupSignalHandler returned nil context")
	}

	// Send first signal
	p, err := os.FindProcess(os.Getpid())
	if err != nil {
		t.Fatalf("failed to find process: %v", err)
	}

	p.Signal(syscall.SIGTERM)

	// Context should be cancelled after signal
	select {
	case <-ctx.Done():
		// Expected behavior
	case <-time.After(2 * time.Second):
		t.Fatal("context was not cancelled after signal")
	}
}

func TestSetupSignalHandlerPanicsOnSecondCall(t *testing.T) {
	// Reset the signal handler
	onlyOneSignalHandler = make(chan struct{})

	defer func() {
		if r := recover(); r == nil {
			t.Fatal("SetupSignalHandler should panic on second call")
		}
	}()

	SetupSignalHandler()
	SetupSignalHandler() // Should panic
}
