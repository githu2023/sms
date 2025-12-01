
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
      
            <el-form-item label="用户名" prop="username">
  <el-input v-model="searchInfo.username" placeholder="请输入" />
</el-form-item>

            <el-form-item label="商户名称" prop="merchantName">
  <el-input v-model="searchInfo.merchantName" placeholder="请输入" />
</el-form-item>

            <el-form-item label="商户编号" prop="merchantNo">
  <el-input v-model="searchInfo.merchantNo" placeholder="请输入" />
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
        
            <el-table-column align="left" label="商户名称" prop="merchantName" width="140" />

            <el-table-column align="left" label="商户编号" prop="merchantNo" width="120" />

            <el-table-column align="left" label="用户名" prop="username" width="110" />

            <el-table-column align="left" label="邮箱" prop="email" width="140" />

            <!-- <el-table-column align="left" label="客户端登录用的密码哈希" prop="passwordHash" width="120" /> -->

            <el-table-column align="left" label="API密钥" prop="apiSecretKey" width="160">
    <template #default="scope">
      <div class="flex items-center gap-2">
        <span class="text-xs font-mono">{{ scope.row.apiSecretKey?.substring(0, 16) }}...</span>
        <el-button 
          v-if="scope.row.apiSecretKey"
          type="primary" 
          link 
          size="small"
          @click="copyToClipboard(scope.row.apiSecretKey)"
        >
          复制
        </el-button>
      </div>
    </template>
</el-table-column>

            <el-table-column align="left" label="余额" prop="balance" width="80" />

            <el-table-column align="left" label="冻结金额" prop="frozenAmount" width="100">
    <template #default="scope">{{ scope.row.frozenAmount || 0 }}</template>
</el-table-column>

            <el-table-column align="left" label="状态" prop="status" width="80">
    <template #default="scope">{{ formatBoolean(scope.row.status) }}</template>
</el-table-column>
            <el-table-column align="left" label="注册IP" prop="registrationIp" width="120" />

            <el-table-column align="left" label="最后登录IP" prop="lastLoginIp" width="120" />

            <el-table-column align="left" label="最后登录时间" prop="lastLoginAt" width="160">
   <template #default="scope">{{ formatDate(scope.row.lastLoginAt) }}</template>
</el-table-column>
        <el-table-column align="left" label="操作" fixed="right" min-width="420">
            <template #default="scope">
            <el-button  type="primary" link class="table-button" @click="getDetails(scope.row)"><el-icon style="margin-right: 5px"><InfoFilled /></el-icon>查看</el-button>
            <el-button  type="primary" link icon="edit" class="table-button" @click="updateSmsCustomersFunc(scope.row)">编辑</el-button>
            <el-button  type="success" link @click="openCreditDebitDialogFromTable(scope.row)">💰上分/下分</el-button>
            <el-button  type="primary" link @click="openBusinessConfigFromTable(scope.row)">⚙️业务配置</el-button>
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
             <el-form-item label="商户名称" prop="merchantName">
    <el-input v-model="formData.merchantName" :clearable="true" placeholder="请输入商户名称" />
</el-form-item>
             <el-form-item label="商户编号" prop="merchantNo">
    <el-input v-model="formData.merchantNo" :clearable="true" placeholder="请输入商户编号" />
</el-form-item>
             <el-form-item label="用户名" prop="username">
    <el-input v-model="formData.username" :clearable="true" placeholder="请输入登录用户名" />
</el-form-item>
             <el-form-item label="邮箱" prop="email">
    <el-input v-model="formData.email" :clearable="true" placeholder="请输入登录邮箱" />
</el-form-item>
             <el-form-item label="登录密码:" prop="password" v-if="type === 'create'">
    <div class="flex gap-2">
      <el-input 
        v-model="formData.password" 
        :clearable="true" 
        placeholder="请输入登录密码" 
        :type="showPassword ? 'text' : 'password'"
        class="flex-1"
      />
      <el-button @click="showPassword = !showPassword" :icon="showPassword ? 'View' : 'Hide'" />
      <el-button type="primary" @click="generatePassword">生成密码</el-button>
    </div>
    <div class="text-xs text-gray-500 mt-1">密码长度至少6位，建议包含字母、数字和特殊字符</div>
