
<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="平台业务ID:" prop="platformBusinessTypeId">
          <el-input v-model.number="formData.platformBusinessTypeId" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="平台业务编码:" prop="platformBusinessCode">
          <el-input v-model="formData.platformBusinessCode" :clearable="true"  placeholder="请输入平台业务编码" />
       </el-form-item>
        <el-form-item label="三方业务ID:" prop="providerBusinessTypeId">
          <el-input v-model.number="formData.providerBusinessTypeId" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="三方编码:" prop="providerCode">
          <el-input v-model="formData.providerCode" :clearable="true"  placeholder="请输入三方编码" />
       </el-form-item>
        <el-form-item label="三方业务编码:" prop="businessCode">
          <el-input v-model="formData.businessCode" :clearable="true"  placeholder="请输入三方业务编码" />
       </el-form-item>
        <el-form-item label="权重:" prop="weight">
          <el-input v-model.number="formData.weight" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="是否启用该映射:" prop="status">
          <el-switch v-model="formData.status" active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" clearable ></el-switch>
       </el-form-item>
        <el-form-item label="备注:" prop="remark">
          <el-input v-model="formData.remark" :clearable="true"  placeholder="请输入备注" />
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
  createSmsPlatformProviderBusinessMapping,
  updateSmsPlatformProviderBusinessMapping,
  findSmsPlatformProviderBusinessMapping
} from '@/plugin/sms/api/smsPlatformProviderBusinessMapping'

defineOptions({
    name: 'SmsPlatformProviderBusinessMappingForm'
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
            platformBusinessTypeId: undefined,
            platformBusinessCode: '',
            providerBusinessTypeId: undefined,
            providerCode: '',
            businessCode: '',
            weight: undefined,
            status: false,
            remark: '',
        })
// 验证规则
const rule = reactive({
})

const elFormRef = ref()

// 初始化方法
const init = async () => {
 // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (route.query.id) {
      const res = await findSmsPlatformProviderBusinessMapping({ ID: route.query.id })
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
