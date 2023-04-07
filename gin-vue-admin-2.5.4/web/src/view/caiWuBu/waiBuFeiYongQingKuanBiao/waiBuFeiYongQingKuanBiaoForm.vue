<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="报销金额:" prop="baoXiaoJinE">
          <el-input-number v-model="formData.baoXiaoJinE" :precision="2" :clearable="true"></el-input-number>
        </el-form-item>
        <el-form-item label="报销类型:" prop="baoXiaoLeiXing">
          <el-select v-model="formData.baoXiaoLeiXing" placeholder="请选择" :clearable="true">
            <el-option v-for="(item,key) in baoXiaoLeiXingOptions" :key="key" :label="item.label" :value="item.value" />
          </el-select>
        </el-form-item>
        <el-form-item label="报销名称:" prop="baoXiaoMingCheng">
          <el-input v-model="formData.baoXiaoMingCheng" :clearable="true" placeholder="请输入" />
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
        <el-form-item label="所属:" prop="group">
          <el-input v-model="formData.group" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="合同ID:" prop="heTongId">
          <el-input v-model.number="formData.heTongId" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="回单地址-可多图片上传NULL:" prop="huiDanDiZhi">
          <el-input v-model="formData.huiDanDiZhi" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="开户行:" prop="kaiHuHang">
          <el-input v-model="formData.kaiHuHang" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="申请日期:" prop="shenQingRiQi">
          <el-date-picker v-model="formData.shenQingRiQi" type="date" placeholder="选择日期" :clearable="true"></el-date-picker>
        </el-form-item>
        <el-form-item label="实付金额:" prop="shiFuJinE">
          <el-input-number v-model="formData.shiFuJinE" :precision="2" :clearable="true"></el-input-number>
        </el-form-item>
        <el-form-item label="首次报销ID:" prop="shouCiBaoXiaoId">
          <el-input v-model.number="formData.shouCiBaoXiaoId" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="收款账号:" prop="shouKuanZhangHao">
          <el-input v-model="formData.shouKuanZhangHao" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="项目名称:" prop="xiangMuMingCheng">
          <el-input v-model.number="formData.xiangMuMingCheng" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="银行名称:" prop="yinHangMingCheng">
          <el-input v-model="formData.yinHangMingCheng" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="状态:" prop="zhuangTai">
          <el-input v-model.number="formData.zhuangTai" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="最晚付款日期:" prop="zuiWanFuKuanRiQi">
          <el-date-picker v-model="formData.zuiWanFuKuanRiQi" type="date" placeholder="选择日期" :clearable="true"></el-date-picker>
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
  name: 'WaiBuFeiYongQingKuanBiao'
}
</script>

<script setup>
import {
  createWaiBuFeiYongQingKuanBiao,
  updateWaiBuFeiYongQingKuanBiao,
  findWaiBuFeiYongQingKuanBiao
} from '@/api/waiBuFeiYongQingKuanBiao'

// 自动获取字典
import { getDictFunc } from '@/utils/format'
import { useRoute, useRouter } from "vue-router"
import { ElMessage } from 'element-plus'
import { ref, reactive } from 'vue'
const route = useRoute()
const router = useRouter()

const type = ref('')
const baoXiaoLeiXingOptions = ref([])
const baoXiaoShiXiangOptions = ref([])
const formData = ref({
            baoXiaoJinE: 0,
            baoXiaoLeiXing: undefined,
            baoXiaoMingCheng: '',
            baoXiaoNeiRong: '',
            baoXiaoRenYuan: 0,
            baoXiaoShiXiang: undefined,
            beiZhuShuoMing: '',
            feiYongFaShengRiQi: new Date(),
            group: '',
            heTongId: 0,
            huiDanDiZhi: '',
            kaiHuHang: '',
            shenQingRiQi: new Date(),
            shiFuJinE: 0,
            shouCiBaoXiaoId: 0,
            shouKuanZhangHao: '',
            xiangMuMingCheng: 0,
            yinHangMingCheng: '',
            zhuangTai: 0,
            zuiWanFuKuanRiQi: new Date(),
        })
// 验证规则
const rule = reactive({
})

const elFormRef = ref()

// 初始化方法
const init = async () => {
 // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (route.query.id) {
      const res = await findWaiBuFeiYongQingKuanBiao({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data.rewaiBuFeiYongQingKuanBiao
        type.value = 'update'
      }
    } else {
      type.value = 'create'
    }
    baoXiaoLeiXingOptions.value = await getDictFunc('baoXiaoLeiXing')
    baoXiaoShiXiangOptions.value = await getDictFunc('baoXiaoShiXiang')
}

init()
// 保存按钮
const save = async() => {
      elFormRef.value?.validate( async (valid) => {
         if (!valid) return
            let res
           switch (type.value) {
             case 'create':
               res = await createWaiBuFeiYongQingKuanBiao(formData.value)
               break
             case 'update':
               res = await updateWaiBuFeiYongQingKuanBiao(formData.value)
               break
             default:
               res = await createWaiBuFeiYongQingKuanBiao(formData.value)
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
