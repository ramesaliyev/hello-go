#### `main redeclared in this block`

    In Go, a directory is a package, and a package can only have one function with a given name (with the exception of init(), which is a special case). You think of all your .go files in a directory as separate, but Go does not; it sees a single package, and that package declares multiple functions called main, which is not permitted.
