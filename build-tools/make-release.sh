#!/bin/bash

TARGET_DIR=$1
OUTPUT_PREFIX=$2
shift 2

mkdir -p $TARGET_DIR

# Add other platforms as needed
supported_oss=("windows" "linux" "darwin")
supported_archs=("amd64")

for os in "${supported_oss[@]}"; do
  GOOS=${os}
  for arch in "${supported_archs[@]}"; do
    GOARCH=${arch}
    OUTPUT="${OUTPUT_PREFIX}-${GOOS}-${GOARCH}"
    if [ ${GOOS} = "windows" ]; then
      OUTPUT+='.exe'
    fi

    echo "Building binary for platform: ${GOOS}/${GOARCH}"
    env GOOS=$GOOS GOARCH=$GOARCH go build $@ -o $TARGET_DIR/$OUTPUT ./
    sha256sum $TARGET_DIR/$OUTPUT > $TARGET_DIR/$OUTPUT.sha256
  done
done

echo "--------------------------------------------------------"
echo "Finished building the release for $OUTPUT_PREFIX"
echo "--------------------------------------------------------"
echo "All release artifacts are under directory $TARGET_DIR/"
echo "--------------------------------------------------------"
ls -la "${TARGET_DIR}"
echo "--------------------------------------------------------"