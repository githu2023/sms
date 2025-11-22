package service

import (
	"context"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/sms/model"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/sms/model/request"
)

var SmsProvidersBusinessTypes = new(smsProvidersBusinessTypes)

type smsProvidersBusinessTypes struct{}

// CreateSmsProvidersBusinessTypes 创建三方业务记录
// Author [yourname](https://github.com/yourname)
func (s *smsProvidersBusinessTypes) CreateSmsProvidersBusinessTypes(ctx context.Context, smsProvidersBusinessTypes *model.SmsProvidersBusinessTypes) (err error) {
	err = global.GVA_DB.WithContext(ctx).Create(smsProvidersBusinessTypes).Error
	return err
}

// DeleteSmsProvidersBusinessTypes 删除三方业务记录
// Author [yourname](https://github.com/yourname)
func (s *smsProvidersBusinessTypes) DeleteSmsProvidersBusinessTypes(ctx context.Context, ID string) (err error) {
	err = global.GVA_DB.WithContext(ctx).Delete(&model.SmsProvidersBusinessTypes{}, "id = ?", ID).Error
	return err
}

// DeleteSmsProvidersBusinessTypesByIds 批量删除三方业务记录
// Author [yourname](https://github.com/yourname)
func (s *smsProvidersBusinessTypes) DeleteSmsProvidersBusinessTypesByIds(ctx context.Context, IDs []string) (err error) {
	err = global.GVA_DB.WithContext(ctx).Delete(&[]model.SmsProvidersBusinessTypes{}, "id in ?", IDs).Error
	return err
}

// UpdateSmsProvidersBusinessTypes 更新三方业务记录
// Author [yourname](https://github.com/yourname)
func (s *smsProvidersBusinessTypes) UpdateSmsProvidersBusinessTypes(ctx context.Context, smsProvidersBusinessTypes model.SmsProvidersBusinessTypes) (err error) {
	err = global.GVA_DB.WithContext(ctx).Model(&model.SmsProvidersBusinessTypes{}).Where("id = ?", smsProvidersBusinessTypes.ID).Updates(&smsProvidersBusinessTypes).Error
	return err
}

// GetSmsProvidersBusinessTypes 根据ID获取三方业务记录
// Author [yourname](https://github.com/yourname)
func (s *smsProvidersBusinessTypes) GetSmsProvidersBusinessTypes(ctx context.Context, ID string) (smsProvidersBusinessTypes model.SmsProvidersBusinessTypes, err error) {
	err = global.GVA_DB.WithContext(ctx).Where("id = ?", ID).First(&smsProvidersBusinessTypes).Error
	return
}

// GetSmsProvidersBusinessTypesInfoList 分页获取三方业务记录
// Author [yourname](https://github.com/yourname)
func (s *smsProvidersBusinessTypes) GetSmsProvidersBusinessTypesInfoList(ctx context.Context, info request.SmsProvidersBusinessTypesSearch) (list []model.SmsProvidersBusinessTypes, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.WithContext(ctx).Model(&model.SmsProvidersBusinessTypes{})
	var smsProvidersBusinessTypess []model.SmsProvidersBusinessTypes
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	if info.ProviderCode != "" {
		db = db.Where("provider_code = ?", info.ProviderCode)
	}
	if info.BusinessCode != "" {
		db = db.Where("business_code = ?", info.BusinessCode)
	}
	if info.Status != nil {
		db = db.Where("status = ?", info.Status)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Find(&smsProvidersBusinessTypess).Error
	return smsProvidersBusinessTypess, total, err
}
func (s *smsProvidersBusinessTypes) GetSmsProvidersBusinessTypesPublic(ctx context.Context) {
	// 此方法为获取数据源定义的数据
	// 请自行实现
}
