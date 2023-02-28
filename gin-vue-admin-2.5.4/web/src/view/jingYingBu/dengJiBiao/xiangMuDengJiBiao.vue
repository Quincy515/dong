<template>
  <div>
    <div class="gva-search-box">
      <el-form :inline="true" :model="searchInfo" class="demo-form-inline">
      <el-form-item label="创建时间">
      <el-date-picker v-model="searchInfo.startCreatedAt" type="datetime" placeholder="开始时间"></el-date-picker>
       —
      <el-date-picker v-model="searchInfo.endCreatedAt" type="datetime" placeholder="结束时间"></el-date-picker>
      </el-form-item>
        <el-form-item label="所属">
         <el-input v-model="searchInfo.group" placeholder="搜索条件" />

        </el-form-item>
        <el-form-item label="项目名称">
         <el-input v-model="searchInfo.xiangMuMingCheng" placeholder="搜索条件" />

        </el-form-item>
        <el-form-item label="项目简介">
         <el-input v-model="searchInfo.xiangMuJianJie" placeholder="搜索条件" />

        </el-form-item>
        <el-form-item label="项目登记进度">
         <el-input v-model="searchInfo.xiangMuDengJiJinDu" placeholder="搜索条件" />

        </el-form-item>
        <el-form-item label="投标单位名称">
         <el-input v-model="searchInfo.touBiaoDanWeiMingCheng" placeholder="搜索条件" />

        </el-form-item>
           <el-form-item label="资质要求" prop="ziZhiYaoQiu">
            <el-select v-model="searchInfo.ziZhiYaoQiu" clearable placeholder="请选择" @clear="()=>{searchInfo.ziZhiYaoQiu=undefined}">
              <el-option v-for="(item,key) in zi_zhi_yao_qiuOptions" :key="key" :label="item.label" :value="item.value" />
            </el-select>
            </el-form-item>
           <el-form-item label="等级要求" prop="dengJiYaoQiu">
            <el-select v-model="searchInfo.dengJiYaoQiu" clearable placeholder="请选择" @clear="()=>{searchInfo.dengJiYaoQiu=undefined}">
              <el-option v-for="(item,key) in deng_ji_yao_qiuOptions" :key="key" :label="item.label" :value="item.value" />
            </el-select>
            </el-form-item>
           <el-form-item label="所属地区" prop="suoShuDiQu">
            <el-select v-model="searchInfo.suoShuDiQu" clearable placeholder="请选择" @clear="()=>{searchInfo.suoShuDiQu=undefined}">
              <el-option v-for="(item,key) in suo_shu_di_quOptions" :key="key" :label="item.label" :value="item.value" />
            </el-select>
            </el-form-item>
        <el-form-item label="业务负责人">
         <el-input v-model="searchInfo.yeWuFuZeRen" placeholder="搜索条件" />

        </el-form-item>
           <el-form-item label="保证金形式" prop="baoZhengJinXingShi">
            <el-select v-model="searchInfo.baoZhengJinXingShi" clearable placeholder="请选择" @clear="()=>{searchInfo.baoZhengJinXingShi=undefined}">
              <el-option v-for="(item,key) in bao_zheng_jin_xing_shiOptions" :key="key" :label="item.label" :value="item.value" />
            </el-select>
            </el-form-item>
           <el-form-item label="评标办法" prop="pingBiaoBanFa">
            <el-select v-model="searchInfo.pingBiaoBanFa" clearable placeholder="请选择" @clear="()=>{searchInfo.pingBiaoBanFa=undefined}">
              <el-option v-for="(item,key) in ping_biao_ban_faOptions" :key="key" :label="item.label" :value="item.value" />
            </el-select>
            </el-form-item>
           <el-form-item label="标书类型" prop="biaoShuLeiXing">
            <el-select v-model="searchInfo.biaoShuLeiXing" clearable placeholder="请选择" @clear="()=>{searchInfo.biaoShuLeiXing=undefined}">
              <el-option v-for="(item,key) in biao_shu_lei_xingOptions" :key="key" :label="item.label" :value="item.value" />
            </el-select>
            </el-form-item>
        <el-form-item label="项目状态">
         <el-input v-model="searchInfo.xiangMuZhuangTai" placeholder="搜索条件" />

        </el-form-item>
        <el-form-item>
          <el-button size="small" type="primary" icon="search" @click="onSubmit">查询</el-button>
          <el-button size="small" icon="refresh" @click="onReset">重置</el-button>
        </el-form-item>
      </el-form>
    </div>
    <div class="gva-table-box">
        <div class="gva-btn-list">
            <el-button size="small" type="primary" icon="plus" @click="openDialog">新增</el-button>
            <el-popover v-model:visible="deleteVisible" placement="top" width="160">
            <p>确定要删除吗？</p>
            <div style="text-align: right; margin-top: 8px;">
                <el-button size="small" type="primary" link @click="deleteVisible = false">取消</el-button>
                <el-button size="small" type="primary" @click="onDelete">确定</el-button>
            </div>
            <template #reference>
                <el-button icon="delete" size="small" style="margin-left: 10px;" :disabled="!multipleSelection.length" @click="deleteVisible = true">删除</el-button>
            </template>
            </el-popover>
        </div>
        <el-table
        ref="multipleTable"
        style="width: 100%"
        tooltip-effect="dark"
        :data="tableData"
        row-key="ID"
        @selection-change="handleSelectionChange"
        >
        <el-table-column type="selection" width="55" />
        <el-table-column align="left" label="日期" width="180">
            <template #default="scope">{{ formatDate(scope.row.CreatedAt) }}</template>
        </el-table-column>
        <el-table-column align="left" label="所属" prop="group" width="120" />
        <el-table-column align="left" label="项目名称" prop="xiangMuMingCheng" width="120" />
        <el-table-column align="left" label="项目简介" prop="xiangMuJianJie" width="120" />
        <el-table-column align="left" label="项目登记进度" prop="xiangMuDengJiJinDu" width="120" />
        <el-table-column align="left" label="投标单位名称" prop="touBiaoDanWeiMingCheng" width="120" />
        <el-table-column align="left" label="资质要求" prop="ziZhiYaoQiu" width="120">
            <template #default="scope">
            {{ filterDict(scope.row.ziZhiYaoQiu,zi_zhi_yao_qiuOptions) }}
            </template>
        </el-table-column>
        <el-table-column align="left" label="等级要求" prop="dengJiYaoQiu" width="120">
            <template #default="scope">
            {{ filterDict(scope.row.dengJiYaoQiu,deng_ji_yao_qiuOptions) }}
            </template>
        </el-table-column>
         <el-table-column align="left" label="开标时间" width="180">
            <template #default="scope">{{ formatDate(scope.row.kaiBiaoShiJian) }}</template>
         </el-table-column>
        <el-table-column align="left" label="开标地点" prop="kaiBiaoDiDian" width="120" />
        <el-table-column align="left" label="所属地区" prop="suoShuDiQu" width="120">
            <template #default="scope">
            {{ filterDict(scope.row.suoShuDiQu,suo_shu_di_quOptions) }}
            </template>
        </el-table-column>
        <el-table-column align="left" label="招标规模" prop="zhaoBiaoGuiMo" width="120" />
        <el-table-column align="left" label="开标人员" prop="kaiBiaoRenYuan" width="120" />
        <el-table-column align="left" label="业务负责人" prop="yeWuFuZeRen" width="120" />
        <el-table-column align="left" label="保证金形式" prop="baoZhengJinXingShi" width="120">
            <template #default="scope">
            {{ filterDict(scope.row.baoZhengJinXingShi,bao_zheng_jin_xing_shiOptions) }}
            </template>
        </el-table-column>
        <el-table-column align="left" label="保证金金额" prop="baoZhengJinJinE" width="120" />
        <el-table-column align="left" label="评标办法" prop="pingBiaoBanFa" width="120">
            <template #default="scope">
            {{ filterDict(scope.row.pingBiaoBanFa,ping_biao_ban_faOptions) }}
            </template>
        </el-table-column>
        <el-table-column align="left" label="项目网址" prop="xiangMuWangZhi" width="120" />
        <el-table-column align="left" label="标书类型" prop="biaoShuLeiXing" width="120">
            <template #default="scope">
            {{ filterDict(scope.row.biaoShuLeiXing,biao_shu_lei_xingOptions) }}
            </template>
        </el-table-column>
         <el-table-column align="left" label="项目登记时间" width="180">
            <template #default="scope">{{ formatDate(scope.row.xiangMuDengJiShiJian) }}</template>
         </el-table-column>
        <el-table-column align="left" label="部门审核" prop="buMenShenHe" width="120" />
         <el-table-column align="left" label="部门审核时间" width="180">
            <template #default="scope">{{ formatDate(scope.row.buMenShenHeShiJian) }}</template>
         </el-table-column>
        <el-table-column align="left" label="CEO审核" prop="ceOshenHe" width="120" />
         <el-table-column align="left" label="CEO审核时间" width="180">
            <template #default="scope">{{ formatDate(scope.row.ceOshenHeShiJian) }}</template>
         </el-table-column>
        <el-table-column align="left" label="投标函制作人" prop="touBiaoHanZhiZuoRen" width="120" />
        <el-table-column align="left" label="投标函提成金额" prop="touBiaoHanTiChengJinE" width="120" />
        <el-table-column align="left" label="投标函制作人确认" prop="touBiaoHanZhiZuoRenQueRen" width="120" />
         <el-table-column align="left" label="投标函制作人确认时间" width="180">
            <template #default="scope">{{ formatDate(scope.row.touBiaoHanZhiZuoRenQueRenShiJian) }}</template>
         </el-table-column>
        <el-table-column align="left" label="商务标制作人" prop="shangWuBiaoZhiZuoRen" width="120" />
        <el-table-column align="left" label="商务标提成金额" prop="shangWuBiaoTiChengJinE" width="120" />
        <el-table-column align="left" label="商务标制作人确认" prop="shangWuBiaoZhiZuoRenQueRen" width="120" />
         <el-table-column align="left" label="商务标制作人确认时间" width="180">
            <template #default="scope">{{ formatDate(scope.row.shangWuBiaoZhiZuoRenQueRenShiJian) }}</template>
         </el-table-column>
        <el-table-column align="left" label="技术标制作人" prop="jiShuBiaoZhiZuoRen" width="120" />
        <el-table-column align="left" label="技术标提成金额" prop="jiShuBiaoTiChengJinE" width="120" />
        <el-table-column align="left" label="技术标制作人确认" prop="jiShuBiaoZhiZuoRenQueRen" width="120" />
         <el-table-column align="left" label="技术标制作人确认时间" width="180">
            <template #default="scope">{{ formatDate(scope.row.jiShuBiaoZhiZuoRenQueRenShiJian) }}</template>
         </el-table-column>
        <el-table-column align="left" label="项目状态" prop="xiangMuZhuangTai" width="120" />
        <el-table-column align="left" label="部门申请下浮点位" prop="buMenShenQingXiaFuDianWei" width="120" />
         <el-table-column align="left" label="部门下浮点位申请时间" width="180">
            <template #default="scope">{{ formatDate(scope.row.buMenXiaFuDianWeiShenQingShiJian) }}</template>
         </el-table-column>
        <el-table-column align="left" label="CEO审核下浮点位" prop="ceOshenHeXiaFuDianWei" width="120" />
         <el-table-column align="left" label="CEO下浮点位审核时间" width="180">
            <template #default="scope">{{ formatDate(scope.row.ceOxiaFuDianWeiShenHeShiJian) }}</template>
         </el-table-column>
        <el-table-column align="left" label="中标金额" prop="zhongBiaoJinE" width="120" />
        <el-table-column align="left" label="备注" prop="beiZhu" width="120" />
        <el-table-column align="left" label="按钮组">
            <template #default="scope">
            <el-button type="primary" link icon="edit" size="small" class="table-button" @click="updateXiangMuDengJiBiaoFunc(scope.row)">变更</el-button>
            <el-button type="primary" link icon="delete" size="small" @click="deleteRow(scope.row)">删除</el-button>
            </template>
        </el-table-column>
        </el-table>
        <div class="gva-pagination">
            <el-pagination
            layout="total, sizes, prev, pager, next, jumper"
            :current-page="page"
            :page-size="pageSize"
            :page-sizes="[10, 30, 50, 100]"
            :total="total"
            @current-change="handleCurrentChange"
            @size-change="handleSizeChange"
            />
        </div>
    </div>
    <el-dialog v-model="dialogFormVisible" :before-close="closeDialog" title="弹窗操作">
      <el-form :model="formData" label-position="right" ref="elFormRef" :rules="rule" label-width="80px">
        <el-form-item label="所属:"  prop="group" >
          <el-input v-model="formData.group" :clearable="true"  placeholder="请输入" />
        </el-form-item>
        <el-form-item label="项目名称:"  prop="xiangMuMingCheng" >
          <el-input v-model="formData.xiangMuMingCheng" :clearable="true"  placeholder="请输入" />
        </el-form-item>
        <el-form-item label="项目简介:"  prop="xiangMuJianJie" >
          <el-input v-model="formData.xiangMuJianJie" :clearable="true"  placeholder="请输入" />
        </el-form-item>
        <el-form-item label="项目登记进度:"  prop="xiangMuDengJiJinDu" >
          <el-input v-model="formData.xiangMuDengJiJinDu" :clearable="true"  placeholder="请输入" />
        </el-form-item>
        <el-form-item label="投标单位名称:"  prop="touBiaoDanWeiMingCheng" >
          <el-input v-model="formData.touBiaoDanWeiMingCheng" :clearable="true"  placeholder="请输入" />
        </el-form-item>
        <el-form-item label="资质要求:"  prop="ziZhiYaoQiu" >
          <el-select v-model="formData.ziZhiYaoQiu" placeholder="请选择" style="width:100%" :clearable="true" >
            <el-option v-for="(item,key) in zi_zhi_yao_qiuOptions" :key="key" :label="item.label" :value="item.value" />
          </el-select>
        </el-form-item>
        <el-form-item label="等级要求:"  prop="dengJiYaoQiu" >
          <el-select v-model="formData.dengJiYaoQiu" placeholder="请选择" style="width:100%" :clearable="true" >
            <el-option v-for="(item,key) in deng_ji_yao_qiuOptions" :key="key" :label="item.label" :value="item.value" />
          </el-select>
        </el-form-item>
        <el-form-item label="开标时间:"  prop="kaiBiaoShiJian" >
          <el-date-picker v-model="formData.kaiBiaoShiJian" type="date" style="width:100%" placeholder="选择日期" :clearable="true"  />
        </el-form-item>
        <el-form-item label="开标地点:"  prop="kaiBiaoDiDian" >
          <el-input v-model="formData.kaiBiaoDiDian" :clearable="true"  placeholder="请输入" />
        </el-form-item>
        <el-form-item label="所属地区:"  prop="suoShuDiQu" >
          <el-select v-model="formData.suoShuDiQu" placeholder="请选择" style="width:100%" :clearable="true" >
            <el-option v-for="(item,key) in suo_shu_di_quOptions" :key="key" :label="item.label" :value="item.value" />
          </el-select>
        </el-form-item>
        <el-form-item label="招标规模:"  prop="zhaoBiaoGuiMo" >
          <el-input v-model="formData.zhaoBiaoGuiMo" :clearable="true"  placeholder="请输入" />
        </el-form-item>
        <el-form-item label="开标人员:"  prop="kaiBiaoRenYuan" >
          <el-input v-model="formData.kaiBiaoRenYuan" :clearable="true"  placeholder="请输入" />
        </el-form-item>
        <el-form-item label="业务负责人:"  prop="yeWuFuZeRen" >
          <el-input v-model="formData.yeWuFuZeRen" :clearable="true"  placeholder="请输入" />
        </el-form-item>
        <el-form-item label="保证金形式:"  prop="baoZhengJinXingShi" >
          <el-select v-model="formData.baoZhengJinXingShi" placeholder="请选择" style="width:100%" :clearable="true" >
            <el-option v-for="(item,key) in bao_zheng_jin_xing_shiOptions" :key="key" :label="item.label" :value="item.value" />
          </el-select>
        </el-form-item>
        <el-form-item label="保证金金额:"  prop="baoZhengJinJinE" >
          <el-input v-model="formData.baoZhengJinJinE" :clearable="true"  placeholder="请输入" />
        </el-form-item>
        <el-form-item label="评标办法:"  prop="pingBiaoBanFa" >
          <el-select v-model="formData.pingBiaoBanFa" placeholder="请选择" style="width:100%" :clearable="true" >
            <el-option v-for="(item,key) in ping_biao_ban_faOptions" :key="key" :label="item.label" :value="item.value" />
          </el-select>
        </el-form-item>
        <el-form-item label="项目网址:"  prop="xiangMuWangZhi" >
          <el-input v-model="formData.xiangMuWangZhi" :clearable="true"  placeholder="请输入" />
        </el-form-item>
        <el-form-item label="标书类型:"  prop="biaoShuLeiXing" >
          <el-select v-model="formData.biaoShuLeiXing" placeholder="请选择" style="width:100%" :clearable="true" >
            <el-option v-for="(item,key) in biao_shu_lei_xingOptions" :key="key" :label="item.label" :value="item.value" />
          </el-select>
        </el-form-item>
        <el-form-item label="项目登记时间:"  prop="xiangMuDengJiShiJian" >
          <el-date-picker v-model="formData.xiangMuDengJiShiJian" type="date" style="width:100%" placeholder="选择日期" :clearable="true"  />
        </el-form-item>
        <el-form-item label="部门审核:"  prop="buMenShenHe" >
          <el-input v-model="formData.buMenShenHe" :clearable="true"  placeholder="请输入" />
        </el-form-item>
        <el-form-item label="部门审核时间:"  prop="buMenShenHeShiJian" >
          <el-date-picker v-model="formData.buMenShenHeShiJian" type="date" style="width:100%" placeholder="选择日期" :clearable="true"  />
        </el-form-item>
        <el-form-item label="CEO审核:"  prop="ceOshenHe" >
          <el-input v-model="formData.ceOshenHe" :clearable="true"  placeholder="请输入" />
        </el-form-item>
        <el-form-item label="CEO审核时间:"  prop="ceOshenHeShiJian" >
          <el-date-picker v-model="formData.ceOshenHeShiJian" type="date" style="width:100%" placeholder="选择日期" :clearable="true"  />
        </el-form-item>
        <el-form-item label="投标函制作人:"  prop="touBiaoHanZhiZuoRen" >
          <el-input v-model="formData.touBiaoHanZhiZuoRen" :clearable="true"  placeholder="请输入" />
        </el-form-item>
        <el-form-item label="投标函提成金额:"  prop="touBiaoHanTiChengJinE" >
          <el-input v-model="formData.touBiaoHanTiChengJinE" :clearable="true"  placeholder="请输入" />
        </el-form-item>
        <el-form-item label="投标函制作人确认:"  prop="touBiaoHanZhiZuoRenQueRen" >
          <el-input v-model="formData.touBiaoHanZhiZuoRenQueRen" :clearable="true"  placeholder="请输入" />
        </el-form-item>
        <el-form-item label="投标函制作人确认时间:"  prop="touBiaoHanZhiZuoRenQueRenShiJian" >
          <el-date-picker v-model="formData.touBiaoHanZhiZuoRenQueRenShiJian" type="date" style="width:100%" placeholder="选择日期" :clearable="true"  />
        </el-form-item>
        <el-form-item label="商务标制作人:"  prop="shangWuBiaoZhiZuoRen" >
          <el-input v-model="formData.shangWuBiaoZhiZuoRen" :clearable="true"  placeholder="请输入" />
        </el-form-item>
        <el-form-item label="商务标提成金额:"  prop="shangWuBiaoTiChengJinE" >
          <el-input v-model="formData.shangWuBiaoTiChengJinE" :clearable="true"  placeholder="请输入" />
        </el-form-item>
        <el-form-item label="商务标制作人确认:"  prop="shangWuBiaoZhiZuoRenQueRen" >
          <el-input v-model="formData.shangWuBiaoZhiZuoRenQueRen" :clearable="true"  placeholder="请输入" />
        </el-form-item>
        <el-form-item label="商务标制作人确认时间:"  prop="shangWuBiaoZhiZuoRenQueRenShiJian" >
          <el-date-picker v-model="formData.shangWuBiaoZhiZuoRenQueRenShiJian" type="date" style="width:100%" placeholder="选择日期" :clearable="true"  />
        </el-form-item>
        <el-form-item label="技术标制作人:"  prop="jiShuBiaoZhiZuoRen" >
          <el-input v-model="formData.jiShuBiaoZhiZuoRen" :clearable="true"  placeholder="请输入" />
        </el-form-item>
        <el-form-item label="技术标提成金额:"  prop="jiShuBiaoTiChengJinE" >
          <el-input v-model="formData.jiShuBiaoTiChengJinE" :clearable="true"  placeholder="请输入" />
        </el-form-item>
        <el-form-item label="技术标制作人确认:"  prop="jiShuBiaoZhiZuoRenQueRen" >
          <el-input v-model="formData.jiShuBiaoZhiZuoRenQueRen" :clearable="true"  placeholder="请输入" />
        </el-form-item>
        <el-form-item label="技术标制作人确认时间:"  prop="jiShuBiaoZhiZuoRenQueRenShiJian" >
          <el-date-picker v-model="formData.jiShuBiaoZhiZuoRenQueRenShiJian" type="date" style="width:100%" placeholder="选择日期" :clearable="true"  />
        </el-form-item>
        <el-form-item label="项目状态:"  prop="xiangMuZhuangTai" >
          <el-input v-model="formData.xiangMuZhuangTai" :clearable="true"  placeholder="请输入" />
        </el-form-item>
        <el-form-item label="部门申请下浮点位:"  prop="buMenShenQingXiaFuDianWei" >
          <el-input v-model="formData.buMenShenQingXiaFuDianWei" :clearable="true"  placeholder="请输入" />
        </el-form-item>
        <el-form-item label="部门下浮点位申请时间:"  prop="buMenXiaFuDianWeiShenQingShiJian" >
          <el-date-picker v-model="formData.buMenXiaFuDianWeiShenQingShiJian" type="date" style="width:100%" placeholder="选择日期" :clearable="true"  />
        </el-form-item>
        <el-form-item label="CEO审核下浮点位:"  prop="ceOshenHeXiaFuDianWei" >
          <el-input v-model="formData.ceOshenHeXiaFuDianWei" :clearable="true"  placeholder="请输入" />
        </el-form-item>
        <el-form-item label="CEO下浮点位审核时间:"  prop="ceOxiaFuDianWeiShenHeShiJian" >
          <el-date-picker v-model="formData.ceOxiaFuDianWeiShenHeShiJian" type="date" style="width:100%" placeholder="选择日期" :clearable="true"  />
        </el-form-item>
        <el-form-item label="中标金额:"  prop="zhongBiaoJinE" >
          <el-input v-model="formData.zhongBiaoJinE" :clearable="true"  placeholder="请输入" />
        </el-form-item>
        <el-form-item label="备注:"  prop="beiZhu" >
          <el-input v-model="formData.beiZhu" :clearable="true"  placeholder="请输入" />
        </el-form-item>
      </el-form>
      <template #footer>
        <div class="dialog-footer">
          <el-button size="small" @click="closeDialog">取 消</el-button>
          <el-button size="small" type="primary" @click="enterDialog">确 定</el-button>
        </div>
      </template>
    </el-dialog>
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
  deleteXiangMuDengJiBiao,
  deleteXiangMuDengJiBiaoByIds,
  updateXiangMuDengJiBiao,
  findXiangMuDengJiBiao,
  getXiangMuDengJiBiaoList
} from '@/api/xiangMuDengJiBiao'

