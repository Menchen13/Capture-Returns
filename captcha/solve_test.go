package captcha

import (
	"Menchen13/Capture-Returns/util"
	"net/http"
	"testing"
)

func Test_getImage(t *testing.T) {
	type args struct {
		resp *http.Response
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "shape",
			//file location matches layout of dev box
			args: args{resp: util.RespFromFile("Responses/circle.html")},
			want: "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",
		},
		{
			name: "term",
			//file location matches layout of dev container
			args: args{resp: util.RespFromFile("Responses/multiplication.html")},
			want: "iVBORw0KGgoAAAANSUhEUgAAAMgAAAAoCAYAAAC7HLUcAAAKoUlEQVR4nO2cfUxb1R/Gn1LawkpnqC0UujkIIzMLi1bMEMZ0vkXYXJkuMRhkxoGJU+Ym0/gSVOKSKWOwscW4OXUuGkVNmExnRnQvaObYUBhGNl42oI4iderEUgod7fP7Y6GR0F4osJX8cj4Jf9zvOc85z7m5D7fn3oKMJCEQCHwSEmwDAsFMRgREIJBABEQgkEAERCCQQAREIJBABEQgkGBSAfn666+RmZkJnU4HhUIBg8GA1atX49ixY341VqsVGzZsQGJiIsLDw2E0GrFixQpJDQC0t7cjPz8f8+bNg1KphE6ng9lsxuHDhydjHbW1tXjooYcmpT1y5AiysrJgMBigVCqh1WrxwAMP4MCBA341V65cwfbt22EymaBWq6HVanHLLbegqKgINpttUj6mmxMnTuCRRx7BnDlzoFQqERsbi6VLl2L37t1wu93BthdcGCDFxcWUyWTcsGEDOzo6ODg4yLa2Nr700ktUKpUsLi4eo2lqamJ0dDSXL1/O+vp6OhwOdnV1sbi4mEqlkjt27PA5V01NDcPDw7lq1SqeOnWKdrudHR0d3Lp1KzUaDQsLCwO1z+PHj3PVqlUB63bu3EkAzMvL49mzZ+l0Otna2srXX3+d4eHhXLdu3RiNw+Fgeno6Y2JiuH//fvb29vLvv/9mdXU1ExISqNPp2NraGrCX6WTPnj0MCQlhbm4um5qa6HQ62dbWxhdffJGhoaFcvXp1UP0Fm4AC0tDQQJlMxrfeestne2lpKQHwxx9/9NacTicXLFjAjIwMDg8Pj9Hs3buXISEhPHHixKi63W6nTqdjVlaWz7m++OILAmBNTc24vpubm7lixQp+8sknPHToELOystjQ0MDCwkIWFRWNq+/q6qJCoeCaNWt8tn/66acEwIMHD46qb9q0iTfccANbWlrGaCwWCyMjI2k2m8ed/1px8eJFKpVK5uTk+GzfsmXLhM/x/ysBBaSgoIB6vZ6Dg4M+210uFyMiIvjss896ax988AFlMhnb29v9jnvbbbcxIyNjVK2qqooAWF9f71cXFxfH3NzccX27XC4ePHiQjz76KDUaDeVyORcvXszy8nJ2d3ePq3/zzTcpk8l8Xugj3HrrrXz44Ye9x+fPn6dcLmdJSYlfTUFBARUKBYeGhsb1cC0oLS1lSEgILRaLz3an00mVSsUXXnjhOjubOQS0B2lsbITJZIJKpfLZrlAoEBMTgz/++MNbq6ysREpKCubPn+933JycHNTU1OCff/7x1kY+nxuNRr+6mJgY/P777+P6VigUWLlyJXbt2oUbb7wRiYmJKCwsxHPPPSc5/ghnzpxBdHQ0FixY4LdPfHw82travMefffYZZDIZ1q5d61ezZcsWtLe3Qy6Xj+vhWtDU1IT4+HjcdNNNPtvDwsIwa9YsuFyu6+xs5hAaSOcdO3YgLCzMb/vw8DBsNhvmzp3rrdXV1eGJJ56QHDc1NRUkUV9fj/vvvx8AEBsbCwDo7OxETEyMT53FYvH2nwj5+flYt24dzGYz7rnnHqSlpY3y6o+FCxf69TBCT08PlEql97impgYmkwk6nc6vRqPRQKPRTNj/dJOfn4/s7Gy/7efOncPly5eRlJR0HV3NMKbzdlReXk4APHPmDMmrm1QALC8vl9RZrVYC4P79+701h8NBnU7HBx980Kdm3759BMDDhw9PyFtnZyfNZjPdbjdJcteuXXz77bcnpB2Ps2fPMiQkhE899ZS3ptfrvXuWQ4cO8d5772VkZCTVajVNJhNLSkrocDimZf5rgcvl4pIlSxgbG0u73R5sO0FjygHxeDzs7OxkUVERw8LCuHPnTm/bb7/9RgB8//33Jcfo7+8nAJaVlY2qHzlyhBEREczJyfE+ObJarSwrK6NarZ7UU6zppq+vjyaTiUql0rvPcrlcBMDCwkK+8sorTE1N5Xfffcf+/n729fWxqqqKiYmJTEpKos1mC/IKxjI8PMzc3FyqVCp+//33wbYTVKYUkOLiYgIgAMrlcr7xxhu8fPmyt/3XX38lAFZWVo47lkwm46uvvjqq5na7+fnnnzMiIsI7z8jP+vXr+e+//07F/pSxWCxctGgRAXDz5s3e+sDAAAHwjjvuYGpqqs9N+KVLlxgVFcXMzMwJzZWUlDTmHEj9HD9+fFJr8ng8zM7Opkql4ldffTWpMf6fmJY7iNVq5YEDB5icnEyj0ciffvqJ5NULaLJ3ELvdzszMTBoMBm7fvp0XLlzw3kGqqqq4ePFiRkdHB+033Jdffkm9Xk8A3LRp06g2j8dDmUxGADx16pTfMSoqKgiAv/zyy7W2O2HKysool8v5zTffBNvKjGBa9yADAwM0mUy8+eabeeXKFe+FP5k9yMaNGzl79mw2Nzf71AwNDfHOO++kwWC4rncSh8PBJ598kgCoVqv54Ycf+uyn0Wg4Z84cybFG1l1RUXEtrAaMx+Oh0WhkXl5esK3MGCb8mNdms6GsrAwXLlzw2yc8PBwvv/wyWlpa0NzcDLVaDY1GA4vFIjl2V1cXAMBgMHhr+/btQ15eHhYuXOhTo1QqsXXrVvT29uLo0aMTXcaUaG1tRXJyMvbu3Ytly5ahqakJjz/+uM++cXFx0Ov1kuONrPevv/6adq+T4fz587BarUhPTw+2lRlDQAF5/vnncfr0acl+I49ne3t7AQApKSmoq6uT1NTV1UEmk+H2228HcPWC6evrQ1xcnKQuPj4eANDd3T2RJUyJxsZGLF26FB0dHdi2bRuOHj2KhIQEv/0XLVqEnp4eyTFH3heNF6TrRX9/P3Q6HaKiooJtZeYw0VuNw+GgQqHg2rVrJfvt2bOHANjR0UHy6ldJJvIm/b777vMeu91uRkREsKCgQHKuH374gQBYXV090WVMCpvNxujoaM6aNYvHjh2bkObjjz8mAL9vqUny3XffJQCeO3du3PGu1yZdMJqA9iCPPfYY5XI5T58+7bO9q6vL+6XEEQYGBjh//nwuX77c+w7iv7z33nuUyWSsra0dVX/66aepVqvZ2Njoc66hoSGmpaXRaDTS6XQGsoyAWb9+PQEEtHEdHBzk3LlzmZ2dTY/HM6a9u7ubsbGxzM7Onk6rU2Lbtm0EwI0bNwbbyowhoID8+eefTEpKYnh4OEtKSnjx4kUODQ2xvb2dFRUVNBgMNJlMY57tNzQ0UK/Xc+XKlfz555/pcDhosVi4efNmKpVKlpaWjpnLbrfzrrvuolarZUlJCVtbWzkwMMCenh5WV1czJSWFWq2WJ0+enNoZGAePx0OdTsdly5YFrK2trWVYWBjNZjNPnjzJgYEB9vX1sbKyknq9nunp6UF/VP1fREDGEvBTLKfTybKyMqalpTEyMpJyuZxarZZ3330333nnHbpcLp86i8XCZ555hgkJCQwLC6PBYGBGRga//fZbv3O53W5+9NFHzMjIYFRUFENDQ6nRaJicnMzXXnuNly5dCtR+wNhstgl/rNHpdGP0LS0tXLNmDefNm0eVSsXZs2dzyZIl3L17t99zFSxEQMYiI8X/xRII/CH+5FYgkEAERCCQQAREIJBABEQgkEAERCCQQAREIJBABEQgkEAERCCQQAREIJBABEQgkEAERCCQQAREIJBABEQgkEAERCCQQAREIJDgf3fnFdYWKAbAAAAAAElFTkSuQmCC",
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

func TestSolver(t *testing.T) {
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
			args: args{u: "http://10.10.11.28"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Solver(tt.args.u)
		})
	}
}
