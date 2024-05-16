<!--分配结果反馈页面-->

<template>
  <div>
    <h1>-Allocation result feedback</h1>

    <el-table :data="tableData" border style="width: 100%">
      <el-table-column prop="name" label="室友名称" ></el-table-column>
      <el-table-column fixed="right" label="操作" width="150">
        <template>
          <el-row>
            <el-button type="primary" @click="checkQuesionnaire">查看问卷</el-button>
          </el-row>
        </template>
      </el-table-column>
    </el-table>

    <el-dialog title="问卷查看" :visible.sync="dialogVisble" width="50%" >
      <el-form :model="form">

        <el-form-item label="学号" :label-width="formLabelWidth">
          <el-input v-model="form.studentId" autocomplete="off" :disabled="true"></el-input>
        </el-form-item>

        <el-form-item label="姓名" :label-width="formLabelWidth">
          <el-input v-model="form.name" autocomplete="off" :disabled="true"></el-input>
        </el-form-item>
        <el-form-item label="性别" :label-width="formLabelWidth" >
          <el-select v-model="form.sex" placeholder="请选择性别" disabled>
            <el-option label="男" value="男" ></el-option>
            <el-option label="女" value="女" ></el-option>
          </el-select>
        </el-form-item>

        <el-form-item label="专业" :label-width="formLabelWidth"  >
          <el-input v-model="form.major" autocomplete="off" :disabled="true"></el-input>
        </el-form-item>

        <el-form-item label="年龄" :label-width="formLabelWidth">
          <el-input v-model="form.age" autocomplete="off" :disabled="true"></el-input>
        </el-form-item>

        <span>籍贯</span><br>
        <el-cascader
            size="large"
            :options="pcTextArr"
            v-model="form.home" placeholder="请选择地区">
        </el-cascader>



        <el-form-item label="民族" :label-width="formLabelWidth">
          <el-input v-model="form.ethnic" autocomplete="off" :disabled="true"></el-input>
        </el-form-item>

        <el-form-item label="睡觉时间" :label-width="formLabelWidth">
          <el-select v-model="form.sleepTime" placeholder="请选择时间" :disabled="true">
            <el-option label="22:30-23:30" value=0></el-option>
            <el-option label="23:30-00:30" value=1></el-option>
            <el-option label="00:30-01:30" value=2></el-option>
            <el-option label="01:30以后" value=3></el-option>
          </el-select>
        </el-form-item>

        <el-form-item label="起床时间" :label-width="formLabelWidth">
          <el-select v-model="form.getupTime" placeholder="请选择时间"  :disabled="true">
            <el-option label="7:00之前" value=0></el-option>
            <el-option label="7:00-8:30" value=1></el-option>
            <el-option label="8:30-10:00" value=2></el-option>
            <el-option label="10:00以后" value=3></el-option>
          </el-select>
        </el-form-item>

        <el-form-item label="是否希望作息时间同步">
          <el-select v-model="form.sameRoutine" :disabled="true">
            <el-option label="是" value=0></el-option>
            <el-option label="否" value=1></el-option>
          </el-select>
        </el-form-item>

        <el-form-item label="学习是否更偏向在寝室">
          <el-select v-model="form.learnInDorm" :disabled="true">
            <el-option label="是" value=0></el-option>
            <el-option label="否" value=1></el-option>
          </el-select>
        </el-form-item>

        <el-form-item label="对寝室的整洁程度要求和期望">
          <el-select v-model="form.neatExpection" :disabled="true">
            <el-option label="有洁癖，需要寝室保持高整洁度" value=0></el-option>
            <el-option label="间歇式洁癖，有的时候特别整洁" value=1></el-option>
            <el-option label="别人整洁我也整洁，别人无所谓我也无所谓" value=2></el-option>
            <el-option label="对整洁度要求不高，无所谓" value=3></el-option>
          </el-select>
        </el-form-item>

        <el-form-item label="打扫寝室周期">
          <el-select v-model="form.cleanPeriod" :disabled="true">
            <el-option label="每天" value=0></el-option>
            <el-option label="一周两-三次" value=1></el-option>
            <el-option label="一周一次" value=2></el-option>
            <el-option label="看心情" value=3></el-option>
            <el-option label="从不" value=4></el-option>
          </el-select>
        </el-form-item>

        <el-form-item label="洗澡周期">
          <el-select v-model="form.bathePeriod" :disabled="true">
            <el-option label="每天" value=0></el-option>
            <el-option label="一周两-三次" value=1></el-option>
            <el-option label="一周一次" value=2></el-option>
            <el-option label="看心情，可能很久" value=3></el-option>
          </el-select>
        </el-form-item>

        <el-form-item label="月生活费大概多少">
          <el-select v-model="form.expense" :disabled="true">
            <el-option label="0k-1.5k" value=0></el-option>
            <el-option label="1.5k-2.5k" value=1></el-option>
            <el-option label="2.5k-3.5k" value=2></el-option>
            <el-option label="3.5k以上" value=3></el-option>
          </el-select>
        </el-form-item>

        <el-form-item label="花销主体">
          <el-checkbox-group v-model="form.costType" :disabled="true">
            <el-checkbox label="吃饭，买零食" name=0></el-checkbox>
            <el-checkbox label="娱乐活动" name=1></el-checkbox>
            <el-checkbox label="社交" name=2></el-checkbox>
            <el-checkbox label="提升自己" name=3></el-checkbox>
          </el-checkbox-group>
        </el-form-item>

        <el-form-item label="是否愿意与室友一起外出消费">
          <el-select v-model="form.outCost" :disabled="true">
            <el-option label="经常" value=0></el-option>
            <el-option label="偶尔" value=1></el-option>
            <el-option label="不会" value=2></el-option>
          </el-select>
        </el-form-item>

        <el-form-item label="是否会愿意室友一起共享消费">
          <el-select v-model="form.shareCost" :disabled="true">
            <el-option label="经常" value=0></el-option>
            <el-option label="偶尔" value=1></el-option>
            <el-option label="不会" value=2></el-option>
          </el-select>
        </el-form-item>

        <el-form-item label="兴趣爱好 ">
          <el-checkbox-group v-model="form.hobby" :disabled="true">
            <el-checkbox label="运动" name=0></el-checkbox>
            <el-checkbox label="网上冲浪" name=1></el-checkbox>
            <el-checkbox label="读书学习" name=2></el-checkbox>
            <el-checkbox label="艺术相关" name=3></el-checkbox>
            <el-checkbox label="其他" name=3></el-checkbox>
          </el-checkbox-group>
        </el-form-item>

        <el-form-item label="是否期望与室友爱好相同">
          <el-select v-model="form.hobbySameExpection" :disabled="true">
            <el-option label="非常希望" value=0></el-option>
            <el-option label="最好可以" value=1></el-option>
            <el-option label="都可以接收" value=2></el-option>
          </el-select>
        </el-form-item>

        <el-form-item label="是否会与室友聊天或交流心事">
          <el-select v-model="form.wantCommunicate" :disabled="true">
            <el-option label="经常会" value=0></el-option>
            <el-option label="偶尔会" value=1></el-option>
            <el-option label="不会" value=2></el-option>
          </el-select>
        </el-form-item>

        <el-form-item label="是否抽烟">
          <el-select v-model="form.smoke" :disabled="true">
            <el-option label="是" value=0></el-option>
            <el-option label="否" value=1></el-option>
          </el-select>
        </el-form-item>

        <el-form-item label="是否喝酒">
          <el-select v-model="form.drink" :disabled="true">
            <el-option label="是" value=0></el-option>
            <el-option label="否" value=1></el-option>
          </el-select>
        </el-form-item>

        <el-form-item label="是否打呼噜">
          <el-select v-model="form.snore" :disabled="true">
            <el-option label="是" value=0></el-option>
            <el-option label="否" value=1></el-option>
          </el-select>
        </el-form-item>

        <el-form-item label="睡眠质量">
          <el-select v-model="form.sleepQuality" :disabled="true">
            <el-option label="好" value=0></el-option>
            <el-option label="一般" value=1></el-option>
            <el-option label="差" value=2></el-option>
          </el-select>
        </el-form-item>
      </el-form>
    </el-dialog>

    <el-button type="primary" @click="Reassign">Reassign?</el-button>
  </div>


