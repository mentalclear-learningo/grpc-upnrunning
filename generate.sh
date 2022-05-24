#!/bin/bash

protodir="ecommerce"
protofile="${protodir}/product_info.proto"


# Generate Go code 
protoc --go_out=. --go_opt=paths=source_relative \
--go-grpc_out=. --go-grpc_opt=paths=source_relative \
$protofile

# Generate Python code 
# python3 -m grpc_tools.protoc -I$protodir \
# --python_out=$protodir --grpc_python_out=$protodir \
# $protofile
