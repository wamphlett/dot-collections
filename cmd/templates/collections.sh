COLLECTIONS_PATH=$HOME/.dot-collections-2

# Go through all the collections.sh and load the dot files within
for collection in "$COLLECTIONS_PATH"/*; do
  if [ -d "$collection" ]; then
    if [ -f "$collection"/collection.sh ]; then
      . "$collection"/collection.sh
    else
      echo "missing collection file for $(basename "$collection")"
    fi
  fi
done