package captcha

import (
	"Menchen13/Capture-Returns/util"
	"testing"
)

func Test_shape(t *testing.T) {
	type args struct {
		b64encoded string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name:    "square",
			args:    args{b64encoded: getImage(util.RespFromFile("Responses/square.html"))},
			want:    "square",
			wantErr: false,
		},
		{
			name:    "triangle",
			args:    args{b64encoded: getImage(util.RespFromFile("Responses/triangle.html"))},
			want:    "triangle",
			wantErr: false,
		},
		{
			name:    "circle",
			args:    args{b64encoded: getImage(util.RespFromFile("Responses/circle.html"))},
			want:    "circle",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := shape(tt.args.b64encoded)
			if (err != nil) != tt.wantErr {
				t.Errorf("shape() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("shape() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_term(t *testing.T) {
	type args struct {
		base64encoded string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name:    "multiplication",
			args:    args{base64encoded: getImage(util.RespFromFile("Responses/multiplication.html"))},
			want:    "8008",
			wantErr: false,
		},
		{
			name:    "addition",
			args:    args{base64encoded: "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"},
			want:    "380",
			wantErr: false,
		},
		{
			name:    "subtraction",
			args:    args{base64encoded: "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"},
			want:    "388",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := term(tt.args.base64encoded)
			if (err != nil) != tt.wantErr {
				t.Errorf("term() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("term() = %v, want %v", got, tt.want)
			}
		})
	}
}
