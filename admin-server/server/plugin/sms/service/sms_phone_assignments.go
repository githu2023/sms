
package service

import (
	"context"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/sms/model"
    "github.com/flipped-aurora/gin-vue-admin/server/plugin/sms/model/request"
)

var SmsPhoneAssignments = new(smsPhoneAssignments)

type smsPhoneAssignments struct {}
// CreateSmsPhoneAssignments 创建号码记录记录
// Author [yourname](https://github.com/yourname)
func (s *smsPhoneAssignments) CreateSmsPhoneAssignments(ctx context.Context, smsPhoneAssignments *model.SmsPhoneAssignments) (err error) {
	err = global.GVA_DB.Create(smsPhoneAssignments).Error
	return err
}

// DeleteSmsPhoneAssignments 删除号码记录记录
// Author [yourname](https://github.com/yourname)
func (s *smsPhoneAssignments) DeleteSmsPhoneAssignments(ctx context.Context, ID string) (err error) {
	err = global.GVA_DB.Delete(&model.SmsPhoneAssignments{},"id = ?",ID).Error
	return err
}

// DeleteSmsPhoneAssignmentsByIds 批量删除号码记录记录
// Author [yourname](https://github.com/yourname)
func (s *smsPhoneAssignments) DeleteSmsPhoneAssignmentsByIds(ctx context.Context, IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]model.SmsPhoneAssignments{},"id in ?",IDs).Error
	return err
}

// UpdateSmsPhoneAssignments 更新号码记录记录
// Author [yourname](https://github.com/yourname)
func (s *smsPhoneAssignments) UpdateSmsPhoneAssignments(ctx context.Context, smsPhoneAssignments model.SmsPhoneAssignments) (err error) {
	err = global.GVA_DB.Model(&model.SmsPhoneAssignments{}).Where("id = ?",smsPhoneAssignments.ID).Updates(&smsPhoneAssignments).Error
	return err
}

// GetSmsPhoneAssignments 根据ID获取号码记录记录
// Author [yourname](https://github.com/yourname)
func (s *smsPhoneAssignments) GetSmsPhoneAssignments(ctx context.Context, ID string) (smsPhoneAssignments model.SmsPhoneAssignments, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&smsPhoneAssignments).Error
	return
}
// GetSmsPhoneAssignmentsInfoList 分页获取号码记录记录
// Author [yourname](https://github.com/yourname)
func (s *smsPhoneAssignments) GetSmsPhoneAssignmentsInfoList(ctx context.Context, info request.SmsPhoneAssignmentsSearch) (list []model.SmsPhoneAssignments, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&model.SmsPhoneAssignments{})
    var smsPhoneAssignmentss []model.SmsPhoneAssignments
    // 如果有条件搜索 下方会自动创建搜索语句
    if len(info.CreatedAtRange) == 2 {
     db = db.Where("created_at BETWEEN ? AND ?", info.CreatedAtRange[0], info.CreatedAtRange[1])
    }
  
    if info.CustomerId != nil {
        db = db.Where("customer_id = ?", *info.CustomerId)
    }
    if info.BusinessTypeId != nil {
        db = db.Where("business_type_id = ?", *info.BusinessTypeId)
    }
    if info.PhoneNumber != nil && *info.PhoneNumber != "" {
        db = db.Where("phone_number = ?", *info.PhoneNumber)
    }
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }

	if limit != 0 {
       db = db.Limit(limit).Offset(offset)
    }
	err = db.Find(&smsPhoneAssignmentss).Error
	return  smsPhoneAssignmentss, total, err
}

func (s *smsPhoneAssignments)GetSmsPhoneAssignmentsPublic(ctx context.Context) {

}
