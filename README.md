# Wind Me Up, Chuck

[![Build Status](https://travis-ci.org/mfinelli/wmuc.svg?branch=master)](https://travis-ci.org/mfinelli/wmuc)

A git repository manager.


## Installation

```shell
$ go get github.com/mfinelli/wmuc
```

## Usage

Specify the repos that you want in a directory with a `chuckfile`:

```ruby
# chuckfile

# clone repositories into the same directory
# as the chuckfile
repo "https://github.com/mfinelli/wmuc.git"

# clone repositories into a directory
project "wmuc" do
    repo "https://github.com/mfinelli/wmuc.git"

    # optionally, specify a branch to checkout
    repo "https://github.com/mfinelli/wmuc.git", branch: "dev"
end
```

Now, to clone any missing repositories run `wmuc sync`.
