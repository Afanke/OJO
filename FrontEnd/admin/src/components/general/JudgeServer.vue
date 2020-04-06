<template>

  <div>
    <transition name="slide-fade">

      <div class="center-box" v-if="show">
        <el-row style="height:60px;line-height:60px;">
          <span style="font-size:20px;margin-left:20px">Judge Server</span>
        </el-row>
        <el-row style="height:1px;float:top;border-top:1px solid rgb(233, 233, 235);"></el-row>
        <el-row :gutter="20" style="margin-left:20px;margin-right:20px">
          <el-col :span="8" v-for="(item,index) in judgeServer" :key="index">
            <el-row class="card" :style="item.style">
              <el-row style="height:50px;line-height:50px;font-size:18px">
                <span style="margin-left:5%">{{item.name}}</span>
              </el-row>
              <el-row style="height:1px;float:top;border-top:1px solid rgb(233, 233, 235);"></el-row>
              <el-row style="height:30px;line-height:30px">
                <span style="margin-left:5%;">Id:</span><span style="float:right;margin-right:5%">{{item.id}}</span>
              </el-row>
              <el-row style="height:30px;line-height:30px">
                <span style="margin-left:5%">Weight:</span><span
                  style="float:right;margin-right:5%">{{item.weight}}</span>
              </el-row>
              <el-row style="height:30px;line-height:30px">
                <span style="margin-left:5%">Address:</span><span
                  style="float:right;margin-right:5%">{{item.address}}</span>
              </el-row>
              <el-row style="height:30px;line-height:30px">
                <span style="margin-left:5%">Port:</span><span style="float:right;margin-right:5%">{{item.port}}</span>
              </el-row>
               <el-row style="height:30px;line-height:30px">
                <span style="margin-left:5%">Enabled:</span><span style="float:right;margin-right:5%">{{item.enabled}}</span>
              </el-row>
            </el-row>
          </el-col>
          <el-col :span="8">
            <el-row class="create-card">

            </el-row>
          </el-col>
        </el-row>
        <el-row style="height:30px">

        </el-row>
      </div>
    </transition>
  </div>

</template>
<script>
  export default {
    data() {
      return {
        show: false,
        judgeServer: [],
        timeout: null
      }
    },
    beforeDestroy() {
      clearTimeout(this.timeout)
    },
    created() {
      this.$bus.emit("changeHeader", "2-2")
      this.show = false
    },
    mounted() {
      this.show = true
      this.getAllInfo()
    },
    methods: {
      async getAllInfo() {
        try {
          const {
            data: res
          } = await this.$http.get('/admin/jsp/getAllInfo');
          if (res.error) {
            this.$message.error(res.error)
          } else {
            for (let i = 0; i < res.data.length; i++) {
              if (res.data[i].status) {
                res.data[i].style = "background: rgb(240, 249, 235);"
              } else {
                res.data[i].style = "background: rgb(254, 240, 240);"
              }
            }
            this.judgeServer = res.data
          }
          this.timeout = setTimeout(
            this.getAllInfo
          , 5000);
        } catch (err) {
          console.log(err);
          alert(err)
        }
      }
    },
    components: {}
  };
</script>

<style scoped>
  .create-card {
    border: 1px solid rgb(233, 233, 235);
    margin-top: 30px;
    border-radius: 4px;
    box-shadow: 0 2px 12px 0 rgba(0, 0, 0, 0.1)
  }

  .card {

    border: 1px solid rgb(233, 233, 235);
    margin-top: 30px;
    border-radius: 4px;
    box-shadow: 0 2px 12px 0 rgba(0, 0, 0, 0.1)
  }

  .center-box {
    background-color: #ffffff;
    border-radius: 10px;
    /* min-height: 600px; */
    box-shadow: 0 2px 12px 0 rgba(0, 0, 0, 0.1);
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