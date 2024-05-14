<template>
  <div class="main-content">

    <div style="display: flex; margin-bottom: 10px">
      <div style="flex: 1">
        <el-button type="primary" class="btn-green" size="medium" @click="handLeAdd">创建新的问卷</el-button>
      </div>
      <div>
        <el-input v-model="name" style="width: 300px; margin-right: 10px" size="medium" placeholder="请输入名称搜索问卷" clearable></el-input>
        <el-button type="primary" size="medium" class="btn-green" @click="load(1)">搜 索</el-button>
      </div>
    </div>

    <div class="pagination">
      <el-pagination>
          background
          @current-change="handleCurrentChange'
          :current-page="pageNum"
          :page-sizes="[5,10,20]
          :page-size="pageSize"
          layout="total, prev, pager,next"
          :total="total">
      </el-pagination>
    </div>

    <div style="margin-bottom: 10px">
      <div class="card" v-for="item in tableData" :key="item.id" style="margin-bottom: 5px">
        <div style="display: flex;">
          <div style="flex: 1; width: 0;">
            <div style="display: flex; align-items: center; margin-bottom: 10px">
              <el-tag v-if="item.saved === '否'" type="warning">未发布</el-tag>
              <el-tag v-if="item.saved === '是'" type="success">已发布</el-tag>
              <span style="font-size: 20px; margin-left: 10px">{{ item.name }}</span>
            </div>

            <div>
              <el-button type="text" style="color: #2a60c9; font-size: 16px"><i class="el-icon-edit"></i> 编辑</el-button>
              <el-button type="text" style="color: #2a60c9; font-size: 16px"><i class="el-icon-setting"></i> 设计</el-button>
              <el-button type="text" style="color: #2a60c9; font-size: 16px"><i class="el-icon-document-copy"></i> 复制</el-button>
              <el-button type="text" style="color: #2a60c9; font-size: 16px"><i class="el-icon-share"></i> 分享</el-button>
              <el-button type="text" style="color: #fc4b4b; font-size: 16px"><i class="el-icon-delete"></i> 删除</el-button>
              <el-button type="text" style="color: orange; font-size: 16px"><i class="el-icon-s-marketing"></i> 数据统计</el-button>
            </div>
          </div>

          <div style="width: 100px">
            <img :src="item.img" alt="" style="width: 100%; border-radius: 5px; display: block">
          </div>
        </div>
      </div>
    </div>

    <div>
      <el-pagination
          background
          @current-change="handleCurrentChange"
          :current-page="pageNum"
          :page-size="pageSize"
          layout="total, prev, pager, next"
          :total="total">
      </el-pagination>
    </div>

  </div>
</template>

<script>
export default {
  name: "FPages",
  data() {
    return {
      tableData:[],  // 所有的数据
      pageNum: 1,   // 当前的页码
      pageSize: 4,  // 每页显示的个数
      total: 0,
      name: null
    }
  },
  created() {
    this.load(1)
  },
  methods: {
    handLeAdd()//新增数据
    { this.form=0//新增数据的时候清空数据
      this.fromVisible = true//打开弹窗
    },
    // handleEdit(row)
    // {//编辑数据this.form = JsoN.parse（JsoN.stringify（row））// 给form对象赋值注意要深拷贝数据
    //   this.fromVisible = true//打开弹窗
    // },
    load(pageNum) {  // 分页查询
      if (pageNum) this.pageNum = pageNum
      this.$request.get('/pages/selectPage', {
        params: {
          pageNum: this.pageNum,
          pageSize: this.pageSize,
          name: this.name,
        }
      }).then(res => {
        this.tableData = res.data?.list
        this.total = res.data?.total
      })
    },
    reset() {
      this.name = null
      this.load(1)
    },
    handleCurrentChange(pageNum) {
      this.load(pageNum)
    },
  }
}
</script>

<style scoped>

</style>