</el-form-item>
             <!-- <el-form-item label="客户端登录用的密码哈希:" prop="passwordHash">
    <el-input v-model="formData.passwordHash" :clearable="true" placeholder="请输入客户端登录用的密码哈希" />
</el-form-item> -->
             <el-form-item label="API密钥" prop="apiSecretKey">
    <el-input v-model="formData.apiSecretKey" :clearable="true" placeholder="可选，自动生成" />
</el-form-item>
             <el-form-item label="余额" prop="balance">
    <el-input-number v-model="formData.balance" style="width:100%" :precision="2" :clearable="true" :disabled="true" />
    <div class="text-xs text-gray-500 mt-1">余额只能通过"充值/扣费"操作修改，不能在此直接编辑</div>
</el-form-item>
             <el-form-item label="冻结金额" prop="frozenAmount">
    <el-input-number v-model="formData.frozenAmount" style="width:100%" :precision="2" :clearable="true" :disabled="true" />
    <div class="text-xs text-gray-500 mt-1">冻结金额只能通过"冻结/解冻"操作修改</div>
</el-form-item>
             <el-form-item label="备注" prop="remark">
    <el-input v-model="formData.remark" type="textarea" :rows="3" :clearable="true" placeholder="可选，输入备注信息" />
</el-form-item>
             <el-form-item label="状态" prop="status">
    <el-switch v-model="formData.status" active-color="#13ce66" inactive-color="#ff4949" active-text="正常" inactive-text="冻结" clearable ></el-switch>
</el-form-item>
             <!-- <el-form-item label="注册时的IP地址:" prop="registrationIp">
    <el-input v-model="formData.registrationIp" :clearable="true" placeholder="请输入注册时的IP地址" />
</el-form-item>
             <el-form-item label="上次登录IP" prop="lastLoginIp">
    <el-input v-model="formData.lastLoginIp" :clearable="true" placeholder="请输入最后一次登录的IP地址" />
</el-form-item>
             <el-form-item label="最后登录时间" prop="lastLoginAt">
    <el-date-picker v-model="formData.lastLoginAt" type="date" style="width:100%" placeholder="选择日期" :clearable="true" />
</el-form-item> -->
          </el-form>
    </el-drawer>

    <el-drawer destroy-on-close size="800" v-model="detailShow" :show-close="true" :before-close="closeDetailShow" title="查看">
            <el-descriptions :column="1" border>
                 <el-descriptions-item label="商户名称">
    {{ detailForm.merchantName }}
</el-descriptions-item>
                 <el-descriptions-item label="商户编号">
    {{ detailForm.merchantNo }}
</el-descriptions-item>
                 <el-descriptions-item label="用户名">
    {{ detailForm.username }}
</el-descriptions-item>
                 <el-descriptions-item label="邮箱">
    {{ detailForm.email }}
</el-descriptions-item>
                 <!-- <el-descriptions-item label="客户端登录用的密码哈希">
    {{ detailForm.passwordHash }}
</el-descriptions-item> -->
                 <el-descriptions-item label="API密钥">
    <div class="flex items-center gap-2">
      <el-input 
        :model-value="detailForm.apiSecretKey"
        :readonly="true"
        class="flex-1"
        type="password"
      />
      <el-button 
        v-if="detailForm.apiSecretKey"
        type="primary" 
        @click="copyToClipboard(detailForm.apiSecretKey)"
        size="small"
      >
        复制
      </el-button>
      <el-button 
        v-if="detailForm.apiSecretKey"
        type="primary" 
        link
        @click="showApiKey = !showApiKey"
        size="small"
      >
        {{ showApiKey ? '隐藏' : '显示' }}
      </el-button>
    </div>
    <div v-if="showApiKey" class="mt-2 p-2 bg-gray-100 rounded font-mono text-sm break-all">
      {{ detailForm.apiSecretKey }}
    </div>
