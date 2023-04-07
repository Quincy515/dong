<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="备注:" prop="beiZhu">
          <el-input v-model="formData.beiZhu" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="采购计划ID:" prop="caiGouJiHuaId">
          <el-input v-model.number="formData.caiGouJiHuaId" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="采购量:" prop="caiGouLiang">
          <el-input-number v-model="formData.caiGouLiang" :precision="2" :clearable="true"></el-input-number>
        </el-form-item>
        <el-form-item label="材料ID:" prop="caiLiaoId">
          <el-input v-model.number="formData.caiLiaoId" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="单价:" prop="danJia">
          <el-input v-model="formData.danJia" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="供应商ID:" prop="gongYingShangId">
          <el-input v-model.number="formData.gongYingShangId" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="所属:" prop="group">
          <el-input v-model="formData.group" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="入库量:" prop="ruKuLiang">
          <el-input-number v-model="formData.ruKuLiang" :precision="2" :clearable="true"></el-input-number>
        </el-form-item>
        <el-form-item label="状态:" prop="zhuangTai">
          <el-input v-model.number="formData.zhuangTai" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="总价:" prop="zongJia">
          <el-input-number v-model="formData.zongJia" :precision="2" :clearable="true"></el-input-number>
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
  name: 'CaiGouJiHuaMingXiBiao'
}
</script>

<script setup>
import {
  createCaiGouJiHuaMingXiBiao,
  updateCaiGouJiHuaMingXiBiao,
  findCaiGouJiHuaMingXiBiao
} from '@/api/caiGouJiHuaMingXiBiao'

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
            caiGouJiHuaId: 0,
            caiGouLiang: 0,
            caiLiaoId: 0,
            danJia: '',
            gongYingShangId: 0,
            group: '',
            ruKuLiang: 0,
            zhuangTai: 0,
            zongJia: 0,
        })
// 验证规则
const rule = reactive({
})

const elFormRef = ref()

// 初始化方法
const init = async () => {
 // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (route.query.id) {
      const res = await findCaiGouJiHuaMingXiBiao({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data.recaiGouJiHuaMingXiBiao
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
               res = await createCaiGouJiHuaMingXiBiao(formData.value)
               break
             case 'update':
               res = await updateCaiGouJiHuaMingXiBiao(formData.value)
               break
             default:
               res = await createCaiGouJiHuaMingXiBiao(formData.value)
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
