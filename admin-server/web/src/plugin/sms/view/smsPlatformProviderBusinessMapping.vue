
<template>
  <div>
    <div class="gva-search-box">
      <el-form ref="elSearchFormRef" :inline="true" :model="searchInfo" class="demo-form-inline" @keyup.enter="onSubmit">
      <el-form-item label="创建日期" prop="createdAtRange">
      <template #label>
        <span>
          创建日期
          <el-tooltip content="搜索范围是开始日期（包含）至结束日期（不包含）">
            <el-icon><QuestionFilled /></el-icon>
          </el-tooltip>
        </span>
      </template>
         <el-date-picker
                  v-model="searchInfo.createdAtRange"
                  class="!w-380px"
                  type="datetimerange"
                  range-separator="至"
                  start-placeholder="开始时间"
                  end-placeholder="结束时间"
                />
       </el-form-item>
      
            <el-form-item label="平台业务ID" prop="platformBusinessTypeId">
  <el-input v-model.number="searchInfo.platformBusinessTypeId" placeholder="搜索条件" />
</el-form-item>
           
            <el-form-item label="平台业务编码" prop="platformBusinessCode">
  <el-input v-model="searchInfo.platformBusinessCode" placeholder="搜索条件" />
</el-form-item>
           
            <el-form-item label="三方业务ID" prop="providerBusinessTypeId">
  <el-input v-model.number="searchInfo.providerBusinessTypeId" placeholder="搜索条件" />
</el-form-item>
           
            <el-form-item label="三方编码" prop="providerCode">
  <el-input v-model="searchInfo.providerCode" placeholder="搜索条件" />
</el-form-item>
           
        <template v-if="showAllQuery">
          <!-- 将需要控制显示状态的查询条件添加到此范围内 -->
        </template>

        <el-form-item>
          <el-button type="primary" icon="search" @click="onSubmit">查询</el-button>
          <el-button icon="refresh" @click="onReset">重置</el-button>
          <el-button link type="primary" icon="arrow-down" @click="showAllQuery=true" v-if="!showAllQuery">展开</el-button>
          <el-button link type="primary" icon="arrow-up" @click="showAllQuery=false" v-else>收起</el-button>
        </el-form-item>
      </el-form>
    </div>
    <div class="gva-table-box">
        <div class="gva-btn-list">
            <el-button  type="primary" icon="plus" @click="openDialog()">新增</el-button>
            <el-button  icon="delete" style="margin-left: 10px;" :disabled="!multipleSelection.length" @click="onDelete">删除</el-button>
            
        </div>
        <el-table
        ref="multipleTable"
        style="width: 100%"
        tooltip-effect="dark"
        :data="tableData"
        row-key="ID"
        @selection-change="handleSelectionChange"
        >
        <el-table-column type="selection" width="55" />
        
        <el-table-column sortable align="left" label="日期" prop="CreatedAt" width="180">
            <template #default="scope">{{ formatDate(scope.row.CreatedAt) }}</template>
        </el-table-column>
        
            <el-table-column align="left" label="平台业务" prop="platformBusinessTypeId" width="150">
              <template #default="scope">
                {{ getPlatformBusinessName(scope.row.platformBusinessTypeId) }}
              </template>
            </el-table-column>

            <el-table-column align="left" label="平台业务编码" prop="platformBusinessCode" width="120" />

            <el-table-column align="left" label="三方业务" prop="providerBusinessTypeId" width="200">
              <template #default="scope">
                {{ getProviderBusinessName(scope.row.providerBusinessTypeId) }}
              </template>
            </el-table-column>

            <el-table-column align="left" label="三方编码" prop="providerCode" width="120" />

            <el-table-column align="left" label="三方业务编码" prop="businessCode" width="120" />

            <el-table-column align="left" label="权重" prop="weight" width="120" />

            <el-table-column align="left" label="是否启用该映射" prop="status" width="120">
    <template #default="scope">{{ formatBoolean(scope.row.status) }}</template>
