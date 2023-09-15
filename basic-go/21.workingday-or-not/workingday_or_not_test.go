package main

import (
	"basic-go/date"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestIsHoliday(t *testing.T) {
	parseTime := func(ti string) time.Time {

		time, err := time.Parse(string(date.DATE_MONTH), ti)
		require.NoError(t, err)
		return time
	}

	tests := []struct {
		name             string
		dateTocheck      time.Time
		holiDayList      map[time.Time]bool
		exPectedResponse string
	}{

		{
			name:        "result- not a working day resultDate is in HolidayList",
			dateTocheck: time.Date(2023, time.January, 10, 0, 0, 0, 0, time.Local),
			holiDayList: map[time.Time]bool{
				parseTime("02 Jan"): true,
				parseTime("10 Jan"): true,
			},
			exPectedResponse: "date is not a working day",
		},
		{
			name:        "result- not a working day resultDate is a WeekendDay",
			dateTocheck: time.Date(2023, time.January, 8, 0, 0, 0, 0, time.Local),
			holiDayList: map[time.Time]bool{
				parseTime("02 Jan"): true,
			},
			exPectedResponse: "date is not a working day",
		},
		{
			name:        "result- working day",
			dateTocheck: time.Date(2023, time.January, 12, 0, 0, 0, 0, time.Local),
			holiDayList: map[time.Time]bool{

				parseTime("11 Jan"): true,
			},
			exPectedResponse: "date is a working day",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {

			if isHoliday(tc.dateTocheck, tc.holiDayList) {
				require.Equal(t, tc.exPectedResponse, "date is not a working day")
			} else {
				require.Equal(t, tc.exPectedResponse, "date is a working day")
			}
		})

	}
}
