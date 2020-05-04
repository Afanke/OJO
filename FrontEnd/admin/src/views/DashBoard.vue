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
                    {{totalUsers}}
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
                    {{totalProblems}}
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
                    {{todaySubmissions}}
                  </div>
                  <div class="label-small-text">Today Submissions</div>
                </div>
              </el-row>
            </el-col>
            <el-col :span="6">
              <el-row class="label">
                <div class="label-left" style="background-color:#ff9c6e;">
                  <i class="el-icon-trophy" style="color:#fff;"></i>
                </div>
                <div class="label-right">
                  <div class="label-big-text" style="color:#ff9c6e;">
                    {{recentContests}}
                  </div>
                  <div class="label-small-text">Recent Contests</div>
                </div>
              </el-row>
            </el-col>
          </el-row>
          <el-row style="margin-top:20px" class="card">
            <el-row style="height:60px;line-height:60px">
              <span style="font-size:20px;margin-left:20px">Submissions</span>
              <el-radio-group v-model="range" style="float:right;margin-right:30px;margin-top:10px" size="medium"
                @change="handleChange">
                <el-radio-button label="Today"></el-radio-button>
                <el-radio-button label="Week"></el-radio-button>
                <el-radio-button label="Month"></el-radio-button>
              </el-radio-group>
            </el-row>
            <el-row style="height:1px;float:top;border-top:1px solid rgb(233, 233, 235);"></el-row>
            <el-row style="margin-top:20px;margin-left:5%;margin-right:5%">
                <ve-line :data="chartData" :settings="chartSettings"></ve-line>
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
        range: "Today",
        totalUsers: "--",
        totalProblems: "--",
        todaySubmissions: "--",
        recentContests: "--",
        chartSettings: {
          labelMap: {
            'practice': 'Practice',
            'contest': 'Contest'
          },
          legendName: {
            'practice': 'Practice',
            'contest': 'Contest'
          },
          xAxisType: 'time',
          area: true
        },
        chartData: {
          columns: ['hour', 'practice', 'contest'],
          rows: []
        },
      }
    },
    created() {
      this.$bus.emit("changeHeader", "1")
      this.show = false
    },
    mounted() {
      this.show = true
      this.getData()
      this.getTodayCount()
    },
    methods: {
      async getData() {
        try {
          const {
            data: res0
          } = await this.$http.post('/admin/user/getCount',{});
          if (res0.error) {
            this.$message.error(res0.error)
            return
          }
          this.totalUsers=res0.data
          const {
            data: res1
          } = await this.$http.post('/admin/problem/getCount', {});
          if (res1.error) {
            this.$message.error(res1.error)
            return
          }
          this.totalProblems=res1.data 
          const {
            data: res2
          } = await this.$http.post('/admin/contest/getRecentCount', {});
          if (res2.error) {
            this.$message.error(res2.error)
            return
          }
          this.recentContests=res2.data 
        } catch (err) {
          console.log(err);
          alert(err)
        }
      },
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
        this.chartData.columns = ['hour', 'practice', 'contest']
        this.chartData.rows = []
        let now = new Date();
        let s = "" + now.getFullYear() + "-" + (now.getMonth() + 1) + "-" + now.getDate() + " "
        for (let i = 0; i < 24; i++) {
          this.chartData.rows.push({
            'hour': s + (i < 10 ? "0" + i : "" + i) + ":00",
            'practice': 0,
            'contest': 0,
          })
        }
        let total=0
        if (pct) {
          for (let i = 0; i < pct.length; i++) {
            this.chartData.rows[Number(pct[i].hour)].practice = pct[i].count
            total+=pct[i].count
          }
        }
        if (cts) {
          for (let i = 0; i < cts.length; i++) {
            this.chartData.rows[Number(cts[i].hour)].contest = cts[i].count
            total+=cts[i].count
          }
        }
        this.todaySubmissions=total
      },
      async getWeekCount() {
        try {
          const {
            data: res0
          } = await this.$http.post('/admin/practice/getWeekCount');
          if (res0.error) {
            this.$message.error(res0.error)
            return
          }
          const {
            data: res1
          } = await this.$http.post('/admin/contest/getWeekCount');
          if (res1.error) {
            this.$message.error(res1.error)
            return
          }
          console.log(res0.data)
          this.freshWeekCount(res0.data, res1.data)
        } catch (err) {
          console.log(err);
          alert(err)
        }
      },
      freshWeekCount(pct, cts) {
        this.chartData.columns = ['day', 'practice', 'contest']
        this.chartData.rows = []
        let now = new Date(pct.today)
        let pctn = 0
        let ctsn = 0
        let tempPctDate = null
        let tempCtsDate = null
        if (pct.dayCount) {
          tempPctDate = new Date(pct.dayCount[pctn].day).getTime()
        }
        if (cts.dayCount) {
          tempCtsDate = new Date(cts.dayCount[ctsn].day).getTime()
        }
        for (let i = 7; i >= 0; i--) {
          let d = new Date(now.getTime() - i * 24 * 60 * 60 * 1000)
          let pctCount = 0
          let ctsCount = 0
          if (tempPctDate && d.getTime() === tempPctDate) {
            pctCount = pct.dayCount[pctn].count
            pctn++
            if (pctn < pct.dayCount.length) {
              tempPctDate = new Date(pct.dayCount[pctn].day).getTime()
            }
          }
          if (tempCtsDate && d.getTime() === tempCtsDate) {
            ctsCount = cts.dayCount[ctsn].count
            ctsn++
            if (ctsn < cts.dayCount.length) {
              tempCtsDate = new Date(cts.dayCount[ctsn].day).getTime()
            }
          }
          this.chartData.rows.push({
            'day': d.getFullYear() + "-" + (d.getMonth() + 1) + "-" + d.getDate(),
            'practice': pctCount,
            'contest': ctsCount,
          })
        }
      },
      async getMonthCount() {
        try {
          const {
            data: res0
          } = await this.$http.post('/admin/practice/getMonthCount');
          if (res0.error) {
            this.$message.error(res0.error)
            return
          }
          const {
            data: res1
          } = await this.$http.post('/admin/contest/getMonthCount');
          if (res1.error) {
            this.$message.error(res1.error)
            return
          }
          console.log(res0.data)
          this.freshMonthCount(res0.data, res1.data)
        } catch (err) {
          console.log(err);
          alert(err)
        }
      },
      freshMonthCount(pct, cts) {
        this.chartData.columns = ['day', 'practice', 'contest']
        this.chartData.rows = []
        let now = new Date(pct.today)
        let pctn = 0
        let ctsn = 0
        let tempPctDate = null
        let tempCtsDate = null
        if (pct.dayCount) {
          tempPctDate = new Date(pct.dayCount[pctn].day).getTime()
        }
        if (cts.dayCount) {
          tempCtsDate = new Date(cts.dayCount[ctsn].day).getTime()
        }
        for (let i = 30; i >= 0; i--) {
          let d = new Date(now.getTime() - i * 24 * 60 * 60 * 1000)
          let pctCount = 0
          let ctsCount = 0
          if (tempPctDate && d.getTime() === tempPctDate) {
            pctCount = pct.dayCount[pctn].count
            pctn++
            if (pctn < pct.dayCount.length) {
              tempPctDate = new Date(pct.dayCount[pctn].day).getTime()
            }
          }
          if (tempCtsDate && d.getTime() === tempCtsDate) {
            ctsCount = cts.dayCount[ctsn].count
            ctsn++
            if (ctsn < cts.dayCount.length) {
              tempCtsDate = new Date(cts.dayCount[ctsn].day).getTime()
            }
          }
          this.chartData.rows.push({
            'day': d.getFullYear() + "-" + (d.getMonth() + 1) + "-" + d.getDate(),
            'practice': pctCount,
            'contest': ctsCount,
          })
        }
      },
      handleChange(val) {
        switch (val) {
          case "Today":
            this.getTodayCount()
            break
          case "Week":
            this.getWeekCount()
            break
          case "Month":
            this.getMonthCount()
            break
        }
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