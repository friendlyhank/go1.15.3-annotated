package runtime

import "runtime/internal/sys"

// GOARCH is the running program's architecture target:
// one of 386, amd64, arm, s390x, and so on.
const GOARCH string = sys.GOARCH
