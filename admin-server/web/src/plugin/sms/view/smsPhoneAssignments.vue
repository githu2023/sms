
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
      
            <el-form-item label="商户号" prop="merchantNo">
  <el-input v-model="searchInfo.merchantNo" placeholder="搜索条件" />
</el-form-item>
           
            <el-form-item label="业务编码" prop="businessCode">
  <el-input v-model="searchInfo.businessCode" placeholder="搜索条件" />
</el-form-item>
           
            <el-form-item label="手机号" prop="phoneNumber">
  <el-input v-model="searchInfo.phoneNumber" placeholder="搜索条件" />
</el-form-item>

            <el-form-item label="状态" prop="status">
  <el-select v-model="searchInfo.status" placeholder="请选择" clearable>
    <el-option label="待取码" value="pending" />
    <el-option label="已完成" value="completed" />
    <el-option label="已过期" value="expired" />
    <el-option label="失败" value="failed" />
  </el-select>
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
            <!-- 号码记录页面不提供新增和删除功能，由系统自动创建 -->
        </div>
        <el-table
        ref="multipleTable"
        style="width: 100%"
        tooltip-effect="dark"
        :data="tableData"
        row-key="ID"
        >
        
        <el-table-column sortable align="left" label="创建时间" prop="CreatedAt" width="180">
            <template #default="scope">{{ formatDate(scope.row.CreatedAt) }}</template>
        </el-table-column>
        
            <el-table-column align="left" label="运营商" prop="providerId" width="120">
              <template #default="scope">
                {{ getProviderName(scope.row.providerId) }}
              </template>
            </el-table-column>
        
            <el-table-column align="left" label="业务名称" prop="businessName" width="140" />

            <el-table-column align="left" label="业务编码" prop="businessCode" width="120" />

            <el-table-column align="left" label="商户号" prop="merchantNo" width="120" />

            <el-table-column align="left" label="商户名称" prop="merchantName" width="140" />

            <el-table-column align="left" label="手机号" prop="phoneNumber" width="130" />

            <el-table-column align="left" label="验证码" prop="verificationCode" width="100" />

            <el-table-column align="left" label="取码次数" prop="fetchCount" width="90" />

            <el-table-column align="left" label="状态" prop="status" width="100">
    <template #default="scope">
      <el-tag v-if="scope.row.status === 'pending'" type="info">待取码</el-tag>
      <el-tag v-else-if="scope.row.status === 'completed'" type="success">已完成</el-tag>
      <el-tag v-else-if="scope.row.status === 'expired'" type="warning">已过期</el-tag>
      <el-tag v-else-if="scope.row.status === 'failed'" type="danger">失败</el-tag>
      <el-tag v-else>{{ scope.row.status }}</el-tag>
    </template>
</el-table-column>

            <el-table-column align="left" label="渠道成本" prop="providerCost" width="100">
    <template #default="scope">{{ scope.row.providerCost ? '¥' + scope.row.providerCost : '-' }}</template>
</el-table-column>

            <el-table-column align="left" label="商户费用" prop="merchantFee" width="100">
    <template #default="scope">{{ scope.row.merchantFee ? '¥' + scope.row.merchantFee : '-' }}</template>
</el-table-column>

            <el-table-column align="left" label="代理费用" prop="agentFee" width="100">
    <template #default="scope">{{ scope.row.agentFee ? '¥' + scope.row.agentFee : '-' }}</template>
</el-table-column>

            <el-table-column align="left" label="利润" prop="profit" width="100">
    <template #default="scope">
      <span :class="scope.row.profit > 0 ? 'text-green-600' : scope.row.profit < 0 ? 'text-red-600' : ''">
        {{ scope.row.profit ? '¥' + scope.row.profit : '-' }}
      </span>
    </template>
</el-table-column>

            <el-table-column align="left" label="备注" prop="remark" width="150" show-overflow-tooltip />
        <el-table-column align="left" label="操作" fixed="right" min-width="300">
            <template #default="scope">
            <el-button type="primary" link class="table-button" @click="getDetails(scope.row)"><el-icon style="margin-right: 5px"><InfoFilled /></el-icon>查看</el-button>
            <el-button 
              type="success" 
              link 
              @click="getVerificationCode(scope.row)"
              :disabled="scope.row.status === 'completed' || scope.row.status === 'expired'"
            >
              📱获取验证码
            </el-button>
            <el-button type="warning" link @click="releasePhone(scope.row)">🔓释放</el-button>
            <el-button type="info" link @click="expirePhone(scope.row)">⏱️过期</el-button>
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
             <el-form-item label="业务名称" prop="businessName">
    <el-input v-model="formData.businessName" :clearable="true" placeholder="请输入业务名称" />
</el-form-item>
             <el-form-item label="业务编码" prop="businessCode">
    <el-input v-model="formData.businessCode" :clearable="true" placeholder="请输入业务编码" />
</el-form-item>
             <el-form-item label="商户号" prop="merchantNo">
    <el-input v-model="formData.merchantNo" :clearable="true" placeholder="请输入商户号" />
