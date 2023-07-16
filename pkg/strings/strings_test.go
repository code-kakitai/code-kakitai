package strings

import "testing"

func TestRemoveHyphen(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "正常系",
			args: args{
				s: "123-456-789",
			},
			want: "123456789",
		},
		{
			name: "正常系: ハイフンなし",
			args: args{
				s: "123456789",
			},
			want: "123456789",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RemoveHyphen(tt.args.s); got != tt.want {
				t.Errorf("RemoveHyphen() = %v, want %v", got, tt.want)
			}
		})
	}
}