// 全量引入格式化工具 请按需保留
import { getDictFunc, formatDate, formatBoolean, filterDict } from '@/utils/format'
import { ElMessage, ElMessageBox } from 'element-plus'
import { ref, reactive } from 'vue'

// 自动化生成的字典（可能为空）以及字段
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


// =========== 表格控制部分 ===========
const page = ref(1)
const total = ref(0)
const pageSize = ref(10)
const tableData = ref([])
const searchInfo = ref({})

// 重置
const onReset = () => {
  searchInfo.value = {}
}

// 搜索
const onSubmit = () => {
  page.value = 1
  pageSize.value = 10
  getTableData()
}

// 分页
const handleSizeChange = (val) => {
  pageSize.value = val
  getTableData()
}

// 修改页面容量
const handleCurrentChange = (val) => {
  page.value = val
  getTableData()
}

// 查询
const getTableData = async() => {
  const table = await getXiangMuDengJiBiaoList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
  if (table.code === 0) {
    tableData.value = table.data.list
    total.value = table.data.total
    page.value = table.data.page
    pageSize.value = table.data.pageSize
  }
}

getTableData()

// ============== 表格控制部分结束 ===============

// 获取需要的字典 可能为空 按需保留
const setOptions = async () =>{
    zi_zhi_yao_qiuOptions.value = await getDictFunc('zi_zhi_yao_qiu')
    deng_ji_yao_qiuOptions.value = await getDictFunc('deng_ji_yao_qiu')
    suo_shu_di_quOptions.value = await getDictFunc('suo_shu_di_qu')
    bao_zheng_jin_xing_shiOptions.value = await getDictFunc('bao_zheng_jin_xing_shi')
    ping_biao_ban_faOptions.value = await getDictFunc('ping_biao_ban_fa')
    biao_shu_lei_xingOptions.value = await getDictFunc('biao_shu_lei_xing')
}