</el-form-item>
             <el-form-item label="商户名称" prop="merchantName">
    <el-input v-model="formData.merchantName" :clearable="true" placeholder="请输入商户名称" />
</el-form-item>
             <el-form-item label="手机号" prop="phoneNumber">
    <el-input v-model="formData.phoneNumber" :clearable="true" placeholder="请输入手机号" />
</el-form-item>
             <el-form-item label="验证码" prop="verificationCode">
    <el-input v-model="formData.verificationCode" :clearable="true" placeholder="请输入验证码" />
</el-form-item>
             <el-form-item label="获取验证码次数" prop="fetchCount">
    <el-input-number v-model="formData.fetchCount" style="width:100%" :min="0" :clearable="true" />
</el-form-item>
             <el-form-item label="状态" prop="status">
    <el-select v-model="formData.status" placeholder="请选择状态" style="width:100%">
      <el-option label="待取码" value="pending" />
      <el-option label="已完成" value="completed" />
      <el-option label="已过期" value="expired" />
      <el-option label="失败" value="failed" />
    </el-select>
</el-form-item>
             <el-form-item label="渠道成本" prop="providerCost">
    <el-input-number v-model="formData.providerCost" style="width:100%" :precision="4" :min="0" :clearable="true" />
</el-form-item>
             <el-form-item label="商户费用" prop="merchantFee">
    <el-input-number v-model="formData.merchantFee" style="width:100%" :precision="4" :min="0" :clearable="true" />
</el-form-item>
             <el-form-item label="代理费用" prop="agentFee">
    <el-input-number v-model="formData.agentFee" style="width:100%" :precision="4" :min="0" :clearable="true" placeholder="预留字段，暂未启用" />
</el-form-item>
             <el-form-item label="利润" prop="profit">
    <el-input-number v-model="formData.profit" style="width:100%" :precision="4" :clearable="true" />
</el-form-item>
             <el-form-item label="备注" prop="remark">
    <el-input v-model="formData.remark" type="textarea" :rows="3" :clearable="true" placeholder="可选，输入备注信息" />
</el-form-item>
          </el-form>
    </el-drawer>

    <el-drawer destroy-on-close size="800" v-model="detailShow" :show-close="true" :before-close="closeDetailShow" title="查看">
            <el-descriptions :column="1" border>
                 <el-descriptions-item label="运营商">
    {{ getProviderName(detailForm.providerId) }}
</el-descriptions-item>
                 <el-descriptions-item label="业务名称">
    {{ detailForm.businessName }}
</el-descriptions-item>
                 <el-descriptions-item label="业务编码">
    {{ detailForm.businessCode }}
</el-descriptions-item>
                 <el-descriptions-item label="商户号">
    {{ detailForm.merchantNo }}
</el-descriptions-item>
                 <el-descriptions-item label="商户名称">
    {{ detailForm.merchantName }}
</el-descriptions-item>
                 <el-descriptions-item label="手机号">
    {{ detailForm.phoneNumber }}
</el-descriptions-item>
                 <el-descriptions-item label="验证码">
    {{ detailForm.verificationCode }}
</el-descriptions-item>
                 <el-descriptions-item label="获取验证码次数">
    {{ detailForm.fetchCount || 0 }}
</el-descriptions-item>
                 <el-descriptions-item label="状态">
    <el-tag v-if="detailForm.status === 'pending'" type="info">待取码</el-tag>
    <el-tag v-else-if="detailForm.status === 'completed'" type="success">已完成</el-tag>
    <el-tag v-else-if="detailForm.status === 'expired'" type="warning">已过期</el-tag>
    <el-tag v-else-if="detailForm.status === 'failed'" type="danger">失败</el-tag>
    <span v-else>{{ detailForm.status }}</span>
</el-descriptions-item>
                 <el-descriptions-item label="渠道成本">
    {{ detailForm.providerCost ? '¥' + detailForm.providerCost : '-' }}
</el-descriptions-item>
                 <el-descriptions-item label="商户费用">
    {{ detailForm.merchantFee ? '¥' + detailForm.merchantFee : '-' }}
</el-descriptions-item>
                 <el-descriptions-item label="代理费用">
    {{ detailForm.agentFee ? '¥' + detailForm.agentFee : '-' }}
</el-descriptions-item>
                 <el-descriptions-item label="利润">
    <span :class="detailForm.profit > 0 ? 'text-green-600' : detailForm.profit < 0 ? 'text-red-600' : ''">
      {{ detailForm.profit ? '¥' + detailForm.profit : '-' }}
    </span>
</el-descriptions-item>
                 <el-descriptions-item label="备注">
    {{ detailForm.remark || '-' }}
</el-descriptions-item>
                 <el-descriptions-item label="创建时间">
    {{ formatDate(detailForm.CreatedAt) }}
</el-descriptions-item>
                 <el-descriptions-item label="更新时间">
    {{ formatDate(detailForm.UpdatedAt) }}
