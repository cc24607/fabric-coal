<template>
	<div>
		<i-divider color="#2d8cf0" lineColor="#2d8cf0">通话人员</i-divider>
		<i-input :value="productData.product_id" title="通话人员1的ID" disabled />
		<i-input :value="productData.product_name" title="通话人员1的姓名" disabled />
		<i-input :value="productData.retailer" title="通话时间段" disabled />
		<i-input :value="productData.retailer_name" title="通话人员2的姓名" disabled />
		<i-input :value="productData.retailer_tel" title="通话人员2的ID" disabled />

		<i-divider color="#2d8cf0" lineColor="#2d8cf0">产品基本信息</i-divider>

		<i-input :value="productData.factory" title="所属煤矿企业" disabled />
		<i-input :value="productData.leader" title="企业负责人" disabled />
		<i-input :value="productData.leader_tel" title="企业负责人电话" disabled />
		<i-input :value="productData.Net_Content" title="产量" disabled />
		<i-input :value="productData.workshop" title="开采车间" disabled />
		<i-input :value="productData.work_hours" title="开采工时" disabled />
		<i-input :value="productData.keep_mathod" title="保存方法" disabled />
		<i-input :value="productData.cooking_recommend" title="额外说明" disabled />
		<i-input :value="productData.remarks" title="溯源码" disabled />

		<i-divider color="#2d8cf0" lineColor="#2d8cf0">溯源更多信息</i-divider>

		<i-row>
			<i-col span="5" i-class="col-class"><i-tab-bar-item color="#f759ab" @click="tracePlant" icon="homepage" current-icon="activity_fill" title="开采过程"></i-tab-bar-item></i-col>
			<i-col span="5" i-class="col-class"><i-tab-bar-item color="#2d8cf0" @click="traceDriver" icon="coordinates_fill" current-icon="coordinates_fill" title="人员行为"></i-tab-bar-item></i-col>
			<i-col span="5" i-class="col-class"><i-tab-bar-item color="#19be6b" @click="traceCheck" icon="task_fill" current-icon="task_fill" title="质检信息"></i-tab-bar-item></i-col>
			<i-col span="5" i-class="col-class"><i-tab-bar-item color="#ff9900" @click="productProcess" icon="createtask_fill" current-icon="create_fill" title="生产行为"></i-tab-bar-item></i-col>
			<i-col span="4" i-class="col-class"><i-tab-bar-item color="#80848f" @click="traceCropsInfo" icon="shop_fill" current-icon="shop_fill" title="简要信息"></i-tab-bar-item></i-col>
		</i-row>

		<!-- 开采环境 -->
		<div v-show="showPlant" style="padding-top: 1.25rem;" v-for="(detail,index) in cropsProcessDetailsArray" :key="index">
			<i-card :title="detail.crops_grow_id" :thumb="detail.crops_grow_photo_url">
				<l-album :urls='photo'></l-album>
				<view slot="footer"><i-input :value="detail.grow_status" title="设备情况" disabled /></view>
				<view slot="footer"><i-input :value="detail.illumination_status" title="光线情况" disabled /></view>
				<view slot="footer"><i-input :value="detail.water_content" title="湿度" disabled /></view>
				<view slot="footer"><i-input :value="detail.temperature" title="温度" disabled /></view>
				<view slot="footer"><i-input :value="detail.record_time" title="记录时间" disabled /></view>
				<view slot="footer"><i-input :value="detail.remarks" title="溯源码" disabled /></view>
			</i-card>
		</div>
		<br>

		<div v-show="showDriver" style="padding-top: 1.25rem;">
			<i-steps :current="len" direction="vertical">
				<i-step v-for="(detail,index) in cropsDriverArray" :key="index">
					<view slot="content">物流ID:{{detail.transport_id}}</view>
					<view slot="content">经过地:{{detail.transport_to_address}}</view>
					<view slot="content">时间:{{detail.transport_to_chain_time}}</view>
					<view slot="content">司机:{{detail.driver_name}}</view>
					<view slot="content">电话:{{detail.driver_tel}}</view>
					<view slot="content">备注:{{detail.remarks}}</view>
				</i-step>
			</i-steps>
		</div>

		<div v-show="showMaching" style="padding-top: 1.25rem;">
			<i-card :title="machingInfo.machining_id" :thumb="machingInfo.testing_photo_url">
				<view slot="footer"><i-input :value="machingInfo.leader" title="管理员姓名" disabled /></view>
				<view slot="footer"><i-input :value="machingInfo.leader_tel" title="电话" disabled /></view>
				<view slot="footer"><i-input :value="machingInfo.factory_name" title="监管机构" disabled /></view>
				<view slot="footer"><i-input :value="machingInfo.in_factory_time" title="读取时间" disabled /></view>
				<view slot="footer"><i-input :value="machingInfo.out_factory_time" title="完成审核时间" disabled /></view>
				<view slot="footer"><i-input :value="machingInfo.testing_result" title="审核结果" disabled /></view>
				<view slot="footer"><i-input :value="machingInfo.remarks" title="溯源码" disabled /></view>
			</i-card>
		</div>

		<div v-show="showProcess" style="padding-top: 1.25rem;" v-for="(detail,index) in operationArray" :key="index">
			<i-card :title="detail.operation_id">
				<view slot="footer"><i-input :value="detail.operation_people_name" title="操作人" disabled /></view>
				<view slot="footer"><i-input :value="detail.operation_people_tel" title="操作人电话" disabled /></view>
				<view slot="footer"><i-input :value="detail.work_content" title="操作内容" disabled /></view>
				<view slot="footer"><i-input :value="detail.time" title="时间" disabled /></view>
				<view slot="footer"><i-input :value="detail.remarks" title="溯源码" disabled /></view>
			</i-card>
		</div>

		<div v-show="showCropsInfo" style="padding-top: 1.25rem;">
			<i-card :title="cropsDetails.crops_id">
				<view slot="footer"><i-input :value="cropsDetails.crops_name" title="操作员(语音转文本)姓名" disabled /></view>
				<view slot="footer"><i-input :value="cropsDetails.farmer_name" title="通话人员1" disabled /></view>
				<view slot="footer"><i-input :value="cropsDetails.farmer_tel" title="通话人员1电话" disabled /></view>
				<view slot="footer"><i-input :value="cropsDetails.address" title="地址" disabled /></view>
				<view slot="footer"><i-input :value="cropsDetails.plant_mode" title="通话方式" disabled /></view>
				<view slot="footer"><i-input :value="cropsDetails.year" title="通话时长" disabled /></view>
				<view slot="footer"><i-input :value="cropsDetails.bagging_status" title="处理编号" disabled /></view>
				<view slot="footer"><i-input :value="cropsDetails.fertilizer_name" title="通话人员2" disabled /></view>
				<view slot="footer"><i-input :value="cropsDetails.grow_seedlings_cycle" title="通话人员2电话" disabled /></view>
				<view slot="footer"><i-input :value="cropsDetails.irrigation_cycle" title="原始语音哈希" disabled /></view>
				<view slot="footer"><i-input :value="cropsDetails.weed_cycle" title="关键信息摘要" disabled /></view>
				<view slot="footer"><i-input :value="cropsDetails.apply_fertilizer_cycle" title="审查结果" disabled /></view>
				<view slot="footer"><i-input :value="cropsDetails.register_time" title="审查时间" disabled /></view>
				<view slot="footer"><i-input :value="cropsDetails.remarks" title="溯源码" disabled /></view>
			</i-card>
		</div>


	</div>
