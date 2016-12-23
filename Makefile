GOPATH = $(PWD)

all: socket client 

socket:
	cd bin && go build -gcflags "-N -l" socket;

client:
	cd bin && go build -gcflags "-N -l" client;

clean:
	-rm bin/socket
	-rm bin/client