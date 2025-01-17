# modelGenerator
Generate EF entities from postgres create table script - add needed data attributes.

For example we have sql script like

CREATE TABLE my_table (
    my_test_field2 UUID NOT NULL,
    another_field2 VARCHAR(255),
    my_date_field2 DATE NOT NULL
);

After running script we should have 

public class MyTable
{
    [Column("my_test_field2")]
    [Required]
    public Guid MyTestField2 { get; set; }

    [Column("another_field2")]
    public string AnotherField2 { get; set; }

    [Column("my_date_field2")]
    [Required]
    public DateTime MyDateField2 { get; set; }

}

So script first version 1.0.0 do the next:
    For each script in folder create .cs file with PascalCase Name convention.
    Add if need [Required] attributes
    Add [Column] attribute

File examples could be find in examples folder
Executable files in executables folder


The plans are the next:
    Write tests for current functionality
    Write tests for the new features
    Add new features
    Add Jenkinsfile with next stages:
        build
        run test
        copy executable with updated version in executable folder
    
    
