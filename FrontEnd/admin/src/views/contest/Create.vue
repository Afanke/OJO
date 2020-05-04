<template>
  <div>
    <transition name="slide-fade">
      <div class="center-box" v-if="show">
        <el-page-header style="height:60px;line-height:60px;margin-left:20px" title="Back" @back="goBack"
          content="Create Contest">
        </el-page-header>
        <el-row style="height:1px;float:top;border-top:1px solid rgb(233, 233, 235);">
        </el-row>
        <el-row style="margin-left:20px;margin-right:20px">
          <el-row :gutter="20" style="margin-top:20px">
            <el-col :span="24">
              <span style="color:red">*</span>
              <span>&nbsp;Title</span>
              <el-input v-model="title" placeholder="Title" style="margin-top:20px"></el-input>
            </el-col>
          </el-row>
          <el-row style="margin-top:30px">
            <span style="color:red">*</span>
            <span>&nbsp;Description</span>
            <editor style="margin-top:20px" v-model="description"></editor>
          </el-row>
          <el-row :gutter="20" style="margin-top:40px">
            <el-col :span="8">
              <span style="color:red">*</span>
              <span>&nbsp;Start Time</span>
              <el-row class="small-element">
                <el-date-picker v-model="startTime" type="datetime" placeholder="Start Time">
                </el-date-picker>
              </el-row>

            </el-col>
            <el-col :span="8">
              <span style="color:red">*</span>
              <span>&nbsp;End Time</span>
              <el-row class="small-element">
                <el-date-picker v-model="endTime" type="datetime" placeholder="End Time">
                </el-date-picker>
              </el-row>
            </el-col>
            <el-col :span="8">
              <span style="color:red">*</span>
              <span>&nbsp;Password</span>
              <el-row class="small-element">
                <el-input placeholder="Password" v-model="password" show-password clearable></el-input>
              </el-row>
            </el-col>
          </el-row>
          <el-row :gutter="20" style="margin-top:40px">
            <el-col :span="2" style="padding-top:3px">
              <span style="color:red">*</span>
              <span>&nbsp;Visible</span>
              <el-row style="margin-top:29px;margin-left:10px">
                <el-switch v-model="visible" active-color="#13ce66" inactive-color="#ff4949">
                </el-switch>
              </el-row>
            </el-col>
            <el-col :span="5" style="padding-top:3px">
              <span style="color:red">*</span>
              <span>&nbsp;Rule</span>
              <el-row class="small-element" style="margin-top:30px;margin-left:10px">
                <el-radio v-model="rule" label="OI">OI</el-radio>
                <el-radio v-model="rule" label="ACM">ACM</el-radio>
              </el-row>
            </el-col>
            <el-col :span="8" :offset="1">
              <span style="color:#ffffff">*</span>
              <span v-if="rule==='OI'">&nbsp;Punish Score</span>
              <span v-if="rule==='ACM'">&nbsp;Punish Time (second)</span>
              <el-row class="small-element">
                <el-input-number v-model="punish" controls-position="right" :min="0"></el-input-number>
              </el-row>
            </el-col>
            <el-col :span="8">
              <span style="color:#ffffff">*</span>
              <span>&nbsp;Submit Limit</span>
              <el-row class="small-element">
                <el-input-number v-model="submitLimit" controls-position="right" :min="0"></el-input-number>
              </el-row>
            </el-col>
          </el-row>
          <el-row style="margin-top:30px">
            <span style="color:#ffffff">*</span>
            <span>&nbsp;Allowed IP Ranges</span>
            <div style="height:10px"></div>
            <el-row style="width:40%;margin-top:15px;margin-left:10px;display:flex" v-for="(item,index) in IPRange"
              :key="index">
              <el-input style="flex:1" placeholder="CIDR Network" v-model="item.address" clearable>
              </el-input>
              <el-button plain style="float:right;margin-left:11px" icon="el-icon-plus" @click="addIPRange"></el-button>
              <el-button plain style="float:right" icon="el-icon-minus" @click="deleteIPRange(index)"></el-button>
            </el-row>
          </el-row>
          <el-row style="margin-top:30px;margin-bottom:20px;">
            <div style="text-align: center;">
              <el-button type="primary" style="width:200px" @click="save">Save</el-button>
            </div>
          </el-row>
        </el-row>
      </div>
    </transition>
  </div>