</el-table-column>
            <el-table-column align="left" label="备注" prop="remark" width="120" />

        <el-table-column align="left" label="操作" fixed="right" min-width="240">
            <template #default="scope">
            <el-button  type="primary" link class="table-button" @click="getDetails(scope.row)"><el-icon style="margin-right: 5px"><InfoFilled /></el-icon>查看</el-button>
            <el-button  type="primary" link icon="edit" class="table-button" @click="updateSmsPlatformProviderBusinessMappingFunc(scope.row)">编辑</el-button>
            <el-button   type="primary" link icon="delete" @click="deleteRow(scope.row)">删除</el-button>
            </template>
        </el-table-column>
        </el-table>
        <div class="gva-pagination">
            <el-pagination
            layout="total, sizes, prev, pager, next, jumper"
            :current-page="page"
            :page-size="pageSize"
            :page-sizes="[10, 30, 50, 100]"
            :total="total"
            @current-change="handleCurrentChange"
            @size-change="handleSizeChange"
            />
        </div>
    </div>
    <el-drawer destroy-on-close size="800" v-model="dialogFormVisible" :show-close="false" :before-close="closeDialog">
       <template #header>
              <div class="flex justify-between items-center">
                <span class="text-lg">{{type==='create'?'新增':'编辑'}}</span>
                <div>
                  <el-button :loading="btnLoading" type="primary" @click="enterDialog">确 定</el-button>
                  <el-button @click="closeDialog">取 消</el-button>
                </div>
              </div>
            </template>

          <el-form :model="formData" label-position="top" ref="elFormRef" :rules="rule" label-width="80px">
             <el-form-item label="平台业务类型:" prop="platformBusinessTypeId">
    <el-select v-model="formData.platformBusinessTypeId" placeholder="请选择平台业务类型" clearable @change="onPlatformBusinessChange">
      <el-option
        v-for="item in platformBusinessTypesList"
        :key="item.ID"
        :label="item.name"
        :value="item.ID">
        <span>{{ item.name }} ({{ item.code }})</span>
      </el-option>
    </el-select>
</el-form-item>
             <el-form-item label="平台业务编码:" prop="platformBusinessCode">
    <el-input v-model="formData.platformBusinessCode" :clearable="true" placeholder="自动填充" :disabled="true" />
</el-form-item>
             <el-form-item label="三方业务:" prop="providerBusinessTypeId">
    <el-select v-model="formData.providerBusinessTypeId" placeholder="请选择三方业务" clearable @change="onProviderBusinessChange">
      <el-option
        v-for="item in providerBusinessTypesList"
        :key="item.ID"
        :label="`${getProviderName(item.providerId)} - ${item.businessName}`"
        :value="item.ID">
        <span>{{ getProviderName(item.providerId) }} - {{ item.businessName }} ({{ item.businessCode }})</span>
      </el-option>
    </el-select>
</el-form-item>
             <el-form-item label="三方编码:" prop="providerCode">
    <el-input v-model="formData.providerCode" :clearable="true" placeholder="自动填充" :disabled="true" />
</el-form-item>
             <el-form-item label="三方业务编码:" prop="businessCode">
    <el-input v-model="formData.businessCode" :clearable="true" placeholder="自动填充" :disabled="true" />
</el-form-item>
             <el-form-item label="权重:" prop="weight">
    <el-input v-model.number="formData.weight" :clearable="true" placeholder="请输入权重" />
</el-form-item>
             <el-form-item label="是否启用该映射:" prop="status">
    <el-switch v-model="formData.status" active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" clearable ></el-switch>
</el-form-item>
             <el-form-item label="备注:" prop="remark">
    <el-input v-model="formData.remark" :clearable="true" placeholder="请输入备注" />
</el-form-item>
          </el-form>
    </el-drawer>

    <el-drawer destroy-on-close size="800" v-model="detailShow" :show-close="true" :before-close="closeDetailShow" title="查看">
            <el-descriptions :column="1" border>
                 <el-descriptions-item label="平台业务ID">
    {{ detailForm.platformBusinessTypeId }}
</el-descriptions-item>
                 <el-descriptions-item label="平台业务编码">
    {{ detailForm.platformBusinessCode }}
</el-descriptions-item>
                 <el-descriptions-item label="三方业务ID">
    {{ detailForm.providerBusinessTypeId }}
</el-descriptions-item>
                 <el-descriptions-item label="三方编码">
    {{ detailForm.providerCode }}
</el-descriptions-item>
                 <el-descriptions-item label="三方业务编码">
    {{ detailForm.businessCode }}
</el-descriptions-item>
                 <el-descriptions-item label="权重">
    {{ detailForm.weight }}
</el-descriptions-item>
                 <el-descriptions-item label="是否启用该映射">
    {{ detailForm.status }}
</el-descriptions-item>
                 <el-descriptions-item label="备注">
    {{ detailForm.remark }}
</el-descriptions-item>
            </el-descriptions>
        </el-drawer>

  </div>
</template>

