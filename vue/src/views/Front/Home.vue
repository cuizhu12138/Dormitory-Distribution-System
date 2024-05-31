<template>
  <div>
  <el-container>
    <el-header style="background-color:#687178">
      <img src="../../assets/hrbust.jpg" style="width:40px; position:relative; top:15px">

        <span style="font-size:20px;margin-left:15px;color:white">Dormitory-Smart-Matching-System</span>
        <el-button type="primary" @click="loginOut" plain style="margin-left: 20px; border: 1px solid #fff; padding: 8px 10px;">Exit</el-button>
    </el-header>
  </el-container>

    <el-container>
      <el-aside style="overflow: hidden;min-height:100vh;background-color:#545c64" >
      <el-menu
          default-active="$route.path" router
          class="el-menu-vertical-demo"
          @open="handleOpen"
          @close="handleClose"
          background-color="#545c64"
          text-color="#fff"
          active-text-color="#ffd04b">
        <el-menu-item index="/front/home">
          <i class="el-icon-s-home"></i>
          <span slot="title">home</span>
        </el-menu-item>
        <el-submenu index="2">
          <template slot="title">
            <i class="el-icon-location"></i>
            <span>Profile</span>
          </template>
          <el-menu-item-group>
            <el-menu-item index="/front/person">self-Profile</el-menu-item>
          </el-menu-item-group>

        </el-submenu>
        <el-menu-item index="/front/question">
          <i class="el-icon-menu"></i>
          <span slot="title">Question</span>
        </el-menu-item>
        <el-menu-item index="/front/result">
          <i class="el-icon-news"></i>
          <span slot="title">Results</span>
        </el-menu-item>
        <el-menu-item index="/front/feedback">
          <i class="el-icon-message"></i>
          <span slot="title">Feedback</span>
        </el-menu-item>
        <el-menu-item index="/front/test">
          <i class="el-icon-message"></i>
          <span slot="title">Test</span>
        </el-menu-item>
      </el-menu>

    </el-aside>
    <el-main>
      <router-view/>
    </el-main>
  </el-container>

  </div>
</template>

<script>

export default {
  name: "FrontView",
  methods:{
    loginOut(){
      localStorage.removeItem('token')
      this.$message({
        message: '成功退出',
        type: 'success'
      });
      this.$router.push({path:"/login"})
    }
  }
}
</script>


<style>
.el-menu{
  border-right:none!important;
}
</style>