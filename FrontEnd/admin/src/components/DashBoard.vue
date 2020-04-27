<template>
  <div>
    <transition name="slide-fade">
      <div v-if="show">
        <div class="content">
          <el-row :gutter="20" style="height:100px">
            <el-col :span="6">
              <el-row class="label">
                <div class="label-left" style="background-color:#69c0ff;">
                  <i class="el-icon-user-solid" style="color:#fff;"></i>
                </div>
                <div class="label-right">
                  <div class="label-big-text" style="color:#69c0ff;">
                    {{totalUser}}
                  </div>
                  <div class="label-small-text">Total Users</div>
                </div>
              </el-row>
            </el-col>
            <el-col :span="6">
              <el-row class="label">
                <div class="label-left" style="background-color:#b37feb;">
                  <i class="el-icon-s-grid" style="color:#fff;"></i>
                </div>
                <div class="label-right">
                  <div class="label-big-text" style="color:#b37feb;">
                    {{totalUser}}
                  </div>
                  <div class="label-small-text">Total Problems</div>
                </div>
              </el-row>
            </el-col>
            <el-col :span="6">
              <el-row class="label">
                <div class="label-left" style="background-color:#95de64;">
                  <i class="el-icon-s-order" style="color:#fff;"></i>
                </div>
                <div class="label-right">
                  <div class="label-big-text" style="color:#95de64;">
                    {{totalUser}}
                  </div>
                  <div class="label-small-text">Today Submissions</div>
                </div>
              </el-row>
            </el-col>
            <el-col :span="6">
              <el-row class="label">
                <div class="label-left" style="background-color:#ff9c6e;">
                  <i class="el-icon-user-solid" style="color:#fff;"></i>
                </div>
                <div class="label-right">
                  <div class="label-big-text" style="color:#ff9c6e;">
                    {{totalUser}}
                  </div>
                  <div class="label-small-text">Recent Contests</div>
                </div>
              </el-row>
            </el-col>
          </el-row>
          <el-row style="margin-top:20px" class="card">
            <el-row style="height:60px;line-height:60px">
              <span style="font-size:20px;margin-left:20px">Submissions</span>
              <el-radio-group v-model="practiceRange" style="float:right;margin-right:30px;margin-top:10px"
                size="medium">
                <el-radio-button label="Today"></el-radio-button>
                <el-radio-button label="Week"></el-radio-button>
                <el-radio-button label="Month"></el-radio-button>
              </el-radio-group>
            </el-row>
            <el-row style="height:1px;float:top;border-top:1px solid rgb(233, 233, 235);"></el-row>
            <el-row style="margin-top:20px">
              <el-col :span="22" :offset="1">
                <ve-line :data="todayCount" :settings="chartSettings"></ve-line>
              </el-col>
            </el-row>

          </el-row>
        </div>
      </div>

    </transition>
  </div>

</template>
<script>
  import VeLine from 'v-charts/lib/line.common'
  export default {
    data() {
      return {
        show: false,
        totalUser: 100,
        practiceRange: "",
        chartSettings: {
          xAxisType: 'time',
          area: true
        },
        todayCount: {
          columns: ['hour', 'practice', 'contest'],
          rows: []
        }
      }
    },
    created() {
      this.$bus.emit("changeHeader", "1")
      this.show = false
    },
    mounted() {
      this.show = true
      this.getTodayCount()
    },
    methods: {
      async getTodayCount() {
        try {
          const {
            data: res0
          } = await this.$http.post('/admin/practice/getTodayCount', {});
          if (res0.error) {
            this.$message.error(res0.error)
            return
          }
          const {
            data: res1
          } = await this.$http.post('/admin/contest/getTodayCount', {});
          if (res1.error) {
            this.$message.error(res1.error)
            return
          }
          this.freshTodayCount(res0.data, res1.data)
        } catch (err) {
          console.log(err);
          alert(err)
        }
      },
      freshTodayCount(pct, cts) {
        this.todayCount.rows = []
        let now=new Date();
        let s=""+now.getFullYear()+"-"+(now.getMonth()+1)+"-"+now.getDate()+" "
        for (let i = 0; i < 24; i++) {
          this.todayCount.rows.push({
            'hour': s+ (i < 10 ? "0" + i : "" + i) + ":00",
            'practice': 0,
            'contest': 0,
          })
        }
        console.log(pct)
        if (pct) {
          for (let i = 0; i < pct.length; i++) {
            this.todayCount.rows[Number(pct[i].hour)].practice=pct[i].count
          }
        }
        if(cts){
           for (let i = 0; i < cts.length; i++) {
            this.todayCount.rows[Number(cts[i].hour)].contest=cts[i].count
          }
        }
        console.log(this.todayCount)
      }
    },
    components: {
      VeLine
    }
  };
</script>

<style scoped>
  .label {
    display: flex;
    text-align: center;
    height: 88px;
    background-color: #fff;
    box-shadow: 0 2px 4px rgba(0, 0, 0, .12), 0 0 6px rgba(0, 0, 0, .04);
    border-radius: 4px;
    overflow: hidden;
  }

  .card {
    background-color: #fff;
    box-shadow: 0 2px 4px rgba(0, 0, 0, .12), 0 0 6px rgba(0, 0, 0, .04);
    border-radius: 4px;
  }

  .label-big-text {
    font-size: 25px;
    font-weight: 700;
    margin-top: 22px;
  }

  .label-small-text {
    font-size: 14px;
    font-weight: 300;
  }

  .label-left {
    line-height: 88px;
    width: 80px;
    font-size: 38px;
  }

  .label-right {
    flex: 1;
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