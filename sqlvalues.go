package sqlvalues

import (
	"database/sql"
	"encoding/json"
	"strings"
	"time"
)

func SqlString(s string) sql.NullString {
	return sql.NullString{
		String: s,
		Valid:  strings.TrimSpace(s) != "",
	}
}

func SqlStringPtr(t *string) sql.NullString {
	if t == nil {
		return sql.NullString{}
	}

	return sql.NullString{
		String: *t,
		Valid:  strings.TrimSpace(*t) != "",
	}
}

func SqlInt32(i int) sql.NullInt32 {
	return sql.NullInt32{
		Int32: int32(i),
		Valid: true,
	}
}

func sqlNullJSONB(b []byte) sql.NullString {
	if len(b) == 0 {
		return sql.NullString{}
	}

	if string(b) == "null" {
		return sql.NullString{}
	}

	if string(b) == "{}" {
		return sql.NullString{}
	}

	return sql.NullString{
		String: string(b),
		Valid:  true,
	}
}

func SqlNullJSONB(x any) sql.NullString {
	b, _ := json.Marshal(x)

	return sqlNullJSONB(b)
}

func SqlTime(t time.Time) sql.NullTime {
	return sql.NullTime{
		Time:  t,
		Valid: !t.IsZero(),
	}
}

func SqlTimePtr(t *time.Time) sql.NullTime {
	if t == nil {
		return sql.NullTime{}
	}

	return sql.NullTime{
		Time:  *t,
		Valid: !t.IsZero(),
	}
}

func SqlNullFloat64(f float64) sql.NullFloat64 {
	return sql.NullFloat64{
		Float64: f,
		Valid:   true,
	}
}
