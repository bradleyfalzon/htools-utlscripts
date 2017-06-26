#!/bin/bash

buildVars(){
  export GOPATH="${PWD}:${GOPATH}"
}

buildAll(){
  (
    set -ex
    buildVars
    local oldDir=${PWD}
    cd ./src/github.com/hypersuite/htools-utlscripts
    dep ensure

    cd ${oldDir}
    go install github.com/hypersuite/htools-utlscripts/cmd/...
  )  
}

buildAll
