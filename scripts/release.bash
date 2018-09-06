#!/bin/bash -e

# start fresh
make clean
git stash --all

# tests pass?
make test

# build it once: make sure any `go run`
# that we issue runs with the current OS
# (go run with GOOS or GOARCH will try
# and use those settings which might not
# work depending on the order below and
# our actual system)
make

# remove the artifact
rm wmuc

# loop through what we offer and build it
for os in darwin linux windows; do
  for arch in amd64; do
    GOOS=$os GOARCH=$arch make

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

# generate a source tarball and sign it
tag=$(git describe --tags $(git rev-list --tags --max-count=1))
git archive -o "wmuc-$tag.tar.gz" HEAD
gpg -ba "wmuc-$tag.tar.gz"

git stash pop

exit 0