</template>
<script>
     import Editor from '@/components/Editor'
  export default {
    data() {
      return {
        show: false,
        input: "",
        title: '',
        password: "",
        startTime: new Date(),
        endTime: new Date(),
        description: '',
        visible: false,
        rule: "OI",
        submitLimit: 0,
        punish: 0,
        IPRange: [{
          address: ""
        }]
      }
    },
    created() {
      this.$bus.emit("changeHeader", "4-2")
      this.show = false
    },
    mounted() {
      this.show = true
    },
    methods: {
      check() {
        if (this.title === "") {
          this.$message.error("title is required")
          return false
        }
        if (this.description === "") {
          this.$message.error("description is required")
          return false
        }
        if (this.endTime.getTime()<=this.startTime.getTime()) {
          this.$message.error("End time cannot be earlier than or equal to start time")
          return false
        }
        for (let i = 0; i < this.IPRange.length; i++) {
          if (this.IPRange[i].address === "") {
            continue
          }
          if (/^\d{1,3}[.]\d{1,3}[.]\d{1,3}[.]\d{1,3}$/.test(this.IPRange[i].address)) {
            console.log(this.IPRange[i].address, 1)
            continue
          }
          if (/^\d{1,3}[.]\d{1,3}[.]\d{1,3}[.]\d{1,3}[/]\d{1,2}$/.test(this.IPRange[i].address)) {
            let mask = Number(this.IPRange[i].address.split("/")[1])
            if (mask >= 0 && mask <= 32) {
              console.log(this.IPRange[i].address, 2)
              continue
            }
          }
          this.$message.error("ip address " + (i + 1) + " is not legal")
          return false
        }
        return true
      },
      goBack() {
        this.$router.go(-1)
      },
      async save() {
        if (!this.check()) {
          return
        }
        console.log(this.dateFormat("YYYY-mm-dd HH:MM:SS", this.startTime))
        console.log(this.endTime)
        let obj = {
          title: this.title,
          password: this.password,
          startTime: this.dateFormat("YYYY-mm-dd HH:MM:SS", this.startTime),
          endTime: this.dateFormat("YYYY-mm-dd HH:MM:SS", this.endTime),
          description: this.description,
          visible: this.visible,
          rule: this.rule,
          submitLimit: this.submitLimit,
          punish: this.punish,
          IPLimit: this.IPRange
        }
        try {
          const {
            data: res
          } = await this.$http.post('/admin/contest/addContest', obj);
          if (res.error) {
            this.$message.error(res.error)
            return
          }
          this.$message({
            message: res.data,
            type: 'success'
          });
          this.$router.push("/contest")
        } catch (err) {
          console.log(err);
          alert(err)
        }
      },
      addIPRange() {
        this.IPRange.push({
          address: ""
        })
      },
      dateFormat(fmt, date) {
        let ret;
        const opt = {
          "Y+": date.getFullYear().toString(), // 年
          "m+": (date.getMonth() + 1).toString(), // 月
          "d+": date.getDate().toString(), // 日
          "H+": date.getHours().toString(), // 时
          "M+": date.getMinutes().toString(), // 分
          "S+": date.getSeconds().toString() // 秒
          // 有其他格式化字符需求可以继续添加，必须转化成字符串
        };
        for (let k in opt) {
          ret = new RegExp("(" + k + ")").exec(fmt);
          if (ret) {
            fmt = fmt.replace(ret[1], (ret[1].length == 1) ? (opt[k]) : (opt[k].padStart(ret[1].length, "0")))
          };
        };
        return fmt;
      },
      deleteIPRange(index) {
        if (this.IPRange.length === 1) {
          this.$message.error("At least one is needed. If you don't want any ip limit, just keep it empty")
          return
        }
        this.IPRange.splice(index, 1);
      },
    },
    components: {
      editor: Editor
    }
  };
</script>

<style scoped>
  .center-box {
    width: 100%;
    background-color: #ffffff;
    border-radius: 10px;
    box-shadow: 0 2px 12px 0 rgba(0, 0, 0, 0.1);
  }

  .center-box>>>.w-e-text {
    overflow: visible !important
  }

  #add-button {
    border: 1px solid rgb(233, 233, 235);
    border-radius: 5px;
    margin-top: 30px;
    text-align: center;
    height: 40px;
    line-height: 40px;
  }

  #add-button:hover {
    background-color: #FAFDFF;
    cursor: pointer;
  }

  .small-element {
    margin-top: 20px;
    margin-left: 10px
  }

  .slide-fade-enter-active {
    transition: all 0.8s ease;
  }

  .slide-fade-leave-active {
    transition: all .8s cubic-bezier(1.0, 0.5, 0.8, 1.0);
  }

  .slide-fade-enter,
  .slide-fade-leave-to

  /* .slide-fade-leave-active for below version 2.1.8 */
    {
    transform: translateY(40px);
    opacity: 0;
  }
</style>