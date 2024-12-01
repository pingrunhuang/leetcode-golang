# Summary
Golang is great, leetcode is awesome. Learning Golang by solve interesting algorithms in leetcode is such a geek way to go!!

## project structure
To be honest, as I'm new to golang, I ask gpt for help to generate a resonable structure of such project. Here is the explanation of each file directory:

```
leetcode-golang/
├── cmd/                    # (Optional) For CLI tools or utilities
├── solutions/              # Solutions grouped by problem or category
│   ├── arrays/
│   │   ├── two_sum.go      # Solution for Two Sum problem
│   │   ├── two_sum_test.go # Unit test for Two Sum problem
│   ├── strings/
│   │   ├── longest_substring.go
│   │   ├── longest_substring_test.go
│   └── trees/
│       ├── binary_tree.go
│       ├── binary_tree_test.go
├── pkg/                    # Reusable helper functions or data structures
│   ├── utils.go            # Utility functions (e.g., input parsing)
│   ├── utils_test.go
├── go.mod                  # Go module file
├── go.sum                  # Dependencies file
└── README.md               # Documentation for the project
```

## How to run all test?
run all tests recursively
```bash
go test ./...
```

## How to run the specific test?
use `go test` to run unit test for specific test of solution. 
For example:
```bash
go test ./solutions/arrays -run TwoSumTest
```
