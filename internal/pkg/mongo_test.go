package pkg

import "testing"

func TestGetDBNameFromURI(t *testing.T) {
	type args struct {
		uri string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			"OK",
			args{"mongodb://192.168.0.5:9999/foo"},
			"foo",
			false,
		},
		{
			"Fail[no dbname]",
			args{"mongodb://192.168.0.5:9999/"},
			"",
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetDBNameFromURI(tt.args.uri)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetDBNameFromURI() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GetDBNameFromURI() = %v, want %v", got, tt.want)
			}
		})
	}
}
