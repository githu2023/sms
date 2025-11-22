
package service

import (
	"context"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/sms/model"
    "github.com/flipped-aurora/gin-vue-admin/server/plugin/sms/model/request"
)

var SmsPlatformProviderBusinessMapping = new(smsPlatformProviderBusinessMapping)

type smsPlatformProviderBusinessMapping struct {}
// CreateSmsPlatformProviderBusinessMapping 创建平台子业务记录
// Author [yourname](https://github.com/yourname)
func (s *smsPlatformProviderBusinessMapping) CreateSmsPlatformProviderBusinessMapping(ctx context.Context, smsPlatformProviderBusinessMapping *model.SmsPlatformProviderBusinessMapping) (err error) {
	err = global.GVA_DB.Create(smsPlatformProviderBusinessMapping).Error
	return err
}

// DeleteSmsPlatformProviderBusinessMapping 删除平台子业务记录
// Author [yourname](https://github.com/yourname)
func (s *smsPlatformProviderBusinessMapping) DeleteSmsPlatformProviderBusinessMapping(ctx context.Context, ID string) (err error) {
	err = global.GVA_DB.Delete(&model.SmsPlatformProviderBusinessMapping{},"id = ?",ID).Error
	return err
}

// DeleteSmsPlatformProviderBusinessMappingByIds 批量删除平台子业务记录
// Author [yourname](https://github.com/yourname)
func (s *smsPlatformProviderBusinessMapping) DeleteSmsPlatformProviderBusinessMappingByIds(ctx context.Context, IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]model.SmsPlatformProviderBusinessMapping{},"id in ?",IDs).Error
	return err
}

// UpdateSmsPlatformProviderBusinessMapping 更新平台子业务记录
// Author [yourname](https://github.com/yourname)
func (s *smsPlatformProviderBusinessMapping) UpdateSmsPlatformProviderBusinessMapping(ctx context.Context, smsPlatformProviderBusinessMapping model.SmsPlatformProviderBusinessMapping) (err error) {
	err = global.GVA_DB.Model(&model.SmsPlatformProviderBusinessMapping{}).Where("id = ?",smsPlatformProviderBusinessMapping.ID).Updates(&smsPlatformProviderBusinessMapping).Error
	return err
}

// GetSmsPlatformProviderBusinessMapping 根据ID获取平台子业务记录
// Author [yourname](https://github.com/yourname)
func (s *smsPlatformProviderBusinessMapping) GetSmsPlatformProviderBusinessMapping(ctx context.Context, ID string) (smsPlatformProviderBusinessMapping model.SmsPlatformProviderBusinessMapping, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&smsPlatformProviderBusinessMapping).Error
	return
}
// GetSmsPlatformProviderBusinessMappingInfoList 分页获取平台子业务记录
// Author [yourname](https://github.com/yourname)
func (s *smsPlatformProviderBusinessMapping) GetSmsPlatformProviderBusinessMappingInfoList(ctx context.Context, info request.SmsPlatformProviderBusinessMappingSearch) (list []model.SmsPlatformProviderBusinessMapping, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&model.SmsPlatformProviderBusinessMapping{})
    var smsPlatformProviderBusinessMappings []model.SmsPlatformProviderBusinessMapping
    // 如果有条件搜索 下方会自动创建搜索语句
    if len(info.CreatedAtRange) == 2 {
     db = db.Where("created_at BETWEEN ? AND ?", info.CreatedAtRange[0], info.CreatedAtRange[1])
    }
  
    if info.PlatformBusinessTypeId != nil {
        db = db.Where("platform_business_type_id = ?", *info.PlatformBusinessTypeId)
    }
    if info.PlatformBusinessCode != nil && *info.PlatformBusinessCode != "" {
        db = db.Where("platform_business_code = ?", *info.PlatformBusinessCode)
    }
    if info.ProviderBusinessTypeId != nil {
        db = db.Where("provider_business_type_id = ?", *info.ProviderBusinessTypeId)
    }
    if info.ProviderCode != nil && *info.ProviderCode != "" {
        db = db.Where("provider_code = ?", *info.ProviderCode)
    }
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }

	if limit != 0 {
       db = db.Limit(limit).Offset(offset)
    }
	err = db.Find(&smsPlatformProviderBusinessMappings).Error
	return  smsPlatformProviderBusinessMappings, total, err
}

func (s *smsPlatformProviderBusinessMapping)GetSmsPlatformProviderBusinessMappingPublic(ctx context.Context) {

}