</el-descriptions-item>
                 <el-descriptions-item label="余额">
    {{ detailForm.balance }}
</el-descriptions-item>
                 <el-descriptions-item label="冻结金额">
    {{ detailForm.frozenAmount || 0 }}
</el-descriptions-item>
                 <el-descriptions-item label="备注">
    {{ detailForm.remark || '-' }}
</el-descriptions-item>
                 <el-descriptions-item label="状态">
    {{ detailForm.status }}
</el-descriptions-item>
                 <el-descriptions-item label="注册IP">
    {{ detailForm.registrationIp }}
</el-descriptions-item>
                 <el-descriptions-item label="最后登录IP">
    {{ detailForm.lastLoginIp }}
</el-descriptions-item>
                 <el-descriptions-item label="最后登录时间">
    {{ detailForm.lastLoginAt }}
</el-descriptions-item>
                <el-divider />
                <div class="flex gap-2 mt-4">
                  <el-button type="success" size="large" @click="openCreditDebitDialog">💰 上分/下分</el-button>
                  <el-button type="primary" size="large" @click="openBusinessConfigDialog">⚙️ 业务配置</el-button>
                </div>
            </el-descriptions>
        </el-drawer>

    <!-- 充值/扣费弹窗 -->
    <el-drawer destroy-on-close size="600" v-model="creditDebitDialogVisible" :show-close="false" :before-close="closeCreditDebitDialog">
       <template #header>
              <div class="flex justify-between items-center">
                <span class="text-lg">上分/下分</span>
                <div>
                  <el-button :loading="creditDebitLoading" type="primary" @click="submitCreditDebit">确 定</el-button>
                  <el-button @click="closeCreditDebitDialog">取 消</el-button>
                </div>
              </div>
            </template>

          <el-form :model="creditDebitForm" label-position="top" ref="creditDebitFormRef" :rules="creditDebitRule" label-width="100px">
             <el-form-item label="当前余额">
    <div class="flex gap-4">
      <span class="text-lg font-semibold">余额: {{ detailForm.balance || 0 }}</span>
      <span class="text-lg font-semibold text-orange-500">冻结: {{ detailForm.frozenAmount || 0 }}</span>
    </div>
</el-form-item>
             <el-form-item label="操作类型" prop="type">
    <el-select v-model="creditDebitForm.type" placeholder="请选择操作类型" style="width:100%">
      <el-option label="上分" value="4" />
      <el-option label="下分" value="5" />
      <el-option label="冻结金额" value="6" />
      <el-option label="冻结金额返回" value="7" />
    </el-select>
</el-form-item>
             <el-form-item label="金额" prop="amount">
    <el-input-number 
      v-model="creditDebitForm.amount" 
      style="width:100%" 
      :precision="2"
      :min="0.01"
      :clearable="true"
      placeholder="请输入金额"
    />
</el-form-item>
             <el-form-item label="备注" prop="notes">
    <el-input 
      v-model="creditDebitForm.notes" 
      type="textarea" 
      rows="3"
      :clearable="true"
      placeholder="可选，输入操作备注"
    />
