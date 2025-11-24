
<template>
  <div>
    <div class="gva-search-box">
      <el-form ref="elSearchFormRef" :inline="true" :model="searchInfo" class="demo-form-inline" @keyup.enter="onSubmit">
      <el-form-item label="三方名称" prop="name">
        <el-input v-model="searchInfo.name" placeholder="请输入三方名称" />
      </el-form-item>
      
      <el-form-item label="三方编码" prop="code">
        <el-input v-model="searchInfo.code" placeholder="请输入三方编码" />
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
        
            <el-table-column align="left" label="三方名称" prop="name" width="120" />

            <el-table-column align="left" label="三方编码" prop="code" width="100" />

            <el-table-column align="left" label="API网关" prop="apiGateway" width="200">
    <template #default="scope">
      <span class="text-xs truncate" :title="scope.row.apiGateway">{{ scope.row.apiGateway }}</span>
    </template>
</el-table-column>
            <el-table-column align="left" label="商户号" prop="merchantId" width="120" />

            <el-table-column align="left" label="启用状态" prop="status" width="100">
    <template #default="scope">{{ formatBoolean(scope.row.status) }}</template>
</el-table-column>
        <el-table-column align="left" label="操作" fixed="right" min-width="240">
            <template #default="scope">
            <el-button  type="primary" link class="table-button" @click="getDetails(scope.row)"><el-icon style="margin-right: 5px"><InfoFilled /></el-icon>查看</el-button>
            <el-button  type="primary" link icon="edit" class="table-button" @click="updateSmsProvidersFunc(scope.row)">编辑</el-button>
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
             <el-form-item label="三方名称" prop="name">
    <el-input v-model="formData.name" :clearable="true" placeholder="请输入三方名称" />
</el-form-item>
             <el-form-item label="三方编码" prop="code">
    <el-input v-model="formData.code" :clearable="true" placeholder="请输入三方编码" />
</el-form-item>
             <el-form-item label="API网关" prop="apiGateway">
    <el-input v-model="formData.apiGateway" type="textarea" rows="3" placeholder="请输入API网关地址" />
</el-form-item>
             <el-form-item label="商户号" prop="merchantId">
    <el-input v-model="formData.merchantId" :clearable="true" placeholder="请输入三方商户号" />
</el-form-item>
             <el-form-item label="商户Key" prop="merchantKey">
    <el-input v-model="formData.merchantKey" type="password" show-password :clearable="true" placeholder="请输入三方商户key" />
</el-form-item>
             <el-form-item label="启用状态" prop="status">
    <el-switch v-model="formData.status" active-color="#13ce66" inactive-color="#ff4949" active-text="启用" inactive-text="停用" />
</el-form-item>
             <el-form-item label="备注" prop="remark">
    <el-input v-model="formData.remark" type="textarea" rows="3" placeholder="请输入备注信息" />
</el-form-item>
             <el-form-item label="额外配置(JSON)" prop="extraConfig">
    <el-input 
      v-model="formData.extraConfigText" 
      type="textarea" 
      rows="6" 
      placeholder='请输入JSON格式配置，例如: {"projectName": "hema"}' 
      @blur="validateJSON"
    />
    <div v-if="jsonError" class="text-red-500 text-xs mt-1">{{ jsonError }}</div>
    <div v-else class="text-gray-500 text-xs mt-1">支持JSON格式，用于存储特殊配置如projectName等</div>
</el-form-item>
          </el-form>
    </el-drawer>

    <el-drawer destroy-on-close size="800" v-model="detailShow" :show-close="true" :before-close="closeDetailShow" title="查看">
            <el-descriptions :column="1" border>
                 <el-descriptions-item label="三方名称">
    {{ detailForm.name }}
</el-descriptions-item>
                 <el-descriptions-item label="三方编码">
    {{ detailForm.code }}
</el-descriptions-item>
                 <el-descriptions-item label="API网关">
    <div class="whitespace-pre-wrap font-mono text-sm">{{ detailForm.apiGateway }}</div>
</el-descriptions-item>
                 <el-descriptions-item label="商户号">
    {{ detailForm.merchantId }}
</el-descriptions-item>
                 <el-descriptions-item label="启用状态">
    {{ formatBoolean(detailForm.status) }}
