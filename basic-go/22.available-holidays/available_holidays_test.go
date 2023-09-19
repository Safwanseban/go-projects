package main

import (
	"basic-go/date"
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestCheckDate(t *testing.T) {
	parseTime := func(ti string) time.Time {

		time, err := time.Parse(string(date.DATE_MONTH), ti)
		require.NoError(t, err)
		return time
	}

	type args struct {
		startDate  time.Time
		endDate    time.Time
		holidayMap map[time.Time]bool
	}

	tests := []struct {
		name             string
		args             args
		expectedError    error
		expectedResponse int
	}{
		{

			name: "error- date not provided",
			args: args{

				endDate: time.Date(2023, time.January, 11, 0, 0, 0, 0, time.Local),
				holidayMap: map[time.Time]bool{
					parseTime("10 Jan"): true,
				},
			},
			expectedError: errors.New("please provide valid time"),
		},
		{

			name: "error- endDate is before afterDate",
			args: args{
				startDate: time.Date(2023, time.January, 12, 0, 0, 0, 0, time.Local),
				endDate:   time.Date(2023, time.January, 11, 0, 0, 0, 0, time.Local),
				holidayMap: map[time.Time]bool{
					parseTime("10 Jan"): true,
				},
			},
			expectedError:    errors.New("endDate is before start date"),
			expectedResponse: 0,
		},
		{

			name: "error- not an expected response",
			args: args{
				startDate: time.Date(2023, time.January, 6, 0, 0, 0, 0, time.Local),
				endDate:   time.Date(2023, time.January, 11, 0, 0, 0, 0, time.Local),
				holidayMap: map[time.Time]bool{
					parseTime("10 Jan"): true,
				},
			},
			expectedResponse: 2,
		},
		{

			name: "success",
			args: args{
				startDate: time.Date(2023, time.January, 6, 0, 0, 0, 0, time.Local),
				endDate:   time.Date(2023, time.January, 11, 0, 0, 0, 0, time.Local),
				holidayMap: map[time.Time]bool{
					parseTime("10 Jan"): true,
				},
			},
			expectedResponse: 3,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {

			holidays, err := CheckDates(tc.args.startDate, tc.args.endDate, tc.args.holidayMap)
			if err != nil {
				require.EqualError(t, tc.expectedError, err.Error())
			} else if holidays != tc.expectedResponse {
				require.NotEqual(t, tc.expectedResponse, holidays)
			} else {
				require.Equal(t, tc.expectedResponse, holidays)
			}
		})
	}

}
