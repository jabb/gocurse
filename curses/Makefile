include $(GOROOT)/src/Make.inc

TARG=curses

CGOFILES=curses.go\
		curses_defs.go

CGO_LDFLAGS=-lncurses

CLEANFILES+=

include $(GOROOT)/src/Make.pkg

# Simple test programs

%: install %.go
	$(GC) $*.go
	$(LD) -o $@ $*.$O
