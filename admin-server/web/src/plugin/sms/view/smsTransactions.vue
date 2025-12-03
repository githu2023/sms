
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
      
            <el-form-item label="客户ID" prop="customerId">
  <el-input v-model.number="searchInfo.customerId" placeholder="搜索条件" />
</el-form-item>
           
            <el-form-item label="交易类型" prop="type">
  <el-select v-model="searchInfo.type" placeholder="请选择交易类型" clearable style="width:200px">
    <el-option
      v-for="item in transactionTypeOptions"
      :key="item.value"
      :label="item.label"
      :value="item.value"
    />
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
        
            <el-table-column align="left" label="客户ID" prop="customerId" width="100" />

            <el-table-column align="left" label="变动金额" prop="amount" width="100" />

            <el-table-column align="left" label="变动前余额" prop="balanceBefore" width="120" />

            <el-table-column align="left" label="变动后余额" prop="balanceAfter" width="120" />

            <el-table-column align="left" label="变动前冻结" prop="frozenBefore" width="120">
        <template #default="scope">{{ scope.row.frozenBefore || 0 }}</template>
      </el-table-column>

            <el-table-column align="left" label="变动后冻结" prop="frozenAfter" width="120">
        <template #default="scope">{{ scope.row.frozenAfter || 0 }}</template>
      </el-table-column>

            <el-table-column align="left" label="交易类型" prop="type" width="100">
        <template #default="scope">{{ transactionTypeLabel(scope.row.type) }}</template>
      </el-table-column>
            <el-table-column align="left" label="引用ID" prop="referenceId" width="100" />

            <el-table-column align="left" label="备注" prop="notes" width="250" show-overflow-tooltip />

        <el-table-column align="left" label="操作" fixed="right" min-width="240">
            <template #default="scope">
            <el-button  type="primary" link class="table-button" @click="getDetails(scope.row)"><el-icon style="margin-right: 5px"><InfoFilled /></el-icon>查看</el-button>
            <!-- <el-button  type="primary" link icon="edit" class="table-button" @click="updateSmsTransactionsFunc(scope.row)">编辑</el-button> -->
            <!-- <el-button   type="primary" link icon="delete" @click="deleteRow(scope.row)">删除</el-button> -->
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
             <el-form-item label="客户ID" prop="customerId">
    <el-input v-model.number="formData.customerId" :clearable="true" placeholder="请输入" />
</el-form-item>
             <el-form-item label="变动金额" prop="amount">
    <el-input-number v-model="formData.amount" style="width:100%" :precision="2" :clearable="true" />
</el-form-item>
             <el-form-item label="变动前余额" prop="balanceBefore">
    <el-input-number v-model="formData.balanceBefore" style="width:100%" :precision="2" :clearable="true" />
</el-form-item>
             <el-form-item label="变动后余额" prop="balanceAfter">
    <el-input-number v-model="formData.balanceAfter" style="width:100%" :precision="2" :clearable="true" />
</el-form-item>
             <el-form-item label="交易类型" prop="type">
    <el-select v-model="formData.type" placeholder="请选择交易类型" style="width:100%">
      <el-option
        v-for="item in transactionTypeOptions"
        :key="item.value"
        :label="item.label"
        :value="item.value"
      />
    </el-select>
</el-form-item>
             <el-form-item label="引用ID" prop="referenceId">
    <el-input v-model.number="formData.referenceId" :clearable="true" placeholder="请输入" />
</el-form-item>
             <el-form-item label="备注" prop="notes">
    <el-input v-model="formData.notes" :clearable="true" placeholder="请输入备注" />
</el-form-item>
          </el-form>
    </el-drawer>

    <el-drawer destroy-on-close size="800" v-model="detailShow" :show-close="true" :before-close="closeDetailShow" title="查看">
            <el-descriptions :column="1" border>
                 <el-descriptions-item label="客户ID">
    {{ detailForm.customerId }}
