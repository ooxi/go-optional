// SPDX-License-Identifier: Zlib
package optional

import (
	"errors"
)

var NoSuchElementError = errors.New("The element being requested does not exist.")
