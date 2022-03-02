#!/bin/sh

# The main collection file. You can include additional files here that your collection might require

# Uncomment the following code to prevent your collection being available until the configuration has been run
#
# VARFILE=$(dirname "$0")/.vars
# if [ ! -f "$VARFILE" ]; then
#     echo "it doesn't look like you have configured the collection, run 'dotc configure <collection>' before continuing"
#     exit
# fi

. "$(dirname "$0")"/vars.sh
. "$(dirname "$0")"/aliases.sh
. "$(dirname "$0")"/functions.sh