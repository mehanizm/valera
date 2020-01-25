package main

import (
	"reflect"
	"testing"
)

func Test_parseConfigFromFile(t *testing.T) {
	type args struct {
		filePath string
	}
	tests := []struct {
		name     string
		filePath string
		want     *config
		wantErr  bool
	}{
		{
			name:     "Case 1. Not existing file",
			filePath: "test/test_config.yaml",
			want:     nil,
			wantErr:  true,
		},
		{
			name:     "Case 2. Not valid yaml file",
			filePath: "test/test_config_1.yaml",
			want:     nil,
			wantErr:  true,
		},
		{
			name:     "Case 3. Positive",
			filePath: "test/test_config_2.yaml",
			want: &config{
				ProxyURL: "url:port",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := parseConfigFromFile(tt.filePath)
			if (err != nil) != tt.wantErr {
				t.Errorf("parseConfigFromFile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parseConfigFromFile() = %v, want %v", got, tt.want)
			}
		})
	}
}
