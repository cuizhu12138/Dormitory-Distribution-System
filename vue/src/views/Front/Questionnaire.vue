<template>
  <div>
    <div>
      <h1>Quesionnaire</h1>

    </div>

    <div>
      <el-table :data="tableData" border style="width: 100%">
        <el-table-column prop="title" label="Name" ></el-table-column>
        <el-table-column prop="state" label="State" ></el-table-column>
        <el-table-column fixed="right" label="option" width="150">
          <template slot-scope="scope">
            <el-row>
              <el-button @click="checkQuesionnaire(scope.row.qid)" type="primary" size="small">Check</el-button>
              <el-button @click="writeQuesionnaire" type="success" size="small" :disabled="scope.row.state==='Disabled'" >Fill in</el-button>
            </el-row>
          </template>
        </el-table-column>

      </el-table>

    <div>
      <!--      问卷填写弹窗-->
      <el-dialog title="问卷填写" :visible.sync="dialogVisble" width="50%" >
        <el-form :model="form">

          <el-form-item label="学号" :label-width="formLabelWidth">
            <el-input v-model="form.studentId" autocomplete="off"></el-input>
          </el-form-item>

          <el-form-item label="姓名" :label-width="formLabelWidth">
            <el-input v-model="form.name" autocomplete="off"></el-input>
          </el-form-item>
          <el-form-item label="性别" :label-width="formLabelWidth">
            <el-select v-model="form.sex" placeholder="请选择性别">
              <el-option label="男" value="男"></el-option>
              <el-option label="女" value="女"></el-option>
            </el-select>
          </el-form-item>

          <el-form-item label="专业" :label-width="formLabelWidth">
            <el-input v-model="form.major" autocomplete="off"></el-input>
          </el-form-item>

          <el-form-item label="年龄" :label-width="formLabelWidth">
            <el-input v-model="form.age" autocomplete="off"></el-input>
          </el-form-item>

          <span>籍贯</span><br>
          <el-cascader
              size="large"
              :options="pcTextArr"
              v-model="form.home" placeholder="请选择地区">
          </el-cascader>


          <el-form-item label="民族" :label-width="formLabelWidth">
            <el-input v-model="form.ethnic" autocomplete="off"></el-input>
          </el-form-item>

          <el-form-item label="睡觉时间" :label-width="formLabelWidth">
            <el-select v-model="form.sleepTime" placeholder="请选择时间">
              <el-option label="22:30-23:30" value=0></el-option>
              <el-option label="23:30-00:30" value=1></el-option>
              <el-option label="00:30-01:30" value=2></el-option>
              <el-option label="01:30以后" value=3></el-option>
            </el-select>
          </el-form-item>

          <el-form-item label="起床时间" :label-width="formLabelWidth">
            <el-select v-model="form.getupTime" placeholder="请选择时间">
              <el-option label="7:00之前" value=0></el-option>
              <el-option label="7:00-8:30" value=1></el-option>
              <el-option label="8:30-10:00" value=2></el-option>
              <el-option label="10:00以后" value=3></el-option>
            </el-select>
          </el-form-item>

          <el-form-item label="是否希望作息时间同步">
            <el-select v-model="form.sameRoutine">
              <el-option label="是"></el-option>
              <el-option label="否"></el-option>
            </el-select>
          </el-form-item>

          <el-form-item label="学习是否更偏向在寝室">
            <el-select v-model="form.learnInDorm">
              <el-option label="是" value=0></el-option>
              <el-option label="否" value=1></el-option>
            </el-select>
          </el-form-item>

          <el-form-item label="对寝室的整洁程度要求和期望">
            <el-select v-model="form.neatExpection">
              <el-option label="有洁癖，需要寝室保持高整洁度" value=0></el-option>
              <el-option label="间歇式洁癖，有的时候特别整洁" value=1></el-option>
              <el-option label="别人整洁我也整洁，别人无所谓我也无所谓" value=2></el-option>
              <el-option label="对整洁度要求不高，无所谓" value=3></el-option>
            </el-select>
          </el-form-item>

          <el-form-item label="打扫寝室周期">
            <el-select v-model="form.cleanPeriod">
              <el-option label="每天" value=0></el-option>
              <el-option label="一周两-三次" value=1></el-option>
              <el-option label="一周一次" value=2></el-option>
              <el-option label="看心情" value=3></el-option>
              <el-option label="从不" value=4></el-option>
            </el-select>
          </el-form-item>

          <el-form-item label="洗澡周期">
            <el-select v-model="form.bathePeriod">
              <el-option label="每天" value=0></el-option>
              <el-option label="一周两-三次" value=1></el-option>
              <el-option label="一周一次" value=2></el-option>
              <el-option label="看心情，可能很久" value=3></el-option>
            </el-select>
          </el-form-item>

          <el-form-item label="月生活费大概多少">
            <el-select v-model="form.expense">
              <el-option label="0k-1.5k" value=0></el-option>
              <el-option label="1.5k-2.5k" value=1></el-option>
              <el-option label="2.5k-3.5k" value=2></el-option>
              <el-option label="3.5k以上" value=3></el-option>
            </el-select>
          </el-form-item>

          <el-form-item label="花销主体">
            <el-checkbox-group v-model="form.costType">
              <el-checkbox label="吃饭，买零食" name=0></el-checkbox>
              <el-checkbox label="娱乐活动" name=1></el-checkbox>
              <el-checkbox label="社交" name=2></el-checkbox>
              <el-checkbox label="提升自己" name=3></el-checkbox>
            </el-checkbox-group>
          </el-form-item>

          <el-form-item label="是否愿意与室友一起外出消费">
            <el-select v-model="form.outCost">
              <el-option label="经常" value=0></el-option>
              <el-option label="偶尔" value=1></el-option>
              <el-option label="不会" value=2></el-option>
            </el-select>
          </el-form-item>

          <el-form-item label="是否会愿意室友一起共享消费">
            <el-select v-model="form.shareCost">
              <el-option label="经常" value=0></el-option>
              <el-option label="偶尔" value=1></el-option>
              <el-option label="不会" value=2></el-option>
            </el-select>
          </el-form-item>

          <el-form-item label="兴趣爱好 ">
            <el-checkbox-group v-model="form.hobby">
              <el-checkbox label="运动" name=0></el-checkbox>
              <el-checkbox label="网上冲浪" name=1></el-checkbox>
              <el-checkbox label="读书学习" name=2></el-checkbox>
              <el-checkbox label="艺术相关" name=3></el-checkbox>
              <el-checkbox label="其他" name=3></el-checkbox>
            </el-checkbox-group>
          </el-form-item>

          <el-form-item label="是否期望与室友爱好相同">
            <el-select v-model="form.hobbySameExpection">
              <el-option label="非常希望" value=0></el-option>
              <el-option label="最好可以" value=1></el-option>
              <el-option label="都可以接收" value=2></el-option>
            </el-select>
          </el-form-item>

          <el-form-item label="是否会与室友聊天或交流心事">
            <el-select v-model="form.wantCommunicate">
              <el-option label="经常会" value=0></el-option>
              <el-option label="偶尔会" value=1></el-option>
              <el-option label="不会" value=2></el-option>
            </el-select>
          </el-form-item>

          <el-form-item label="是否抽烟">
            <el-select v-model="form.smoke">
              <el-option label="是" value=0></el-option>
              <el-option label="否" value=1></el-option>
            </el-select>
          </el-form-item>

          <el-form-item label="是否喝酒">
            <el-select v-model="form.drink">
              <el-option label="是" value=0></el-option>
              <el-option label="否" value=1></el-option>
            </el-select>
          </el-form-item>

          <el-form-item label="是否打呼噜">
            <el-select v-model="form.snore">
              <el-option label="是" value=0></el-option>
              <el-option label="否" value=1></el-option>
            </el-select>
          </el-form-item>

          <el-form-item label="睡眠质量">
            <el-select v-model="form.sleepQuality">
              <el-option label="好" value=0></el-option>
              <el-option label="一般" value=1></el-option>
              <el-option label="差" value=2></el-option>
            </el-select>
          </el-form-item>


        </el-form>
        <div slot="footer" class="dialog-footer">
          <el-button @click="dialogVisble = false">取 消</el-button>
          <el-button type="primary" @click="saveQuetionnaire">保存</el-button>
          <el-button type="success" @click="updateQuetionnaire">提交</el-button>
        </div>
      </el-dialog>

