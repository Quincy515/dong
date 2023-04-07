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
        <el-form-item label="节点ID只保存最近的节点:" prop="jieDianId">
          <el-input v-model.number="formData.jieDianId" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="上传人:" prop="shangChuanRen">
          <el-input v-model.number="formData.shangChuanRen" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="上传日期:" prop="shangChuanRiQi">
          <el-date-picker v-model="formData.shangChuanRiQi" type="date" placeholder="选择日期" :clearable="true"></el-date-picker>
        </el-form-item>
        <el-form-item label="审核人:" prop="shenHeRen">
          <el-input v-model.number="formData.shenHeRen" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="审核日期:" prop="shenHeRiQi">
          <el-date-picker v-model="formData.shenHeRiQi" type="date" placeholder="选择日期" :clearable="true"></el-date-picker>
        </el-form-item>
        <el-form-item label="是否审核:" prop="shiFouShenHe">
          <el-input v-model.number="formData.shiFouShenHe" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="文件id:" prop="wenJianId">
          <el-input v-model.number="formData.wenJianId" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="文件类型:" prop="wenJianLeiXing">
          <el-input v-model.number="formData.wenJianLeiXing" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="文件名称:" prop="wenJianMingCheng">
          <el-input v-model="formData.wenJianMingCheng" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="阅读次数:" prop="yueDuCiShu">
          <el-input v-model.number="formData.yueDuCiShu" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="状态:" prop="zhuangTai">
          <el-input v-model.number="formData.zhuangTai" :clearable="true" placeholder="请输入" />
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
  name: 'WenJianZiLiaoBiao'
}
</script>

<script setup>
import {
  createWenJianZiLiaoBiao,
  updateWenJianZiLiaoBiao,
  findWenJianZiLiaoBiao
} from '@/api/wenJianZiLiaoBiao'

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
            group: '',
            jieDianId: 0,
            shangChuanRen: 0,
            shangChuanRiQi: new Date(),
            shenHeRen: 0,
            shenHeRiQi: new Date(),
            shiFouShenHe: 0,
            wenJianId: 0,
            wenJianLeiXing: 0,
            wenJianMingCheng: '',
            yueDuCiShu: 0,
            zhuangTai: 0,
        })
// 验证规则
const rule = reactive({
})

const elFormRef = ref()

// 初始化方法
const init = async () => {
 // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (route.query.id) {
      const res = await findWenJianZiLiaoBiao({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data.rewenJianZiLiaoBiao
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
               res = await createWenJianZiLiaoBiao(formData.value)
               break
             case 'update':
               res = await updateWenJianZiLiaoBiao(formData.value)
               break
             default:
               res = await createWenJianZiLiaoBiao(formData.value)
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
