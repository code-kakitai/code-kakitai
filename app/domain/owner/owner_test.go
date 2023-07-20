package owner

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
)

func TestNewOwner(t *testing.T) {
	email := "test@example.com"
	type args struct {
		name  string
		email string
	}
	tests := []struct {
		name    string
		args    args
		want    *Owner
		wantErr bool
	}{
		{
			name: "正常系",
			args: args{
				name:  "山田",
				email: email,
			},
			want: &Owner{
				name:  "山田",
				email: email,
			},
			wantErr: false,
		},
		{
			name: "異常系: 名前が空文字",
			args: args{
				name:  "",
				email: email,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "異常系: メールアドレスが不正",
			args: args{
				name:  "山田",
				email: "test",
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewOwner(tt.args.name, tt.args.email)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewOwner() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			diff := cmp.Diff(
				got, tt.want,
				cmp.AllowUnexported(Owner{}),
				cmpopts.IgnoreFields(Owner{}, "id"),
			)
			if diff != "" {
				t.Errorf("NewOwner() = %v, want %v. error is %s", got, tt.want, err)
			}
		})
	}
}
