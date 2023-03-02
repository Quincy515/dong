<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="所属:" prop="group">
          <el-input v-model="formData.group" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="证书名称:" prop="zhengShuMingCheng">
          <el-input v-model="formData.zhengShuMingCheng" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="证书编号:" prop="zhengShuBianHao">
          <el-input v-model="formData.zhengShuBianHao" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="证书ID:" prop="zhengShuId">
          <el-input v-model.number="formData.zhengShuId" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="借出人ID:" prop="jieChuRen">
          <el-input v-model.number="formData.jieChuRen" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="借出日期:" prop="jieChuRiQi">
          <el-date-picker v-model="formData.jieChuRiQi" type="date" placeholder="选择日期" :clearable="true"></el-date-picker>
        </el-form-item>
        <el-form-item label="借出事由:" prop="jieChuShiYou">
          <el-input v-model="formData.jieChuShiYou" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="预计归还日期:" prop="yuJiGuiHaiRiQi">
          <el-date-picker v-model="formData.yuJiGuiHaiRiQi" type="date" placeholder="选择日期" :clearable="true"></el-date-picker>
        </el-form-item>
        <el-form-item label="实际归还日期:" prop="shiJiGuiHaiRiQi">
          <el-date-picker v-model="formData.shiJiGuiHaiRiQi" type="date" placeholder="选择日期" :clearable="true"></el-date-picker>
        </el-form-item>
        <el-form-item label="添加人:" prop="tianJiaRen">
          <el-input v-model.number="formData.tianJiaRen" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="借出备注:" prop="jieChuBeiZhu">
          <el-input v-model="formData.jieChuBeiZhu" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="归还备注:" prop="guiHaiBeiZhu">
          <el-input v-model="formData.guiHaiBeiZhu" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item>
          <el-button size="small" type="primary" @click="save">保存</el-button>
          <el-button size="small" type="primary" @click="back">返回</el-button>
        </el-form-item>
      </el-form>
    </div>
  </div>
</template>

<script>
export default {
  name: 'ZhengShuJieChuJiLuBiao'
}
</script>

<script setup>
import {
  createZhengShuJieChuJiLuBiao,
  updateZhengShuJieChuJiLuBiao,
  findZhengShuJieChuJiLuBiao
} from '@/api/zhengShuJieChuJiLuBiao'

// 自动获取字典
import { getDictFunc } from '@/utils/format'
import { useRoute, useRouter } from "vue-router"
import { ElMessage } from 'element-plus'
import { ref, reactive } from 'vue'
const route = useRoute()
const router = useRouter()

const type = ref('')
const formData = ref({
            group: '',
            zhengShuMingCheng: '',
            zhengShuBianHao: '',
            zhengShuId: 0,
            jieChuRen: 0,
            jieChuRiQi: new Date(),
            jieChuShiYou: '',
            yuJiGuiHaiRiQi: new Date(),
            shiJiGuiHaiRiQi: new Date(),
            tianJiaRen: 0,
            jieChuBeiZhu: '',
            guiHaiBeiZhu: '',
        })
// 验证规则
const rule = reactive({
})

const elFormRef = ref()

// 初始化方法
const init = async () => {
 // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (route.query.id) {
      const res = await findZhengShuJieChuJiLuBiao({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data.rezhengShuJieChuJiLuBiao
        type.value = 'update'
      }
    } else {
      type.value = 'create'
    }
}

init()
// 保存按钮
const save = async() => {
      elFormRef.value?.validate( async (valid) => {
         if (!valid) return
            let res
           switch (type.value) {
             case 'create':
               res = await createZhengShuJieChuJiLuBiao(formData.value)
               break
             case 'update':
               res = await updateZhengShuJieChuJiLuBiao(formData.value)
               break
             default:
               res = await createZhengShuJieChuJiLuBiao(formData.value)
               break
           }
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
