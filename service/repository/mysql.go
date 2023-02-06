package repository

// import (
// 	"context"
// 	"database/sql"
// 	"fmt"
// 	"strings"

// 	"github.com/win30221/core/http/catch"
// 	"github.com/win30221/core/syserrno"
// 	"github.com/win30221/ichiban-kuji/domain"
// )

// type MysqlRepo struct {
// 	master *sql.DB
// 	slave  *sql.DB
// }

// func NewMysqlRepo(master, slave *sql.DB) *MysqlRepo {
// 	return &MysqlRepo{master, slave}
// }

// func (r *MysqlRepo) GetUsersIchibanKuji(ctx context.Context, userID string) (result []domain.IchibanKuji, err error) {
// 	query := "SELECT `id`, `userID`, `name`, `isLast`, `lastType`, `cost`, `rewardCount`, `status`,`sellAt`, `createdAt`, `updatedAt` FROM ichibanKuji WHERE userID = ?;"

// 	rows, err := r.slave.QueryContext(ctx, query, userID)
// 	if err != nil {
// 		err = catch.New(syserrno.MySQL, "query data error", err.Error())
// 		return
// 	}
// 	defer rows.Close()

// 	for rows.Next() {
// 		r := domain.IchibanKuji{}
// 		err = rows.Scan(
// 			&r.ID,
// 			&r.UserID,
// 			&r.Name,
// 			&r.LastType,
// 			&r.Cost,
// 			&r.RewardCount,
// 			&r.Status,
// 			&r.SellAt,
// 			&r.CreatedAt,
// 			&r.UpatedAt,
// 		)
// 		if err != nil {
// 			err = catch.New(syserrno.MySQL, "scan data error", err.Error())
// 			return
// 		}
// 		result = append(result, r)
// 	}

// 	return
// }

// func (r *MysqlRepo) CreateIchibanKuji(ctx context.Context, data domain.CreateIchibanKujiReq) (err error) {
// 	query := "INSERT INTO ichibanKuji(`userID`, `name`, `isLast`, `lastType`, `cost`, `rewardCount`) VALUES (?, ?, ?, ?, ?, ?);"

// 	_, err = r.master.ExecContext(ctx, query, data.UserID, data.Name, data.IsLast, data.Cost, data.RewardCount)
// 	if err != nil {
// 		err = catch.New(syserrno.MySQL, "insert data error", err.Error())
// 		return
// 	}

// 	return
// }

// func (r *MysqlRepo) GetIchibanKuji(ctx context.Context, ichibanKujiID string) (result domain.IchibanKuji, err error) {
// 	query := "SELECT `id`, `userID`, `name`, `isLast`, `lastType`, `cost`, `rewardCount`, `status`,`sellAt`, `createdAt`, `updatedAt` FROM ichibanKuji WHERE id = ?;"

// 	err = r.slave.QueryRowContext(ctx, query, ichibanKujiID).Scan(
// 		&result.ID,
// 		&result.UserID,
// 		&result.Name,
// 		&result.IsLast,
// 		&result.LastType,
// 		&result.Cost,
// 		&result.RewardCount,
// 		&result.Status,
// 		&result.SellAt,
// 		&result.CreatedAt,
// 		&result.UpatedAt,
// 	)
// 	if err != nil {
// 		err = catch.New(syserrno.MySQL, "query data error", err.Error())
// 		return
// 	}

// 	return
// }

// func (r *MysqlRepo) UpdateIchibanKuji(ctx context.Context, ichibanKujiID string, data domain.UpdateIchibanKujiReq) (err error) {

// 	values := []interface{}{}
// 	set := ""

