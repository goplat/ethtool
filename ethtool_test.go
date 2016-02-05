/*
 *
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 *  http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on an
 * "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 * KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations
 * under the License.
 *
 */

package ethtool

import (
    "fmt"
	"net"
    "os/exec"
	"testing"
)

func TestStats(t *testing.T) {
	intfs, err := net.Interfaces()
	if err != nil {
		t.Fatal(err)
	}

	// we expected to have at least one success
	empty := true
	for _, intf := range intfs {
		stats, err := Stats(intf.Name)
		if err != nil {
			continue
		}

		for _, _ = range stats {
			empty = false
			break
		}
	}

	if empty {
		t.Fatal("Unable to retrieve stats from any interface of this system.")
	}
}

func TestDriverName(t *testing.T) {
	intfs, err := net.Interfaces()
	if err != nil {
		t.Fatal(err)
	}

	// we expected to have at least one success
	success := false
	for _, intf := range intfs {
		drv, err := DriverName(intf.Name)
        fmt.Println("-----------------")
        fmt.Println(intf.Name)
        fmt.Println(drv)
        out, err := exec.Command("sudo", "ethtool", "-i", intf.Name).Output()
        fmt.Println(err)
        fmt.Println(string(out))

		if err == nil {
			success = true
		}
	}

	if !success {
		t.Fatal("Unable to retrieve driver from any interface of this system.")
	}
}
