<!--分配结果反馈页面-->

<template>
  <div>
    <h1>Suggestions and Advices</h1>

    <el-input v-model="input" placeholder="Provide some suggestions?"></el-input>
    <el-button type="primary">Submit!</el-button>

  </div>



</template>

<script>

import axios from "axios"
export default {
  name: "FeedbackView",
  data(){
    return{
      input:"",
    }
  },

  created(){

  }
}
</script>

<style scoped>

</style>