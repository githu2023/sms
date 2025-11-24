
<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="服务商名称" prop="name">
          <el-input v-model="formData.name" :clearable="true"  placeholder="请输入" />
       </el-form-item>
        <el-form-item label="API配置" prop="apiConfig">
          <el-input v-model="formData.apiConfig" type="textarea" rows="4" placeholder="请输入API配置信息，如URL、密钥等" />
       </el-form-item>
        <el-form-item label="是否启用" prop="isEnabled">
          <el-switch v-model="formData.isEnabled" active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" clearable ></el-switch>
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
  createSmsProviders,
  updateSmsProviders,
  findSmsProviders
} from '@/plugin/sms/api/smsProviders'

defineOptions({
    name: 'SmsProvidersForm'
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
            name: '',
            apiConfig: '',
            isEnabled: false,
            extraConfig: null,
            extraConfigText: '', // 用于编辑的文本字段
        })

// JSON验证错误
const jsonError = ref('')

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

// 验证规则
const rule = reactive({
})

const elFormRef = ref()

// 初始化方法
const init = async () => {
 // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (route.query.id) {
      const res = await findSmsProviders({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data
        // 将extraConfig对象转换为文本用于编辑
        formData.value.extraConfigText = extraConfigToText(res.data.extraConfig)
        type.value = 'update'
      }
    } else {
      type.value = 'create'
    }
}

init()
// 保存按钮
const save = async() => {
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
