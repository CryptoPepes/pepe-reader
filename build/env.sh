#!/bin/sh

set -e

if [ ! -f "build/env.sh" ]; then
    echo "$0 must be run from the root of the repository."
    exit 2
fi

# Create fake Go workspace if it doesn't exist yet.
workspace="$PWD/build/_workspace"
root="$PWD"
namespacedir="$workspace/src/cryptopepe.io"
reponame="cryptopepe-reader"
if [ ! -L "$namespacedir/$reponame" ]; then
    mkdir -p "$namespacedir"
    cd "$namespacedir"
    # Link to repo in PWD
    ln -s ../../../../. "$reponame"
    cd "$root"
fi

# Set up the environment to use the workspace.
GOPATH="$workspace"
export GOPATH

# Run the command inside the workspace.
cd "$namespacedir/$reponame"
PWD="$namespacedir/$reponame"

# Launch the arguments with the configured environment.
exec "$@"
