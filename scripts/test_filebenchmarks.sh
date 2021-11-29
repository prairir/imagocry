#!/bin/bash

set -euxo

# benchmarking script
# creates a large number of files

# filename is of
# file size is (bs * count)

dd if=/dev/random of=~/file00 bs=1M count=2

dd if=/dev/random of=~/file01 bs=1M count=4

dd if=/dev/random of=~/file02 bs=1M count=8

dd if=/dev/random of=~/file03 bs=1M count=16

dd if=/dev/random of=~/file04 bs=1M count=32

dd if=/dev/random of=~/file05 bs=1M count=64

dd if=/dev/random of=~/file06 bs=1M count=128

dd if=/dev/random of=~/file07 bs=1M count=256

dd if=/dev/random of=~/file08 bs=1M count=512

dd if=/dev/random of=~/file09 bs=1M count=1024

dd if=/dev/random of=~/file10 bs=1M count=2048
