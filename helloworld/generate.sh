#/bin/bash

protoc -I helloworldpb --go_out=plugins=grpc:helloworldpb --go_opt=paths=source_relative helloworldpb/helloworld.proto
