package test

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"gopkg.in/DATA-DOG/go-sqlmock.v1"
	"net/url"
	"strings"
)

func Test_handler(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	rows := sqlmock.NewRows([]string{"id", "title", "description"}).
		AddRow(1, "one", "hello")

	mock.ExpectQuery(".*").WillReturnRows(rows)

	handler := getHandler(db)
	type args struct {
		res *httptest.ResponseRecorder
		req *http.Request
	}
	type Input struct {
		name     string
		args     args
		expected string
	}
	tests := []Input{
		{"function to handle request",
			args{
				httptest.NewRecorder(),
				httptest.NewRequest("", "localhost:8000", nil)},
			`[{"id":1,"title":"one","description":"hello"}]`},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			handler(tt.args.res, tt.args.req)
			w := tt.args.res
			body, _ := ioutil.ReadAll(w.Body)
			if string(body) != tt.expected {
				t.Errorf("got %v, want %v", string(body), tt.expected)
			}
		})
	}
}

func Test_getCreateHandler(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	createHandler := create(db)

	mock.ExpectExec("/write/")
	type args struct {
		res *httptest.ResponseRecorder
		req *http.Request
	}
	type Input struct {
		name     string
		args     args
	}

	data := url.Values{}
	data.Set("title", "foo")
	data.Add("description", "bar")

	tests := []Input{
		{"function to handle request",
			args{
				httptest.NewRecorder(),
				httptest.NewRequest("", "localhost:8000",strings.NewReader(data.Encode()))},
	},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			createHandler(tt.args.res,tt.args.req)
			if tt.args.res.Body.String() != "description=bar&title=foo" {
				t.Fatal("body not found in request")
			}
		})
	}
}
