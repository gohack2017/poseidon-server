#!/usr/bin/env bash

export APPROOT=$(pwd)

# adjust GOPATH
case ":$GOPATH:" in
    *":$APPROOT:"*) :;;
    *) GOPATH=$APPROOT:$GOPATH;;
esac
export GOPATH


# adjust PATH
readopts="ra"
if [ -n "$ZSH_VERSION" ]; then
    readopts="rA";
fi
while IFS=':' read -$readopts ARR; do
    for i in "${ARR[@]}"; do
        case ":$PATH:" in
            *":$i/bin:"*) :;;
            *) PATH=$i/bin:$PATH
        esac
    done
done <<< "$GOPATH"
export PATH


# mock development && test envs
if [ ! -d "$APPROOT/src/github.com/poseidon" ];
then
    mkdir -p "$APPROOT/src/github.com"
    ln -s "$APPROOT/gogo/" "$APPROOT/src/github.com/poseidon"
fi
