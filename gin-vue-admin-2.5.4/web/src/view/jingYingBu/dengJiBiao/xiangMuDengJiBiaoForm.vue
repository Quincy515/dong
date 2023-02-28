<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="所属:" prop="group">
          <el-input v-model="formData.group" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="项目名称:" prop="xiangMuMingCheng">
          <el-input v-model="formData.xiangMuMingCheng" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="项目简介:" prop="xiangMuJianJie">
          <el-input v-model="formData.xiangMuJianJie" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="项目登记进度:" prop="xiangMuDengJiJinDu">
          <el-input v-model="formData.xiangMuDengJiJinDu" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="投标单位名称:" prop="touBiaoDanWeiMingCheng">
          <el-input v-model="formData.touBiaoDanWeiMingCheng" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="资质要求:" prop="ziZhiYaoQiu">
          <el-select v-model="formData.ziZhiYaoQiu" placeholder="请选择" :clearable="true">
            <el-option v-for="(item,key) in zi_zhi_yao_qiuOptions" :key="key" :label="item.label" :value="item.value" />
          </el-select>
        </el-form-item>
        <el-form-item label="等级要求:" prop="dengJiYaoQiu">
          <el-select v-model="formData.dengJiYaoQiu" placeholder="请选择" :clearable="true">
            <el-option v-for="(item,key) in deng_ji_yao_qiuOptions" :key="key" :label="item.label" :value="item.value" />
          </el-select>
        </el-form-item>
        <el-form-item label="开标时间:" prop="kaiBiaoShiJian">
          <el-date-picker v-model="formData.kaiBiaoShiJian" type="date" placeholder="选择日期" :clearable="true"></el-date-picker>
        </el-form-item>
        <el-form-item label="开标地点:" prop="kaiBiaoDiDian">
          <el-input v-model="formData.kaiBiaoDiDian" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="所属地区:" prop="suoShuDiQu">
          <el-select v-model="formData.suoShuDiQu" placeholder="请选择" :clearable="true">
            <el-option v-for="(item,key) in suo_shu_di_quOptions" :key="key" :label="item.label" :value="item.value" />
          </el-select>
        </el-form-item>
        <el-form-item label="招标规模:" prop="zhaoBiaoGuiMo">
          <el-input v-model="formData.zhaoBiaoGuiMo" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="开标人员:" prop="kaiBiaoRenYuan">
          <el-input v-model="formData.kaiBiaoRenYuan" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="业务负责人:" prop="yeWuFuZeRen">
          <el-input v-model="formData.yeWuFuZeRen" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="保证金形式:" prop="baoZhengJinXingShi">
          <el-select v-model="formData.baoZhengJinXingShi" placeholder="请选择" :clearable="true">
            <el-option v-for="(item,key) in bao_zheng_jin_xing_shiOptions" :key="key" :label="item.label" :value="item.value" />
          </el-select>
        </el-form-item>
        <el-form-item label="保证金金额:" prop="baoZhengJinJinE">
          <el-input v-model="formData.baoZhengJinJinE" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="评标办法:" prop="pingBiaoBanFa">
          <el-select v-model="formData.pingBiaoBanFa" placeholder="请选择" :clearable="true">
            <el-option v-for="(item,key) in ping_biao_ban_faOptions" :key="key" :label="item.label" :value="item.value" />
          </el-select>
        </el-form-item>
        <el-form-item label="项目网址:" prop="xiangMuWangZhi">
          <el-input v-model="formData.xiangMuWangZhi" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="标书类型:" prop="biaoShuLeiXing">
          <el-select v-model="formData.biaoShuLeiXing" placeholder="请选择" :clearable="true">
            <el-option v-for="(item,key) in biao_shu_lei_xingOptions" :key="key" :label="item.label" :value="item.value" />
          </el-select>
        </el-form-item>
        <el-form-item label="项目登记时间:" prop="xiangMuDengJiShiJian">
          <el-date-picker v-model="formData.xiangMuDengJiShiJian" type="date" placeholder="选择日期" :clearable="true"></el-date-picker>
        </el-form-item>
        <el-form-item label="部门审核:" prop="buMenShenHe">
          <el-input v-model="formData.buMenShenHe" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="部门审核时间:" prop="buMenShenHeShiJian">
          <el-date-picker v-model="formData.buMenShenHeShiJian" type="date" placeholder="选择日期" :clearable="true"></el-date-picker>
        </el-form-item>
        <el-form-item label="CEO审核:" prop="ceOshenHe">
          <el-input v-model="formData.ceOshenHe" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="CEO审核时间:" prop="ceOshenHeShiJian">
          <el-date-picker v-model="formData.ceOshenHeShiJian" type="date" placeholder="选择日期" :clearable="true"></el-date-picker>
        </el-form-item>
        <el-form-item label="投标函制作人:" prop="touBiaoHanZhiZuoRen">
          <el-input v-model="formData.touBiaoHanZhiZuoRen" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="投标函提成金额:" prop="touBiaoHanTiChengJinE">
          <el-input v-model="formData.touBiaoHanTiChengJinE" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="投标函制作人确认:" prop="touBiaoHanZhiZuoRenQueRen">
          <el-input v-model="formData.touBiaoHanZhiZuoRenQueRen" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="投标函制作人确认时间:" prop="touBiaoHanZhiZuoRenQueRenShiJian">
          <el-date-picker v-model="formData.touBiaoHanZhiZuoRenQueRenShiJian" type="date" placeholder="选择日期" :clearable="true"></el-date-picker>
        </el-form-item>
        <el-form-item label="商务标制作人:" prop="shangWuBiaoZhiZuoRen">
          <el-input v-model="formData.shangWuBiaoZhiZuoRen" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="商务标提成金额:" prop="shangWuBiaoTiChengJinE">
          <el-input v-model="formData.shangWuBiaoTiChengJinE" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="商务标制作人确认:" prop="shangWuBiaoZhiZuoRenQueRen">
          <el-input v-model="formData.shangWuBiaoZhiZuoRenQueRen" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="商务标制作人确认时间:" prop="shangWuBiaoZhiZuoRenQueRenShiJian">
          <el-date-picker v-model="formData.shangWuBiaoZhiZuoRenQueRenShiJian" type="date" placeholder="选择日期" :clearable="true"></el-date-picker>
        </el-form-item>
        <el-form-item label="技术标制作人:" prop="jiShuBiaoZhiZuoRen">
          <el-input v-model="formData.jiShuBiaoZhiZuoRen" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="技术标提成金额:" prop="jiShuBiaoTiChengJinE">
          <el-input v-model="formData.jiShuBiaoTiChengJinE" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="技术标制作人确认:" prop="jiShuBiaoZhiZuoRenQueRen">
          <el-input v-model="formData.jiShuBiaoZhiZuoRenQueRen" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="技术标制作人确认时间:" prop="jiShuBiaoZhiZuoRenQueRenShiJian">
          <el-date-picker v-model="formData.jiShuBiaoZhiZuoRenQueRenShiJian" type="date" placeholder="选择日期" :clearable="true"></el-date-picker>
        </el-form-item>
        <el-form-item label="项目状态:" prop="xiangMuZhuangTai">
          <el-input v-model="formData.xiangMuZhuangTai" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="部门申请下浮点位:" prop="buMenShenQingXiaFuDianWei">
          <el-input v-model="formData.buMenShenQingXiaFuDianWei" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="部门下浮点位申请时间:" prop="buMenXiaFuDianWeiShenQingShiJian">
          <el-date-picker v-model="formData.buMenXiaFuDianWeiShenQingShiJian" type="date" placeholder="选择日期" :clearable="true"></el-date-picker>
        </el-form-item>
        <el-form-item label="CEO审核下浮点位:" prop="ceOshenHeXiaFuDianWei">
          <el-input v-model="formData.ceOshenHeXiaFuDianWei" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="CEO下浮点位审核时间:" prop="ceOxiaFuDianWeiShenHeShiJian">
          <el-date-picker v-model="formData.ceOxiaFuDianWeiShenHeShiJian" type="date" placeholder="选择日期" :clearable="true"></el-date-picker>
        </el-form-item>
        <el-form-item label="中标金额:" prop="zhongBiaoJinE">
          <el-input v-model="formData.zhongBiaoJinE" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="备注:" prop="beiZhu">
          <el-input v-model="formData.beiZhu" :clearable="true" placeholder="请输入" />
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
  name: 'XiangMuDengJiBiao'
}
</script>

