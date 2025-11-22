
package service

import (
	"context"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/sms/model"
    "github.com/flipped-aurora/gin-vue-admin/server/plugin/sms/model/request"
)

var SmsPlatformBusinessTypes = new(smsPlatformBusinessTypes)

type smsPlatformBusinessTypes struct {}
// CreateSmsPlatformBusinessTypes 创建平台业务记录
// Author [yourname](https://github.com/yourname)
func (s *smsPlatformBusinessTypes) CreateSmsPlatformBusinessTypes(ctx context.Context, smsPlatformBusinessTypes *model.SmsPlatformBusinessTypes) (err error) {
	err = global.GVA_DB.Create(smsPlatformBusinessTypes).Error
	return err
}

// DeleteSmsPlatformBusinessTypes 删除平台业务记录
// Author [yourname](https://github.com/yourname)
func (s *smsPlatformBusinessTypes) DeleteSmsPlatformBusinessTypes(ctx context.Context, ID string) (err error) {
	err = global.GVA_DB.Delete(&model.SmsPlatformBusinessTypes{},"id = ?",ID).Error
	return err
}

// DeleteSmsPlatformBusinessTypesByIds 批量删除平台业务记录
// Author [yourname](https://github.com/yourname)
func (s *smsPlatformBusinessTypes) DeleteSmsPlatformBusinessTypesByIds(ctx context.Context, IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]model.SmsPlatformBusinessTypes{},"id in ?",IDs).Error
	return err
}

// UpdateSmsPlatformBusinessTypes 更新平台业务记录
// Author [yourname](https://github.com/yourname)
func (s *smsPlatformBusinessTypes) UpdateSmsPlatformBusinessTypes(ctx context.Context, smsPlatformBusinessTypes model.SmsPlatformBusinessTypes) (err error) {
	err = global.GVA_DB.Model(&model.SmsPlatformBusinessTypes{}).Where("id = ?",smsPlatformBusinessTypes.ID).Updates(&smsPlatformBusinessTypes).Error
	return err
}

// GetSmsPlatformBusinessTypes 根据ID获取平台业务记录
// Author [yourname](https://github.com/yourname)
func (s *smsPlatformBusinessTypes) GetSmsPlatformBusinessTypes(ctx context.Context, ID string) (smsPlatformBusinessTypes model.SmsPlatformBusinessTypes, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&smsPlatformBusinessTypes).Error
	return
}
// GetSmsPlatformBusinessTypesInfoList 分页获取平台业务记录
// Author [yourname](https://github.com/yourname)
func (s *smsPlatformBusinessTypes) GetSmsPlatformBusinessTypesInfoList(ctx context.Context, info request.SmsPlatformBusinessTypesSearch) (list []model.SmsPlatformBusinessTypes, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&model.SmsPlatformBusinessTypes{})
    var smsPlatformBusinessTypess []model.SmsPlatformBusinessTypes
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
	err = db.Find(&smsPlatformBusinessTypess).Error
	return  smsPlatformBusinessTypess, total, err
}

func (s *smsPlatformBusinessTypes)GetSmsPlatformBusinessTypesPublic(ctx context.Context) {

}
