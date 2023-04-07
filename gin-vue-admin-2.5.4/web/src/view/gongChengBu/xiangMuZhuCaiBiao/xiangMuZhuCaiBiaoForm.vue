<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="备注:" prop="beiZhu">
          <el-input v-model="formData.beiZhu" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="材料编号:" prop="caiLiaoBianHao">
          <el-input v-model="formData.caiLiaoBianHao" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="材料名称:" prop="caiLiaoMingCheng">
          <el-input v-model="formData.caiLiaoMingCheng" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="单价:" prop="danJia">
          <el-input-number v-model="formData.danJia" :precision="2" :clearable="true"></el-input-number>
        </el-form-item>
        <el-form-item label="单位:" prop="danWei">
          <el-input v-model="formData.danWei" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="所属:" prop="group">
          <el-input v-model="formData.group" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="合计:" prop="heJi">
          <el-input-number v-model="formData.heJi" :precision="2" :clearable="true"></el-input-number>
        </el-form-item>
        <el-form-item label="价差合计:" prop="jiaChaHeJi">
          <el-input-number v-model="formData.jiaChaHeJi" :precision="2" :clearable="true"></el-input-number>
        </el-form-item>
        <el-form-item label="现行价格:" prop="xianHangJiaGe">
          <el-input-number v-model="formData.xianHangJiaGe" :precision="2" :clearable="true"></el-input-number>
        </el-form-item>
        <el-form-item label="项目名称:" prop="xiangMuMingCheng">
          <el-input v-model.number="formData.xiangMuMingCheng" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="用量:" prop="yongLiang">
          <el-input-number v-model="formData.yongLiang" :precision="2" :clearable="true"></el-input-number>
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
  name: 'XiangMuZhuCaiBiao'
}
</script>

<script setup>
import {
  createXiangMuZhuCaiBiao,
  updateXiangMuZhuCaiBiao,
  findXiangMuZhuCaiBiao
} from '@/api/xiangMuZhuCaiBiao'

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
            caiLiaoBianHao: '',
            caiLiaoMingCheng: '',
            danJia: 0,
            danWei: '',
            group: '',
            heJi: 0,
            jiaChaHeJi: 0,
            xianHangJiaGe: 0,
            xiangMuMingCheng: 0,
            yongLiang: 0,
        })
// 验证规则
const rule = reactive({
})

const elFormRef = ref()

// 初始化方法
const init = async () => {
 // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (route.query.id) {
      const res = await findXiangMuZhuCaiBiao({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data.rexiangMuZhuCaiBiao
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
               res = await createXiangMuZhuCaiBiao(formData.value)
               break
             case 'update':
               res = await updateXiangMuZhuCaiBiao(formData.value)
               break
             default:
               res = await createXiangMuZhuCaiBiao(formData.value)
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
