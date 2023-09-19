package main

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"regexp"
	"testing"
	"text/template"

	"github.com/stretchr/testify/require"
)

func TestReadAndCompress(t *testing.T) {

	tests := []struct {
		Name string

		ResultPath     string
		ExpectedError  error
		expectedResult string
	}{

		{
			Name:          "err- no file found",
			ResultPath:    "testdata/tst.html",
			ExpectedError: errors.New("open testdata/tst.html: no such file or directory"),
		},
		{
			Name:           "success",
			ResultPath:     "testdata/test_response.html",
			ExpectedError:  nil,
			expectedResult: "testdata/test_response.html.gz",
		},
	}

	for _, tc := range tests {
		t.Run(tc.Name, func(t *testing.T) {
			file, err := ReadAndCompress(tc.ResultPath)
			if err != nil {
				require.EqualError(t, tc.ExpectedError, err.Error())
			} else {
				require.Equal(t, tc.expectedResult, file)
			}
		})
	}

}

func TestFetchAndSetOuput(t *testing.T) {
	//creates a mockServer and a Handler function,Handler function mocks the output
	//given by actual server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//parsing html file and creating output
		tpl := template.Must(template.ParseFiles("testdata/test.html"))
		tpl.ExecuteTemplate(w, "test.html", nil)
	}))
	defer server.Close()
	tests := []struct {
		name             string
		inputUrl         string
		outPutFile       string
		expectedError    error
		expectStatusCode int
	}{
		{

			name:             "error- not a valid url",
			inputUrl:         "notValid",
			expectedError:    errors.New(regexp.QuoteMeta("Get \"/notValid\": unsupported protocol scheme \"\"")),
			expectStatusCode: http.StatusNotFound,
		},
		{

			name:             "error- can't find file",
			inputUrl:         server.URL,
			expectedError:    errors.New("open .: is a directory"),
			expectStatusCode: http.StatusBadRequest,
		},
		{

			name:             "success",
			inputUrl:         server.URL,
			outPutFile:       "testdata/test_response.html",
			expectedError:    nil,
			expectStatusCode: http.StatusOK,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {

			code, err := FetchAndSetOutput(context.Background(), tc.inputUrl, tc.outPutFile)
			if err != nil {
				require.EqualError(t, tc.expectedError, err.Error())
			} else {
				require.Equal(t, tc.expectStatusCode, code)
			}
		})
	}

}
