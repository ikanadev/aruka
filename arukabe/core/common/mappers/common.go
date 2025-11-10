package mappers

import (
	"time"

	"google.golang.org/protobuf/types/known/timestamppb"
)

func TimestampFromTime(t *time.Time) *timestamppb.Timestamp {
	var timestamp *timestamppb.Timestamp
	if t != nil {
		timestamp = timestamppb.New(*t)
	}
	return timestamp
}