<script setup>
import {
  createXiangMuDengJiBiao,
  updateXiangMuDengJiBiao,
  findXiangMuDengJiBiao
} from '@/api/xiangMuDengJiBiao'

// 自动获取字典
import { getDictFunc } from '@/utils/format'
import { useRoute, useRouter } from "vue-router"
import { ElMessage } from 'element-plus'
import { ref, reactive } from 'vue'
const route = useRoute()
const router = useRouter()

const type = ref('')
const zi_zhi_yao_qiuOptions = ref([])
const deng_ji_yao_qiuOptions = ref([])
const suo_shu_di_quOptions = ref([])
const bao_zheng_jin_xing_shiOptions = ref([])
const ping_biao_ban_faOptions = ref([])
const biao_shu_lei_xingOptions = ref([])
const formData = ref({
            group: '',
            xiangMuMingCheng: '',
            xiangMuJianJie: '',
            xiangMuDengJiJinDu: '',
            touBiaoDanWeiMingCheng: '',
            ziZhiYaoQiu: undefined,
            dengJiYaoQiu: undefined,
            kaiBiaoShiJian: new Date(),
            kaiBiaoDiDian: '',
            suoShuDiQu: undefined,
            zhaoBiaoGuiMo: '',
            kaiBiaoRenYuan: '',
            yeWuFuZeRen: '',
            baoZhengJinXingShi: undefined,
            baoZhengJinJinE: '',
            pingBiaoBanFa: undefined,
            xiangMuWangZhi: '',
            biaoShuLeiXing: undefined,
            xiangMuDengJiShiJian: new Date(),
            buMenShenHe: '',
            buMenShenHeShiJian: new Date(),
            ceOshenHe: '',
            ceOshenHeShiJian: new Date(),
            touBiaoHanZhiZuoRen: '',
            touBiaoHanTiChengJinE: '',
            touBiaoHanZhiZuoRenQueRen: '',
            touBiaoHanZhiZuoRenQueRenShiJian: new Date(),
            shangWuBiaoZhiZuoRen: '',
            shangWuBiaoTiChengJinE: '',
            shangWuBiaoZhiZuoRenQueRen: '',
            shangWuBiaoZhiZuoRenQueRenShiJian: new Date(),
            jiShuBiaoZhiZuoRen: '',
            jiShuBiaoTiChengJinE: '',
            jiShuBiaoZhiZuoRenQueRen: '',
            jiShuBiaoZhiZuoRenQueRenShiJian: new Date(),
            xiangMuZhuangTai: '',
            buMenShenQingXiaFuDianWei: '',
            buMenXiaFuDianWeiShenQingShiJian: new Date(),
            ceOshenHeXiaFuDianWei: '',
            ceOxiaFuDianWeiShenHeShiJian: new Date(),
            zhongBiaoJinE: '',
            beiZhu: '',
        })
