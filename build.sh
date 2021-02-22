#! /bin/bash

function log() {
    echo "【go-disk build】"====\> $@
}
go_path=$(go env GOPATH)
build_dir=$(pwd)
build_dir_name=$(basename $build_dir)
correct_build_dir=$go_path/src/$build_dir_name
client_dir=$(pwd)/front
server_dir=$(pwd)/back
ENV=$1
ENV="${ENV:-production}"
log  corrent build dir is $correct_build_dir, current build dir is $build_dir

if [ "$build_dir" != "$correct_build_dir" ]; then
    log build failed, you must place project at $correct_build_dir
    exit
fi



function beforeBuild() {
    log "before build, cleaning..."
    rm -rf $build_dir/build_dir_name
    rm -rf $server_dir/static
    rm -rf $server_dir/template
}

function buildClient() {
  log "start building client for go-disk..."
    cd $client_dir
    npm install && npm run build
}

function buildServer() {
    log "start build server for go-disk"
    mv $client_dir/dist/static $server_dir/
    mv $client_dir/dist/index.html $server_dir/static/

    cd $server_dir
    go get -v -t -d
    go build -o go-disk
    ./go-disk
}

function afterBuild() {
    log "after build, cleaning..."
    rm -rf $client_dir/dist
    du -ah ./
}

beforeBuild
buildClient
buildServer
afterBuild

