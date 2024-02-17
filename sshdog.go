// Copyright 2016 Google Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"fmt"
	"os"
	"strconv"
)

type Debugger bool

func (d Debugger) Debug(format string, args ...interface{}) {
	if d {
		msg := fmt.Sprintf(format, args...)
		fmt.Fprintf(os.Stderr, "[DEBUG] %s\n", msg)
	}
}
var dbg Debugger = true
func main() {

	server := NewServer()

	if host_key, ok := os.LookupEnv("SSHDOG_HOST_KEY"); ok {
		if err := server.AddHostkey([]byte(host_key)); err != nil {
			return
		}
	} else {
		return
	}

	if public_key, ok := os.LookupEnv("SSHDOG_AUTH_KEY"); ok {
		server.AddAuthorizedKeys([]byte(public_key))
	} else {
		return
	}

	if port, err := strconv.Atoi(os.Getenv("SSHDOG_PORT")); err != nil {
		server.ListenAndServe(int16(port))
	} else {
		return
	}
	
	server.Wait()
}
