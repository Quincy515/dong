<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="备注:" prop="beiZhu">
          <el-input v-model="formData.beiZhu" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="出勤天数:" prop="chuQinTianShu">
          <el-input v-model.number="formData.chuQinTianShu" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="所属:" prop="group">
          <el-input v-model="formData.group" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="结束日期:" prop="jieShuRiQi">
          <el-input v-model="formData.jieShuRiQi" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="开始日期:" prop="kaiShiRiQi">
          <el-input v-model="formData.kaiShiRiQi" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="请假天数:" prop="qingJiaTianShu">
          <el-input v-model.number="formData.qingJiaTianShu" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="详情:" prop="xiangQing">
          <el-input v-model="formData.xiangQing" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="姓名:" prop="xingMing">
          <el-input v-model="formData.xingMing" :clearable="true" placeholder="请输入" />
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
  name: 'KaoQinTongJi'
}
</script>

<script setup>
import {
  createKaoQinTongJi,
  updateKaoQinTongJi,
  findKaoQinTongJi
} from '@/api/kaoQinTongJi'

// 自动获取字典
import { getDictFunc } from '@/utils/format'
import { useRoute, useRouter } from "vue-router"
import { ElMessage } from 'element-plus'
import { ref, reactive } from 'vue'
const route = useRoute()
const router = useRouter()

const type = ref('')
const formData = ref({
            beiZhu: '',
            chuQinTianShu: 0,
            group: '',
            jieShuRiQi: '',
            kaiShiRiQi: '',
            qingJiaTianShu: 0,
            xiangQing: '',
            xingMing: '',
        })
// 验证规则
const rule = reactive({
})

const elFormRef = ref()

// 初始化方法
const init = async () => {
 // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (route.query.id) {
      const res = await findKaoQinTongJi({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data.rekaoQinTongJi
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
               res = await createKaoQinTongJi(formData.value)
               break
             case 'update':
               res = await updateKaoQinTongJi(formData.value)
               break
             default:
               res = await createKaoQinTongJi(formData.value)
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
