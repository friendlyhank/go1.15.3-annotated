// Copyright 2013 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build amd64 amd64p32 386

package runtime

import (
	"runtime/internal/sys"
	"unsafe"
)

// adjust Gobuf as if it executed a call to fn with context ctxt
// and then did an immediate gosave.
func gostartcall(buf *gobuf, fn, ctxt unsafe.Pointer) {
	//调整sp
	sp := buf.sp
	if sys.RegSize > sys.PtrSize {
		sp -= sys.PtrSize
		*(*uintptr)(unsafe.Pointer(sp)) = 0
	}
	sp -= sys.PtrSize
	//将buf.pc也就是goexit入栈
	*(*uintptr)(unsafe.Pointer(sp)) = buf.pc

	//然后再设置sp和pc,此时pc才是G任务函数
	buf.sp = sp
	buf.pc = uintptr(fn)
	buf.ctxt = ctxt
}
