package service

import (
	"context"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/sms/model"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/sms/model/request"
)

var SmsTransactions = new(smsTransactions)

type smsTransactions struct{}

// CreateSmsTransactions 创建交易记录记录
// Author [yourname](https://github.com/yourname)
func (s *smsTransactions) CreateSmsTransactions(ctx context.Context, smsTransactions *model.SmsTransactions) (err error) {
	err = global.GVA_DB.Create(smsTransactions).Error
	return err
}

// DeleteSmsTransactions 删除交易记录记录
// Author [yourname](https://github.com/yourname)
func (s *smsTransactions) DeleteSmsTransactions(ctx context.Context, ID string) (err error) {
	err = global.GVA_DB.Delete(&model.SmsTransactions{}, "id = ?", ID).Error
	return err
}

// DeleteSmsTransactionsByIds 批量删除交易记录记录
// Author [yourname](https://github.com/yourname)
func (s *smsTransactions) DeleteSmsTransactionsByIds(ctx context.Context, IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]model.SmsTransactions{}, "id in ?", IDs).Error
	return err
}

// UpdateSmsTransactions 更新交易记录记录
// Author [yourname](https://github.com/yourname)
func (s *smsTransactions) UpdateSmsTransactions(ctx context.Context, smsTransactions model.SmsTransactions) (err error) {
	err = global.GVA_DB.Model(&model.SmsTransactions{}).Where("id = ?", smsTransactions.ID).Updates(&smsTransactions).Error
	return err
}

// GetSmsTransactions 根据ID获取交易记录记录
// Author [yourname](https://github.com/yourname)
func (s *smsTransactions) GetSmsTransactions(ctx context.Context, ID string) (smsTransactions model.SmsTransactions, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&smsTransactions).Error
	return
}

// GetSmsTransactionsInfoList 分页获取交易记录记录
// Author [yourname](https://github.com/yourname)
func (s *smsTransactions) GetSmsTransactionsInfoList(ctx context.Context, info request.SmsTransactionsSearch) (list []model.SmsTransactions, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&model.SmsTransactions{})
	var smsTransactionss []model.SmsTransactions
	// 如果有条件搜索 下方会自动创建搜索语句
	if len(info.CreatedAtRange) == 2 {
		db = db.Where("created_at BETWEEN ? AND ?", info.CreatedAtRange[0], info.CreatedAtRange[1])
	}

	if info.CustomerId != nil {
		db = db.Where("customer_id = ?", *info.CustomerId)
	}
	if info.Type != nil {
		db = db.Where("type = ?", *info.Type)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}
	err = db.Find(&smsTransactionss).Error
	return smsTransactionss, total, err
}

func (s *smsTransactions) GetSmsTransactionsPublic(ctx context.Context) {

}
