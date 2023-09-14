package main

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestReadAndPut(t *testing.T) {
	tests := []struct {
		Name         string
		FileInput    string
		ExpectError  error
		ExpectResult map[string]int
	}{
		{
			Name:         "error- file opening",
			FileInput:    "./file",
			ExpectError:  errors.New("open ./file: no such file or directory"),
			ExpectResult: nil,
		},
		{
			Name:        "success",
			FileInput:   "./file.txt",
			ExpectError: nil,
			ExpectResult: map[string]int{

				"beautiful,": 1, "gift": 1, "is": 2, "nature": 1, "testok": 1, "tree": 1,
			},
		},
	}

	for _, tc := range tests {

		t.Run(tc.Name, func(t *testing.T) {

			data, err := readAndPut(tc.FileInput)
			if err != nil {

				require.EqualError(t, tc.ExpectError, err.Error())
			} else {

				require.Equal(t, tc.ExpectResult, data)
			}
		})
	}
}
