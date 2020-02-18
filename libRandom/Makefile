all: clean init-path build copy-artifacts

init-path:
	mkdir -p ../bin

SRC    = $(wildcard *.c)
LIBS   = 
CFLAGS = -Wall -std=gnu99 -fPIC

build: $(SRC)
	gcc -shared $^ $(CFLAGS) $(LIBS) -o libRandom.so 

copy-artifacts:
	mkdir -p ../bin
	cp libRandom.so ../bin
	
clean:
	rm -rf libRandom.so
	rm -f ../bin/libRandom.so
