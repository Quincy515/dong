<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="备注:" prop="beiZhu">
          <el-input v-model="formData.beiZhu" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="工作名称:" prop="gongZuoMingCheng">
          <el-input v-model="formData.gongZuoMingCheng" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="所属:" prop="group">
          <el-input v-model="formData.group" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="计划开始日期:" prop="jiHuaKaiShiRiQi">
          <el-date-picker v-model="formData.jiHuaKaiShiRiQi" type="date" placeholder="选择日期" :clearable="true"></el-date-picker>
        </el-form-item>
        <el-form-item label="计划天数:" prop="jiHuaTianShu">
          <el-input v-model.number="formData.jiHuaTianShu" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="计划完成日期:" prop="jiHuaWanChengRiQi">
          <el-date-picker v-model="formData.jiHuaWanChengRiQi" type="date" placeholder="选择日期" :clearable="true"></el-date-picker>
        </el-form-item>
        <el-form-item label="顺序:" prop="shunXu">
          <el-input v-model.number="formData.shunXu" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="完成比:" prop="wanChengBi">
          <el-input-number v-model="formData.wanChengBi" :precision="2" :clearable="true"></el-input-number>
        </el-form-item>
        <el-form-item label="项目ID:" prop="xiangMuId">
          <el-input v-model.number="formData.xiangMuId" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="占总工作量百分比:" prop="zhanZongGongZuoLiangBaiFenBi">
          <el-input-number v-model="formData.zhanZongGongZuoLiangBaiFenBi" :precision="2" :clearable="true"></el-input-number>
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
  name: 'XiangMuShiGongBiao'
}
</script>

<script setup>
import {
  createXiangMuShiGongBiao,
  updateXiangMuShiGongBiao,
  findXiangMuShiGongBiao
} from '@/api/xiangMuShiGongBiao'

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
            gongZuoMingCheng: '',
            group: '',
            jiHuaKaiShiRiQi: new Date(),
            jiHuaTianShu: 0,
            jiHuaWanChengRiQi: new Date(),
            shunXu: 0,
            wanChengBi: 0,
            xiangMuId: 0,
            zhanZongGongZuoLiangBaiFenBi: 0,
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
      const res = await findXiangMuShiGongBiao({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data.rexiangMuShiGongBiao
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
               res = await createXiangMuShiGongBiao(formData.value)
               break
             case 'update':
               res = await updateXiangMuShiGongBiao(formData.value)
               break
             default:
               res = await createXiangMuShiGongBiao(formData.value)
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