</template>

<script>
export default {
	data() {
		return {
			productData: '',
			current: 'homepage',
			cropsProcessDetailsArray:[],
			cropsDriverArray:[],
			operationArray:[],
			machingInfo:'',
			cropsDetails:'',
		    verticalCurrent : 2,
			len:'',
			showDriver:false,
			showPlant:false,
			showMaching:false,
			showProcess:false,
			showCropsInfo:false,
			photo:['https://ss0.bdstatic.com/70cFvHSh_Q1YnxGkpoWK1HF6hhy/it/u=154562898,2539807908&fm=15&gp=0.jpg']
		};
	},

	methods: {
		traceCropsInfo(){
			this.showCropsInfo = true
			this.showDriver = false
			this.showMaching = false
			this.showPlant = false
			this.showProcess = false
			this.$httpBlock
				.get(this.$httpUrl + '/farmerapi/queryCropsById?id='+this.productData.crops_id)
				.then(res => {
					this.cropsDetails = res.data;
				})
				.catch(err => {});
		},

		productProcess(){
			this.showProcess = true
			this.showCropsInfo = false
			this.showDriver = false
			this.showMaching = false
			this.showPlant = false
			this.$httpBlock
			    	.get(this.$httpUrl + '/productapi/queryOperationByCropsId?id='+this.productData.crops_id)
			    	.then(res => {
						const array = [];
						for (let i = 0; i < res.data.length; i++) {
							array.push(res.data[i].Record);
						}
						this.operationArray = array;
			    	})
			    	.catch(err => {
			    		this.msgError('查询异常 ' + err);
			    	});
		},

		traceCheck(){
			this.showMaching = true
			this.showCropsInfo = false
			this.showDriver = false
			this.showPlant = false
			this.showProcess = false
			this.$httpBlock
				.get(this.$httpUrl + '/materialapi/queryMachiningByCropsId?id='+this.productData.crops_id)
				.then(res => {
					this.machingInfo = res.data[0].Record
				})
				.catch(err => {});
		},
		traceDriver(){
			this.showDriver = true
			this.showCropsInfo = false
			this.showMaching = false
			this.showPlant = false
			this.showProcess = false
			this.$httpBlock
				.get(this.$httpUrl + '/driverapi/queryTransportByCropsId?id='+this.productData.crops_id)
				.then(res => {
					const array = [];
					this.len = res.data.length
					for (let i = 0; i < res.data.length; i++) {
						array.push(res.data[i].Record);
					}
					this.cropsDriverArray = array
					console.log("array "+JSON.stringify(array))
				})
				.catch(err => {});
		},

		tracePlant() {
			this.showPlant = true
			this.showCropsInfo = false
			this.showDriver = false
			this.showMaching = false
			this.showProcess = false
			this.$httpBlock
				.get(this.$httpUrl + '/farmerapi/queryCropsProcessByCropsId?id='+this.productData.crops_id)
				.then(res => {
					const array = [];
					for (let i = 0; i < res.data.length; i++) {
						array.push(res.data[i].Record);
					}
					this.cropsProcessDetailsArray = array
					console.log("array "+JSON.stringify(array))
				})
				.catch(err => {});
		},

		handleChange({ detail }) {
			current: detail.key;
		},
		query() {
			this.$httpBlock
				.get(this.$httpUrl + '/retailerapi/queryRetailerById?id='+this.productData.crops_id)
				.then(res => {
					wx.navigateTo({
						url: '/pages/product/main'
					});
					console.log('res ' + JSON.stringify(res.data));
				})
				.catch(err => {
					console.log('调用失败 ' + JSON.stringify(err));
				});
		}
	},

	created() {

	},

	mounted() {
		this.productData = JSON.parse(decodeURIComponent(this.$root.$mp.query.data));
	}
};
</script>

<style></style>
