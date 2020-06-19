#!/bin/sh

set -e
set -u
set -x

# Generates a collection of release tarballs

ARCH="$(uname -p)"

cd "$(dirname "$0")"


make clean
make build

mkdir -p ./release

VERSION="$(./build/bin/wavegen --version | cut -d' ' -f 2)"

RELEASENAME="wavegen-$VERSION-$ARCH"

cd build
tar cvfz "../release/$RELEASENAME.tar.gz" .
cd ..

PROJ="$(pwd)"
TEMP="$(mktemp -d)"

cp -R ./ "$TEMP"
cd "$TEMP"

printf "Synthetic wave generation tool\n" | sudo checkinstall -D --install=no --gzman --strip --nodoc --pkgrelease "$VERSION" --pkgname wavegen
sudo chown "$(whoami)" *.deb
mv *.deb "$PROJ/release"

ls -lah

cd "$PROJ"
rm -rf "$TEMP"
