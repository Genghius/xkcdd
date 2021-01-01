.POSIX: 

PREFIX = ~/.local

all: xkcdd 

xkcdd: xkcdd.go 
	go build xkcdd.go

clean: 
	rm -f xkcdd 

install: xkcdd 
	mkdir -p ${DESTDIR}${PREFIX}/bin 
	cp xkcdd ${DESTDIR}${PREFIX}/bin/ 
	mkdir -p ${DESTDIR}${PREFIX}/share/man/man1 
	cp xkcdd.1 ${DESTDIR}${PREFIX}/share/man/man1/xkcdd.1 

uninstall: 
	rm -f ${DESTDIR}${PREFIX}/bin/xkcdd
	rm -f ${DESTDIR}${PREFIX}/share/man/man1/xkcdd.1 

.PHONY: all clean install uninstall 

