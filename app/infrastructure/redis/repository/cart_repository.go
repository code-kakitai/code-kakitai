package repository

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/redis/go-redis/v9"

	domainCart "github/code-kakitai/code-kakitai/domain/cart"
	infraRedis "github/code-kakitai/code-kakitai/infrastructure/redis"
)

type cartRepository struct {
}

func NewCartRepository() domainCart.CartRepository {
	return &cartRepository{}
}

type cartProduct struct {
	ProductID string `json:"product_id"`
	Quantity  int    `json:"quantity"`
}

func (r *cartRepository) FindByUserID(ctx context.Context, userID string) (*domainCart.Cart, error) {
	// userIDをキーにしたカート情報をRedisから取得
	rdb := infraRedis.GetRedisClient()
	cart, err := domainCart.NewCart(userID)
	if err != nil {
		return nil, err
	}

	// userIDをキーにしたカート情報がエラー
	jsonData, err := rdb.Get(ctx, cartKey(userID)).Result()
	if err != nil {
		if err == redis.Nil {
			// キーがなかった場合は空のカートを返す
			return cart, nil
		}
		return nil, err
	}

	// 取得した JSON データを CartProduct のスライスにデシリアル化
	var products []cartProduct
	err = json.Unmarshal([]byte(jsonData), &products)
	if err != nil {
		return nil, err
	}

	// Redisから取得したカート情報をCartドメインに変換
	for _, p := range products {
		cart.AddProduct(p.ProductID, p.Quantity)
	}
	return cart, nil
}

func (r *cartRepository) Save(ctx context.Context, cart *domainCart.Cart) error {
	rdb := infraRedis.GetRedisClient()
	// カート情報をRedisに保存
	cps := make([]*cartProduct, 0, len(cart.Products()))
	for _, cp := range cart.Products() {
		cps = append(cps, &cartProduct{
			ProductID: cp.ProductID(),
			Quantity:  cp.Quantity(),
		})
	}
	j, err := json.Marshal(cps)
	if err != nil {
		return err
	}
	key := cartKey(cart.UserID())
	if _, err := rdb.Set(ctx, key, j, domainCart.CartTimeOut).Result(); err != nil {
		return err
	}
	return nil
}

func cartKey(userID string) string {
	return fmt.Sprintf("cart-userID-%s", userID)
}
