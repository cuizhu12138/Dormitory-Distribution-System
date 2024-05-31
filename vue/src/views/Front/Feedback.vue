<!--分配结果反馈页面-->

<template>
  <div>
    <h1>Suggestions and Advices</h1>

    <el-input v-model="input" placeholder="Provide some suggestions?"></el-input>
    <el-button @click="submit" type="primary">Submit!</el-button>

  </div>



</template>

<script>

import request from "@/utils/request";
import axios from "axios"
import { Message } from 'element-ui';
export default {
  name: "FeedbackView",
  data(){
    return{
      input:"",
    }
  },
  methods:{
    submit(){
      request.post("/feedback",{"content":this.input}).then(res=>{
        this.$message({
          message: '反馈成功',
          type: 'success'
        })
      },err=>{
        Message.error(err)
      })
    }
  },

  created(){

  }
}
</script>

<style scoped>

</style>