package sqlx

import (
	"database/sql"
	"time"
)

func NewNullString(val string) sql.NullString {
	return sql.NullString{
		String: val,
		Valid:  val != "",
	}
}

func NewNullInt64(val int64) sql.NullInt64 {
	return sql.NullInt64{
		Int64: val,
		Valid: val != 0,
	}
}

func NewNullFloat64(val float64) sql.NullFloat64 {
	return sql.NullFloat64{
		Float64: val,
		Valid:   val != 0,
	}
}

func NewNullBool(val bool) sql.NullBool {
	return sql.NullBool{
		Bool:  val,
		Valid: val,
	}
}

func NewNullTime(val time.Time) sql.NullTime {
	return sql.NullTime{
		Time:  val,
		Valid: !val.IsZero(),
	}
}

func NewNullBytes(val []byte) sql.NullString {
	return sql.NullString{
		String: string(val),
		Valid:  len(val) > 0,
	}
}
