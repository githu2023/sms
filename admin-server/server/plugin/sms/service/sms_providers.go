
package service

import (
	"context"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/sms/model"
    "github.com/flipped-aurora/gin-vue-admin/server/plugin/sms/model/request"
)

var SmsProviders = new(smsProviders)

type smsProviders struct {}
// CreateSmsProviders 创建服务端记录
// Author [yourname](https://github.com/yourname)
func (s *smsProviders) CreateSmsProviders(ctx context.Context, smsProviders *model.SmsProviders) (err error) {
	err = global.GVA_DB.Create(smsProviders).Error
	return err
}

// DeleteSmsProviders 删除服务端记录
// Author [yourname](https://github.com/yourname)
func (s *smsProviders) DeleteSmsProviders(ctx context.Context, ID string) (err error) {
	err = global.GVA_DB.Delete(&model.SmsProviders{},"id = ?",ID).Error
	return err
}

// DeleteSmsProvidersByIds 批量删除服务端记录
// Author [yourname](https://github.com/yourname)
func (s *smsProviders) DeleteSmsProvidersByIds(ctx context.Context, IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]model.SmsProviders{},"id in ?",IDs).Error
	return err
}

// UpdateSmsProviders 更新服务端记录
// Author [yourname](https://github.com/yourname)
func (s *smsProviders) UpdateSmsProviders(ctx context.Context, smsProviders model.SmsProviders) (err error) {
	err = global.GVA_DB.Model(&model.SmsProviders{}).Where("id = ?",smsProviders.ID).Updates(&smsProviders).Error
	return err
}

// GetSmsProviders 根据ID获取服务端记录
// Author [yourname](https://github.com/yourname)
func (s *smsProviders) GetSmsProviders(ctx context.Context, ID string) (smsProviders model.SmsProviders, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&smsProviders).Error
	return
}
// GetSmsProvidersInfoList 分页获取服务端记录
// Author [yourname](https://github.com/yourname)
func (s *smsProviders) GetSmsProvidersInfoList(ctx context.Context, info request.SmsProvidersSearch) (list []model.SmsProviders, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&model.SmsProviders{})
    var smsProviderss []model.SmsProviders
    // 如果有条件搜索 下方会自动创建搜索语句
    if len(info.CreatedAtRange) == 2 {
     db = db.Where("created_at BETWEEN ? AND ?", info.CreatedAtRange[0], info.CreatedAtRange[1])
    }
  
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }

	if limit != 0 {
       db = db.Limit(limit).Offset(offset)
    }
	err = db.Find(&smsProviderss).Error
	return  smsProviderss, total, err
}

func (s *smsProviders)GetSmsProvidersPublic(ctx context.Context) {

}
