package captcha

import (
	"Menchen13/Capture-Returns/brute"
	"testing"
)

func TestSolver(t *testing.T) {
	// t.Skip("Skippting till external server available")
	type args struct {
		u string
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{
			name: "Solver Test",
			//this address is temporary
			args: args{u: "http://10.10.3.117/login"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			toCaptcha(t, tt.args.u)
			Solver(tt.args.u)
		})
	}
}

func toCaptcha(t *testing.T, u string) {
	t.Helper()
	for i := 0; i < 3; i++ {
		brute.Orca(u, "a", "a")

	}

}
