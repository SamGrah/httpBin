#!/bin/bash

# This script generates the protobufs for the project.
( # within a subshell so not to change the current directory
  cd ./bin-manager || exit ; ##  

  # Generate the protobufs 
  protoc --go_out=./pkg/generated --go_opt=Mpkg/bin.proto=./ --go-grpc_out=./pkg/generated --go-grpc_opt=Mpkg/bin.proto=./ ./pkg/bin.proto
)