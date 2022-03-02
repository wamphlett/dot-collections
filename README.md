# DOT COLLECTIONS
Ever needed to manage a collection of dot files for a project? share a set of shell functions between a team? all while leaving your personal dot files untainted. Dot collections makes this really easy by letting you create configurable collections for each project you work on.    

This is not intended to be a replacement for dot file managers, there are many out there which are much better suited. Instead, this is intended to be used in addition.  

## Install
```
go install github.com/wamphlett/dot-collections/dotc@latest
```
Run `dotc install` to install dot collections. (running any other command before this will also trigger the install). This will create a directory in your home directory called `.dot-collections` which is where all collections will be added.

## Usage
### Add an existing collection
```
dotc add
```
The `add` method will ask you for a URL and an identifier. The URL should be a git repository of a valid collection and the identifier can be any string.

### Create a new collection
```
dotc new
```
The `new` method will create new collection using the template files and is the recommended way of making a new collection. This function will not initialise a git repo, if you want to manage your collection in a repository, this will need to be done manually.

### Configure collection
```
dotc configure [collection]
```
Dot collections provides the ability to configure all the required environment variables. Variables should be defined in the collection detail yaml file. Users can then use `confgure` to run through the setup. N.B. In order to make use of this functionality, the collection must make use of the variable parser. 

### Bootstrap
```
dotc bootstrap [collection]
```
Each collection can define a bootstrap file in the [collection details](#collection-details). This allows collections to define code which should be run only once - good for writing install scripts. `bootstrap` will invoke the specified shell script.

### List collections
```
dotc list
```
`list` will just show a list of all the installed collections.

### Get collection
```
dotc get [collection]
```
`get` will return information about the given collection. It's a good way to see what variables you have configured! 

### Purge
```
dotc purge
```
Had enough? Just want it all gone? `purge` will completely from all dot collections. Careful, this can't be undone. 

## Collections
Collections can be installed without the needing to use this manager as they are simply a "collection" file which just imports a number of other shell scripts.

### Valid collections
There are a few files which are required in order for a collection to be successfully managed by this tool.

#### Collection file 
The collection file must be named `collection.sh` and in the root of the project. In theory, you could place all of your code in here but it's recommended to use this file to import other files such as aliases and functions etc

#### Collection details
A yaml file named `collection.yaml` must be present in the root of the project to provide information about the collection. This file is used to define the variables which can be configured by the user.

#### Variable parser
After configuration, this tool saves variables in a file called `.vars` in the root of the project. In order for these to be included in your collection, the `vars.sh` file should be included in your [collection file](#collection-file).

### Example Collection
https://github.com/wamphlett/dot-collection

## TODO
- [ ] Better error handling
- [ ] Better logging
