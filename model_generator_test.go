package main

import (
	"fmt"
	"testing"
)

func Test_toPascalCase(t *testing.T) {
	fmt.Println(toPascalCase("my_test_table"))
	// Output: MyTestTable
}

func Test_processFilesInDirectory(t *testing.T) {
	dir := "..\\examples"
	count, _ := processFilesInDirectory(dir)
	fmt.Println(count)
	// Output: 1
}
