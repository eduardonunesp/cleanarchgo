package repository

import (
	"fmt"
	"time"

	"github.com/jackc/pgx/v5/pgtype"
)

func fromPgTypeUUIDToString(uuid pgtype.UUID) string {
	return fmt.Sprintf("%x-%x-%x-%x-%x", uuid.Bytes[0:4], uuid.Bytes[4:6], uuid.Bytes[6:8], uuid.Bytes[8:10], uuid.Bytes[10:16])
}

func mapStringToPgTypeUUID(strUUID string) (pgtype.UUID, error) {
	var uuid pgtype.UUID
	if err := uuid.Scan(strUUID); err != nil {
		return uuid, err
	}
	return uuid, nil
}

func fromStringToPgTypeText(str string) pgtype.Text {
	return pgtype.Text{String: str, Valid: true}
}

func fromPgTypeNumericToString(numeric pgtype.Numeric) string {
	if !numeric.Valid {
		return ""
	}
	v, _ := numeric.MarshalJSON()
	return string(v)
}

func mapStringToPgTypeNumeric(strNumeric string) (pgtype.Numeric, error) {
	if strNumeric == "" {
		return pgtype.Numeric{}, nil
	}
	numeric := pgtype.Numeric{}
	if err := numeric.Scan(strNumeric); err != nil {
		return numeric, err
	}
	return numeric, nil
}

func fromTimeToPgTypeTimestamp(d int64) pgtype.Timestamp {
	return pgtype.Timestamp{
		Time:  time.Unix(d, 0),
		Valid: true,
	}
}
