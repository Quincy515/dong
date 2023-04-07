<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="单价:" prop="danJia">
          <el-input-number v-model="formData.danJia" :precision="2" :clearable="true"></el-input-number>
        </el-form-item>
        <el-form-item label="所属:" prop="group">
          <el-input v-model="formData.group" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="计划退租日期:" prop="jiHuaTuiZuRiQi">
          <el-date-picker v-model="formData.jiHuaTuiZuRiQi" type="date" placeholder="选择日期" :clearable="true"></el-date-picker>
        </el-form-item>
        <el-form-item label="计划租赁日期:" prop="jiHuaZuLinRiQi">
          <el-date-picker v-model="formData.jiHuaZuLinRiQi" type="date" placeholder="选择日期" :clearable="true"></el-date-picker>
        </el-form-item>
        <el-form-item label="设备id:" prop="sheBeiId">
          <el-input v-model.number="formData.sheBeiId" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="实际退租日期:" prop="shiJiTuiZuRiQi">
          <el-date-picker v-model="formData.shiJiTuiZuRiQi" type="date" placeholder="选择日期" :clearable="true"></el-date-picker>
        </el-form-item>
        <el-form-item label="实际租赁日期:" prop="shiJiZuLinRiQi">
          <el-date-picker v-model="formData.shiJiZuLinRiQi" type="date" placeholder="选择日期" :clearable="true"></el-date-picker>
        </el-form-item>
        <el-form-item label="状态:" prop="zhuangTai">
          <el-input v-model.number="formData.zhuangTai" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="租赁单据ID:" prop="zuLinDanJuId">
          <el-input v-model.number="formData.zuLinDanJuId" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="租赁数量:" prop="zuLinShuLiang">
          <el-input v-model.number="formData.zuLinShuLiang" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="租赁天数:" prop="zuLinTianShu">
          <el-input v-model.number="formData.zuLinTianShu" :clearable="true" placeholder="请输入" />
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
  name: 'SheBeiMingXiBiao'
}
</script>

<script setup>
import {
  createSheBeiMingXiBiao,
  updateSheBeiMingXiBiao,
  findSheBeiMingXiBiao
} from '@/api/sheBeiMingXiBiao'

// 自动获取字典
import { getDictFunc } from '@/utils/format'
import { useRoute, useRouter } from "vue-router"
import { ElMessage } from 'element-plus'
import { ref, reactive } from 'vue'
const route = useRoute()
const router = useRouter()

const type = ref('')
const formData = ref({
            danJia: 0,
            group: '',
            jiHuaTuiZuRiQi: new Date(),
            jiHuaZuLinRiQi: new Date(),
            sheBeiId: 0,
            shiJiTuiZuRiQi: new Date(),
            shiJiZuLinRiQi: new Date(),
            zhuangTai: 0,
            zuLinDanJuId: 0,
            zuLinShuLiang: 0,
            zuLinTianShu: 0,
        })
// 验证规则
const rule = reactive({
})

const elFormRef = ref()

// 初始化方法
const init = async () => {
 // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (route.query.id) {
      const res = await findSheBeiMingXiBiao({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data.resheBeiMingXiBiao
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
               res = await createSheBeiMingXiBiao(formData.value)
               break
             case 'update':
               res = await updateSheBeiMingXiBiao(formData.value)
               break
             default:
               res = await createSheBeiMingXiBiao(formData.value)
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
