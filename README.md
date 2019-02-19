# script2matrix

Runs a script and sends its stdout to a Matrix chat.

## Usage

Build with `make`, install with `make install`. Run like this:

```sh
script2matrix <command> [arg...]
# for example:
script2matrix ls -l ./releases
```

The credentials for Matrix are given in environment variables:

```sh
export MATRIX_USER=@myuser:example.org       # user account used by script2matrix
export MATRIX_PASSWORD=aegh3Iif0Ge5epeivaih  # password for that user account
export MATRIX_TARGET=#channel:example.org    # must be a room alias starting with "#"
```

All environment variables given to `script2matrix` are passed on to the command that it runs, except for those starting with `MATRIX_`.
