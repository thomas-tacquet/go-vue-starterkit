package helpers

import "fmt"

func DatabaseFormat(host, port, user, password, dbname, sslMode, schema string) string {
	const toBeFormatted = "host=%s port=%s user=%s password=%s dbname=%s sslmode=%s search_path=%s"
	return fmt.Sprintf(toBeFormatted, host, port, user, password, dbname, sslMode, schema)
}
