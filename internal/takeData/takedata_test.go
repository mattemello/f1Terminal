package takedata

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestTakeCar(t *testing.T) {

	t1, _ := time.Parse("2006-01-02T15:04:05.199000+00:00", "2023-08-26T09:30:47.199000+00:00")
	t2, _ := time.Parse("2006-01-02T15:04:05.199000+00:00", "2023-08-26T09:30:47.199000+00:00")
	t3, _ := time.Parse("2006-01-02T15:04:05.199000+00:00", "2023-08-26T09:31:47.199000+00:00")

	pos := []Position{
		{
			Date:         t1,
			DriverNumber: 40,
			MeetingKey:   1217,
			Position:     2,
			SessionKey:   9144,
		},
		{
			Date:         t2,
			DriverNumber: 39,
			MeetingKey:   1217,
			Position:     2,
			SessionKey:   9144,
		},
		{
			Date:         t3,
			DriverNumber: 40,
			MeetingKey:   1217,
			Position:     1,
			SessionKey:   9144,
		},
	}

	theMap := make(map[int]Position)
	theMap[40] = Position{
		Date:         t3,
		DriverNumber: 40,
		MeetingKey:   1217,
		Position:     1,
		SessionKey:   9144,
	}

	theMap[39] = Position{
		Date:         t2,
		DriverNumber: 39,
		MeetingKey:   1217,
		Position:     2,
		SessionKey:   9144,
	}

	clPos := cleanSession(pos)
	assert.Equal(t, theMap, clPos, "The clean does not work")
}
