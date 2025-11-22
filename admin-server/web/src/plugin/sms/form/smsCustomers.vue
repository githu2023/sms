
<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="用户名" prop="username">
          <el-input v-model="formData.username" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="邮箱" prop="email">
          <el-input v-model="formData.email" :clearable="true"  placeholder="请输入邮箱" />
       </el-form-item>
        <!-- <el-form-item label="密码哈希" prop="passwordHash">
          <el-input v-model="formData.passwordHash" :clearable="true"  placeholder="系统自动生成" />
       </el-form-item> -->
        <el-form-item label="API密钥" prop="apiSecretKey">
          <el-input v-model="formData.apiSecretKey" :clearable="true"  placeholder="可选，系统自动生成" />
       </el-form-item>
        <el-form-item label="余额" prop="balance">
          <el-input-number v-model="formData.balance" :precision="2" :clearable="true"></el-input-number>
       </el-form-item>
        <el-form-item label="状态" prop="status">
          <el-switch v-model="formData.status" active-color="#13ce66" inactive-color="#ff4949" active-text="正常" inactive-text="冻结" clearable ></el-switch>
       </el-form-item>
        <el-form-item label="注册IP" prop="registrationIp">
          <el-input v-model="formData.registrationIp" :clearable="true"  placeholder="请输入注册时的IP地址" />
       </el-form-item>
        <el-form-item label="最后登录IP" prop="lastLoginIp">
          <el-input v-model="formData.lastLoginIp" :clearable="true"  placeholder="请输入" />
       </el-form-item>
        <el-form-item label="最后登录时间" prop="lastLoginAt">
          <el-date-picker v-model="formData.lastLoginAt" type="date" placeholder="选择日期" :clearable="true"></el-date-picker>
       </el-form-item>
        <el-form-item>
          <el-button :loading="btnLoading" type="primary" @click="save">保存</el-button>
          <el-button type="primary" @click="back">返回</el-button>
        </el-form-item>
      </el-form>
    </div>
  </div>
</template>

<script setup>
import {
  createSmsCustomers,
  updateSmsCustomers,
  findSmsCustomers
} from '@/plugin/sms/api/smsCustomers'

defineOptions({
    name: 'SmsCustomersForm'
})

// 自动获取字典
import { getDictFunc } from '@/utils/format'
import { useRoute, useRouter } from "vue-router"
import { ElMessage } from 'element-plus'
import { ref, reactive } from 'vue'


const route = useRoute()
const router = useRouter()

// 提交按钮loading
const btnLoading = ref(false)

const type = ref('')
const formData = ref({
            username: '',
            email: '',
            passwordHash: '',
            apiSecretKey: '',
            balance: 0,
            status: false,
            registrationIp: '',
            lastLoginIp: '',
            lastLoginAt: new Date(),
        })
// 验证规则
const rule = reactive({
})

const elFormRef = ref()

// 初始化方法
const init = async () => {
 // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (route.query.id) {
      const res = await findSmsCustomers({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data
        type.value = 'update'
      }
    } else {
      type.value = 'create'
    }
}

init()
// 保存按钮
const save = async() => {
      btnLoading.value = true
      elFormRef.value?.validate( async (valid) => {
         if (!valid) return btnLoading.value = false
            let res
           switch (type.value) {
             case 'create':
               res = await createSmsCustomers(formData.value)
               break
             case 'update':
               res = await updateSmsCustomers(formData.value)
               break
             default:
               res = await createSmsCustomers(formData.value)
               break
           }
           btnLoading.value = false
           if (res.code === 0) {
             ElMessage({
               type: 'success',
               message: '创建/更改成功'
             })
           }
       })
}

// 返回按钮
const back = () => {
    router.go(-1)
}

</script>

<style>
</style>
