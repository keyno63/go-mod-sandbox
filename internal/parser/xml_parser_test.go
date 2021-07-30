package parser

import "testing"

func Test_isParsableXml(t *testing.T) {
	type args struct {
		value string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "正常なXML",
			args: args{
				value: `<?xml version="1.0"?><Sample>value_</Sample>`,
			},
			want: true,
		},
		{
			name: "異常なXML.要素違いで閉じていない",
			args: args{
				value: `<?xml version="1.0"?>
<Sample>value_</Sample1>`,
			},
			want: false,
		},
		{
			name: "異常なXML. 空文字の場合",
			args: args{
				value: "",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isParsableXml(tt.args.value); got != tt.want {
				t.Errorf("isParsableXml() = %v, want %v", got, tt.want)
			}
		})
	}
}
