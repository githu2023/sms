
<template>
  <div>
    <div class="gva-search-box">
      <el-form ref="elSearchFormRef" :inline="true" :model="searchInfo" class="demo-form-inline" @keyup.enter="onSubmit">
      <el-form-item label="三方编码" prop="providerCode">
        <el-input v-model="searchInfo.providerCode" placeholder="请输入三方编码" />
      </el-form-item>
      
      <el-form-item label="业务编码" prop="businessCode">
        <el-input v-model="searchInfo.businessCode" placeholder="请输入业务编码" />
      </el-form-item>
      
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
      
        <template v-if="showAllQuery">
          <el-form-item label="启用状态" prop="status">
            <el-select v-model="searchInfo.status" placeholder="请选择启用状态" clearable>
              <el-option label="启用" :value="true" />
              <el-option label="停用" :value="false" />
            </el-select>
          </el-form-item>
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
        
            <el-table-column align="left" label="三方渠道" prop="providerId" width="150">
              <template #default="scope">
                {{ getProviderName(scope.row.providerId) }}
              </template>
            </el-table-column>

            <el-table-column align="left" label="三方编码" prop="providerCode" width="120" />

            <el-table-column align="left" label="业务名称" prop="businessName" width="150" />

            <el-table-column align="left" label="业务编码" prop="businessCode" width="120" />

            <el-table-column align="left" label="价格" prop="price" width="120" />

            <el-table-column align="left" label="启用状态" prop="status" width="120">
    <template #default="scope">{{ formatBoolean(scope.row.status) }}</template>
</el-table-column>
            <el-table-column align="left" label="备注" prop="remark" width="120" />

        <el-table-column align="left" label="操作" fixed="right" min-width="240">
            <template #default="scope">
            <el-button  type="primary" link class="table-button" @click="getDetails(scope.row)"><el-icon style="margin-right: 5px"><InfoFilled /></el-icon>查看</el-button>
            <el-button  type="primary" link icon="edit" class="table-button" @click="updateSmsProvidersBusinessTypesFunc(scope.row)">编辑</el-button>
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
             <el-form-item label="三方渠道:" prop="providerId">
    <el-select v-model="formData.providerId" placeholder="请选择三方渠道" clearable @change="onProviderChange">
      <el-option
        v-for="item in providersList"
        :key="item.ID"
        :label="item.name"
        :value="item.ID">
        <span>{{ item.name }} ({{ item.code }})</span>
      </el-option>
    </el-select>
</el-form-item>
             <el-form-item label="三方编码:" prop="providerCode">
    <el-input v-model="formData.providerCode" :clearable="true" placeholder="自动填充" :disabled="true" />
</el-form-item>
             <el-form-item label="业务名称:" prop="businessName">
    <el-input v-model="formData.businessName" :clearable="true" placeholder="请输入业务名称" />
</el-form-item>
             <el-form-item label="业务编码:" prop="businessCode">
    <el-input v-model="formData.businessCode" :clearable="true" placeholder="请输入业务编码" />
</el-form-item>
             <el-form-item label="价格:" prop="price">
    <el-input-number v-model="formData.price" :precision="4" :step="0.0001" :clearable="true" placeholder="请输入价格" />
</el-form-item>
             <el-form-item label="启用状态:" prop="status">
    <el-switch v-model="formData.status" active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" clearable ></el-switch>
</el-form-item>
             <el-form-item label="备注:" prop="remark">
    <el-input v-model="formData.remark" :clearable="true" placeholder="请输入备注" />
</el-form-item>
          </el-form>
    </el-drawer>

    <el-drawer destroy-on-close size="800" v-model="detailShow" :show-close="true" :before-close="closeDetailShow" title="查看">
            <el-descriptions :column="1" border>
                 <el-descriptions-item label="三方渠道ID">
    {{ detailForm.providerId }}
</el-descriptions-item>
                 <el-descriptions-item label="三方编码">
    {{ detailForm.providerCode }}
</el-descriptions-item>
                 <el-descriptions-item label="业务名称">
    {{ detailForm.businessName }}
</el-descriptions-item>
                 <el-descriptions-item label="业务编码">
    {{ detailForm.businessCode }}
</el-descriptions-item>
                 <el-descriptions-item label="价格">
    {{ detailForm.price }}
</el-descriptions-item>
                 <el-descriptions-item label="启用状态">
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
  createSmsProvidersBusinessTypes,
  deleteSmsProvidersBusinessTypes,
  deleteSmsProvidersBusinessTypesByIds,
  updateSmsProvidersBusinessTypes,
  findSmsProvidersBusinessTypes,
  getSmsProvidersBusinessTypesList
} from '@/plugin/sms/api/smsProvidersBusinessTypes'

import { getSmsProvidersList } from '@/plugin/sms/api/smsProviders'

// 全量引入格式化工具 请按需保留
import { getDictFunc, formatDate, formatBoolean, filterDict ,filterDataSource, returnArrImg, onDownloadFile } from '@/utils/format'
import { ElMessage, ElMessageBox } from 'element-plus'
import { ref, reactive } from 'vue'




defineOptions({
    name: 'SmsProvidersBusinessTypes'
})

// 提交按钮loading
const btnLoading = ref(false)

// 控制更多查询条件显示/隐藏状态
const showAllQuery = ref(false)

// 三方渠道列表
const providersList = ref([])

// 自动化生成的字典（可能为空）以及字段
const formData = ref({
            providerId: 0,
            providerCode: '',
            businessName: '',
            businessCode: '',
            price: 0,
            status: false,
            remark: '',
        })

// 加载三方渠道列表
const loadProvidersList = async () => {
  const res = await getSmsProvidersList({ page: 1, pageSize: 999 })
  if (res.code === 0) {
    providersList.value = res.data.list || []
  }
}

// 选择三方渠道时自动填充编码
const onProviderChange = (providerId) => {
  const provider = providersList.value.find(p => p.ID === providerId)
  if (provider) {
    formData.value.providerCode = provider.code
  } else {
    formData.value.providerCode = ''
  }
}

// 根据ID获取三方渠道名称
const getProviderName = (providerId) => {
  const provider = providersList.value.find(p => p.ID === providerId)
  return provider ? provider.name : providerId
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
  const table = await getSmsProvidersBusinessTypesList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
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
  await loadProvidersList()
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
            deleteSmsProvidersBusinessTypesFunc(row)
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
      const res = await deleteSmsProvidersBusinessTypesByIds({ IDs })
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
const updateSmsProvidersBusinessTypesFunc = async(row) => {
    const res = await findSmsProvidersBusinessTypes({ ID: row.ID })
    type.value = 'update'
    if (res.code === 0) {
        formData.value = res.data
        dialogFormVisible.value = true
    }
}


// 删除行
const deleteSmsProvidersBusinessTypesFunc = async (row) => {
    const res = await deleteSmsProvidersBusinessTypes({ ID: row.ID })
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
        providerId: 0,
        providerCode: '',
        businessName: '',
        businessCode: '',
        price: 0,
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
                  res = await createSmsProvidersBusinessTypes(formData.value)
                  break
                case 'update':
                  res = await updateSmsProvidersBusinessTypes(formData.value)
                  break
                default:
                  res = await createSmsProvidersBusinessTypes(formData.value)
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
  const res = await findSmsProvidersBusinessTypes({ ID: row.ID })
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
