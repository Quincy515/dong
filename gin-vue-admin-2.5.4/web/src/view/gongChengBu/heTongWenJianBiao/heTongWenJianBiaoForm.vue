<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="备注:" prop="beiZhu">
          <el-input v-model="formData.beiZhu" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="发票类型:" prop="faPiaoLeiXing">
          <el-select v-model="formData.faPiaoLeiXing" placeholder="请选择" :clearable="true">
            <el-option v-for="(item,key) in faPiaoLeiXingOptions" :key="key" :label="item.label" :value="item.value" />
          </el-select>
        </el-form-item>
        <el-form-item label="付款方式:" prop="fuKuanFangShi">
          <el-select v-model="formData.fuKuanFangShi" placeholder="请选择" :clearable="true">
            <el-option v-for="(item,key) in fuKuanFangShiOptions" :key="key" :label="item.label" :value="item.value" />
          </el-select>
        </el-form-item>
        <el-form-item label="所属:" prop="group">
          <el-input v-model="formData.group" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="合同编号:" prop="heTongBianHao">
          <el-input v-model="formData.heTongBianHao" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="合同金额:" prop="heTongJinE">
          <el-input-number v-model="formData.heTongJinE" :precision="2" :clearable="true"></el-input-number>
        </el-form-item>
        <el-form-item label="合同类别 1.主合同 2.补充合同:" prop="heTongLeiBie">
          <el-select v-model="formData.heTongLeiBie" placeholder="请选择" :clearable="true">
            <el-option v-for="(item,key) in heTongLeiBieOptions" :key="key" :label="item.label" :value="item.value" />
          </el-select>
        </el-form-item>
        <el-form-item label="合同类型 1.承包 2.分包 3.采购 4.租聘:" prop="heTongLeiXing">
          <el-select v-model="formData.heTongLeiXing" placeholder="请选择" :clearable="true">
            <el-option v-for="(item,key) in heTongLeiXingOptions" :key="key" :label="item.label" :value="item.value" />
          </el-select>
        </el-form-item>
        <el-form-item label="合同名称:" prop="heTongMingCheng">
          <el-input v-model="formData.heTongMingCheng" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="合同下载地址:" prop="heTongXiaZaiDiZhi">
          <el-input v-model="formData.heTongXiaZaiDiZhi" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="合同状态:" prop="heTongZhuangTai">
          <el-select v-model="formData.heTongZhuangTai" placeholder="请选择" :clearable="true">
            <el-option v-for="(item,key) in heTongZhuangTaiOptions" :key="key" :label="item.label" :value="item.value" />
          </el-select>
        </el-form-item>
        <el-form-item label="甲方公司:" prop="jiaFangGongSi">
          <el-input v-model.number="formData.jiaFangGongSi" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="结算方式:" prop="jieSuanFangShi">
          <el-select v-model="formData.jieSuanFangShi" placeholder="请选择" :clearable="true">
            <el-option v-for="(item,key) in jieSuanFangShiOptions" :key="key" :label="item.label" :value="item.value" />
          </el-select>
        </el-form-item>
        <el-form-item label="开户行:" prop="kaiHuHang">
          <el-input v-model="formData.kaiHuHang" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="开票单位:" prop="kaiPiaoDanWei">
          <el-input v-model="formData.kaiPiaoDanWei" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="签订人:" prop="qianDingRen">
          <el-input v-model.number="formData.qianDingRen" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="签订日期:" prop="qianDingRiQi">
          <el-date-picker v-model="formData.qianDingRiQi" type="date" placeholder="选择日期" :clearable="true"></el-date-picker>
        </el-form-item>
        <el-form-item label="收付款条件:" prop="shouFuKuanTiaoJian">
          <el-input v-model="formData.shouFuKuanTiaoJian" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="收款账号:" prop="shouKuanZhangHao">
          <el-input v-model="formData.shouKuanZhangHao" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="添加人:" prop="tianJiaRen">
          <el-input v-model.number="formData.tianJiaRen" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="添加时间:" prop="tianJiaShiJian">
          <el-date-picker v-model="formData.tianJiaShiJian" type="date" placeholder="选择日期" :clearable="true"></el-date-picker>
        </el-form-item>
        <el-form-item label="项目ID:" prop="xiangMuId">
          <el-input v-model.number="formData.xiangMuId" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="乙方公司:" prop="yiFangGongSi">
          <el-input v-model.number="formData.yiFangGongSi" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="银行名称:" prop="yinHangMingCheng">
          <el-input v-model="formData.yinHangMingCheng" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="主要条款:" prop="zhuYaoTiaoKuan">
          <el-input v-model="formData.zhuYaoTiaoKuan" :clearable="true" placeholder="请输入" />
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
  name: 'HeTongWenJianBiao'
}
</script>

<script setup>
import {
  createHeTongWenJianBiao,
  updateHeTongWenJianBiao,
  findHeTongWenJianBiao
} from '@/api/heTongWenJianBiao'

// 自动获取字典
import { getDictFunc } from '@/utils/format'
import { useRoute, useRouter } from "vue-router"
import { ElMessage } from 'element-plus'
import { ref, reactive } from 'vue'
const route = useRoute()
const router = useRouter()

const type = ref('')
const faPiaoLeiXingOptions = ref([])
const fuKuanFangShiOptions = ref([])
const heTongLeiBieOptions = ref([])
const heTongLeiXingOptions = ref([])
const heTongZhuangTaiOptions = ref([])
const jieSuanFangShiOptions = ref([])
const formData = ref({
            beiZhu: '',
            faPiaoLeiXing: undefined,
            fuKuanFangShi: undefined,
            group: '',
            heTongBianHao: '',
            heTongJinE: 0,
            heTongLeiBie: undefined,
            heTongLeiXing: undefined,
            heTongMingCheng: '',
            heTongXiaZaiDiZhi: '',
            heTongZhuangTai: undefined,
            jiaFangGongSi: 0,
            jieSuanFangShi: undefined,
            kaiHuHang: '',
            kaiPiaoDanWei: '',
            qianDingRen: 0,
            qianDingRiQi: new Date(),
            shouFuKuanTiaoJian: '',
            shouKuanZhangHao: '',
            tianJiaRen: 0,
            tianJiaShiJian: new Date(),
            xiangMuId: 0,
            yiFangGongSi: 0,
            yinHangMingCheng: '',
            zhuYaoTiaoKuan: '',
        })
// 验证规则
const rule = reactive({
})

const elFormRef = ref()

// 初始化方法
const init = async () => {
 // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (route.query.id) {
      const res = await findHeTongWenJianBiao({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data.reheTongWenJianBiao
        type.value = 'update'
      }
    } else {
      type.value = 'create'
    }
    faPiaoLeiXingOptions.value = await getDictFunc('faPiaoLeiXing')
    fuKuanFangShiOptions.value = await getDictFunc('fuKuanFangShi')
    heTongLeiBieOptions.value = await getDictFunc('heTongLeiBie')
    heTongLeiXingOptions.value = await getDictFunc('heTongLeiXing')
    heTongZhuangTaiOptions.value = await getDictFunc('heTongZhuangTai')
    jieSuanFangShiOptions.value = await getDictFunc('jieSuanFangShi')
}

init()
// 保存按钮
const save = async() => {
      elFormRef.value?.validate( async (valid) => {
         if (!valid) return
            let res
           switch (type.value) {
             case 'create':
               res = await createHeTongWenJianBiao(formData.value)
               break
             case 'update':
               res = await updateHeTongWenJianBiao(formData.value)
               break
             default:
               res = await createHeTongWenJianBiao(formData.value)
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