</el-descriptions-item>
                 <el-descriptions-item label="变动金额">
    {{ detailForm.amount }}
</el-descriptions-item>
                 <el-descriptions-item label="变动前余额">
    {{ detailForm.balanceBefore }}
</el-descriptions-item>
                 <el-descriptions-item label="变动后余额">
    {{ detailForm.balanceAfter }}
</el-descriptions-item>
                 <el-descriptions-item label="变动前冻结金额">
    {{ detailForm.frozenBefore || 0 }}
</el-descriptions-item>
                 <el-descriptions-item label="变动后冻结金额">
    {{ detailForm.frozenAfter || 0 }}
</el-descriptions-item>
                 <el-descriptions-item label="交易类型">
          {{ transactionTypeLabel(detailForm.type) }}
        </el-descriptions-item>
                 <el-descriptions-item label="引用ID">
    {{ detailForm.referenceId }}
</el-descriptions-item>
                 <el-descriptions-item label="备注">
    {{ detailForm.notes }}
</el-descriptions-item>
            </el-descriptions>
        </el-drawer>

  </div>
</template>

<script setup>
import {
  createSmsTransactions,
  deleteSmsTransactions,
  deleteSmsTransactionsByIds,
  updateSmsTransactions,
  findSmsTransactions,
  getSmsTransactionsList
} from '@/plugin/sms/api/smsTransactions'

// 全量引入格式化工具 请按需保留
import { getDictFunc, formatDate, formatBoolean, filterDict ,filterDataSource, returnArrImg, onDownloadFile } from '@/utils/format'
import { ElMessage, ElMessageBox } from 'element-plus'
import { ref, reactive } from 'vue'




defineOptions({
    name: 'SmsTransactions'
})

// 交易类型字典
const transactionTypeOptions = ref([])

const transactionTypeLabel = (value) => {
  // 直接查找，因为类型已统一为字符串
  const option = transactionTypeOptions.value.find(item => item.value === value)
  return option ? option.label : '未知'
}

// 提交按钮loading
const btnLoading = ref(false)

// 控制更多查询条件显示/隐藏状态
const showAllQuery = ref(false)

// 自动化生成的字典（可能为空）以及字段
const formData = ref({
            customerId: undefined,
            amount: 0,
            balanceBefore: 0,
            balanceAfter: 0,
      type: '1',
            referenceId: undefined,
            notes: '',
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
    // 将搜索条件中的type转换为数字，便于后端处理
    if (searchInfo.value.type === ""){
        searchInfo.value.type = null
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
  const table = await getSmsTransactionsList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
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
  transactionTypeOptions.value = await getDictFunc('transaction_type')
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
            deleteSmsTransactionsFunc(row)
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
      const res = await deleteSmsTransactionsByIds({ IDs })
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
const updateSmsTransactionsFunc = async(row) => {
    const res = await findSmsTransactions({ ID: row.ID })
    type.value = 'update'
    if (res.code === 0) {
        formData.value = res.data
        dialogFormVisible.value = true
    }
}


// 删除行
const deleteSmsTransactionsFunc = async (row) => {
    const res = await deleteSmsTransactions({ ID: row.ID })
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
        amount: 0,
        balanceBefore: 0,
        balanceAfter: 0,
  type: '1',
        referenceId: undefined,
        notes: '',
        }
}
// 弹窗确定
const enterDialog = async () => {
     btnLoading.value = true
     elFormRef.value?.validate( async (valid) => {
             if (!valid) return btnLoading.value = false
              // type已统一为字符串，直接发送
              let res
              switch (type.value) {
                case 'create':
                        res = await createSmsTransactions(formData.value)
                  break
                case 'update':
                        res = await updateSmsTransactions(formData.value)
                  break
                default:
                        res = await createSmsTransactions(formData.value)
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
  const res = await findSmsTransactions({ ID: row.ID })
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
