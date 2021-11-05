package parser

import "testing"

func Test_isParsableJson(t *testing.T) {
	type args struct {
		value string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "正常なJSON",
			args: args{
				value: `{
  "sample": 1
}`,
			},
			want: true,
		},
		{
			name: "正常なJSON. 空のJSONの場合",
			args: args{
				value: `{}`,
			},
			want: true,
		},
		{
			name: "異常なJSON、value がない",
			args: args{
				value: `{"invalid"}`,
			},
			want: false,
		},
		{
			name: "異常なJSON、空文字の場合",
			args: args{
				value: "",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isParsableJSON(tt.args.value); got != tt.want {
				t.Errorf("isParsableJSON() = %v, want %v", got, tt.want)
			}
		})
	}
}
