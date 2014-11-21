gowin32
=======

This library provides wrappers to facilitate calling the Win32 API from Go.  The `wrappers` package contains wrappers
that directly expose certain portions of the Win32 API in Go, similar to what is provided by the `syscall` package in
the Go runtime.  The `gowin32` package contains helper functions and data structures that encapsulate Win32
functionality in a more Go-friendly manner.  Developers may elect to use either package or to combine both of them as
they see fit.

This library is based on the Windows SDK 7.1.
