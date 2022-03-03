COLLECTIONS_PATH=$HOME/.dot-collections

# Go through all the collections and load the dot files within
for collection in "$COLLECTIONS_PATH"/*; do
  if [ -d "$collection" ]; then
    if [ -f "$collection"/collection.sh ]; then
      . "$collection"/collection.sh
    else
      echo "missing collection file for $(basename "$collection")"
    fi
  fi
done