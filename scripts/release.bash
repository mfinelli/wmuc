#!/bin/bash -e

# start fresh
make clean

# tests pass?
make test

# build it once
make

# remove the artifact
rm wmuc

# loop through what we offer and build it
for os in darwin linux windows; do
  for arch in amd64; do
    env GOOS=$os GOARCH=$arch make

    if [[ $os == windows ]]; then
      exe=wmuc.exe
    else
      exe=wmuc
    fi

    mv ${exe} ${exe}-${os}_${arch}
    xz ${exe}-${os}_${arch}
    chmod 644 ${exe}-${os}_${arch}.xz
    gpg -ba ${exe}-${os}_${arch}.xz
  done
done

# capture the sources we used
make third-party.tar.gz
gpg -ba third-party.tar.gz

exit 0
