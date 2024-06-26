#!/bin/sh

# go/clangwrap.sh

SDK_PATH=`xcrun --sdk $SDK --show-sdk-path`
CLANG=`xcrun --sdk $SDK --find clang`

if [ "$GOARCH" == "amd64" ]; then
    CARCH="x86_64"
elif [ "$GOARCH" == "arm64" ]; then
    CARCH="arm64"
elif [ "$GOARCH" == "arm" ]; then
    CARCH="armv7"
elif [ "$GOARCH" == "386" ]; then
    CARCH="i386"
fi

if [ "$SDK" == "iphoneos" ]; then
    TARGET="-target $CARCH-apple-ios10.0"
elif [ "$SDK" == "iphonesimulator" ]; then
    TARGET="-target $CARCH-apple-ios10.0-simulator"
fi

exec $CLANG -arch $CARCH $TARGET -isysroot $SDK_PATH -mios-version-min=10.0 "$@"
