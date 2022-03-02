# Bootstrap file which is executed when "dotc bootstrap <collection-name>"

# Check that the var file exists to determine if configuration has been run or not yet
# Remove this if your collection does not require configuration before being bootstrapped
VARFILE=$(dirname "$0")/.vars
if [ ! -f "$VARFILE" ]; then
    echo "it doesn't look like you have configured the collection, run 'dotc configure <collection>' before continuing"
    exit
fi