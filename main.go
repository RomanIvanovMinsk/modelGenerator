package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

// Mapping SQL types to C# types
var sqlToCSharpType = map[string]string{
	"uuid":      "Guid",
	"int":       "int",
	"bigint":    "long",
	"smallint":  "short",
	"decimal":   "decimal",
	"numeric":   "decimal",
	"varchar":   "string",
	"nvarchar":  "string",
	"text":      "string",
	"boolean":   "bool",
	"date":      "DateTime",
	"timestamp": "DateTime",
	"double":    "double",
}

// Function to convert snake_case to PascalCase
func toPascalCase(input string) string {
	parts := strings.Split(input, "_")
	for i := range parts {
		parts[i] = strings.Title(parts[i])
	}
	return strings.Join(parts, "")
}

func processFilesInDirectory(dir string) error {
	// Проходим по всем файлам и поддиректориям в заданной директории
	return filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Если это не .cs файл, пропускаем его
		if !info.IsDir() && strings.HasSuffix(path, ".sql") {
			fmt.Println("Обрабатываем файл:", path)
			return processFile(path)
		}
		return nil
	})
}

func processFile(path string) error {
	// Open the SQL script file
	file, err := os.Open(path)
	if err != nil {
		fmt.Println("Ошибка открытия файла:", err)
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var tableName string
	var fields []string

	reTable := regexp.MustCompile(`(?i)^create\s+table\s+(\w+)`)
	reField := regexp.MustCompile(`^\s*(\w+)\s+(\w+.*)`)

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		// Match table name
		if matches := reTable.FindStringSubmatch(line); matches != nil {
			tableName = matches[1]
			continue
		}

		// Match fields
		if matches := reField.FindStringSubmatch(line); matches != nil {
			fields = append(fields, matches[0])
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Ошибка чтения файла:", err)
		return err
	}

	// Generate and print the C# model
	if tableName != "" {
		os.WriteFile(tableName+".cs", []byte(generateCSharpModel(tableName, fields)), 0644)
	} else {
		fmt.Println("Не удалось найти таблицу в SQL скрипте.")
	}

	return nil
}

// Function to generate C# model from SQL field definitions
func generateCSharpModel(tableName string, fields []string) string {
	var sb strings.Builder
	sb.WriteString(fmt.Sprintf("public class %s\n{\n", toPascalCase(tableName)))

	for _, field := range fields {
		re := regexp.MustCompile(`(\w+)\s+(\w+)`)
		matches := re.FindStringSubmatch(field)

		if len(matches) == 3 {
			fieldName := toPascalCase(matches[1])
			sqlType := strings.ToLower(matches[2])
			csharpType := sqlToCSharpType[sqlType]

			sb.WriteString(fmt.Sprintf("    [Column(\"%s\")]\n", matches[1]))
			if strings.Contains(strings.ToLower(field), "not null") {
				sb.WriteString(fmt.Sprintf("    [Required]\n"))
			}
			sb.WriteString(fmt.Sprintf("    public %s %s { get; set; }\n\n", csharpType, fieldName))
		}
	}

	sb.WriteString("}")
	return sb.String()
}

func main() {

	// Указываем директорию для обработки, например, текущую директорию
	dir := "."

	if err := processFilesInDirectory(dir); err != nil {
		fmt.Println("Ошибка при обработке файлов:", err)
	}
}
