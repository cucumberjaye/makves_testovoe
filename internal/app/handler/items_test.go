package handler

import (
	"encoding/json"
	makves "github.com/cucumberjaye/makves_testovoe"
	"github.com/cucumberjaye/makves_testovoe/internal/app/service/mock"
	"github.com/stretchr/testify/require"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandler_GetItem(t *testing.T) {
	type want struct {
		code int
		body any
	}
	tests := []struct {
		name string
		way  string
		want want
	}{
		{
			name: "ok",
			way:  "/get-items/1",
			want: want{
				code: 200,
				body: makves.Item{},
			},
		},
		{
			name: "fail_400",
			way:  "/get-items/0",
			want: want{
				code: 400,
				body: makves.Item{},
			},
		},
		{
			name: "fail_404",
			way:  "/get-items/",
			want: want{
				code: 404,
				body: makves.Item{},
			},
		},
	}

	svc := &mock.Mock{}
	handler := New(svc)

	r := handler.InitRoutes()
	ts := httptest.NewServer(r)
	defer ts.Close()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Log(ts.URL + tt.way)
			request := httptest.NewRequest(http.MethodGet, ts.URL+tt.way, nil)
			request.RequestURI = ""

			resp, err := http.DefaultClient.Do(request)
			require.NoError(t, err)

			require.Equal(t, tt.want.code, resp.StatusCode)

			if tt.want.code == 200 {
				defer resp.Body.Close()
				resBody, err := io.ReadAll(resp.Body)
				require.NoError(t, err)

				testBody, err := json.Marshal(tt.want.body)
				require.NoError(t, err)
				require.Equal(t, testBody, resBody)
			}
		})
	}
}

func TestHandler_GetItems(t *testing.T) {
	type want struct {
		code int
		body any
	}
	tests := []struct {
		name string
		way  string
		want want
	}{
		{
			name: "ok",
			way:  "/get-items/1-10",
			want: want{
				code: 200,
				body: []makves.Item{},
			},
		},
		{
			name: "fail_400",
			way:  "/get-items/0-10",
			want: want{
				code: 400,
				body: []makves.Item{},
			},
		},
		{
			name: "fail_empty_400",
			way:  "/get-items/-",
			want: want{
				code: 400,
				body: []makves.Item{},
			},
		},
		{
			name: "fail_404",
			way:  "/get-items/",
			want: want{
				code: 404,
				body: []makves.Item{},
			},
		},
	}

	svc := &mock.Mock{}
	handler := New(svc)

	r := handler.InitRoutes()
	ts := httptest.NewServer(r)
	defer ts.Close()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Log(ts.URL + tt.way)
			request := httptest.NewRequest(http.MethodGet, ts.URL+tt.way, nil)
			request.RequestURI = ""

			resp, err := http.DefaultClient.Do(request)
			require.NoError(t, err)

			require.Equal(t, tt.want.code, resp.StatusCode)

			if tt.want.code == 200 {
				defer resp.Body.Close()
				resBody, err := io.ReadAll(resp.Body)
				require.NoError(t, err)

				testBody, err := json.Marshal(tt.want.body)
				require.NoError(t, err)
				require.Equal(t, testBody, resBody)
			}
		})
	}
}