// 验证规则
const rule = reactive({
})

const elFormRef = ref()

// 初始化方法
const init = async () => {
 // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (route.query.id) {
      const res = await findXiangMuDengJiBiao({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data.rexiangMuDengJiBiao
        type.value = 'update'
      }
    } else {
      type.value = 'create'
    }
    zi_zhi_yao_qiuOptions.value = await getDictFunc('zi_zhi_yao_qiu')
    deng_ji_yao_qiuOptions.value = await getDictFunc('deng_ji_yao_qiu')
    suo_shu_di_quOptions.value = await getDictFunc('suo_shu_di_qu')
    bao_zheng_jin_xing_shiOptions.value = await getDictFunc('bao_zheng_jin_xing_shi')
    ping_biao_ban_faOptions.value = await getDictFunc('ping_biao_ban_fa')
    biao_shu_lei_xingOptions.value = await getDictFunc('biao_shu_lei_xing')
}

init()
// 保存按钮
const save = async() => {
      elFormRef.value?.validate( async (valid) => {
         if (!valid) return
            let res
           switch (type.value) {
             case 'create':
               res = await createXiangMuDengJiBiao(formData.value)
               break
             case 'update':
               res = await updateXiangMuDengJiBiao(formData.value)
               break
             default:
               res = await createXiangMuDengJiBiao(formData.value)
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
