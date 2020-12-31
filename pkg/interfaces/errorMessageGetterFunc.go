// Copyright ©2021 by Randy R Bell. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package interfaces

// ErrorMessageGetterFunc defines a function that, given error message context, return an error string
type ErrorMessageGetterFunc func(ctx ValidatorContextGetter) string