</el-descriptions-item>
            </el-descriptions>
        </el-drawer>

  </div>
</template>

<script setup>
import {
  createSmsPhoneAssignments,
  deleteSmsPhoneAssignments,
  deleteSmsPhoneAssignmentsByIds,
  updateSmsPhoneAssignments,
  findSmsPhoneAssignments,
  getSmsPhoneAssignmentsList
} from '@/plugin/sms/api/smsPhoneAssignments'
import { getSmsProvidersList } from '@/plugin/sms/api/smsProviders'

// 全量引入格式化工具 请按需保留
import { getDictFunc, formatDate, formatBoolean, filterDict ,filterDataSource, returnArrImg, onDownloadFile } from '@/utils/format'
import { ElMessage, ElMessageBox } from 'element-plus'
import { ref, reactive } from 'vue'




defineOptions({
    name: 'SmsPhoneAssignments'
})

// 提交按钮loading
const btnLoading = ref(false)

// 控制更多查询条件显示/隐藏状态
const showAllQuery = ref(false)

// 运营商列表
const providersList = ref([])

// 自动化生成的字典（可能为空）以及字段
const formData = ref({
            businessName: '',
            businessCode: '',
            merchantNo: '',
            merchantName: '',
            phoneNumber: '',
            verificationCode: '',
            fetchCount: 0,
            status: 'pending',
            providerCost: undefined,
            merchantFee: undefined,
            agentFee: undefined,
            profit: undefined,
            remark: '',
        })



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
  const table = await getSmsPhoneAssignmentsList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
  if (table.code === 0) {
    tableData.value = table.data.list
    total.value = table.data.total
    page.value = table.data.page
    pageSize.value = table.data.pageSize
  }
}

getTableData()

// ============== 表格控制部分结束 ===============

// 加载运营商列表
const loadProvidersList = async () => {
  const res = await getSmsProvidersList({ page: 1, pageSize: 999 })
  if (res.code === 0) {
    providersList.value = res.data.list || []
  }
}

// 根据ID获取运营商名称
const getProviderName = (providerId) => {
  if (!providerId) return '-'
  const provider = providersList.value.find(p => p.ID === providerId)
  return provider ? provider.name : providerId
}

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
            deleteSmsPhoneAssignmentsFunc(row)
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
      const res = await deleteSmsPhoneAssignmentsByIds({ IDs })
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
const updateSmsPhoneAssignmentsFunc = async(row) => {
    const res = await findSmsPhoneAssignments({ ID: row.ID })
    type.value = 'update'
    if (res.code === 0) {
        formData.value = res.data
        dialogFormVisible.value = true
    }
}


// 删除行
const deleteSmsPhoneAssignmentsFunc = async (row) => {
    const res = await deleteSmsPhoneAssignments({ ID: row.ID })
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
        customerId: undefined,
        providerId: undefined,
        businessTypeId: undefined,
        cardType: '',
        phoneNumber: '',
        verificationCode: '',
        cost: 0,
        status: false,
        expiresAt: new Date(),
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
                  res = await createSmsPhoneAssignments(formData.value)
                  break
                case 'update':
                  res = await updateSmsPhoneAssignments(formData.value)
                  break
                default:
                  res = await createSmsPhoneAssignments(formData.value)
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
  const res = await findSmsPhoneAssignments({ ID: row.ID })
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

// 获取验证码
const getVerificationCode = (row) => {
  ElMessageBox.confirm(
    `确定要获取手机号 ${row.phoneNumber} 的验证码吗？`,
    '获取验证码',
    {
      confirmButtonText: '确定',
      cancelButtonText: '取消',
      type: 'info'
    }
  ).then(() => {
    // TODO: 调用获取验证码的 API
    ElMessage({
      type: 'info',
      message: '获取验证码功能待实现'
    })
    // 刷新列表
    // getTableData()
  }).catch(() => {
    // 用户取消
  })
}

// 释放号码
const releasePhone = (row) => {
  ElMessageBox.confirm(
    `确定要释放手机号 ${row.phoneNumber} 吗？释放后该号码可被其他用户使用。`,
    '释放号码',
    {
      confirmButtonText: '确定',
      cancelButtonText: '取消',
      type: 'warning'
    }
  ).then(() => {
    // TODO: 调用释放号码的 API
    ElMessage({
      type: 'info',
      message: '释放号码功能待实现'
    })
    // 刷新列表
    // getTableData()
  }).catch(() => {
    // 用户取消
  })
}

// 设置号码为过期
const expirePhone = (row) => {
  ElMessageBox.confirm(
    `确定要将手机号 ${row.phoneNumber} 设置为过期吗？`,
    '设置过期',
    {
      confirmButtonText: '确定',
      cancelButtonText: '取消',
      type: 'warning'
    }
  ).then(() => {
    // TODO: 调用设置过期的 API
    ElMessage({
      type: 'info',
      message: '设置过期功能待实现'
    })
    // 刷新列表
    // getTableData()
  }).catch(() => {
    // 用户取消
  })
}


</script>

<style>

</style>
