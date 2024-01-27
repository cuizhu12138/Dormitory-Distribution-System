<template>

  <div>
    <div>
      <h1> 调查问卷页面</h1>
    </div>

    <div>
      <el-form ref="form" :model="form" label-width="200px" label-position="top" >
        <el-form-item label="性别">
          <el-select v-model="form.sex" placeholder="请选择性别">
            <el-option label="男" value="man"></el-option>
            <el-option label="女" value="woman"></el-option>
          </el-select>
        </el-form-item>

        <el-form-item label="专业">
          <el-input v-model="form.major"></el-input>
        </el-form-item>

        <el-form-item label="籍贯">
          <el-input v-model="form.home"></el-input>
        </el-form-item>

        <el-form-item label="你通常几点睡觉?">
          <el-radio-group v-model="form.sleepTime">
            <el-radio label="22:00以前"></el-radio>
            <el-radio label="22:00-23:00"></el-radio>
            <el-radio label="23:00-24:00"></el-radio>
            <el-radio label="24:00以后"></el-radio>
          </el-radio-group>
        </el-form-item>

        <el-form-item label="你通常几点起床?">
          <el-radio-group v-model="form.getupTime">
            <el-radio label="6:00以前"></el-radio>
            <el-radio label="6:00-7:00"></el-radio>
            <el-radio label="7:00-8:00"></el-radio>
            <el-radio label="8:00以后"></el-radio>
          </el-radio-group>
        </el-form-item>

        <el-form-item label="你对室友的作息时间有什么要求或期望?">
          <el-radio-group v-model="form.timeExpectation">
            <el-radio label="与我相同或相近"></el-radio>
            <el-radio label="不影响我的休息就行"></el-radio>
            <el-radio label="没有要求或期望"></el-radio>
          </el-radio-group>
        </el-form-item>

        <el-form-item label="你是否会在寝室学习或工作?">
          <el-radio-group v-model="form.dormitoryActivity">
            <el-radio label="是，经常"></el-radio>
            <el-radio label="偶尔"></el-radio>
            <el-radio label="否，从不"></el-radio>
          </el-radio-group>
        </el-form-item>

        <el-form-item label="你是否会在寝室娱乐或放松?">
          <el-radio-group v-model="form.dormitoryEntertainment">
            <el-radio label="是，经常"></el-radio>
            <el-radio label="偶尔"></el-radio>
            <el-radio label="否，从不"></el-radio>
          </el-radio-group>
        </el-form-item>

        <el-form-item label="你对寝室的整洁程度有什么要求或期望?">
          <el-radio-group v-model="form.dormitoryCleanExpectation">
            <el-radio label="非常干净"></el-radio>
            <el-radio label="比较干净"></el-radio>
            <el-radio label="一般"></el-radio>
            <el-radio label="不太在意"></el-radio>
          </el-radio-group>
        </el-form-item>

        <el-form-item label="你通常多久打扫一次寝室?">
          <el-radio-group v-model="form.dormitoryCleanFrequency">
            <el-radio label="每天"></el-radio>
            <el-radio label="每周"></el-radio>
            <el-radio label="每月"></el-radio>
            <el-radio label="不定期"></el-radio>
          </el-radio-group>
        </el-form-item>

        <el-form-item label="你是否会主动参与寝室的打扫工作?">
          <el-radio-group v-model="form.dormitoryCleanWork">
            <el-radio label="是，主动"></el-radio>
            <el-radio label="是，被提醒后"></el-radio>
            <el-radio label="否，不愿意"></el-radio>
          </el-radio-group>
        </el-form-item>

        <el-form-item label="你是否会在寝室吃东西或喝饮料?">
          <el-radio-group v-model="form.dormitoryFood">
            <el-radio label="是，经常"></el-radio>
            <el-radio label="是，偶尔"></el-radio>
            <el-radio label="否，从不"></el-radio>
          </el-radio-group>
        </el-form-item>

        <el-form-item label="你是否会及时处理寝室的垃圾?">
          <el-radio-group v-model="form.dormitoryRubbish">
            <el-radio label="是，每次"></el-radio>
            <el-radio label="是，大部分时候"></el-radio>
            <el-radio label="否，很少"></el-radio>
          </el-radio-group>
        </el-form-item>

        <el-form-item label="你是否会与室友共用一些物品，如卫生纸、洗衣液、电饭煲等?">
          <el-radio-group v-model="form.shareThings">
            <el-radio label="是，愿意"></el-radio>
            <el-radio label="是，但要征得对方同意"></el-radio>
            <el-radio label="否，不愿意"></el-radio>
          </el-radio-group>
        </el-form-item>

        <el-form-item label="你你是否会借用室友的一些物品，如书籍、衣服、电脑等?">
            <el-radio-group v-model="form.borrowThings">
              <el-radio label="是，愿意"></el-radio>
              <el-radio label="是，但要征得对方同意"></el-radio>
              <el-radio label="否，不愿意"></el-radio>
            </el-radio-group>
        </el-form-item>

        <el-form-item label="你是否会主动购买或补充寝室的一些物品，如卫生纸、洗衣液、电饭煲等?">
          <el-radio-group v-model="form.buyThings">
            <el-radio label="是，愿意"></el-radio>
            <el-radio label="是，但要与室友商量"></el-radio>
            <el-radio label="否，不愿意"></el-radio>
          </el-radio-group>
        </el-form-item>

        <el-form-item label="你是否会及时归还或清洗寝室的一些物品，如书籍、衣服、电饭煲等?">
          <el-radio-group v-model="form.returnThings">
            <el-radio label="是，每次"></el-radio>
            <el-radio label="是，大部分时候"></el-radio>
            <el-radio label="否，很少"></el-radio>
          </el-radio-group>
        </el-form-item>

        <el-form-item label="你的月生活费大概是多少?">
          <el-radio-group v-model="form.money">
            <el-radio label="500元以下"></el-radio>
            <el-radio label="500-1000元"></el-radio>
            <el-radio label="1000-1500元"></el-radio>
            <el-radio label="1500元以上"></el-radio>
          </el-radio-group>
        </el-form-item>


        <el-form-item label="你通常会在什么方面花钱?">
          <el-checkbox-group v-model="form.type">
            <el-checkbox label="饮食" name="type"></el-checkbox>
            <el-checkbox label="衣服" name="type"></el-checkbox>
            <el-checkbox label="娱乐" name="type"></el-checkbox>
            <el-checkbox label="学习" name="type"></el-checkbox>
            <el-checkbox label="其他" name="type"></el-checkbox>
          </el-checkbox-group>
        </el-form-item>

        <el-form-item label="你是否会与室友一起外出消费，如吃饭、逛街、看电影等?">
          <el-radio-group v-model="form.outside">
            <el-radio label="是，经常"></el-radio>
            <el-radio label="是，偶尔"></el-radio>
            <el-radio label="否，从不"></el-radio>
          </el-radio-group>
        </el-form-item>

        <el-form-item label="你是否会与室友一起共享消费，如拼单、拼车、拼房等?">
          <el-radio-group v-model="form.shareConsume">
            <el-radio label="是，经常"></el-radio>
            <el-radio label="是，偶尔"></el-radio>
            <el-radio label="否，从不"></el-radio>
          </el-radio-group>
        </el-form-item>

        <el-form-item label="你的兴趣爱好有哪些?">
          <el-checkbox-group v-model="form.hobby">
            <el-checkbox label="阅读" name="hobby"></el-checkbox>
            <el-checkbox label="写作" name="hobby"></el-checkbox>
            <el-checkbox label="绘画" name="hobby"></el-checkbox>
            <el-checkbox label="音乐" name="hobby"></el-checkbox>
            <el-checkbox label="电影" name="hobby"></el-checkbox>
            <el-checkbox label="体育" name="hobby"></el-checkbox>
            <el-checkbox label="游戏" name="hobby"></el-checkbox>
            <el-checkbox label="旅行" name="hobby"></el-checkbox>
            <el-checkbox label="其他" name="hobby"></el-checkbox>
          </el-checkbox-group>
        </el-form-item>

        <el-form-item label="你是否会在寝室进行你的兴趣爱好?">
          <el-radio-group v-model="form.dormitoryHobby">
            <el-radio label="是，经常"></el-radio>
            <el-radio label="是，偶尔"></el-radio>
            <el-radio label="否，从不"></el-radio>
          </el-radio-group>
        </el-form-item>

        <el-form-item label="你是否会与室友分享你的兴趣爱好?">
          <el-radio-group v-model="form.shareHobby">
            <el-radio label="是，愿意"></el-radio>
            <el-radio label="是，但要看对方的反应"></el-radio>
            <el-radio label="否，不愿意"></el-radio>
          </el-radio-group>
        </el-form-item>

        <el-form-item label="你是否会尊重和支持室友的兴趣爱好?">
          <el-radio-group v-model="form.respectHobby">
            <el-radio label="是，总是"></el-radio>
            <el-radio label="是，但要看具体情况"></el-radio>
            <el-radio label="否，不会"></el-radio>
          </el-radio-group>
        </el-form-item>

        <el-form-item label="你是否会帮助室友拿快递或外卖?">
          <el-radio-group v-model="form.takeCourier">
            <el-radio label="是，愿意"></el-radio>
            <el-radio label="是，但要看具体情况"></el-radio>
            <el-radio label="否，不愿意"></el-radio>
          </el-radio-group>
        </el-form-item>

        <el-form-item label="你是否会与室友聊天或交流心事?">
          <el-radio-group v-model="form.communication">
            <el-radio label="是，愿意"></el-radio>
            <el-radio label="是，但要看具体情况"></el-radio>
            <el-radio label="否，不愿意"></el-radio>
          </el-radio-group>
        </el-form-item>

        <el-form-item label="你是否会与室友发生冲突或矛盾?">
          <el-radio-group v-model="form.conflict">
            <el-radio label="是，经常"></el-radio>
            <el-radio label="是，偶尔"></el-radio>
            <el-radio label="否，从不"></el-radio>
          </el-radio-group>
        </el-form-item>

        <el-form-item label="你是否会与室友协商或解决冲突或矛盾?">
          <el-radio-group v-model="form.mediation">
            <el-radio label="是，主动"></el-radio>
            <el-radio label="是，被动"></el-radio>
            <el-radio label="否，不会"></el-radio>
          </el-radio-group>
        </el-form-item>

        <el-form-item label="是否喝酒">
          <el-select v-model="form.drink" placeholder="请选择">
            <el-option label="是" value=true></el-option>
            <el-option label="否" value=false></el-option>
          </el-select>
        </el-form-item>

        <el-form-item label="是否抽烟">
          <el-select v-model="form.smoke" placeholder="请选择">
            <el-option label="是" value=true></el-option>
            <el-option label="否" value=false></el-option>
          </el-select>
        </el-form-item>

        <el-form-item>
          <el-button type="primary" @click="onSubmit">提交</el-button>
          <el-button>取消</el-button>
        </el-form-item>

      </el-form>
    </div>




  </div>

</template>

<script>
export default {
  name:'QuestionnaireView',

  data() {
    return {
      form: {
        sex:'',
        major:'',
        home:'',
        sleepTime:'',
        getupTime:'',
        timeExpectation:'',
        dormitoryActivity:'',
        dormitoryEntertainment:'',
        dormitoryCleanExpectation:'',
        dormitoryCleanWork:'',
        dormitoryFood:'',
        dormitoryRubbish:'',
        shareThings:'',
        borrowThings:'',
        buyThings:'',
        returnThings:'',
        money:'',
        type:[],
        outside:'',
        shareConsume:'',
        hobby:[],
        dormitoryHobby:'',
        shareHobby:'',
        respectHobby:'',
        takeCourier:'',
        communication:'',
        conflict:'',
        mediation:'',
        drink:'',
        smoke:'',
      },
    }
  },
  methods: {
    onSubmit() {
      console.log(this.form);
    }
  }


}

</script>



<style scoped>


</style>