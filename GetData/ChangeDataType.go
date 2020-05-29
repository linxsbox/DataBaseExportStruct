package GetData

import (
	"strings"
)

func dataBaseChangeType(val string) string {
	if strings.Contains(val, "varchar") {
		return "string"
	}

	if strings.Contains(val, "text") {
		return "string"
	}

	if strings.Contains(val, "tinyint") {
		if strings.Contains(val, "unsigned") {
			return "uint8"
		}
		return "int8"
	}

	if strings.Contains(val, "smallint") {
		if strings.Contains(val, "unsigned") {
			return "uint16"
		}
		return "int16"
	}

	if strings.Contains(val, "mediumint") {
		if strings.Contains(val, "unsigned") {
			return "uint32"
		}
		return "int32"
	}

	if strings.Contains(val, "bigint") {
		if strings.Contains(val, "unsigned") {
			return "uint64"
		}
		return "int64"
	}

	if strings.Contains(val, "int") {
		if strings.Contains(val, "unsigned") {
			return "uint"
		}
		return "int"
	}

	if strings.Contains(val, "blob") {
		return "byte"
	}

	if strings.Contains(val, "doubel") {
		return "float64"
	}

	if strings.Contains(val, "decimal") {
		return "float64"
	}

	if strings.Contains(val, "float") {
		return "float32"
	}

	if strings.Contains(val, "real") {
		return "float32"
	}

	if strings.Contains(val, "bit") {
		return "bool"
	}

	if strings.Contains(val, "time") {
		IsTime = "\"time\""
		return "time.Time"
	}

	if strings.Contains(val, "date") {
		IsTime = "\"time\""
		return "time.Time"
	}

	if strings.Contains(val, "year") {
		IsTime = "\"time\""
		return "time.Time"
	}
	return val
}

func tableGetSetChangeToLower(ts string) string {
	if ts != "" {
		ctos := string(ts[0])
		stol := strings.ToLower(ctos)
		return strings.Replace(
			ts, ctos, strings.ToLower(stol), 1)
	}
	return ""
}
