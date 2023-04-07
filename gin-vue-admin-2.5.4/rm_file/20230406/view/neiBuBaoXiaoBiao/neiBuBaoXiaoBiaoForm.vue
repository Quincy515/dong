<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="报销部门:" prop="baoXiaoBuMen">
          <el-select v-model="formData.baoXiaoBuMen" placeholder="请选择" :clearable="true">
            <el-option v-for="(item,key) in baoXiaoShiXiangOptions" :key="key" :label="item.label" :value="item.value" />
          </el-select>
        </el-form-item>
        <el-form-item label="报销金额:" prop="baoXiaoJinE">
          <el-input-number v-model="formData.baoXiaoJinE" :precision="2" :clearable="true"></el-input-number>
        </el-form-item>
        <el-form-item label="报销内容:" prop="baoXiaoNeiRong">
          <el-input v-model="formData.baoXiaoNeiRong" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="报销人员:" prop="baoXiaoRenYuan">
          <el-input v-model.number="formData.baoXiaoRenYuan" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="报销事项:" prop="baoXiaoShiXiang">
          <el-select v-model="formData.baoXiaoShiXiang" placeholder="请选择" :clearable="true">
            <el-option v-for="(item,key) in baoXiaoShiXiangOptions" :key="key" :label="item.label" :value="item.value" />
          </el-select>
        </el-form-item>
        <el-form-item label="备注说明:" prop="beiZhuShuoMing">
          <el-input v-model="formData.beiZhuShuoMing" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="费用发生日期:" prop="feiYongFaShengRiQi">
          <el-date-picker v-model="formData.feiYongFaShengRiQi" type="date" placeholder="选择日期" :clearable="true"></el-date-picker>
        </el-form-item>
        <el-form-item label="公司名称:" prop="gongSiMingCheng">
          <el-select v-model="formData.gongSiMingCheng" placeholder="请选择" :clearable="true">
            <el-option v-for="(item,key) in gongSiMingChengOptions" :key="key" :label="item.label" :value="item.value" />
          </el-select>
        </el-form-item>
        <el-form-item label="所属:" prop="group">
          <el-input v-model="formData.group" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="回单地址-可多图片上传:" prop="huiDanDiZhi">
          <el-input v-model="formData.huiDanDiZhi" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="类型 1. 内部报销 2. 项目经理报销:" prop="baoXiaoLeiXing">
          <el-select v-model="formData.baoXiaoLeiXing" placeholder="请选择" :clearable="true">
            <el-option v-for="(item,key) in baoXiaoLeiXingOptions" :key="key" :label="item.label" :value="item.value" />
          </el-select>
        </el-form-item>
        <el-form-item label="票据地址-可多图片上传:" prop="piaoJuDiZhi">
          <el-input v-model="formData.piaoJuDiZhi" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="申请日期:" prop="shenQingRiQi">
          <el-date-picker v-model="formData.shenQingRiQi" type="date" placeholder="选择日期" :clearable="true"></el-date-picker>
        </el-form-item>
        <el-form-item label="首次报销ID-再次申请时关联的ID:" prop="shouCiBaoXiaoId">
          <el-input v-model.number="formData.shouCiBaoXiaoId" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="项目名称:" prop="xiangMuMingCheng">
          <el-input v-model.number="formData.xiangMuMingCheng" :clearable="true" placeholder="请输入" />
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
  name: 'NeiBuBaoXiaoBiao'
}
</script>

<script setup>
import {
  createNeiBuBaoXiaoBiao,
  updateNeiBuBaoXiaoBiao,
  findNeiBuBaoXiaoBiao
} from '@/api/neiBuBaoXiaoBiao'

// 自动获取字典
import { getDictFunc } from '@/utils/format'
import { useRoute, useRouter } from "vue-router"
import { ElMessage } from 'element-plus'
import { ref, reactive } from 'vue'
const route = useRoute()
const router = useRouter()

const type = ref('')
const baoXiaoShiXiangOptions = ref([])
const gongSiMingChengOptions = ref([])
const baoXiaoLeiXingOptions = ref([])
const formData = ref({
            baoXiaoBuMen: undefined,
            baoXiaoJinE: 0,
            baoXiaoNeiRong: '',
            baoXiaoRenYuan: 0,
            baoXiaoShiXiang: undefined,
            beiZhuShuoMing: '',
            feiYongFaShengRiQi: new Date(),
            gongSiMingCheng: undefined,
            group: '',
            huiDanDiZhi: '',
            baoXiaoLeiXing: undefined,
            piaoJuDiZhi: '',
            shenQingRiQi: new Date(),
            shouCiBaoXiaoId: 0,
            xiangMuMingCheng: 0,
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
      const res = await findNeiBuBaoXiaoBiao({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data.reneiBuBaoXiaoBiao
        type.value = 'update'
      }
    } else {
      type.value = 'create'
    }
    baoXiaoShiXiangOptions.value = await getDictFunc('baoXiaoShiXiang')
    gongSiMingChengOptions.value = await getDictFunc('gongSiMingCheng')
    baoXiaoLeiXingOptions.value = await getDictFunc('baoXiaoLeiXing')
}

init()
// 保存按钮
const save = async() => {
      elFormRef.value?.validate( async (valid) => {
         if (!valid) return
            let res
           switch (type.value) {
             case 'create':
               res = await createNeiBuBaoXiaoBiao(formData.value)
               break
             case 'update':
               res = await updateNeiBuBaoXiaoBiao(formData.value)
               break
             default:
               res = await createNeiBuBaoXiaoBiao(formData.value)
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
