package main

import (
	"basic-go/date"
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestCheckDates(t *testing.T) {
	parseTime := func(ti string) time.Time {

		time, err := time.Parse(string(date.DATE_MONTH), ti)
		require.NoError(t, err)
		return time
	}
	type args struct {
		inputDate      time.Time
		BuissinessDays int
		holidayMap     map[time.Time]bool
	}

	tests := []struct {
		name           string
		args           args
		expectedError  error
		expectedResult string
	}{
		{
			name: "error- not a valid time",
			args: args{

				holidayMap: map[time.Time]bool{
					parseTime("18 sep"): true,
				},
				BuissinessDays: 2,
			},
			expectedError: errors.New("provide valid time"),
		},
		{
			name: "result- no buissiness days passed",
			args: args{
				inputDate: time.Date(2023, time.September, 15, 0, 0, 0, 0, time.Local),
				holidayMap: map[time.Time]bool{
					parseTime("18 sep"): true,
				},
			},
			expectedResult: "15-09-2023",
		},
		{
			name: "result- no only holiday provided",
			args: args{
				inputDate: time.Date(2023, time.September, 15, 0, 0, 0, 0, time.Local),
				holidayMap: map[time.Time]bool{
					parseTime("15 sep"): true,
				},
				BuissinessDays: 1,
			},
			expectedResult: "18-09-2023",
		},
		{
			name: "result- no only holiday provided",
			args: args{
				inputDate: time.Date(2023, time.September, 15, 0, 0, 0, 0, time.Local),
				holidayMap: map[time.Time]bool{
					parseTime("15 sep"): true,
				},
				BuissinessDays: 12,
			},
			expectedResult: "18-09-2023",
		},

		{
			name: "success- postive buisiness days",
			args: args{
				inputDate: time.Date(2023, time.September, 15, 0, 0, 0, 0, time.Local),
				holidayMap: map[time.Time]bool{
					parseTime("18 sep"): true,
				},
				BuissinessDays: 2,
			},
			expectedError:  nil,
			expectedResult: "20-09-2023",
		},

		{
			name: "success- negative buisiness days",
			args: args{
				inputDate: time.Date(2023, time.September, 15, 0, 0, 0, 0, time.Local),
				holidayMap: map[time.Time]bool{
					parseTime("14 sep"): true,
				},
				BuissinessDays: -2,
			},
			expectedError:  nil,
			expectedResult: "12-09-2023",
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {

			result, err := CheckDates(tc.args.inputDate, tc.args.BuissinessDays, tc.args.holidayMap)
			if err != nil {
				require.EqualError(t, tc.expectedError, err.Error())
			} else {
				require.Equal(t, tc.expectedResult, result)
			}
		})
	}
}
