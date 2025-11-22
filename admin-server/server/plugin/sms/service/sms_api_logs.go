
package service

import (
	"context"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/sms/model"
    "github.com/flipped-aurora/gin-vue-admin/server/plugin/sms/model/request"
)

var SmsApiLogs = new(smsApiLogs)

type smsApiLogs struct {}
// CreateSmsApiLogs 创建访问日志记录
// Author [yourname](https://github.com/yourname)
func (s *smsApiLogs) CreateSmsApiLogs(ctx context.Context, smsApiLogs *model.SmsApiLogs) (err error) {
	err = global.GVA_DB.Create(smsApiLogs).Error
	return err
}

// DeleteSmsApiLogs 删除访问日志记录
// Author [yourname](https://github.com/yourname)
func (s *smsApiLogs) DeleteSmsApiLogs(ctx context.Context, ID string) (err error) {
	err = global.GVA_DB.Delete(&model.SmsApiLogs{},"id = ?",ID).Error
	return err
}

// DeleteSmsApiLogsByIds 批量删除访问日志记录
// Author [yourname](https://github.com/yourname)
func (s *smsApiLogs) DeleteSmsApiLogsByIds(ctx context.Context, IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]model.SmsApiLogs{},"id in ?",IDs).Error
	return err
}

// UpdateSmsApiLogs 更新访问日志记录
// Author [yourname](https://github.com/yourname)
func (s *smsApiLogs) UpdateSmsApiLogs(ctx context.Context, smsApiLogs model.SmsApiLogs) (err error) {
	err = global.GVA_DB.Model(&model.SmsApiLogs{}).Where("id = ?",smsApiLogs.ID).Updates(&smsApiLogs).Error
	return err
}

// GetSmsApiLogs 根据ID获取访问日志记录
// Author [yourname](https://github.com/yourname)
func (s *smsApiLogs) GetSmsApiLogs(ctx context.Context, ID string) (smsApiLogs model.SmsApiLogs, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&smsApiLogs).Error
	return
}
// GetSmsApiLogsInfoList 分页获取访问日志记录
// Author [yourname](https://github.com/yourname)
func (s *smsApiLogs) GetSmsApiLogsInfoList(ctx context.Context, info request.SmsApiLogsSearch) (list []model.SmsApiLogs, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&model.SmsApiLogs{})
    var smsApiLogss []model.SmsApiLogs
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
	err = db.Find(&smsApiLogss).Error
	return  smsApiLogss, total, err
}

func (s *smsApiLogs)GetSmsApiLogsPublic(ctx context.Context) {

}