<script setup>
import {
  createSmsPlatformProviderBusinessMapping,
  deleteSmsPlatformProviderBusinessMapping,
  deleteSmsPlatformProviderBusinessMappingByIds,
  updateSmsPlatformProviderBusinessMapping,
  findSmsPlatformProviderBusinessMapping,
  getSmsPlatformProviderBusinessMappingList
} from '@/plugin/sms/api/smsPlatformProviderBusinessMapping'

import { getSmsPlatformBusinessTypesList } from '@/plugin/sms/api/smsPlatformBusinessTypes'
import { getSmsProvidersList } from '@/plugin/sms/api/smsProviders'
import { getSmsProvidersBusinessTypesList } from '@/plugin/sms/api/smsProvidersBusinessTypes'

// 全量引入格式化工具 请按需保留
import { getDictFunc, formatDate, formatBoolean, filterDict ,filterDataSource, returnArrImg, onDownloadFile } from '@/utils/format'
import { ElMessage, ElMessageBox } from 'element-plus'
import { ref, reactive } from 'vue'




defineOptions({
    name: 'SmsPlatformProviderBusinessMapping'
})

// 提交按钮loading
const btnLoading = ref(false)

// 控制更多查询条件显示/隐藏状态
const showAllQuery = ref(false)

// 下拉列表数据
const platformBusinessTypesList = ref([])
const providersList = ref([])
const providerBusinessTypesList = ref([])

// 自动化生成的字典（可能为空）以及字段
const formData = ref({
            platformBusinessTypeId: undefined,
            platformBusinessCode: '',
            providerBusinessTypeId: undefined,
            providerCode: '',
            businessCode: '',
            weight: undefined,
            status: false,
            remark: '',
        })

// 加载平台业务类型列表
const loadPlatformBusinessTypesList = async () => {
  const res = await getSmsPlatformBusinessTypesList({ page: 1, pageSize: 999 })
  if (res.code === 0) {
    platformBusinessTypesList.value = res.data.list || []
  }
}

// 加载三方渠道列表
const loadProvidersList = async () => {
  const res = await getSmsProvidersList({ page: 1, pageSize: 999 })
  if (res.code === 0) {
    providersList.value = res.data.list || []
  }
}

// 加载三方业务列表
const loadProviderBusinessTypesList = async () => {
  const res = await getSmsProvidersBusinessTypesList({ page: 1, pageSize: 999 })
  if (res.code === 0) {
    providerBusinessTypesList.value = res.data.list || []
  }
}

// 选择平台业务时自动填充编码
const onPlatformBusinessChange = (businessId) => {
  const business = platformBusinessTypesList.value.find(b => b.ID === businessId)
  if (business) {
    formData.value.platformBusinessCode = business.code
  } else {
    formData.value.platformBusinessCode = ''
  }
}

// 选择三方业务时自动填充编码
const onProviderBusinessChange = (businessId) => {
  const business = providerBusinessTypesList.value.find(b => b.ID === businessId)
  if (business) {
    formData.value.providerCode = business.providerCode
    formData.value.businessCode = business.businessCode
  } else {
    formData.value.providerCode = ''
    formData.value.businessCode = ''
  }
}

// 根据ID获取三方渠道名称
const getProviderName = (providerId) => {
  const provider = providersList.value.find(p => p.ID === providerId)
  return provider ? provider.name : providerId
}

// 根据ID获取平台业务名称
const getPlatformBusinessName = (businessId) => {
  const business = platformBusinessTypesList.value.find(b => b.ID === businessId)
  return business ? business.name : businessId
}

// 根据ID获取三方业务名称
const getProviderBusinessName = (businessId) => {
  const business = providerBusinessTypesList.value.find(b => b.ID === businessId)
  return business ? `${getProviderName(business.providerId)} - ${business.businessName}` : businessId
}


// 验证规则
const rule = reactive({
})

const elFormRef = ref()
const elSearchFormRef = ref()

// =========== 表格控制部分 ===========
const page = ref(1)
const total = ref(0)
const pageSize = ref(10)
const tableData = ref([])
const searchInfo = ref({})
// 重置
const onReset = () => {
  searchInfo.value = {}
  getTableData()
}

// 搜索
const onSubmit = () => {
  elSearchFormRef.value?.validate(async(valid) => {
    if (!valid) return
    page.value = 1
    if (searchInfo.value.status === ""){
        searchInfo.value.status=null
    }
    getTableData()
  })
}

// 分页
const handleSizeChange = (val) => {
  pageSize.value = val
  getTableData()
}

