#!/usr/bin/env bash

cd proto

for file in `\find . -name '*.proto'`; do
    protos="$protos $file"
done

protoc --go_out=plugins=grpc:. $protos

cd ../

mv ./proto/*.pb.go ./protobuf/

cd protobuf

for file in `\find . -name '*.pb.go'`; do
    gsed -i 's/package cqrs-es/package protobuf/g' $file
done
