package usecase

import (
	"github.com/win30221/core/http/ctx"
	"github.com/win30221/ichiban-kuji/domain"
	"github.com/win30221/ichiban-kuji/service/repository"
	"go.uber.org/zap"
)

type UseCase struct {
	mongoRepo *repository.MongoRepo
	redisRepo *repository.RedisRepo
}

func New(
	mongoRepo *repository.MongoRepo,
	redisRepo *repository.RedisRepo,
) *UseCase {
	return &UseCase{
		mongoRepo: mongoRepo,
		redisRepo: redisRepo,
	}
}

func (u *UseCase) GetUsersIchibanKuji(ctx ctx.Context, userID string) (result []domain.IchibanKuji, err error) {
	result, err = u.redisRepo.GetUsersIchibanKuji(ctx.Context, userID)
	if err != nil {
		result, err = u.reloadUsersIchibanKuji(ctx, userID)
	}
	return
}

func (u *UseCase) reloadUsersIchibanKuji(ctx ctx.Context, userID string) (result []domain.IchibanKuji, err error) {
	result, err = u.mongoRepo.GetUsersIchibanKuji(ctx.Context, userID)
	if err != nil {
		return
	}

	noneReturnErr := u.redisRepo.SetUsersIchibanKuji(ctx.Context, userID, result)
	if noneReturnErr != nil {
		zap.L().Error(noneReturnErr.Error())
	}

	return
}

func (u *UseCase) CreateIchibanKuji(ctx ctx.Context, req domain.CreateIchibanKujiReq) (result string, err error) {
	result, err = u.mongoRepo.CreateIchibanKuji(ctx.Context, req)
	return
}

func (u *UseCase) GetIchibanKuji(ctx ctx.Context, ichibanKujiID string) (result domain.IchibanKuji, err error) {
	result, err = u.redisRepo.GetIchibanKuji(ctx.Context, ichibanKujiID)
	if err != nil {
		result, err = u.reloadIchibanKuji(ctx, false, ichibanKujiID)
	}
	return
}

func (u *UseCase) reloadIchibanKuji(ctx ctx.Context, readMaster bool, ichibanKujiID string) (result domain.IchibanKuji, err error) {
	result, err = u.mongoRepo.GetIchibanKuji(ctx.Context, readMaster, ichibanKujiID)
	if err != nil {
		return
	}

	noneReturnErr := u.redisRepo.SetIchibanKuji(ctx.Context, ichibanKujiID, result)
	if noneReturnErr != nil {
		zap.L().Error(noneReturnErr.Error())
	}

	return
}

func (u *UseCase) UpdateIchibanKuji(ctx ctx.Context, ichibanKujiID string, req domain.UpdateIchibanKujiReq) (err error) {
	err = u.mongoRepo.UpdateIchibanKuji(ctx.Context, ichibanKujiID, req)
	if err != nil {
		return
	}

	u.reloadIchibanKuji(ctx, true, ichibanKujiID)
	ichibanKuji, noneReturnErr := u.GetIchibanKuji(ctx, ichibanKujiID)
	if noneReturnErr != nil {
		return
	}
	u.reloadUsersIchibanKuji(ctx, ichibanKuji.UserID)
	return
}

func (u *UseCase) DeleteIchibanKuji(ctx ctx.Context, ichibanKujiID string) (err error) {

	err = u.mongoRepo.DeleteIchibanKuji(ctx.Context, ichibanKujiID)
	if err != nil {
		return
	}

	u.redisRepo.DeleteIchibanKuji(ctx.Context, ichibanKujiID)

	return
}

func (u *UseCase) CreateIchibanKujiReward(ctx ctx.Context, ichibanKujiID string, req domain.CreateIchibanKujiRewardReq) (err error) {
	err = u.mongoRepo.CreateIchibanKujiReward(ctx.Context, ichibanKujiID, req)
	return
}

func (u *UseCase) GetIchibanKujiRewards(ctx ctx.Context, ichibanKujiID string) (result []domain.IchibanKujiReward, err error) {
	result, err = u.redisRepo.GetIchibanKujiRewards(ctx.Context, ichibanKujiID)
	if err != nil {
		result, err = u.reloadIchibanKujiRewards(ctx, false, ichibanKujiID)
	}
	return
}

func (u *UseCase) reloadIchibanKujiRewards(ctx ctx.Context, readMaster bool, ichibanKujiID string) (result []domain.IchibanKujiReward, err error) {
	result, err = u.mongoRepo.GetIchibanKujiRewards(ctx.Context, readMaster, ichibanKujiID)
	if err != nil {
		return
	}

	noneReturnErr := u.redisRepo.SetIchibanKujiRewards(ctx.Context, ichibanKujiID, result)
	if noneReturnErr != nil {
		zap.L().Error(noneReturnErr.Error())
	}

	return
}

func (u *UseCase) GetIchibanKujiReward(ctx ctx.Context, ichibanKujiRewardID string) (result domain.IchibanKujiReward, err error) {
	result, err = u.redisRepo.GetIchibanKujiReward(ctx.Context, ichibanKujiRewardID)
	if err != nil {
		result, err = u.reloadIchibanKujiReward(ctx, false, ichibanKujiRewardID)
	}
	return
}

func (u *UseCase) reloadIchibanKujiReward(ctx ctx.Context, readMaster bool, ichibanKujiRewardID string) (result domain.IchibanKujiReward, err error) {
	result, err = u.mongoRepo.GetIchibanKujiReward(ctx.Context, readMaster, ichibanKujiRewardID)
	if err != nil {
		return
	}

	noneReturnErr := u.redisRepo.SetIchibanKujiReward(ctx.Context, ichibanKujiRewardID, result)
	if noneReturnErr != nil {
		zap.L().Error(noneReturnErr.Error())
	}

	return
}

func (u *UseCase) DeleteIchibanKujiRewards(ctx ctx.Context, ichibanKujiID string) (err error) {

	err = u.mongoRepo.DeleteIchibanKujiRewards(ctx.Context, ichibanKujiID)
	if err != nil {
		return
	}

	u.redisRepo.DeleteIchibanKujiRewards(ctx.Context, ichibanKujiID)

	return
}

func (u *UseCase) UpdateIchibanKujiReward(ctx ctx.Context, ichibanKujiRewardID string, req domain.UpdateIchibanKujiRewardReq) (err error) {

	reward, err := u.GetIchibanKujiReward(ctx, ichibanKujiRewardID)
	if err != nil {
		return
	}

	err = u.mongoRepo.UpdateIchibanKujiReward(ctx.Context, ichibanKujiRewardID, req)
	if err != nil {
		return
	}

	u.reloadIchibanKujiRewards(ctx, true, reward.IchibanKujiID)
	u.reloadIchibanKujiReward(ctx, true, ichibanKujiRewardID)

	return
}

func (u *UseCase) DeleteIchibanKujiReward(ctx ctx.Context, ichibanKujiRewardID string) (err error) {

	reward, err := u.GetIchibanKujiReward(ctx, ichibanKujiRewardID)
	if err != nil {
		return
	}

	err = u.mongoRepo.DeleteIchibanKujiReward(ctx.Context, ichibanKujiRewardID)
	if err != nil {
		return
	}

	u.reloadIchibanKujiRewards(ctx, true, reward.IchibanKujiID)
	u.redisRepo.DeleteIchibanKujiReward(ctx.Context, ichibanKujiRewardID)

	return
}
