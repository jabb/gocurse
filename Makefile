# Copyright 2009 The Go Authors.  All rights reserved.
# Use of this source code is governed by a BSD-style
# license that can be found in the LICENSE file.

include $(GOROOT)/src/Make.inc

TARG=curses

# Can have plain GOFILES too, but this example doesn't.

GOFILES=curses_defs.go
CGOFILES=curses.go

CGO_LDFLAGS=-lncurses

# To add flags necessary for locating the library or its include files,
# set CGO_CFLAGS or CGO_LDFLAGS.  For example, to use an
# alternate installation of the library:
#	CGO_CFLAGS=-I/home/rsc/gmp32/include
#	CGO_LDFLAGS+=-L/home/rsc/gmp32/lib
# Note the += on the second line.

CLEANFILES+=

include $(GOROOT)/src/Make.pkg

curses_defs.go: curses.c
	godefs -g curses curses.c > curses_defs.go

# Simple test programs

%: install %.go
	$(GC) $*.go
	$(LD) -o $@ $*.$O
