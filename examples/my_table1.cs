public class MyTable1
{
    [Column("my_test_field")]
    [Required]
    public Guid MyTestField { get; set; }

    [Column("another_field")]
    public string AnotherField { get; set; }

    [Column("my_date_field")]
    [Required]
    public DateTime MyDateField { get; set; }

}