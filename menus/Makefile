include $(GOROOT)/src/Make.inc

TARG=curses/menus

CGOFILES=\
		menus.go\
		utils.go

CGO_LDFLAGS=-lmenu -lncurses

CLEANFILES+=

include $(GOROOT)/src/Make.pkg

%: install %.go
	$(GC) $*.go
	$(LD) -o $@ $*.$O
