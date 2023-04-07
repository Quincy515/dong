<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="备注说明:" prop="beiZhuShuoMing">
          <el-input v-model="formData.beiZhuShuoMing" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="垫资款:" prop="dianZiKuan">
          <el-input-number v-model="formData.dianZiKuan" :precision="2" :clearable="true"></el-input-number>
        </el-form-item>
        <el-form-item label="发票类型:" prop="faPiaoLeiXing">
          <el-select v-model="formData.faPiaoLeiXing" placeholder="请选择" :clearable="true">
            <el-option v-for="(item,key) in faPiaoLeiXingOptions" :key="key" :label="item.label" :value="item.value" />
          </el-select>
        </el-form-item>
        <el-form-item label="发起人员:" prop="faQiRenYuan">
          <el-input v-model.number="formData.faQiRenYuan" :clearable="true" placeholder="请输入" />
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
        <el-form-item label="回单地址:" prop="huiDanDiZhi">
          <el-input v-model="formData.huiDanDiZhi" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="开户行:" prop="kaiHuHang">
          <el-input v-model="formData.kaiHuHang" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="开票单位:" prop="kaiPiaoDanWei">
          <el-input v-model="formData.kaiPiaoDanWei" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="开票金额:" prop="kaiPiaoJinE">
          <el-input-number v-model="formData.kaiPiaoJinE" :precision="2" :clearable="true"></el-input-number>
        </el-form-item>
        <el-form-item label="履约保证金:" prop="luYueBaoZhengJin">
          <el-input-number v-model="formData.luYueBaoZhengJin" :precision="2" :clearable="true"></el-input-number>
        </el-form-item>
        <el-form-item label="申请日期:" prop="shenQingRiQi">
          <el-date-picker v-model="formData.shenQingRiQi" type="date" placeholder="选择日期" :clearable="true"></el-date-picker>
        </el-form-item>
        <el-form-item label="实付金额:" prop="shiFuJinE">
          <el-input-number v-model="formData.shiFuJinE" :precision="2" :clearable="true"></el-input-number>
        </el-form-item>
        <el-form-item label="实收金额:" prop="shiShouJinE">
          <el-input-number v-model="formData.shiShouJinE" :precision="2" :clearable="true"></el-input-number>
        </el-form-item>
        <el-form-item label="首次收款ID:" prop="shouCiShouKuanId">
          <el-input v-model.number="formData.shouCiShouKuanId" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="收款名称:" prop="shouKuanMingCheng">
          <el-input v-model="formData.shouKuanMingCheng" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="收款账号:" prop="shouKuanZhangHao">
          <el-input v-model="formData.shouKuanZhangHao" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="投标保证金:" prop="touBiaoBaoZhengJin">
          <el-input-number v-model="formData.touBiaoBaoZhengJin" :precision="2" :clearable="true"></el-input-number>
        </el-form-item>
        <el-form-item label="项目名称:" prop="xiangMuMingCheng">
          <el-input v-model.number="formData.xiangMuMingCheng" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="银行名称:" prop="yinHangMingCheng">
          <el-input v-model="formData.yinHangMingCheng" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="应收金额:" prop="yingShouJinE">
          <el-input-number v-model="formData.yingShouJinE" :precision="2" :clearable="true"></el-input-number>
        </el-form-item>
        <el-form-item label="应收类型:" prop="yingShouLeiXing">
          <el-select v-model="formData.yingShouLeiXing" placeholder="请选择" :clearable="true">
            <el-option v-for="(item,key) in yingShouLeiXingOptions" :key="key" :label="item.label" :value="item.value" />
          </el-select>
        </el-form-item>
        <el-form-item label="应收事项:" prop="yingShouShiXiang">
          <el-select v-model="formData.yingShouShiXiang" placeholder="请选择" :clearable="true">
            <el-option v-for="(item,key) in yingShouShiXiangOptions" :key="key" :label="item.label" :value="item.value" />
          </el-select>
        </el-form-item>
        <el-form-item label="预计到账日期:" prop="yuJiDaoZhangRiQi">
          <el-date-picker v-model="formData.yuJiDaoZhangRiQi" type="date" placeholder="选择日期" :clearable="true"></el-date-picker>
        </el-form-item>
        <el-form-item label="质保金:" prop="zhiBaoJin">
          <el-input-number v-model="formData.zhiBaoJin" :precision="2" :clearable="true"></el-input-number>
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
  name: 'YingShouKuanBiao'
}
</script>

<script setup>
import {
  createYingShouKuanBiao,
  updateYingShouKuanBiao,
  findYingShouKuanBiao
} from '@/api/yingShouKuanBiao'

// 自动获取字典
import { getDictFunc } from '@/utils/format'
import { useRoute, useRouter } from "vue-router"
import { ElMessage } from 'element-plus'
import { ref, reactive } from 'vue'
const route = useRoute()
const router = useRouter()

const type = ref('')
const faPiaoLeiXingOptions = ref([])
const yingShouLeiXingOptions = ref([])
const yingShouShiXiangOptions = ref([])
const formData = ref({
            beiZhuShuoMing: '',
            dianZiKuan: 0,
            faPiaoLeiXing: undefined,
            faQiRenYuan: 0,
            feiYongFaShengRiQi: new Date(),
            group: '',
            heTongId: 0,
            huiDanDiZhi: '',
            kaiHuHang: '',
            kaiPiaoDanWei: '',
            kaiPiaoJinE: 0,
            luYueBaoZhengJin: 0,
            shenQingRiQi: new Date(),
            shiFuJinE: 0,
            shiShouJinE: 0,
            shouCiShouKuanId: 0,
            shouKuanMingCheng: '',
            shouKuanZhangHao: '',
            touBiaoBaoZhengJin: 0,
            xiangMuMingCheng: 0,
            yinHangMingCheng: '',
            yingShouJinE: 0,
            yingShouLeiXing: undefined,
            yingShouShiXiang: undefined,
            yuJiDaoZhangRiQi: new Date(),
            zhiBaoJin: 0,
            zhuangTai: 0,
        })
// 验证规则
const rule = reactive({
               shouKuanZhangHao : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
})

const elFormRef = ref()

// 初始化方法
const init = async () => {
 // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (route.query.id) {
      const res = await findYingShouKuanBiao({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data.reyingShouKuanBiao
        type.value = 'update'
      }
    } else {
      type.value = 'create'
    }
    faPiaoLeiXingOptions.value = await getDictFunc('faPiaoLeiXing')
    yingShouLeiXingOptions.value = await getDictFunc('yingShouLeiXing')
    yingShouShiXiangOptions.value = await getDictFunc('yingShouShiXiang')
}

init()
// 保存按钮
const save = async() => {
      elFormRef.value?.validate( async (valid) => {
         if (!valid) return
            let res
           switch (type.value) {
             case 'create':
               res = await createYingShouKuanBiao(formData.value)
               break
             case 'update':
               res = await updateYingShouKuanBiao(formData.value)
               break
             default:
               res = await createYingShouKuanBiao(formData.value)
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
