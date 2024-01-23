package common

import "strings"

func ContainsSQLorXSS(input string) bool {
	input = strings.ToLower(input)

	// Check for common SQL injection patterns
	sqlPatterns := []string{
		";", "--", "/*", "*/",
		"' ", " '", "drop table", "drop database",
		"select *", "select * from", "select count(*) from",
		"<script>", "</script>",
		"<script", "</script",
		"<script/>", "<script />",
		"<script/**/>", "<script/** />",
		"<script/**/**/>", "<script/**/** />",
		"<script/**/ />", "<script/**/ >",
	}

	for _, pattern := range sqlPatterns {
		if strings.Contains(input, pattern) {
			return true
		}
	}

	return false
}
