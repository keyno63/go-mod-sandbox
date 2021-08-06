package gzip

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func TestGzipWrite(t *testing.T) {
	type args struct {
		value string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "test1",
			args: args{
				value: "test1",
			},
			want:    "test1",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GzipWrite(tt.args.value)
			if (err != nil) != tt.wantErr {
				t.Errorf("GzipWrite() error = %+v, wantErr %+v", err, tt.wantErr)
				return
			}

			actual, err := GzipRead(got)
			if err != nil {
				fmt.Println(err.Error())
				t.Errorf("GzipWrite() failed to read ret gzip value. got = %+v, value %+v", got, tt.args.value)
			}
			if actual != tt.want {
				t.Errorf("GzipWrite() actual = %+v, want %+v", actual, tt.want)
			}
		})
	}
}

func TestGzipHttpWriter(t *testing.T) {
	type args struct {
		statusCode int
		headerMap  map[string]string
		body       string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "200 レスポンスなど",
			args: args{
				statusCode: 200,
				headerMap: map[string]string{
					"Content-Type": "plain/text",
				},
				body: "sample body",
			},
		},
		{
			name: "エラーレスポンス",
			args: args{
				statusCode: 200,
				headerMap: map[string]string{
					"Content-Type": "plain/text",
				},
				body: "sample error body",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			GzipHttpWriter(w, tt.args.statusCode, tt.args.headerMap, tt.args.body)
			result := w.Result()
			// status code
			if !reflect.DeepEqual(result.StatusCode, tt.args.statusCode) {
				t.Errorf("failed to compare. got = %+v, want = %+v", result.StatusCode, tt.args.statusCode)
				return
			}
			// header
			if !reflect.DeepEqual(result.Header, createHeader(tt.args.headerMap)) {
				t.Errorf("failed to compare. got = %+v, want = %+v", result.Header, tt.args.headerMap)
				return
			}

			// body
			body, err := GzipRead(w.Body.String())
			if err != nil {
				t.Errorf("failed to read. reason=[%s].", err.Error())
				return
			}
			if !reflect.DeepEqual(body, tt.args.body) {
				t.Errorf("failed to compare body. got = %+v, want = %+v", body, tt.args.body)
				return
			}
		})
	}
}

func createHeader(m map[string]string) http.Header {
	h := http.Header{}
	for k, v := range m {
		h.Add(k, v)
	}
	return h
}
