#!/bin/bash
mkdir -p build
if [ "$1" = "--build" ]; then
    env GO111MODULE=on go mod vendor
    env GO111MODULE=on go build --trimpath --buildmode=plugin -o ./build/backend.so
    COMPILATION_RES=$?
    if [ $COMPILATION_RES -eq 0 ]; then
        echo "Successfully compiled backend modules"
    else
        echo "Error while compiling. Aborting"
        exit $COMPILATION_RES
    fi
fi

nakama migrate up --database.address root@localhost:26257/matchmaker-bug
nakama --runtime.path ./build \
       --name backend \
       --socket.port 5350 \
       --console.port 5351 \
       --database.address root@localhost:26257/matchmaker-bug 