</el-form-item>
             <div class="mt-4 p-3 bg-blue-50 rounded">
               <p class="text-sm">
                 预计变动：
                 <span class="font-semibold">
                   {{ creditDebitForm.type === '4' ? '+' : creditDebitForm.type === '5' ? '-' : creditDebitForm.type === '6' ? '冻结' : '解冻' }}
                   {{ creditDebitForm.amount || 0 }}
                 </span>
               </p>
               <p class="text-sm mt-2">变动后余额：<span class="font-semibold">{{ calculateNewBalance() }}</span></p>
             </div>
          </el-form>
    </el-drawer>

    <!-- 业务配置弹窗 -->
    <el-drawer destroy-on-close size="700" v-model="businessConfigDialogVisible" :show-close="false" :before-close="closeBusinessConfigDialog">
       <template #header>
              <div class="flex justify-between items-center">
                <span class="text-lg">业务配置</span>
                <div>
                  <el-button :loading="businessConfigLoading" type="primary" @click="submitBusinessConfig">确 定</el-button>
                  <el-button @click="closeBusinessConfigDialog">取 消</el-button>
                </div>
              </div>
            </template>

          <el-form :model="businessConfigForm" label-position="top" ref="businessConfigFormRef" label-width="100px">
             <el-form-item label="商户信息">
                <div class="text-sm">
                  <p>用户名: {{ detailForm.username }}</p>
                  <p>邮箱: {{ detailForm.email }}</p>
                </div>
             </el-form-item>
             
             <el-form-item label="选择平台业务类型">
                <el-button type="primary" icon="Plus" @click="openBusinessTypeSelector">+ 添加业务类型</el-button>
                <div class="mt-2" v-if="businessConfigForm.businessConfig.length > 0">
                  <div class="text-sm text-gray-600 mb-2">已选择 {{ businessConfigForm.businessConfig.length }} 个业务类型：</div>
                  <el-tag 
                    v-for="(item, index) in businessConfigForm.businessConfig" 
                    :key="index"
                    closable
                    @close="removeBusinessConfig(index)"
                    class="mr-2 mb-2"
                    type="success"
                  >
                    {{ item.businessName || item.name }}
                  </el-tag>
                </div>
             </el-form-item>

             <el-divider v-if="businessConfigForm.businessConfig.length > 0" content-position="left">
               <span class="text-primary font-semibold">👇 请为每个业务配置价格</span>
             </el-divider>
             
             <div v-for="(item, index) in businessConfigForm.businessConfig" :key="index" class="mb-4 p-4 border-2 rounded-lg shadow-sm transition-all" :class="item.status === 1 ? 'border-blue-200 bg-blue-50' : 'border-gray-300 bg-gray-100 opacity-60'">
                <div class="flex justify-between items-center mb-3">
                  <h4 class="font-bold text-lg" :class="item.status === 1 ? 'text-gray-800' : 'text-gray-400'">{{ item.businessName || item.name }}</h4>
                  <div class="flex gap-2">
                    <el-tag size="small" :type="item.status === 1 ? 'primary' : 'info'">{{ item.businessCode || item.code }}</el-tag>
                    <el-tag size="small" :type="item.status === 1 ? 'success' : 'info'">{{ item.status === 1 ? '启用中' : '已禁用' }}</el-tag>
                  </div>
                </div>
                
                <el-form-item label="💰 业务成本（单价）" :prop="`businessConfig.${index}.cost`" required>
                  <el-input-number 
                    v-model="item.cost" 
                    :min="0"
                    :precision="4"
                    :step="0.01"
                    style="width: 100%"
                    :controls-position="'right'"
                    placeholder="请输入单价"
                  />
                  <div class="text-xs text-orange-600 mt-1 font-semibold">⚠️ 每次使用此业务将扣除的费用</div>
                </el-form-item>
                
                <el-form-item label="🔘 状态">
                  <el-switch 
                    v-model="item.status" 
                    :active-value="1" 
                    :inactive-value="0"
                    active-text="启用" 
                    inactive-text="禁用" 
                  />
                </el-form-item>
             </div>

             <el-empty v-if="businessConfigForm.businessConfig.length === 0" description="暂无业务配置">
               <template #description>
                 <p class="text-gray-500">请点击上方"添加业务类型"按钮选择业务</p>
                 <p class="text-sm text-orange-500 mt-2">选择后需要为每个业务设置价格（cost）</p>
               </template>
             </el-empty>
          </el-form>
    </el-drawer>

    <!-- 业务类型选择器 -->
    <el-dialog v-model="businessTypeSelectorVisible" title="选择业务类型" width="600px">
      <el-table 
        :data="platformBusinessTypes" 
        @selection-change="handleBusinessTypeSelection"
        ref="businessTypeTable"
      >
        <el-table-column type="selection" width="55" :selectable="checkBusinessSelectable" />
        <el-table-column prop="code" label="业务编码" width="120" />
        <el-table-column prop="name" label="业务名称" />
        <el-table-column prop="description" label="业务描述" />
        <el-table-column label="当前配置" width="180">
          <template #default="scope">
            <div v-if="getExistingBusinessConfig(scope.row)" style="display: flex; flex-direction: column; gap: 4px;">
              <div>
                <el-tag type="success" size="small">已配置</el-tag>
                <el-tag :type="getExistingBusinessConfig(scope.row).status === 1 ? 'success' : 'info'" size="small" style="margin-left: 4px;">
                  {{ getExistingBusinessConfig(scope.row).status === 1 ? '启用' : '状态' }}
                </el-tag>
              </div>
              <span style="font-size: 12px; color: #606266;">💰 成本: {{ getExistingBusinessConfig(scope.row).cost.toFixed(4) }} 元</span>
            </div>
            <el-tag v-else type="info" size="small">未配置</el-tag>
          </template>
        </el-table-column>
      </el-table>
      <template #footer>
        <el-button @click="businessTypeSelectorVisible = false">取消</el-button>
        <el-button type="primary" @click="confirmBusinessTypeSelection">确定</el-button>
      </template>
    </el-dialog>

  </div>