// 获取需要的字典 可能为空 按需保留
setOptions()


// 多选数据
const multipleSelection = ref([])
// 多选
const handleSelectionChange = (val) => {
    multipleSelection.value = val
}

// 删除行
const deleteRow = (row) => {
    ElMessageBox.confirm('确定要删除吗?', '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
    }).then(() => {
            deleteXiangMuDengJiBiaoFunc(row)
        })
    }


// 批量删除控制标记
const deleteVisible = ref(false)

// 多选删除
const onDelete = async() => {
      const ids = []
      if (multipleSelection.value.length === 0) {
        ElMessage({
          type: 'warning',
          message: '请选择要删除的数据'
        })
        return
      }
      multipleSelection.value &&
        multipleSelection.value.map(item => {
          ids.push(item.ID)
        })
      const res = await deleteXiangMuDengJiBiaoByIds({ ids })
      if (res.code === 0) {
        ElMessage({
          type: 'success',
          message: '删除成功'
        })
        if (tableData.value.length === ids.length && page.value > 1) {
          page.value--
        }
        deleteVisible.value = false
        getTableData()
      }
    }

// 行为控制标记（弹窗内部需要增还是改）
const type = ref('')

// 更新行
const updateXiangMuDengJiBiaoFunc = async(row) => {
    const res = await findXiangMuDengJiBiao({ ID: row.ID })
    type.value = 'update'
    if (res.code === 0) {
        formData.value = res.data.rexiangMuDengJiBiao
        dialogFormVisible.value = true
    }
}


// 删除行
const deleteXiangMuDengJiBiaoFunc = async (row) => {
    const res = await deleteXiangMuDengJiBiao({ ID: row.ID })
    if (res.code === 0) {
        ElMessage({
                type: 'success',
                message: '删除成功'
            })
            if (tableData.value.length === 1 && page.value > 1) {
            page.value--
        }
        getTableData()
    }
}

// 弹窗控制标记
const dialogFormVisible = ref(false)

// 打开弹窗
const openDialog = () => {
    type.value = 'create'
    dialogFormVisible.value = true
}

// 关闭弹窗
const closeDialog = () => {
    dialogFormVisible.value = false
    formData.value = {
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
        }
}
// 弹窗确定
const enterDialog = async () => {
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
                closeDialog()
                getTableData()
              }
      })
}
</script>

<style>
</style>
