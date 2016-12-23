protoc -I=../protocol --go_out=../src/protocol/out/cs ../protocol/cs.proto
protoc -I=../protocol --go_out=../src/protocol/in/ssrs ../protocol/ssrs.proto

::pause