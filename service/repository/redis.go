package repository

import (
	"context"

	redigo "github.com/gomodule/redigo/redis"
	"github.com/win30221/core/storage/redis"
	"github.com/win30221/ichiban-kuji/domain"
)

const (
	USER_ICHIBAN_KUJI    = "USER_ICHIBAN_KUJI:"
	ICHIBAN_KUJI         = "ICHIBAN_KUJI:"
	ICHIBAN_KUJI_REWARDS = "ICHIBAN_KUJI_REWARDS:"
	ICHIBAN_KUJI_REWARD  = "ICHIBAN_KUJI_REWARD:"
)

type RedisRepo struct {
	pool                  *redigo.Pool
	userIchibanKujiTTL    int
	ichibanKujiTTL        int
	ichibanKujiRewardsTTL int
	ichibanKujiRewardTTL  int
}

func NewRedisRepo(
	pool *redigo.Pool,
	userIchibanKujiTTL int,
	ichibanKujiTTL int,
	ichibanKujiRewardsTTL int,
	ichibanKujiRewardTTL int,
) *RedisRepo {
	return &RedisRepo{
		pool:                  pool,
		userIchibanKujiTTL:    userIchibanKujiTTL,
		ichibanKujiTTL:        ichibanKujiTTL,
		ichibanKujiRewardsTTL: ichibanKujiRewardsTTL,
		ichibanKujiRewardTTL:  ichibanKujiRewardTTL,
	}
}

func (r *RedisRepo) GetUsersIchibanKuji(ctx context.Context, userID string) (result []domain.IchibanKuji, err error) {
	err = redis.GET(r.pool, ctx, USER_ICHIBAN_KUJI+userID, &result)
	return
}

func (r *RedisRepo) SetUsersIchibanKuji(ctx context.Context, userID string, data interface{}) (err error) {
	err = redis.SETEX(r.pool, ctx, USER_ICHIBAN_KUJI+userID, r.userIchibanKujiTTL, data)
	return
}

func (r *RedisRepo) GetIchibanKuji(ctx context.Context, ichibanKujiID string) (result domain.IchibanKuji, err error) {
	err = redis.GET(r.pool, ctx, ICHIBAN_KUJI+ichibanKujiID, &result)
	return
}

func (r *RedisRepo) SetIchibanKuji(ctx context.Context, ichibanKujiID string, data interface{}) (err error) {
	err = redis.SETEX(r.pool, ctx, ICHIBAN_KUJI+ichibanKujiID, r.ichibanKujiTTL, data)
	return
}

func (r *RedisRepo) DeleteIchibanKuji(ctx context.Context, ichibanKujiID string) (err error) {
	err = redis.DEL(r.pool, ctx, ICHIBAN_KUJI+ichibanKujiID)
	return
}

func (r *RedisRepo) GetIchibanKujiRewards(ctx context.Context, ichibanKujiID string) (result []domain.IchibanKujiReward, err error) {
	err = redis.GET(r.pool, ctx, ICHIBAN_KUJI_REWARDS+ichibanKujiID, &result)
	return
}

func (r *RedisRepo) SetIchibanKujiRewards(ctx context.Context, ichibanKujiID string, data interface{}) (err error) {
	err = redis.SETEX(r.pool, ctx, ICHIBAN_KUJI_REWARDS+ichibanKujiID, r.ichibanKujiRewardsTTL, data)
	return
}

func (r *RedisRepo) DeleteIchibanKujiRewards(ctx context.Context, ichibanKujiID string) (err error) {
	err = redis.DEL(r.pool, ctx, ICHIBAN_KUJI_REWARDS+ichibanKujiID)
	return
}

func (r *RedisRepo) GetIchibanKujiReward(ctx context.Context, ichibanKujiRewardID string) (result domain.IchibanKujiReward, err error) {
	err = redis.GET(r.pool, ctx, ICHIBAN_KUJI_REWARD+ichibanKujiRewardID, &result)
	return
}

func (r *RedisRepo) SetIchibanKujiReward(ctx context.Context, ichibanKujiRewardID string, data interface{}) (err error) {
	err = redis.SETEX(r.pool, ctx, ICHIBAN_KUJI_REWARD+ichibanKujiRewardID, r.ichibanKujiRewardTTL, data)
	return
}

func (r *RedisRepo) DeleteIchibanKujiReward(ctx context.Context, ichibanKujiRewardID string) (err error) {
	err = redis.DEL(r.pool, ctx, ICHIBAN_KUJI_REWARD+ichibanKujiRewardID)
	return
}
