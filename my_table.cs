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