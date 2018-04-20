#!/bin/bash

if [ $# -ne 1 ]; then
    echo "please input the number of the leetcode!"
    exit 1
fi

mkdir ./src/$1

echo "package main

func main() {

}" >> ./src/$1/main.go