// 修改页面容量
const handleCurrentChange = (val) => {
  page.value = val
  getTableData()
}

// 查询
const getTableData = async() => {
  const table = await getSmsPlatformProviderBusinessMappingList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
  if (table.code === 0) {
    tableData.value = table.data.list
    total.value = table.data.total
    page.value = table.data.page
    pageSize.value = table.data.pageSize
  }
}

getTableData()

// ============== 表格控制部分结束 ===============

// 获取需要的字典 可能为空 按需保留
const setOptions = async () =>{
  await Promise.all([
    loadPlatformBusinessTypesList(),
    loadProvidersList(),
    loadProviderBusinessTypesList()
  ])
}

// 获取需要的字典 可能为空 按需保留
setOptions()


// 多选数据
const multipleSelection = ref([])
// 多选
const handleSelectionChange = (val) => {
    multipleSelection.value = val
}

// 删除行
const deleteRow = (row) => {
    ElMessageBox.confirm('确定要删除吗?', '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
    }).then(() => {
            deleteSmsPlatformProviderBusinessMappingFunc(row)
        })
    }

// 多选删除
const onDelete = async() => {
  ElMessageBox.confirm('确定要删除吗?', '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning'
  }).then(async() => {
      const IDs = []
      if (multipleSelection.value.length === 0) {
        ElMessage({
          type: 'warning',
          message: '请选择要删除的数据'
        })
        return
      }
      multipleSelection.value &&
        multipleSelection.value.map(item => {
          IDs.push(item.ID)
        })
      const res = await deleteSmsPlatformProviderBusinessMappingByIds({ IDs })
      if (res.code === 0) {
        ElMessage({
          type: 'success',
          message: '删除成功'
        })
        if (tableData.value.length === IDs.length && page.value > 1) {
          page.value--
        }
        getTableData()
      }
      })
    }

// 行为控制标记（弹窗内部需要增还是改）
const type = ref('')

// 更新行
const updateSmsPlatformProviderBusinessMappingFunc = async(row) => {
    const res = await findSmsPlatformProviderBusinessMapping({ ID: row.ID })
    type.value = 'update'
    if (res.code === 0) {
        formData.value = res.data
        dialogFormVisible.value = true
    }
}


// 删除行
const deleteSmsPlatformProviderBusinessMappingFunc = async (row) => {
    const res = await deleteSmsPlatformProviderBusinessMapping({ ID: row.ID })
    if (res.code === 0) {
        ElMessage({
                type: 'success',
                message: '删除成功'
            })
            if (tableData.value.length === 1 && page.value > 1) {
            page.value--
        }
        getTableData()
    }
}

// 弹窗控制标记
const dialogFormVisible = ref(false)

// 打开弹窗
const openDialog = () => {
    type.value = 'create'
    dialogFormVisible.value = true
}

// 关闭弹窗
const closeDialog = () => {
    dialogFormVisible.value = false
    formData.value = {
        platformBusinessTypeId: undefined,
        platformBusinessCode: '',
        providerBusinessTypeId: undefined,
        providerCode: '',
        businessCode: '',
        weight: undefined,
        status: false,
        remark: '',
        }
}
// 弹窗确定
const enterDialog = async () => {
     btnLoading.value = true
     elFormRef.value?.validate( async (valid) => {
             if (!valid) return btnLoading.value = false
              let res
              switch (type.value) {
                case 'create':
                  res = await createSmsPlatformProviderBusinessMapping(formData.value)
                  break
                case 'update':
                  res = await updateSmsPlatformProviderBusinessMapping(formData.value)
                  break
                default:
                  res = await createSmsPlatformProviderBusinessMapping(formData.value)
                  break
              }
              btnLoading.value = false
              if (res.code === 0) {
                ElMessage({
                  type: 'success',
                  message: '创建/更改成功'
                })
                closeDialog()
                getTableData()
              }
      })
}

const detailForm = ref({})

// 查看详情控制标记
const detailShow = ref(false)


// 打开详情弹窗
const openDetailShow = () => {
  detailShow.value = true
}


// 打开详情
const getDetails = async (row) => {
  // 打开弹窗
  const res = await findSmsPlatformProviderBusinessMapping({ ID: row.ID })
  if (res.code === 0) {
    detailForm.value = res.data
    openDetailShow()
  }
}


// 关闭详情弹窗
const closeDetailShow = () => {
  detailShow.value = false
  detailForm.value = {}
}


</script>

<style>

</style>
