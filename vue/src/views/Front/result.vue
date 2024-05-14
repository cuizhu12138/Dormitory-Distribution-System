<!--分配结果反馈页面-->

<template>
  <div>
    <h1>-Allocation result feedback</h1>

    <el-table :data="tableData" border style="width: 100%">
      <el-table-column prop="smoke" label="smoke or not" width="180"></el-table-column>
      <el-table-column prop="home" label="home" width="180"></el-table-column>
      <el-table-column prop="learnInDorm" label="learn in dorm or not"></el-table-column>
      <el-table-column prop="hobby" label="hobby"></el-table-column>
    </el-table>

    <el-button type="primary">Reassign?</el-button>
  </div>



</template>

<script>

import axios from "axios"
export default {
  name: "assignedResults",
  data(){
    return{
      tableData:[]
    }
  },

  created(){
    axios.get("http://localhost:3000/allocationResults").then(res=>{
      this.tableData=res.data;
    })
  }
}
</script>

<style scoped>

</style>