</template>

<script setup>
import {
  createSmsCustomers,
  deleteSmsCustomers,
  deleteSmsCustomersByIds,
  updateSmsCustomers,
  findSmsCustomers,
  getSmsCustomersList,
  creditDebitSmsCustomers,
  configureBusinessSmsCustomers,
  getBusinessConfigSmsCustomers
} from '@/plugin/sms/api/smsCustomers'

import { getSmsPlatformBusinessTypesList } from '@/plugin/sms/api/smsPlatformBusinessTypes'

// 全量引入格式化工具 请按需保留
import { getDictFunc, formatDate, formatBoolean, filterDict ,filterDataSource, returnArrImg, onDownloadFile } from '@/utils/format'
import { ElMessage, ElMessageBox } from 'element-plus'
import { View, Hide } from '@element-plus/icons-vue'
import { ref, reactive } from 'vue'




defineOptions({
    name: 'SmsCustomers'
})

// 提交按钮loading
const btnLoading = ref(false)

// 控制更多查询条件显示/隐藏状态
const showAllQuery = ref(false)

// 密码显示控制
const showPassword = ref(false)

// API密钥显示控制
const showApiKey = ref(false)

// 自动化生成的字典（可能为空）以及字段
const formData = ref({
            merchantName: '',
            merchantNo: '',
            username: '',
            email: '',
            password: '', // 新增明文密码字段
            passwordHash: '',
            apiSecretKey: '',
            balance: 0,
            frozenAmount: 0,
            remark: '',
            status: true, // 默认为正常状态
        })



