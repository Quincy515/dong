<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="备注:" prop="beiZhu">
          <el-input v-model="formData.beiZhu" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="所属:" prop="group">
          <el-input v-model="formData.group" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="结束日期:" prop="jieShuRiQi">
          <el-date-picker v-model="formData.jieShuRiQi" type="date" placeholder="选择日期" :clearable="true"></el-date-picker>
        </el-form-item>
        <el-form-item label="进度:" prop="jinDu">
          <el-input v-model="formData.jinDu" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="开始日期:" prop="kaiShiRiQi">
          <el-date-picker v-model="formData.kaiShiRiQi" type="date" placeholder="选择日期" :clearable="true"></el-date-picker>
        </el-form-item>
        <el-form-item label="请假类型:" prop="leiXing">
          <el-select v-model="formData.leiXing" placeholder="请选择" :clearable="true">
            <el-option v-for="(item,key) in qing_jia_lei_xingOptions" :key="key" :label="item.label" :value="item.value" />
          </el-select>
        </el-form-item>
        <el-form-item label="审核人:" prop="shenHeRen">
          <el-input v-model="formData.shenHeRen" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="审核时间:" prop="shenHeShiJian">
          <el-date-picker v-model="formData.shenHeShiJian" type="date" placeholder="选择日期" :clearable="true"></el-date-picker>
        </el-form-item>
        <el-form-item label="审核详情:" prop="shenHeXiangQing">
          <el-input v-model="formData.shenHeXiangQing" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="审核状态:" prop="shenHeZhuangTai">
          <el-input v-model="formData.shenHeZhuangTai" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="事由:" prop="shiYou">
          <el-input v-model="formData.shiYou" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="天数:" prop="tianShu">
          <el-input v-model.number="formData.tianShu" :clearable="true" placeholder="请输入" />
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
  name: 'QingJiaGuanLi'
}
</script>

<script setup>
import {
  createQingJiaGuanLi,
  updateQingJiaGuanLi,
  findQingJiaGuanLi
} from '@/api/qingJiaGuanLi'

// 自动获取字典
import { getDictFunc } from '@/utils/format'
import { useRoute, useRouter } from "vue-router"
import { ElMessage } from 'element-plus'
import { ref, reactive } from 'vue'
const route = useRoute()
const router = useRouter()

const type = ref('')
const qing_jia_lei_xingOptions = ref([])
const formData = ref({
            beiZhu: '',
            group: '',
            jieShuRiQi: new Date(),
            jinDu: '',
            kaiShiRiQi: new Date(),
            leiXing: undefined,
            shenHeRen: '',
            shenHeShiJian: new Date(),
            shenHeXiangQing: '',
            shenHeZhuangTai: '',
            shiYou: '',
            tianShu: 0,
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
      const res = await findQingJiaGuanLi({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data.reqingJiaGuanLi
        type.value = 'update'
      }
    } else {
      type.value = 'create'
    }
    qing_jia_lei_xingOptions.value = await getDictFunc('qing_jia_lei_xing')
}

init()
// 保存按钮
const save = async() => {
      elFormRef.value?.validate( async (valid) => {
         if (!valid) return
            let res
           switch (type.value) {
             case 'create':
               res = await createQingJiaGuanLi(formData.value)
               break
             case 'update':
               res = await updateQingJiaGuanLi(formData.value)
               break
             default:
               res = await createQingJiaGuanLi(formData.value)
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
