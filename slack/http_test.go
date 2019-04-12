package slack

import (
	"errors"
	"io"
	"io/ioutil"
	"reflect"
	"testing"
)

type MockJsonHttpClient struct {
	PostErr error
}

func (mock MockJsonHttpClient) Post(url string, body io.Reader) ([]byte, error) {
	bytes, _ := ioutil.ReadAll(body)
	return bytes, mock.PostErr
}

func TestNewHttpClient(t *testing.T) {
	if client := NewHttpClient(); client == nil {
		t.Errorf("Expected type: %s to be *HttpClient", reflect.TypeOf(client))
	}
}

func TestHttpClient_Send(t *testing.T) {
	type testFields struct {
		url     string
		message Message
		httpClient HttpContract
	}

	type testResult struct {
		bytes string
		err   error
	}

	tests := []struct {
		name   string
		fields testFields
		want   testResult
	}{
		{
			name:   "Executes Successfully",
			fields: testFields{
				url: "foo",
				message: Message{
					Text: "Example Text",
				},
				httpClient: MockJsonHttpClient{},
			},
			want:   testResult{
				bytes: `{"text":"Example Text","attachments":null}`,
				err: nil,
			},
		},
		{
			name: "Err on Client.Post",
			fields: testFields{
				url: "foo",
				message: Message{
					Text: "Example Text",
				},
				httpClient: MockJsonHttpClient{
					PostErr: errors.New("unexpected error"),
				},
			},
			want:   testResult{
				bytes: `{"text":"Example Text","attachments":null}`,
				err: errors.New("unexpected error"),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client := &HttpClient{
				Client: tt.fields.httpClient,
			}

			b, err := client.Send(tt.fields.url, tt.fields.message)

			if got := string(b); got != tt.want.bytes {
				t.Errorf("Got response body = %v, want %v", got, tt.want)
			}

			if got := err; got != nil && got.Error() != tt.want.err.Error() {
				t.Errorf("Got error = %v, want %v", got, tt.want)
			}
		})
	}
}
