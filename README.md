# . COLLECTIONS
A easy way to manage collections of dot files

## Install
Clone this repo and add the `.collections` file to your `.bashrc` file.

```
source <path to repo>/.collections
```

This will create a directory in your home directory called `.dot-collections`.

## Adding Collections
Collections can be added by creating a directory in the `.dot-collections` directory. The name of the directory is considered the name of the collection. Each collection must include a `.collection` file, this is the file that will be including when loading collections.

_it is worth noting that all collections will be loaded in alphabetical order_