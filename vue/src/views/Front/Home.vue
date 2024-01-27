<script>
export default {
  name: "FrontView"
}
</script>

<template>
  <div>
  <el-container>
    <el-header style="background-color:#687178">
      <img src="../../assets/logo.png" style="width:45px;postion:relative;top:10px">
      <span style="font-size:20px;margin-left:15px;color:white">寝室智能匹配系统</span>
    </el-header>
  </el-container>

  <el-container>
    <el-aside style="overflow: hidden;min-height:100vh;background-color:#545c64">
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
          <span slot="title">系统首页</span>
        </el-menu-item>
        <el-submenu index="2">
          <template slot="title">
            <i class="el-icon-location"></i>
            <span>个人信息</span>
          </template>
          <el-menu-item-group>
            <template slot="title">分组一</template>
            <el-menu-item index="/front/person">查看个人信息</el-menu-item>
          </el-menu-item-group>

        </el-submenu>
        <el-menu-item index="/front/question">
          <i class="el-icon-menu"></i>
          <span slot="title">调查问卷</span>
        </el-menu-item>
      </el-menu>

    </el-aside>
    <el-main>
      <router-view/>
    </el-main>
  </el-container>

  </div>
</template>

<style>
.el-menu{
  border-right:none!important;
}
</style>