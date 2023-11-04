package user

import (
	"context"
	"log"
	"testing"

	"github.com/google/go-cmp/cmp"
	"go.uber.org/mock/gomock"

	userDomain "github/code-kakitai/code-kakitai/domain/user"
)

func TestFindUserUseCase_Run(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockUserRepo := userDomain.NewMockUserRepository(ctrl)
	uc := NewFindUserUseCase(mockUserRepo)

	tests := []struct {
		name     string
		id       string
		mockFunc func()
		want     *FindUseCaseDto
		wantErr  bool
	}{
		{
			name: "IDによってユーザーを取得し、DTOを返却すること",
			id:   "01HCNYK0PKYZWB0ZT1KR0EPWGP",
			mockFunc: func() {
				mockUserRepo.
					EXPECT().
					FindById(gomock.Any(), "01HCNYK0PKYZWB0ZT1KR0EPWGP").
					Return(reconstructUser(
						"01HCNYK0PKYZWB0ZT1KR0EPWGP",
						"example@test.com",
						"08011112222",
						"田中",
						"太郎",
						"東京都",
						"渋谷区",
						"1-1-1",
					), nil)
			},
			want: &FindUseCaseDto{
				ID:          "01HCNYK0PKYZWB0ZT1KR0EPWGP",
				Email:       "example@test.com",
				PhoneNumber: "08011112222",
				LastName:    "田中",
				FirstName:   "太郎",
				Address:     "東京都渋谷区1-1-1",
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt := tt
			t.Parallel()
			tt.mockFunc()
			got, err := uc.Run(context.Background(), tt.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("FindUserUseCase.Run() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("Run() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func reconstructUser(
	id string,
	email string,
	phoneNumber string,
	lastName string,
	firstName string,
	prefecture string,
	city string,
	addressExtra string,
) *userDomain.User {
	user, err := userDomain.Reconstruct(id, email, phoneNumber, lastName, firstName, prefecture, city, addressExtra)
	if err != nil {
		log.Print(err)
	}
	return user
}
