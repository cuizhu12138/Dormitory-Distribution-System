<template>
  <div class="container">
    <div style="width: 350px; padding: 30px; background-color: rgba(255, 255, 255, .6); border-radius: 5px; box-shadow: 0 0 10px rgba(0, 0, 0, .5)">
      <div style="text-align: center; font-size: 26px; margin-bottom: 40px; color: #333">问卷调查管理系统</div>
      <el-form :model="form" :rules="rules" ref="formRef">
        <el-form-item prop="username">
          <el-input prefix-icon="el-icon-user" placeholder="请输入账号" v-model="form.username"></el-input>
        </el-form-item>
        <el-form-item prop="password">
          <el-input prefix-icon="el-icon-lock" placeholder="请输入密码" show-password  v-model="form.password"></el-input>
        </el-form-item>
        <el-form-item prop="role">
          <el-select v-model="form.role" style="width: 100%">
            <el-option value="ADMIN" label="管理员"></el-option>
            <el-option value="USER" label="用户"></el-option>
          </el-select>
        </el-form-item>
        <el-form-item>
          <el-button style="width: 100%; background-color: #333; border-color: #333; color: white" @click="login">登 录</el-button>
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
export default {
  name: "LoginView",
  data() {
    return {
      form: { role: 'USER' },
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
      this.$refs['formRef'].validate((valid) => {
        if (valid) {
          // 验证通过
          this.$request.post('/login', this.form).then(res => {
            if (res.code === '200') {
              localStorage.setItem("xm-user", JSON.stringify(res.data))  // 存储用户数据
              this.$message.success('登录成功')
              setTimeout(() => {  // 延时跳转
                if (res.data.role === 'USER') {
                  location.href = '/front/home'
                } else {
                  location.href = '/manager'
                }
              }, 500)
            } else {
              this.$message.error(res.msg)
            }
          })
        }
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