// 验证规则
const rule = reactive({
  merchantName: [
    { required: true, message: '请输入商户名称', trigger: 'blur' }
  ],
  merchantNo: [
    { required: true, message: '请输入商户编号', trigger: 'blur' }
  ],
  username: [
    { required: true, message: '请输入用户名', trigger: 'blur' },
    { min: 3, message: '用户名至少3个字符', trigger: 'blur' }
  ],
  email: [
    { required: true, message: '请输入邮箱', trigger: 'blur' },
    { type: 'email', message: '请输入正确的邮箱格式', trigger: 'blur' }
  ],
  password: [
    { required: true, message: '请输入密码', trigger: 'blur' },
    { min: 6, message: '密码至少6个字符', trigger: 'blur' }
  ]
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
  const table = await getSmsCustomersList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
  if (table.code === 0) {
    tableData.value = table.data.list
    total.value = table.data.total
    page.value = table.data.page
    pageSize.value = table.data.pageSize
  }
}

getTableData()

// ============== 表格控制部分结束 ===============

// 交易类型字典
const transactionTypeOptions = ref([])

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
            deleteSmsCustomersFunc(row)
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
      const res = await deleteSmsCustomersByIds({ IDs })
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
const updateSmsCustomersFunc = async(row) => {
    const res = await findSmsCustomers({ ID: row.ID })
    type.value = 'update'
    if (res.code === 0) {
        formData.value = res.data
        dialogFormVisible.value = true
    }
}


// 删除行
const deleteSmsCustomersFunc = async (row) => {
    const res = await deleteSmsCustomers({ ID: row.ID })
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
    showPassword.value = false // 重置密码显示状态
    formData.value = {
        username: '',
        email: '',
        password: '', // 重置密码字段
        passwordHash: '',
        apiSecretKey: '',
        balance: 0,
        status: true, // 默认为正常状态
        }
}
// 弹窗确定
const enterDialog = async () => {
     btnLoading.value = true
     elFormRef.value?.validate( async (valid) => {
             if (!valid) return btnLoading.value = false
              
              // 创建要发送的数据
              const submitData = { ...formData.value }
              
              // 新增时需要密码，编辑时不需要
              if (type.value === 'create') {
                if (!submitData.password || submitData.password.length < 6) {
                  ElMessage({
                    type: 'error',
                    message: '请输入至少6位字符的密码'
                  })
                  btnLoading.value = false
                  return
                }
              } else {
                // 编辑时删除密码字段
                delete submitData.password
              }
              
              let res
              switch (type.value) {
                case 'create':
                  res = await createSmsCustomers(submitData)
                  break
                case 'update':
                  res = await updateSmsCustomers(submitData)
                  break
                default:
                  res = await createSmsCustomers(submitData)
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
  const res = await findSmsCustomers({ ID: row.ID })
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

// 生成随机密码
const generatePassword = () => {
  const chars = 'ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789!@#$%^&*'
  let password = ''
  for (let i = 0; i < 12; i++) {
    password += chars.charAt(Math.floor(Math.random() * chars.length))
  }
  formData.value.password = password
  // 复制到剪贴板
  navigator.clipboard.writeText(password).then(() => {
    ElMessage({
      type: 'success',
      message: '密码已生成并复制到剪贴板'
    })
  }).catch(() => {
    ElMessage({
      type: 'success', 
      message: '密码已生成：' + password
    })
  })
}

// 复制到剪贴板
const copyToClipboard = (text) => {
  navigator.clipboard.writeText(text).then(() => {
    ElMessage({
      type: 'success',
      message: '已复制到剪贴板'
    })
  }).catch(() => {
    // 备选方案：使用传统方法
    const textarea = document.createElement('textarea')
    textarea.value = text
    document.body.appendChild(textarea)
    textarea.select()
    document.execCommand('copy')
    document.body.removeChild(textarea)
    ElMessage({
      type: 'success',
      message: '已复制到剪贴板'
    })
  })
}

// ============ 充值/扣费部分 ============
const creditDebitDialogVisible = ref(false)
const creditDebitLoading = ref(false)
const creditDebitFormRef = ref()
const creditDebitForm = ref({
  type: '4',
  amount: undefined,
  notes: ''
})

const creditDebitRule = reactive({
  type: [
    { required: true, message: '请选择操作类型', trigger: 'change' }
  ],
  amount: [
    { required: true, message: '请输入金额', trigger: 'blur' },
    { type: 'number', message: '金额必须是数字', trigger: 'blur' }
  ]
})

// 打开充值/扣费弹窗
const openCreditDebitDialog = () => {
  creditDebitForm.value = {
    type: '1',
    amount: undefined,
    notes: ''
  }
  creditDebitDialogVisible.value = true
}

// 从表格打开充值/扣费弹窗
const openCreditDebitDialogFromTable = async (row) => {
  const res = await findSmsCustomers({ ID: row.ID })
  if (res.code === 0) {
    detailForm.value = res.data
    openCreditDebitDialog()
  }
}

// 关闭充值/扣费弹窗
const closeCreditDebitDialog = () => {
  creditDebitDialogVisible.value = false
  creditDebitFormRef.value?.clearValidate()
}

// 计算变动后的余额
const calculateNewBalance = () => {
  const current = detailForm.value.balance || 0
  const currentFrozen = detailForm.value.frozenAmount || 0
  const amount = creditDebitForm.value.amount || 0
  const type = creditDebitForm.value.type
  
  if (type === '4') { // 上分
    return (current + amount).toFixed(2)
  } else if (type === '5') { // 下分
    return (current - amount).toFixed(2)
  } else if (type === '6') { // 冻结金额
    return `余额: ${(current - amount).toFixed(2)}, 冻结: ${(currentFrozen + amount).toFixed(2)}`
  } else if (type === '7') { // 冻结金额返回
    return `余额: ${(current + amount).toFixed(2)}, 冻结: ${(currentFrozen - amount).toFixed(2)}`
  }
  return current.toFixed(2)
}

// 提交充值/扣费
const submitCreditDebit = async () => {
  creditDebitFormRef.value?.validate(async (valid) => {
    if (!valid) return

    const current = detailForm.value.balance || 0
    const currentFrozen = detailForm.value.frozenAmount || 0
    const type = creditDebitForm.value.type

    // 检查下分时余额是否足够
    if (type === '5') {
      if (current < creditDebitForm.value.amount) {
        ElMessage({
          type: 'error',
          message: '余额不足，无法下分'
        })
        return
      }
    }

    // 检查冻结金额时余额是否足够
    if (type === '6') {
      if (current < creditDebitForm.value.amount) {
        ElMessage({
          type: 'error',
          message: '可用余额不足，无法冻结'
        })
        return
      }
    }

    // 检查冻结金额返回时冻结金额是否足够
    if (type === '7') {
      if (currentFrozen < creditDebitForm.value.amount) {
        ElMessage({
          type: 'error',
          message: '冻结金额不足，无法返回'
        })
        return
      }
    }

    creditDebitLoading.value = true
    const data = {
      customerId: detailForm.value.ID,
      amount: creditDebitForm.value.amount,
      type: creditDebitForm.value.type,
      notes: creditDebitForm.value.notes || undefined
    }

    const res = await creditDebitSmsCustomers(data)
    creditDebitLoading.value = false

    if (res.code === 0) {
      ElMessage({
        type: 'success',
        message: '操作成功'
      })
      closeCreditDebitDialog()
      // 刷新详情
      getDetails(detailForm.value)
      getTableData()
    }
  })
}

// ============ 业务配置部分 ============
const businessConfigDialogVisible = ref(false)
const businessConfigLoading = ref(false)
const businessConfigFormRef = ref()
const businessConfigForm = ref({
  businessConfig: []
})

const businessTypeSelectorVisible = ref(false)
const platformBusinessTypes = ref([])
const businessTypeTable = ref()
const selectedBusinessTypes = ref([])

// 打开业务配置对话框
const openBusinessConfigDialog = async () => {
  // 加载已有的业务配置
  try {
    const res = await getBusinessConfigSmsCustomers({ customerId: detailForm.value.ID })
    if (res.code === 0 && res.data) {
      // 映射后端数据到表单结构
      businessConfigForm.value.businessConfig = res.data.map(config => ({
        platformBusinessTypeId: config.platformBusinessTypeId,
        businessCode: config.businessCode || config.code,
        businessName: config.businessName || config.name,
        cost: Number(config.cost) || 0,
        weight: Number(config.weight) || 0,
        status: Number(config.status) || 1  // 确保是整数类型
      }))
    } else {
      businessConfigForm.value = {
        businessConfig: []
      }
    }
  } catch (error) {
    console.error('加载业务配置失败:', error)
    businessConfigForm.value = {
      businessConfig: []
    }
  }
  
  // 加载平台业务类型列表
  await loadPlatformBusinessTypes()
  businessConfigDialogVisible.value = true
}

// 从表格打开业务配置对话框
const openBusinessConfigFromTable = async (row) => {
  const res = await findSmsCustomers({ ID: row.ID })
  if (res.code === 0) {
    detailForm.value = res.data
    openBusinessConfigDialog()
  }
}

// 关闭业务配置对话框
const closeBusinessConfigDialog = () => {
  businessConfigDialogVisible.value = false
  businessConfigFormRef.value?.clearValidate()
}

// 加载平台业务类型
const loadPlatformBusinessTypes = async () => {
  const res = await getSmsPlatformBusinessTypesList({ page: 1, pageSize: 999 })
  if (res.code === 0) {
    platformBusinessTypes.value = res.data.list || []
  }
}

// 打开业务类型选择器
const openBusinessTypeSelector = () => {
  businessTypeSelectorVisible.value = true
}

// 处理业务类型选择
const handleBusinessTypeSelection = (selection) => {
  selectedBusinessTypes.value = selection
}

// 检查业务是否可选择（未添加的才能选）
const checkBusinessSelectable = (row) => {
  return !isBusinessAdded(row)
}

// 检查业务是否已添加
const isBusinessAdded = (row) => {
  return businessConfigForm.value.businessConfig.some(
    item => item.platformBusinessTypeId === row.ID
  )
}

// 获取已配置业务的详细信息
const getExistingBusinessConfig = (row) => {
  return businessConfigForm.value.businessConfig.find(
    item => item.platformBusinessTypeId === row.ID
  )
}

// 确认业务类型选择
const confirmBusinessTypeSelection = () => {
  selectedBusinessTypes.value.forEach(type => {
    // 检查是否已存在
    const exists = businessConfigForm.value.businessConfig.find(
      item => item.platformBusinessTypeId === type.ID
    )
    if (!exists) {
      businessConfigForm.value.businessConfig.push({
        platformBusinessTypeId: type.ID,
        businessCode: type.code || type.Code,
        businessName: type.name || type.Name,
        cost: 0.0000,
        weight: 1,
        status: true
      })
    }
  })
  businessTypeSelectorVisible.value = false
  // 清空选择
  businessTypeTable.value?.clearSelection()
}

// 提交业务配置
const submitBusinessConfig = async () => {
  if (businessConfigForm.value.businessConfig.length === 0) {
    ElMessage({
      type: 'warning',
      message: '请至少选择一个业务类型'
    })
    return
  }

  businessConfigLoading.value = true
  
  // 数据清洗：确保所有字段类型正确，过滤掉无效数据
  const cleanedConfig = businessConfigForm.value.businessConfig
    .filter(item => item && item.platformBusinessTypeId) // 过滤掉空项和无ID的项
    .map(item => ({
      platformBusinessTypeId: Number(item.platformBusinessTypeId),
      businessCode: String(item.businessCode || ''),
      businessName: String(item.businessName || ''),
      cost: Number(item.cost) || 0,
      weight: Number(item.weight) || 0,
      status: Number(item.status) === 1 ? 1 : 0  // 确保是 0 或 1
    }))

  if (cleanedConfig.length === 0) {
    businessConfigLoading.value = false
    ElMessage({
      type: 'warning',
      message: '没有有效的业务配置数据'
    })
    return
  }

  const data = {
    customerId: detailForm.value.ID,
    businessConfig: cleanedConfig
  }

  console.log('提交业务配置数据:', JSON.stringify(data, null, 2))

  const res = await configureBusinessSmsCustomers(data)
  businessConfigLoading.value = false

  if (res.code === 0) {
    ElMessage({
      type: 'success',
      message: '配置成功'
    })
    closeBusinessConfigDialog()
    // 刷新详情
    getDetails(detailForm.value)
  }
}

</script>

<style>

</style>
