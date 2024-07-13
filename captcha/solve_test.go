package captcha

import (
	"Menchen13/Capture-Returns/util"
	"fmt"
	"net/http"
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
			args: args{u: "http://10.10.83.191/login"},
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
		fmt.Println("a:a was right!")

	}

}

func Test_getImage(t *testing.T) {
	type args struct {
		resp *http.Response
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{
			name: "Multiplication",
			args: args{resp: util.RespFromFile("Responses/multiplication.html")},
			want: "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",
		},

		{
			name: "Square",
			args: args{resp: util.RespFromFile("Responses/square.html")},
			want: "iVBORw0KGgoAAAANSUhEUgAAAHgAAABvCAYAAAAntwTxAAAAAXNSR0IArs4c6QAAAARnQU1BAACxjwv8YQUAAAAJcEhZcwAADvMAAA7zARxTmToAAAJnSURBVHhe7dY/TmpBGEDx7wKWoiS4ARtXYMMa3AwJsAQKXIMLcA92LoCGgrgCGhpsgfvu3IyveALX5gU8Ob9kAhn+NGfmg6KshLBa+VFQjTd4vV7HfD6P/X6fd3Qprq+v4/HxMdrtdhRFkXf/kQKfMhgMyqurq7LVaqWD4Lqg1el0ypeXl7K6fLnWd40jerlcxm63O35CdDZpqqY+Vce8813jiL69vY3Pz8+4u7uLyWSSd3VOq9Uqnp+fo5qqMRqNYjqd1mP6oBT4lG63W4+Dh4eHvKNzWywWf8f0eDzOu4f5LxrOwHAGhjMwnIHhDAxnYDgDwxkYzsBwBoYzMJyB4QwMZ2A4A8MZGM7AcAaGMzCcgeEMDGdgOAPDGRjOwHAGhjMwnIHhDAxnYDgDwxkYzsBwBoYzMJyB4QwMZ2A4A8MZGM7AcAaGMzCcgeEMDGdgOAPDGRjOwHAGhjMwnIHhDAxnYDgDwxkYzsBwBoYzMJyB4QwMZ2A4A8MZGM7AcAaGMzCcgeEMDGdgOAPDGRjOwHAGhjMwnIHhDAxnYDgDwxkYzsBwBoYzMJyB4QwMZ2A4A8MZGM7AcAaGMzCcgeEMDGdgOAPDGfiXKcsy9vt9tFo/S1dUHyjz84Nubm5is9nE/f19vL295V2dS8r18fERT09PdejRaBSz2Sy/ekAKfEq3200HwHUhqyiKst1u1ys9H4/HudRhjfe83+9H9WU/Hgn6v6pmsdvt6pW69Hq9+iYf0zii39/f4/X1Nbbbbd7RpUiXbzgc1pGPXcDGwF+nI70tnRhdjq821aiu1yGNgfW7+cOKFvEHKNmQ9kvbvgIAAAAASUVORK5CYIKMWIgRI0aMGDFixIgRCzFixIgRI0aMGDFiIUaMGDFixIgRI0YsxIgRI0aMGDFixIiFGDFixIgRI0aMGLEQI0aMGDFixIgRIxZixIgRH9CyMs0hxHs92G3VzTyZWfFxVzzrrPh47WpXGWRwALEeaZAuo3SIj9dolMlht1gPN03OMh0G8fHqhpkk4z8gvsg0OSfYMOdJrWrZQHyVLAaDPGXXr/PkXd3WpoH484uc/HqzqE/DPPtyz8m994v6/csLcP07vfpYtd62EO9rd1dV5XHRr6pVbavpFt9UbWr9w/u4Z9dV2/XtpmXF21pX1c6Ke694v65qOhRVVXfrWttnz7+NV1Xbql0jsf5GiBEjFmLE/0E/Ac8azp1AFifiAAAAAElFTkSuQmCC",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getImage(tt.args.resp); got != tt.want {
				t.Errorf("getImage() = %v, want %v", got, tt.want)
			}
		})
	}
}