</template>

<script>

import axios from "axios"
import {
  provinceAndCityData,
  pcTextArr,
  regionData,
  pcaTextArr,
  codeToText,
} from "element-china-area-data";
export default {
  name: "assignedResults",
  data(){
    return{
      dialogVisble:false,
      pcTextArr,
      tableData:[
        
      ],
      form:{
          id: "8bb6",
          questionnaireId: null,
          studentId: "2204020226",
          name: "闫佳安",
          sex: "男",
          major: "网络工程",
          age: "20",
          home: [
            "黑龙江省",
            "哈尔滨市"
          ],
          ethnic: "汉族",
          sleepTime: "0",
          getupTime: "0",
          learnInDorm: "1",
          neatExpection: "0",
          cleanPeriod: "1",
          bathePeriod: "1",
          expense: "1",
          costType: [
            "吃饭，买零食",
            "娱乐活动"
          ],
          outCost: "1",
          shareCost: "1",
          hobby: [
            "运动",
            "网上冲浪"
          ],
          hobbySameExpection: "1",
          wantCommunicate: "1",
          smoke: "0",
          drink: "0",
          snore: "0",
          sleepQuality: "0"
        },

    }
  },

  created(){
    //与后端对接时用,获取室友名称
    axios.get("http://localhost:8080/results").then(res=>{
      this.tableData=res.data
    })
  },
  methods:{
    checkQuesionnaire(){
      this.dialogVisble = true
      // 与后端对接时用,获取室友问卷的具体信息
      // axios.get("").then(res=>{
      //   this.form=res.data
      // })
    },
    Reassign(){
      this.$message({
        message: '收到重新分配需求！',
        type: 'success'
      });
    }
  }
}
</script>

<style scoped>

</style>