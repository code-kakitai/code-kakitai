package notification

import (
	"context"
	"log"
	"sync"

	"go.uber.org/multierr"

	userDomain "github/code-kakitai/code-kakitai/domain/user"
)

type SendSystemMailUseCase struct {
	userRepo     userDomain.UserRepository
	mailNotifier MailNotifier
}

func NewSendSystemMailUseCase(
	userRepo userDomain.UserRepository,
	mailNotifier MailNotifier,
) *SendSystemMailUseCase {
	return &SendSystemMailUseCase{
		userRepo:     userRepo,
		mailNotifier: mailNotifier,
	}
}

func (uc *SendSystemMailUseCase) Run(ctx context.Context) error {
	users, err := uc.userRepo.FindAll(ctx)
	if err != nil {
		return err
	}

	// 一斉送信数で分割する
	var chunkUsers = [][]*userDomain.User{}
	for i := 0; i < len(users); i += emailBatchSize {
		end := i + emailBatchSize
		if end > len(users) {
			end = len(users)
		}
		chunkUsers = append(chunkUsers, users[i:end])
	}

	// メールの内容を生成する
	var allContents = [][]MailContent{}
	for _, chunkUser := range chunkUsers {
		var contents = []MailContent{}
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
	var errs error
	var mu sync.Mutex
	var wg sync.WaitGroup
	for _, v := range allContents {
		wg.Add(1)
		go func(v []MailContent) {
			// goroutine内でpanicが起きた場合には、recoverしてmain goroutineの処理は継続させる
			defer func() {
				if rec := recover(); rec != nil {
					log.Printf("panic: %v", rec)
				}
			}()
			defer wg.Done()
			if err := uc.mailNotifier.Send(ctx, v); err != nil {
				mu.Lock()
				errs = multierr.Append(errs, err)
				mu.Unlock()
				// エラーが起きた際に何かしらの処理を行う場合はここで行う
				return
			}
		}(v)
	}
	wg.Wait()

	return errs
}

// メールの一斉送信数
const emailBatchSize = 1000
