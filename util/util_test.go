package util

import (
	"fmt"
	"testing"
)

func TestRespFromFile(t *testing.T) {
	//path needs to match layout of dev container.
	resp := RespFromFile("Responses/circle.html")
	if resp == nil {
		t.Fatal("resp is nil")
	}

	fmt.Println(resp.Status)
}

func TestB64ToFile(t *testing.T) {
	t.SkipNow()
	type args struct {
		b64encoded string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name:    "Testfile: square.png",
			args:    args{b64encoded: "iVBORw0KGgoAAAANSUhEUgAAAHgAAABvCAYAAAAntwTxAAAAAXNSR0IArs4c6QAAAARnQU1BAACxjwv8YQUAAAAJcEhZcwAADvMAAA7zARxTmToAAAJnSURBVHhe7dY/TmpBGEDx7wKWoiS4ARtXYMMa3AwJsAQKXIMLcA92LoCGgrgCGhpsgfvu3IyveALX5gU8Ob9kAhn+NGfmg6KshLBa+VFQjTd4vV7HfD6P/X6fd3Qprq+v4/HxMdrtdhRFkXf/kQKfMhgMyqurq7LVaqWD4Lqg1el0ypeXl7K6fLnWd40jerlcxm63O35CdDZpqqY+Vce8813jiL69vY3Pz8+4u7uLyWSSd3VOq9Uqnp+fo5qqMRqNYjqd1mP6oBT4lG63W4+Dh4eHvKNzWywWf8f0eDzOu4f5LxrOwHAGhjMwnIHhDAxnYDgDwxkYzsBwBoYzMJyB4QwMZ2A4A8MZGM7AcAaGMzCcgeEMDGdgOAPDGRjOwHAGhjMwnIHhDAxnYDgDwxkYzsBwBoYzMJyB4QwMZ2A4A8MZGM7AcAaGMzCcgeEMDGdgOAPDGRjOwHAGhjMwnIHhDAxnYDgDwxkYzsBwBoYzMJyB4QwMZ2A4A8MZGM7AcAaGMzCcgeEMDGdgOAPDGRjOwHAGhjMwnIHhDAxnYDgDwxkYzsBwBoYzMJyB4QwMZ2A4A8MZGM7AcAaGMzCcgeEMDGdgOAPDGfiXKcsy9vt9tFo/S1dUHyjz84Nubm5is9nE/f19vL295V2dS8r18fERT09PdejRaBSz2Sy/ekAKfEq3200HwHUhqyiKst1u1ys9H4/HudRhjfe83+9H9WU/Hgn6v6pmsdvt6pW69Hq9+iYf0zii39/f4/X1Nbbbbd7RpUiXbzgc1pGPXcDGwF+nI70tnRhdjq821aiu1yGNgfW7+cOKFvEHKNmQ9kvbvgIAAAAASUVORK5CYIKMWIgRI0aMGDFixIgRCzFixIgRI0aMGDFiIUaMGDFixIgRI0YsxIgRI0aMGDFixIiFGDFixIgRI0aMGLEQI0aMGDFixIgRIxZixIgRH9CyMs0hxHs92G3VzTyZWfFxVzzrrPh47WpXGWRwALEeaZAuo3SIj9dolMlht1gPN03OMh0G8fHqhpkk4z8gvsg0OSfYMOdJrWrZQHyVLAaDPGXXr/PkXd3WpoH484uc/HqzqE/DPPtyz8m994v6/csLcP07vfpYtd62EO9rd1dV5XHRr6pVbavpFt9UbWr9w/u4Z9dV2/XtpmXF21pX1c6Ke694v65qOhRVVXfrWttnz7+NV1Xbql0jsf5GiBEjFmLE/0E/Ac8azp1AFifiAAAAAElFTkSuQmCC"},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if _, err := B64ToFile(tt.args.b64encoded); (err != nil) != tt.wantErr {
				t.Errorf("B64encodedToFile() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
