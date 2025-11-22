
package service

import (
	"context"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/sms/model"
    "github.com/flipped-aurora/gin-vue-admin/server/plugin/sms/model/request"
)

var SmsIpWhitelist = new(smsIpWhitelist)

type smsIpWhitelist struct {}
// CreateSmsIpWhitelist 创建白名单记录
// Author [yourname](https://github.com/yourname)
func (s *smsIpWhitelist) CreateSmsIpWhitelist(ctx context.Context, smsIpWhitelist *model.SmsIpWhitelist) (err error) {
	err = global.GVA_DB.Create(smsIpWhitelist).Error
	return err
}

// DeleteSmsIpWhitelist 删除白名单记录
// Author [yourname](https://github.com/yourname)
func (s *smsIpWhitelist) DeleteSmsIpWhitelist(ctx context.Context, ID string) (err error) {
	err = global.GVA_DB.Delete(&model.SmsIpWhitelist{},"id = ?",ID).Error
	return err
}

// DeleteSmsIpWhitelistByIds 批量删除白名单记录
// Author [yourname](https://github.com/yourname)
func (s *smsIpWhitelist) DeleteSmsIpWhitelistByIds(ctx context.Context, IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]model.SmsIpWhitelist{},"id in ?",IDs).Error
	return err
}

// UpdateSmsIpWhitelist 更新白名单记录
// Author [yourname](https://github.com/yourname)
func (s *smsIpWhitelist) UpdateSmsIpWhitelist(ctx context.Context, smsIpWhitelist model.SmsIpWhitelist) (err error) {
	err = global.GVA_DB.Model(&model.SmsIpWhitelist{}).Where("id = ?",smsIpWhitelist.ID).Updates(&smsIpWhitelist).Error
	return err
}

// GetSmsIpWhitelist 根据ID获取白名单记录
// Author [yourname](https://github.com/yourname)
func (s *smsIpWhitelist) GetSmsIpWhitelist(ctx context.Context, ID string) (smsIpWhitelist model.SmsIpWhitelist, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&smsIpWhitelist).Error
	return
}
// GetSmsIpWhitelistInfoList 分页获取白名单记录
// Author [yourname](https://github.com/yourname)
func (s *smsIpWhitelist) GetSmsIpWhitelistInfoList(ctx context.Context, info request.SmsIpWhitelistSearch) (list []model.SmsIpWhitelist, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&model.SmsIpWhitelist{})
    var smsIpWhitelists []model.SmsIpWhitelist
    // 如果有条件搜索 下方会自动创建搜索语句
    if len(info.CreatedAtRange) == 2 {
     db = db.Where("created_at BETWEEN ? AND ?", info.CreatedAtRange[0], info.CreatedAtRange[1])
    }
  
    if info.CustomerId != nil {
        db = db.Where("customer_id = ?", *info.CustomerId)
    }
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }

	if limit != 0 {
       db = db.Limit(limit).Offset(offset)
    }
	err = db.Find(&smsIpWhitelists).Error
	return  smsIpWhitelists, total, err
}

func (s *smsIpWhitelist)GetSmsIpWhitelistPublic(ctx context.Context) {

}
