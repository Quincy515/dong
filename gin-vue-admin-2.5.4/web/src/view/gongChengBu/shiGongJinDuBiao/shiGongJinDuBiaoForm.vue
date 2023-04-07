<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="备注:" prop="beiZhu">
          <el-input v-model="formData.beiZhu" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="工作名称ID:" prop="gongZuoMingChengId">
          <el-input v-model.number="formData.gongZuoMingChengId" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="所属:" prop="group">
          <el-input v-model="formData.group" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="实际开始日期:" prop="shiJiKaiShiRiQi">
          <el-date-picker v-model="formData.shiJiKaiShiRiQi" type="date" placeholder="选择日期" :clearable="true"></el-date-picker>
        </el-form-item>
        <el-form-item label="实际天数:" prop="shiJiTianShu">
          <el-input v-model.number="formData.shiJiTianShu" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="实际完成日期:" prop="shiJiWanChengRiQi">
          <el-date-picker v-model="formData.shiJiWanChengRiQi" type="date" placeholder="选择日期" :clearable="true"></el-date-picker>
        </el-form-item>
        <el-form-item label="添加人:" prop="tianJiaRen">
          <el-input v-model.number="formData.tianJiaRen" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="添加时间:" prop="tianJiaShiJian">
          <el-date-picker v-model="formData.tianJiaShiJian" type="date" placeholder="选择日期" :clearable="true"></el-date-picker>
        </el-form-item>
        <el-form-item label="完成百分比:" prop="wanChengBaiFenBi">
          <el-input-number v-model="formData.wanChengBaiFenBi" :precision="2" :clearable="true"></el-input-number>
        </el-form-item>
        <el-form-item label="项目ID:" prop="xiangMuId">
          <el-input v-model.number="formData.xiangMuId" :clearable="true" placeholder="请输入" />
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
  name: 'ShiGongJinDuBiao'
}
</script>

<script setup>
import {
  createShiGongJinDuBiao,
  updateShiGongJinDuBiao,
  findShiGongJinDuBiao
} from '@/api/shiGongJinDuBiao'

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
            gongZuoMingChengId: 0,
            group: '',
            shiJiKaiShiRiQi: new Date(),
            shiJiTianShu: 0,
            shiJiWanChengRiQi: new Date(),
            tianJiaRen: 0,
            tianJiaShiJian: new Date(),
            wanChengBaiFenBi: 0,
            xiangMuId: 0,
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
      const res = await findShiGongJinDuBiao({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data.reshiGongJinDuBiao
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
               res = await createShiGongJinDuBiao(formData.value)
               break
             case 'update':
               res = await updateShiGongJinDuBiao(formData.value)
               break
             default:
               res = await createShiGongJinDuBiao(formData.value)
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
