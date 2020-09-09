#!/bin/bash

set -e

scriptdir=$(cd $(dirname $0); pwd -P)

# WARNING: build.sh uses CMake to build mcl, but sudo make install in the build/ directory does not actually install the files locally

make
sudo make install
