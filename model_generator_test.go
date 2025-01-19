package main

import (
	"fmt"
	"log"
	"os"
	"testing"
)

func Test_toPascalCase(t *testing.T) {
	fmt.Println(toPascalCase("my_test_table"))
	got := toPascalCase("my_test_table")
	// Output: MyTestTable

	if got != "MyTestTable" {
		t.Errorf("Generated name - %snot expected,MyTestTable", got)
	}
}

func Test_processFilesInDirectory(t *testing.T) {
	dir := "."
	count, _ := processFilesInDirectory(dir)
	fmt.Println(count)
	// Output: 2

	if count != 2 {
		t.Error("Count files in folder not equal expected")
	}
	e := os.Remove("my_table.cs")
	e1 := os.Remove("my_table1.cs")

	if e != nil || e1 != nil {
		log.Fatal(e)
	}
}

func Test_generate_cs_model(t *testing.T) {
	generatedString := generateCSharpModel("my_table1", []string{"my_test_field UUID NOT NULL", "another_field VARCHAR(255)", "my_date_field DATE NOT NULL"})
	expectedString := `public class MyTable1
{
    [Column("my_test_field")]
    [Required]
    public Guid MyTestField { get; set; }

    [Column("another_field")]
    public string AnotherField { get; set; }

    [Column("my_date_field")]
    [Required]
    public DateTime MyDateField { get; set; }

}`
	fmt.Println(generatedString)
	fmt.Println(expectedString)
	if generatedString != expectedString {
		t.Error("Generated string not equal expected")
	}
}
