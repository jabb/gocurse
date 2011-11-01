include $(GOROOT)/src/Make.inc

TARG=curses/panels
CGOFILES=\
		panels.go\
		utils.go

CGO_LDFLAGS=-lpanel -lncurses
CLEANFILES+=

include $(GOROOT)/src/Make.pkg

%: install %.go
	$(GC) $*.go
	$(LD) -o $@ $*.$O
