
<template>

  <div>
    <div>测试页面</div>

    <div>

      <el-table :data="tableData" border style="width: 100%">
        <el-table-column prop="title" label="问卷名称" ></el-table-column>
        <el-table-column prop="state" label="问卷状态" ></el-table-column>
        <el-table-column fixed="right" label="操作" width="150">
          <template slot-scope="scope">
            <el-row>
              <el-button @click="checkQuesionnaire(scope.row.qid)" type="primary" size="small">查看</el-button>
              <el-button @click="writeQuesionnaire" type="success" size="small" :disabled="scope.row.state==='不可填写'" >填写</el-button>
            </el-row>
          </template>
        </el-table-column>

      </el-table>

      <button @click="check">测试按钮</button>
    </div>
  </div>


</template>

<script>
import axios from "axios";

export default {
  name: "testView",

  data(){
    return{
      tableData: []
    }
  },
  created() {
    axios.get("http://localhost:3000/questionnaireInfo").then(res=>{
      this.tableData=res.data;
    })
  },

  methods:{
    check(tmp){
        console.log(tmp)
        if(tmp.state==="可填写"&&!tmp.state)return false
        return true
    },
    checkQuesionnaire(qid){
      // axios.get("http://localhost:3000/questionnaire").then(res=>{
      //   let tmpform={};
      //   for(let i=0;i<res.data.length;i++)
      //   {
      //     let p=res.data[i];
      //     if(p.qid===qid)
      //     {
      //       console.log(res.data[i])
      //     }
      //   }
      // })

      console.log(qid)
    }
  }
}
</script>

<style scoped>

</style>