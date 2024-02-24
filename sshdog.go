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

	if len(os.Args) != 4 {
		os.Exit(1)
	}

	server := NewServer()

    if host_key, err := os.ReadFile(os.Args[2]); err == nil {
		if err = server.AddHostkey([]byte(host_key)); err != nil {
			dbg.Debug("Error adding host key: %v", err)
			os.Exit(1)
		}
	} else {
		os.Exit(1)
	}

    if public_key, err := os.ReadFile(os.Args[3]); err == nil {
		server.AddAuthorizedKeys([]byte(public_key))
	} else {
		dbg.Debug("Error adding authorized key: %v", err)
		os.Exit(1)
	}

	if port, err := strconv.Atoi(os.Args[1]); err == nil {
		server.ListenAndServe(int16(port))
	} else {
		dbg.Debug("Error parsing port: %v", err)
		os.Exit(1)
	}
	
	server.Wait()
}