// 	if data.Name != "" {
// 		set += "name = ?,"
// 		values = append(values, data.Name)
// 	}
// 	// if data.IsLast > 0 {
// 	// 	set += "isLast = ?,"
// 	// 	values = append(values, data.IsLast)
// 	// }
// 	// if data.LastType > 0 {
// 	// 	set += "lastType = ?,"
// 	// 	values = append(values, data.LastType)
// 	// }
// 	if data.Cost > 0 {
// 		set += "cost = ?,"
// 		values = append(values, data.Cost)
// 	}
// 	if data.RewardCount > 0 {
// 		set += "rewardCount = ?,"
// 		values = append(values, data.RewardCount)
// 	}
// 	if data.Status > 0 {
// 		set += "Status = ?,"
// 		values = append(values, data.Status)
// 	}
// 	if data.SellAt.Unix() > 0 {
// 		set += "sellAt = ?,"
// 		values = append(values, data.SellAt)
// 	}

// 	values = append(values, ichibanKujiID)
// 	set = set[:len(set)-1]

// 	query := "UPDATE ichibanKuji SET " + set + " WHERE id = ?;"

// 	_, err = r.master.ExecContext(ctx, query, values...)
// 	if err != nil {
// 		err = catch.New(syserrno.MySQL, "update data error", err.Error())
// 		return
// 	}

// 	return
// }

// func (r *MysqlRepo) CreateIchibanKujiReward(ctx context.Context, ichibanKujiID string, data domain.CreateIchibanKujiRewardReq) (err error) {
// 	values := []interface{}{}

// 	query := fmt.Sprintf(
// 		"INSERT INTO ichibanKujiReward (`ichibanKujiID`, `name`, `description`) VALUES (?, ?, ?) %s;",
// 		strings.Repeat(", (?, ?, ?)", len(data.Name)-1),
// 	)

// 	for i := range data.Name {
// 		values = append(values, ichibanKujiID, data.Name[i], data.Description[i])
// 	}

// 	_, err = r.master.ExecContext(ctx, query, values...)
// 	if err != nil {
// 		err = catch.New(syserrno.MySQL, "insert data error", err.Error())
// 		return
// 	}

// 	return
// }

// func (r *MysqlRepo) GetIchibanKujiRewards(ctx context.Context, ichibanKujiID string) (result []domain.IchibanKujiReward, err error) {
// 	query := "SELECT `id`, `ichibanKujiID`, `userID`, `name`, `description`, `createdAt`, `updatedAt` FROM ichibanKujiReward WHERE ichibankujiID = ?;"

// 	rows, err := r.slave.QueryContext(ctx, query, ichibanKujiID)
// 	if err != nil {
// 		err = catch.New(syserrno.MySQL, "query data error", err.Error())
// 		return
// 	}
// 	defer rows.Close()

// 	for rows.Next() {
// 		r := domain.IchibanKujiReward{}
// 		err = rows.Scan(
// 			&r.ID,
// 			&r.IchibanKujiID,
// 			&r.UserID,
// 			&r.Name,
// 			&r.Description,
// 			&r.CreatedAt,
// 			&r.UpatedAt,
// 		)
// 		if err != nil {
// 			err = catch.New(syserrno.MySQL, "scan data error", err.Error())
// 			return
// 		}
// 		result = append(result, r)
// 	}

// 	return
// }

// func (r *MysqlRepo) UpdateIchibanKujiReward(ctx context.Context, ichibanKujiRewardID string, data domain.UpdateIchibanKujiRewardReq) (err error) {

// 	values := []interface{}{}
// 	set := ""

// 	if data.UserID != "" {
// 		set += "userID = ?,"
// 		values = append(values, data.UserID)
// 	}
// 	if data.Name != "" {
// 		set += "name = ?,"
// 		values = append(values, data.Name)
// 	}
// 	if data.Description != "" {
// 		set += "description = ?,"
// 		values = append(values, data.Description)
// 	}

// 	values = append(values, ichibanKujiRewardID)
// 	set = set[:len(set)-1]

// 	query := "UPDATE ichibanKujiReward SET " + set + " WHERE id = ?;"

// 	_, err = r.master.ExecContext(ctx, query, values...)
// 	if err != nil {
// 		err = catch.New(syserrno.MySQL, "update data error", err.Error())
// 		return
// 	}

// 	return
// }
