
<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="客户ID" prop="customerId">
          <el-input v-model.number="formData.customerId" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="服务商ID" prop="providerId">
          <el-input v-model.number="formData.providerId" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="业务类型ID" prop="businessTypeId">
          <el-input v-model.number="formData.businessTypeId" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="卡类型" prop="cardType">
          <el-input v-model="formData.cardType" :clearable="true"  placeholder="如: physical, virtual" />
       </el-form-item>
        <el-form-item label="手机号" prop="phoneNumber">
          <el-input v-model="formData.phoneNumber" :clearable="true"  placeholder="请输入获取到的手机号" />
       </el-form-item>
        <el-form-item label="验证码" prop="verificationCode">
          <el-input v-model="formData.verificationCode" :clearable="true"  placeholder="请输入获取到的验证码" />
       </el-form-item>
        <el-form-item label="费用" prop="cost">
          <el-input-number v-model="formData.cost" :precision="2" :clearable="true"></el-input-number>
       </el-form-item>
        <el-form-item label="状态" prop="status">
          <el-switch v-model="formData.status" active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" clearable ></el-switch>
       </el-form-item>
        <el-form-item label="过期时间" prop="expiresAt">
          <el-date-picker v-model="formData.expiresAt" type="date" placeholder="选择日期" :clearable="true"></el-date-picker>
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
  createSmsPhoneAssignments,
  updateSmsPhoneAssignments,
  findSmsPhoneAssignments
} from '@/plugin/sms/api/smsPhoneAssignments'

defineOptions({
    name: 'SmsPhoneAssignmentsForm'
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
            customerId: undefined,
            providerId: undefined,
            businessTypeId: undefined,
            cardType: '',
            phoneNumber: '',
            verificationCode: '',
            cost: 0,
            status: false,
            expiresAt: new Date(),
        })
// 验证规则
const rule = reactive({
})

const elFormRef = ref()

// 初始化方法
const init = async () => {
 // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (route.query.id) {
      const res = await findSmsPhoneAssignments({ ID: route.query.id })
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
