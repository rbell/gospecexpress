// Copyright ©2021 by Randy R Bell. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package errorhelpers

// AddMessages adds messages to an error map
func AddMessages(errMap map[string][]string, context string, msg ...string) {
	if _, ok := errMap[context]; !ok {
		errMap[context] = []string{}
	}
	errMap[context] = append(errMap[context], msg...)
}
