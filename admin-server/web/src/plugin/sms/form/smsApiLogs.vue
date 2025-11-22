
<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="客户ID" prop="customerId">
          <el-input v-model.number="formData.customerId" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="请求IP" prop="requestIp">
          <el-input v-model="formData.requestIp" :clearable="true"  placeholder="请输入" />
       </el-form-item>
        <el-form-item label="API路径" prop="requestPath">
          <el-input v-model="formData.requestPath" :clearable="true"  placeholder="请输入" />
       </el-form-item>
        <el-form-item label="请求体" prop="requestBody">
          <el-input v-model="formData.requestBody" :clearable="true"  placeholder="请输入" />
       </el-form-item>
        <el-form-item label="响应码" prop="responseCode">
          <el-input v-model.number="formData.responseCode" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="耗时(毫秒)" prop="durationMs">
          <el-input v-model.number="formData.durationMs" :clearable="true" placeholder="请输入" />
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
  createSmsApiLogs,
  updateSmsApiLogs,
  findSmsApiLogs
} from '@/plugin/sms/api/smsApiLogs'

defineOptions({
    name: 'SmsApiLogsForm'
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
            requestIp: '',
            requestPath: '',
            requestBody: '',
            responseCode: undefined,
            durationMs: undefined,
        })
// 验证规则
const rule = reactive({
})

const elFormRef = ref()

// 初始化方法
const init = async () => {
 // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (route.query.id) {
      const res = await findSmsApiLogs({ ID: route.query.id })
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
               res = await createSmsApiLogs(formData.value)
               break
             case 'update':
               res = await updateSmsApiLogs(formData.value)
               break
             default:
               res = await createSmsApiLogs(formData.value)
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
