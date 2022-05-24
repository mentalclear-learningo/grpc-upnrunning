#!/bin/bash

protodir="ecommerce"
protofile="${protodir}/product_info.proto"

# TODO: Learn: 
# 1. --go_out is it a clinet code? Can it be put into client folder only?
# 2. --go-grpc_out is it a server code? Can it be put into server folder only?

# Generate Go code 
protoc --go_out=. --go_opt=paths=source_relative \
--go-grpc_out=. --go-grpc_opt=paths=source_relative \
$protofile



# Generate Python code 
# python3 -m grpc_tools.protoc -I$protodir \
# --python_out=$protodir --grpc_python_out=$protodir \
# $protofile
