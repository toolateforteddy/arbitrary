# Arbitrary
The goal of this package is to offer utilities for handling well structured
though untyped data structures. These data structures often will be found
when hydrating unknown data from a json blob, or from a structured file.

## Dig
Digging is used for diving into the data to find the value at a specified path.

## Flatten
Flatten will walk the whole tree and create a new map[string]interface{} where
all values in that map are basic datatypes. The path from the original structure
will be converted into a string key. The default joiner will be '.'

## Validate
Validate will walk the unstructured data to check that all leaf nodes are valid
types.

## Hydrate
Hydrate will convert the source arbitrary object into the dest typed object.
