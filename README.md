* put this package inside your GOPATH

* in order to run all tests the sub directory, you will need to use command `go test ./...` from the root directory of this project. 
```
An import path is a pattern if it includes one or more "..." wildcards, each of which can match any string, including the empty string and strings containing slashes.

Such a pattern expands to all package directories found in the GOPATH trees with names matching the patterns.

As a special case, x/... matches x as well as x's subdirectories.
For example, net/... expands to net and packages in its subdirectories.
```