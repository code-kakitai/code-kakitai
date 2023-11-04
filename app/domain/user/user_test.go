package user

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
)

func TestNewUser(t *testing.T) {
	type args struct {
		email        string
		phoneNumber  string
		lastName     string
		firstName    string
		prefecture   string
		city         string
		addressExtra string
	}
	tests := []struct {
		name    string
		args    args
		want    *User
		wantErr bool
	}{
		{
			name: "正常系",
			args: args{
				email:        "test@example.com",
				phoneNumber:  "09012345678",
				lastName:     "山田",
				firstName:    "太郎",
				prefecture:   "東京都",
				city:         "渋谷区",
				addressExtra: "1-1-1",
			},
			want: &User{
				email:       "test@example.com",
				phoneNumber: "09012345678",
				lastName:    "山田",
				firstName:   "太郎",
				address: address{
					prefecture: "東京都",
					city:       "渋谷区",
					extra:      "1-1-1",
				},
			},
			wantErr: false,
		},
		{
			name: "異常系: 名字が空文字",
			args: args{
				email:        "test@example.com",
				phoneNumber:  "09012345678",
				lastName:     "",
				firstName:    "太郎",
				prefecture:   "東京都",
				city:         "渋谷区",
				addressExtra: "1-1-1",
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "異常系: メールアドレスが不正",
			args: args{
				email:        "testcom",
				phoneNumber:  "09012345678",
				lastName:     "山田",
				firstName:    "太郎",
				prefecture:   "東京都",
				city:         "渋谷区",
				addressExtra: "1-1-1",
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "異常系: 電場番号が不正",
			args: args{
				email:        "test@example.com",
				phoneNumber:  "090123456789",
				lastName:     "山田",
				firstName:    "太郎",
				prefecture:   "東京都",
				city:         "渋谷区",
				addressExtra: "1-1-1",
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "異常系: 住所が不正",
			args: args{
				email:        "test@example.com",
				phoneNumber:  "090123456789",
				lastName:     "山田",
				firstName:    "太郎",
				prefecture:   "",
				city:         "渋谷区",
				addressExtra: "1-1-1",
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewUser(tt.args.email, tt.args.phoneNumber, tt.args.lastName, tt.args.firstName, tt.args.prefecture, tt.args.city, tt.args.addressExtra)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			diff := cmp.Diff(
				got, tt.want,
				cmp.AllowUnexported(User{}, address{}),
				cmpopts.IgnoreFields(User{}, "id"),
			)
			if diff != "" {
				t.Errorf("NewUser() = %v, want %v. error is %s", got, tt.want, err)
			}
		})
	}
}
