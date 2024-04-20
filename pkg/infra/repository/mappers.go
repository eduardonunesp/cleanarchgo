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

func fromPgTypeNumericToInt64(numeric pgtype.Numeric) int64 {
	return numeric.Int.Int64()
}

func mapInt64ToPgTypeNumeric(i64 int64) (pgtype.Numeric, error) {
	numeric := pgtype.Numeric{}
	pgInt8 := pgtype.Int8{}
	if err := pgInt8.Scan(i64); err != nil {
		return numeric, err
	}
	if err := numeric.ScanInt64(pgInt8); err != nil {
		return numeric, err
	}
	return numeric, nil
}

func fromTimeToPgTypeTimestamp(d time.Time) pgtype.Timestamp {
	return pgtype.Timestamp{
		Time:  d,
		Valid: true,
	}
}
