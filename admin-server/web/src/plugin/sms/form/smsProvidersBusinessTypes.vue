<template>
  <el-form :model="formData" label-position="top" ref="elFormRef" :rules="rule" label-width="80px">
    <el-form-item label="三方渠道ID:" prop="providerId">
      <el-input-number v-model="formData.providerId" :min="0" :step="1" :clearable="true" placeholder="请输入三方渠道ID" />
    </el-form-item>
    <el-form-item label="三方编码:" prop="providerCode">
      <el-input v-model="formData.providerCode" :clearable="true" placeholder="请输入三方编码" />
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
</template>

<script setup>
import { ref, reactive } from 'vue'

defineOptions({
    name: 'SmsProvidersBusinessTypesForm'
})

const formData = ref({
  providerId: 0,
  providerCode: '',
  businessName: '',
  businessCode: '',
  price: 0,
  status: false,
  remark: '',
})

const rule = reactive({
})

const elFormRef = ref()

</script>

<style>
</style>