<!--      问卷查看弹窗-->
      <el-dialog title="问卷查看" :visible.sync="checkDialogVisble" width="50%" >
        <el-form :model="checkform">

          <el-form-item label="学号" :label-width="formLabelWidth">
            <el-input v-model="checkform.studentId" autocomplete="off"></el-input>
          </el-form-item>

          <el-form-item label="姓名" :label-width="formLabelWidth">
            <el-input v-model="checkform.name" autocomplete="off"></el-input>
          </el-form-item>
          <el-form-item label="性别" :label-width="formLabelWidth">
            <el-select v-model="checkform.sex" placeholder="请选择性别">
              <el-option label="男" value="男"></el-option>
              <el-option label="女" value="女"></el-option>
            </el-select>
          </el-form-item>

          <el-form-item label="专业" :label-width="formLabelWidth">
            <el-input v-model="checkform.major" autocomplete="off"></el-input>
          </el-form-item>

          <el-form-item label="年龄" :label-width="formLabelWidth">
            <el-input v-model="checkform.age" autocomplete="off"></el-input>
          </el-form-item>

          <span>籍贯</span><br>
          <el-cascader
              size="large"
              :options="pcTextArr"
              v-model="checkform.home" placeholder="请选择地区">
          </el-cascader>


          <el-form-item label="民族" :label-width="formLabelWidth">
            <el-input v-model="checkform.ethnic" autocomplete="off"></el-input>
          </el-form-item>

          <el-form-item label="睡觉时间" :label-width="formLabelWidth">
            <el-select v-model="checkform.sleepTime" placeholder="请选择时间">
              <el-option label="22:30-23:30" value=0></el-option>
              <el-option label="23:30-00:30" value=1></el-option>
              <el-option label="00:30-01:30" value=2></el-option>
              <el-option label="01:30以后" value=3></el-option>
            </el-select>
          </el-form-item>

          <el-form-item label="起床时间" :label-width="formLabelWidth">
            <el-select v-model="checkform.getupTime" placeholder="请选择时间">
              <el-option label="7:00之前" value=0></el-option>
              <el-option label="7:00-8:30" value=1></el-option>
              <el-option label="8:30-10:00" value=2></el-option>
              <el-option label="10:00以后" value=3></el-option>
            </el-select>
          </el-form-item>

          <el-form-item label="是否希望作息时间同步">
            <el-select v-model="checkform.sameRoutine">
              <el-option label="是"></el-option>
              <el-option label="否"></el-option>
            </el-select>
          </el-form-item>

          <el-form-item label="学习是否更偏向在寝室">
            <el-select v-model="checkform.learnInDorm">
              <el-option label="是" value=0></el-option>
              <el-option label="否" value=1></el-option>
            </el-select>
          </el-form-item>

          <el-form-item label="对寝室的整洁程度要求和期望">
            <el-select v-model="checkform.neatExpection">
              <el-option label="有洁癖，需要寝室保持高整洁度" value=0></el-option>
              <el-option label="间歇式洁癖，有的时候特别整洁" value=1></el-option>
              <el-option label="别人整洁我也整洁，别人无所谓我也无所谓" value=2></el-option>
              <el-option label="对整洁度要求不高，无所谓" value=3></el-option>
            </el-select>
          </el-form-item>

          <el-form-item label="打扫寝室周期">
            <el-select v-model="checkform.cleanPeriod">
              <el-option label="每天" value=0></el-option>
              <el-option label="一周两-三次" value=1></el-option>
              <el-option label="一周一次" value=2></el-option>
              <el-option label="看心情" value=3></el-option>
              <el-option label="从不" value=4></el-option>
            </el-select>
          </el-form-item>

          <el-form-item label="洗澡周期">
            <el-select v-model="checkform.bathePeriod">
              <el-option label="每天" value=0></el-option>
              <el-option label="一周两-三次" value=1></el-option>
              <el-option label="一周一次" value=2></el-option>
              <el-option label="看心情，可能很久" value=3></el-option>
            </el-select>
          </el-form-item>

          <el-form-item label="月生活费大概多少">
            <el-select v-model="checkform.expense">
              <el-option label="0k-1.5k" value=0></el-option>
              <el-option label="1.5k-2.5k" value=1></el-option>
              <el-option label="2.5k-3.5k" value=2></el-option>
              <el-option label="3.5k以上" value=3></el-option>
            </el-select>
          </el-form-item>

          <el-form-item label="花销主体">
            <el-checkbox-group v-model="checkform.costType">
              <el-checkbox label="吃饭，买零食" name=0></el-checkbox>
              <el-checkbox label="娱乐活动" name=1></el-checkbox>
              <el-checkbox label="社交" name=2></el-checkbox>
              <el-checkbox label="提升自己" name=3></el-checkbox>
            </el-checkbox-group>
          </el-form-item>

          <el-form-item label="是否愿意与室友一起外出消费">
            <el-select v-model="checkform.outCost">
              <el-option label="经常" value=0></el-option>
              <el-option label="偶尔" value=1></el-option>
              <el-option label="不会" value=2></el-option>
            </el-select>
          </el-form-item>

          <el-form-item label="是否会愿意室友一起共享消费">
            <el-select v-model="checkform.shareCost">
              <el-option label="经常" value=0></el-option>
              <el-option label="偶尔" value=1></el-option>
              <el-option label="不会" value=2></el-option>
            </el-select>
          </el-form-item>

          <el-form-item label="兴趣爱好 ">
            <el-checkbox-group v-model="checkform.hobby">
              <el-checkbox label="运动" name=0></el-checkbox>
              <el-checkbox label="网上冲浪" name=1></el-checkbox>
              <el-checkbox label="读书学习" name=2></el-checkbox>
              <el-checkbox label="艺术相关" name=3></el-checkbox>
              <el-checkbox label="其他" name=3></el-checkbox>
            </el-checkbox-group>
          </el-form-item>

          <el-form-item label="是否期望与室友爱好相同">
            <el-select v-model="checkform.hobbySameExpection">
              <el-option label="非常希望" value=0></el-option>
              <el-option label="最好可以" value=1></el-option>
              <el-option label="都可以接收" value=2></el-option>
            </el-select>
          </el-form-item>

          <el-form-item label="是否会与室友聊天或交流心事">
            <el-select v-model="checkform.wantCommunicate">
              <el-option label="经常会" value=0></el-option>
              <el-option label="偶尔会" value=1></el-option>
              <el-option label="不会" value=2></el-option>
            </el-select>
          </el-form-item>

          <el-form-item label="是否抽烟">
            <el-select v-model="checkform.smoke">
              <el-option label="是" value=0></el-option>
              <el-option label="否" value=1></el-option>
            </el-select>
          </el-form-item>

          <el-form-item label="是否喝酒">
            <el-select v-model="checkform.drink">
              <el-option label="是" value=0></el-option>
              <el-option label="否" value=1></el-option>
            </el-select>
          </el-form-item>

          <el-form-item label="是否打呼噜">
            <el-select v-model="checkform.snore">
              <el-option label="是" value=0></el-option>
              <el-option label="否" value=1></el-option>
            </el-select>
          </el-form-item>

          <el-form-item label="睡眠质量">
            <el-select v-model="checkform.sleepQuality">
              <el-option label="好" value=0></el-option>
              <el-option label="一般" value=1></el-option>
              <el-option label="差" value=2></el-option>
            </el-select>
          </el-form-item>


        </el-form>
