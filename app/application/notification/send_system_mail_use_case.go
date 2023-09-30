package notification

import (
	"context"

	"golang.org/x/sync/errgroup"

	userDomain "github/code-kakitai/code-kakitai/domain/user"
)

type SendSystemMailUseCase struct {
	userRepo   userDomain.UserRepository
	mailClient MailClient
}

func NewNotificationUseCase(
	userRepo userDomain.UserRepository,
	mailClient MailClient,
) *SendSystemMailUseCase {
	return &SendSystemMailUseCase{
		userRepo:   userRepo,
		mailClient: mailClient,
	}
}

func (uc *SendSystemMailUseCase) Run(ctx context.Context) error {
	users, err := uc.userRepo.FindAll(ctx)
	if err != nil {
		return err
	}

	// 一斉送信数で分割する
	chunkUsers := [][]*userDomain.User{}
	for i := 0; i < len(users); i += emailBatchSize {
		end := i + emailBatchSize
		if end > len(users) {
			end = len(users)
		}
		chunkUsers = append(chunkUsers, users[i:end])
	}

	// メールの内容を生成する
	allContents := [][]MailContent{}
	for _, chunkUser := range chunkUsers {
		contents := []MailContent{}
		for _, user := range chunkUser {
			// 件名や本文はtemplate等で生成するか永続化層から取得すると思いますが、本筋から外れるため今回は省略します
			contents = append(contents, MailContent{
				To:      user.Email(),
				Subject: "件名",
				Body:    "本文",
			})
		}
		allContents = append(allContents, contents)
	}

	// 一斉送信する
	eg := errgroup.Group{}
	for _, v := range allContents {
		v := v
		eg.Go(func() error {
			return uc.mailClient.Send(ctx, v)
		})
	}

	if err := eg.Wait(); err != nil {
		return err
	}
	return nil
}

// メールの一斉送信数
const emailBatchSize = 1000
