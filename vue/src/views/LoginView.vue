<template>
  <div class="container">
    <div
      style="width: 350px; padding: 30px; background-color: rgba(255, 255, 255, .6{}); border-radius: 5px; box-shadow: 0 0 10px rgba(0, 0, 0, .5)">
      <div style="text-align: center; font-size: 20px; margin-bottom: 40px; color: #333">Dormitory-Smart-Matching-System
      </div>
      <el-form :model="form" :rules="rules" ref="formRef">
        <el-form-item prop="username">
          <el-input prefix-icon="el-icon-user" placeholder="Please enter your account number."
            v-model="loginForm.username"></el-input>
        </el-form-item>
        <el-form-item prop="password">
          <el-input prefix-icon="el-icon-lock" placeholder="Please enter password" show-password
            v-model="loginForm.password"></el-input>
        </el-form-item>
        <el-form-item prop="role">
          <el-select v-model="loginForm.role" style="width: 100%">
            <el-option value="ADMIN" label="Admin"></el-option>
            <el-option value="USER" label="User"></el-option>
          </el-select>
        </el-form-item>
        <el-form-item>
          <el-button style="width: 100%; background-color: #333; border-color: #333; color: white"
            @click="login">Login</el-button>
        </el-form-item>
        <!--        <div style="display: flex; align-items: center">-->
        <!--          <div style="flex: 1"></div>-->
        <!--          <div style="flex: 1; text-align: right">-->
        <!--            还没有账号？请 <a href="/register">注册</a>-->
        <!--          </div>-->
        <!--        </div>-->
      </el-form>
    </div>
  </div>
</template>

<script>
import request from "@/utils/request";
import axios from "axios"
export default {

  name: "LoginView",
  data() {
    return {
      loginForm: {
        role: 'USER',
        username: '',
        password: null
      },
      rules: {
        username: [
          { required: true, message: '请输入账号', trigger: 'blur' },
        ],
        password: [
          { required: true, message: '请输入密码', trigger: 'blur' },
        ]
      }
    }
  },
  created() {

  },
  methods: {
    login() {
      // axios.post("",this.loginForm).then(res=>{
      //   let data=res.data
      //   if(data.success)
      //   {
      //     this.loginForm={}
      //     this.$message({
      //       massage:"成功登录,欢迎来到宿舍分配系统",
      //       type:"success"
      //   })
      //     this.$router.push({path:"/front"})
      //   }
      // })
      request.post("/login", this.loginForm).then(res => {
        this.$message({
          message: '恭喜你，成功登录到寝室智能匹配系统',
          type: 'success'
        });
        localStorage.setItem('token',res.token)
        this.$router.push({ path: "/front" })
      }).catch(res => {
        this.$message({
          message: '登陆失败，账号或密码错误',
          type: 'error'
        });
      })

    }
  }
}
</script>

<style scoped>
.container {
  height: 100vh;
  overflow: hidden;
  background-image: url("@/assets/bg.jpg");
  background-size: 100%;
  display: flex;
  align-items: center;
  justify-content: center;
  color: #666;
}
a {
  color: #2a60c9;
}
</style>