<!--        <div slot="footer" class="dialog-footer">-->
<!--          <el-button @click="dialogVisble = false">取 消</el-button>-->
<!--          <el-button type="primary" @click="saveQuetionnaire">保存</el-button>-->
<!--          <el-button type="success" @click="updateQuetionnaire">提交</el-button>-->
<!--        </div>-->
      </el-dialog>

    </div>

    </div>
  </div>



</template>



<script>
import axios from "axios";
import {
  provinceAndCityData,
  pcTextArr,
  regionData,
  pcaTextArr,
  codeToText,
} from "element-china-area-data";


export default {
  name: 'QuestionnaireView',
  data() {
    return {
      user_id:2204020202,
      qid:null,
      pcTextArr,
      pos:"questionnaire",
      dialogVisble:false,
      checkDialogVisble:false,
      checkform:{},
      form:{
        questionnaireId:null,
        studentId:null,
        name:'',
        sex:'',
        major:'',
        age:'',
        home:[],
        ethnic:'',
        sleepTime:null,
        getupTime:null,
        sameRoutine:null,
        learnInDorm:null,
        neatExpection:null,
        cleanPeriod:null,
        bathePeriod:null,
        expense:null,
        costType:[],
        outCost:null,
        shareCost:null,
        hobby:[],
        hobbySameExpection:null,
        wantCommunicate:null,
        smoke:null,
        drink:null,
        snore:null,
        sleepQuality:null
      },
      tableData: []
    }
  },
  created() {
    axios.get("http://localhost:8080/questionnaireInfo").then(res=>{
      this.tableData=res.data;
    })
  },
  methods:{
    writeQuesionnaire(){
      this.dialogVisble = true
    },
    checkQuesionnaire(qid){
      this.checkDialogVisble=true
      axios.get("http://localhost:8080/questionnaire").then(res=>{
        let tmpform={};
        for(let i=0;i<res.data.length;i++)
        {
          let p=res.data[i];
          if(p.qid===qid&&p.studentId===this.user_id)
          {
            this.checkform=res.data[i]
            console.log(res.data[i])
          }
        }
      })
    },
    saveQuetionnaire(){
      this.dialogVisble = false
      this.$message({
        message: '保存成功',
        type: 'success'
      });
    },
    updateQuetionnaire(){
      axios.post(`http://localhost:8080/${this.pos}`,this.form).then(res=>{
        console.log(res)
      })
      this.dialogVisble=false
      this.$message({
        message: '恭喜你，提交成功',
        type: 'success'
      });
    },
  }



}
</script>