</el-descriptions-item>
                 <el-descriptions-item label="备注">
    {{ detailForm.remark }}
</el-descriptions-item>
                 <el-descriptions-item label="额外配置">
    <div v-if="detailForm.extraConfig" class="whitespace-pre-wrap font-mono text-sm">{{ JSON.stringify(detailForm.extraConfig, null, 2) }}</div>
    <span v-else class="text-gray-400">无</span>
</el-descriptions-item>
            </el-descriptions>
        </el-drawer>

  </div>
</template>

<script setup>
import {
  createSmsProviders,
  deleteSmsProviders,
  deleteSmsProvidersByIds,
  updateSmsProviders,
  findSmsProviders,
  getSmsProvidersList
} from '@/plugin/sms/api/smsProviders'

// 全量引入格式化工具 请按需保留
import { getDictFunc, formatDate, formatBoolean, filterDict ,filterDataSource, returnArrImg, onDownloadFile } from '@/utils/format'
import { ElMessage, ElMessageBox } from 'element-plus'
import { ref, reactive } from 'vue'




defineOptions({
    name: 'SmsProviders'
})

// 提交按钮loading
const btnLoading = ref(false)

// 控制更多查询条件显示/隐藏状态
const showAllQuery = ref(false)

// 自动化生成的字典（可能为空）以及字段
const formData = ref({
            name: '',
            code: '',
            apiGateway: '',
            merchantId: '',
            merchantKey: '',
            status: true,
            remark: '',
            extraConfig: null,
            extraConfigText: '', // 用于编辑的文本字段
        })

// JSON验证错误
const jsonError = ref('')



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
  const table = await getSmsProvidersList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
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
            deleteSmsProvidersFunc(row)
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
      const res = await deleteSmsProvidersByIds({ IDs })
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

// 验证JSON格式
const validateJSON = () => {
    if (!formData.value.extraConfigText || formData.value.extraConfigText.trim() === '') {
        formData.value.extraConfig = null
        jsonError.value = ''
        return true
    }
    try {
        const parsed = JSON.parse(formData.value.extraConfigText)
        formData.value.extraConfig = parsed
        jsonError.value = ''
        return true
    } catch (e) {
        jsonError.value = 'JSON格式错误: ' + e.message
        return false
    }
}

// 将extraConfig对象转换为文本
const extraConfigToText = (config) => {
    if (!config || Object.keys(config).length === 0) {
        return ''
    }
    try {
        return JSON.stringify(config, null, 2)
    } catch (e) {
        return ''
    }
}

// 更新行
const updateSmsProvidersFunc = async(row) => {
    const res = await findSmsProviders({ ID: row.ID })
    type.value = 'update'
    if (res.code === 0) {
        formData.value = res.data
        // 将extraConfig对象转换为文本用于编辑
        formData.value.extraConfigText = extraConfigToText(res.data.extraConfig)
        dialogFormVisible.value = true
    }
}


// 删除行
const deleteSmsProvidersFunc = async (row) => {
    const res = await deleteSmsProviders({ ID: row.ID })
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
    formData.value.extraConfigText = ''
    jsonError.value = ''
    dialogFormVisible.value = true
}

// 关闭弹窗
const closeDialog = () => {
    dialogFormVisible.value = false
    formData.value = {
        name: '',
        code: '',
        apiGateway: '',
        merchantId: '',
        merchantKey: '',
        status: true,
        remark: '',
        extraConfig: null,
        extraConfigText: '',
        }
    jsonError.value = ''
}
// 弹窗确定
const enterDialog = async () => {
     // 先验证JSON
     if (!validateJSON()) {
         ElMessage({
             type: 'error',
             message: '请修正JSON格式错误'
         })
         return
     }
     
     btnLoading.value = true
     elFormRef.value?.validate( async (valid) => {
             if (!valid) return btnLoading.value = false
              
              // 准备提交的数据，移除临时文本字段
              const submitData = { ...formData.value }
              delete submitData.extraConfigText
              
              let res
              switch (type.value) {
                case 'create':
                  res = await createSmsProviders(submitData)
                  break
                case 'update':
                  res = await updateSmsProviders(submitData)
                  break
                default:
                  res = await createSmsProviders(submitData)
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
  const res = await findSmsProviders({ ID: row.ID })
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
