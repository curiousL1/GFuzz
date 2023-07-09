#!/bin/bash -xe
cd "$(dirname "$0")"/..


TARGET_GIT=$1
TARGET_GIT_COMMIT=$2
OUTPUT_DIR=$3
shift 3


docker build \
--build-arg GIT_URL=$TARGET_GIT \
--build-arg GIT_COMMIT=$TARGET_GIT_COMMIT \
-f docker/fuzzer-git/Dockerfile \
-t gfuzzgit:latest .

container_id=$(docker run -d \
-v $(pwd)/tmp/pkgmod:/go/pkg/mod \
-v $OUTPUT_DIR:/fuzz/output \
gfuzzgit:latest /fuzz/output $@)

echo "using command `docker logs $container_id -f` or checking $OUTPUT_DIR/fuzzer.log to get latest log"
