<template>
  <div class="box" v-if="show" v-loading="loading">
    <el-row style="height:60px">
      <span style="float:left;font-size:20px;margin-left:20px;margin-top:15px">Rank</span>
    </el-row>
    <div style="width:85%;float:left;margin-left:7.5%">
      <ve-line style="width:100%" :settings="chartSettings" :legend-visible="true" :extend="chartExtend">
      </ve-line>
    </div>

    <div style="width:100%">
      <el-table :data="tableData" style="width: 100%" v-loading="rankLoading" :cell-style="cellStyle" size="mini"  >
        <el-table-column type="index" label="#" min-width="10" align="center" :index="indexMethod">
        </el-table-column>
        <el-table-column prop="username" label="Username" align="center" min-width="10">
        </el-table-column>
        <el-table-column label="AC/Total" align="center" min-width="10">
          <template slot-scope="scope">
            <span>{{scope.row.ac}}&nbsp;/&nbsp;{{scope.row.total}}</span>
          </template>
        </el-table-column>
        <el-table-column  label="TotalTime" align="center" min-width="10">
          <template slot-scope="scope">
            <span v-if="scope.row.totalTime">{{countDuration(scope.row.totalTime)}}</span>
          </template>
        </el-table-column>
        <el-table-column :label="item.title" min-width="10" v-bind:key="i" v-for="(item, i) in problemList"
          align="center">
          <template slot="header">
            <el-link :underline="false" @click="goProblem(item.id)">{{item.title}}</el-link>
          </template>
          <template slot-scope="scope">
            <span v-if="scope.row.ACMDetail[i]">{{countDuration(scope.row.ACMDetail[i].lastSubmitTime)}}</span><br>
            <span v-if="scope.row.ACMDetail[i] && scope.row.ACMDetail[i].total-scope.row.ACMDetail[i].ac">(-{{scope.row.ACMDetail[i].total-scope.row.ACMDetail[i].ac}})</span>
          </template>
        </el-table-column>
      </el-table>
      <el-pagination style="float:right;margin-top:20px" background layout="prev, pager, next" :page-size="10"
        @current-change="handlePageChange" :current-page="page" :total="count">
      </el-pagination>
    </div>
  </div>
</template>
<script>
  import VeLine from "v-charts/lib/line.common.js";
  import "echarts/lib/component/markPoint";
  import "echarts/lib/component/title";
  export default {
    data() {
      return {
        show: false,
        res: {},
        res1: {},
        tableData: [],
        problemList: [],
        startTime: "",
        page: 1,
        count: 5,
        loading: true,
        rankLoading: true,
        top10: {},
        chartExtend: {
          title: {
            text: "TOP10",
            left: "5%",
            textAlign: "center"
          },
          legend: {
            data: []
          },
          tooltip: {
            trigger: "axis",
            axisPointer: {
              type: "cross",
              animation: true,
              snap: true
            }
          },
          toolbox: {
            show: true,
            feature: {
              dataZoom: {
                //  yAxisIndex: 'none'
              },
              dataView: {
                readOnly: false
              },
              //magicType: {type: ['line']},
              restore: {},
              saveAsImage: {}
            }
          },
          grid: {
            // top: 70,
            // bottom: 50
          },
          xAxis: {
            type: "time",
            scale: true,
            boundaryGap: ["20%", "50%"]
          },
          yAxis: {
            type: "value",
            axisPointer: {
              snap: true,
              type: "none"
            },
            scale: true
          },
          series: []
        },
        chartData: {
          columns: ["date", "11"],
          // columns: ['date', '1', '2'],
          rows: [
            // { date: '2020-01-22 06:01:00', '1': 0 },
            // { date: '2020-01-23 21:01:00', '1': 1 },
            // { date: '2020-01-24 23:01:00', '1': 3 },
            // { date: 12500, '1': 10  },
            // { date: 87980007, '1': 12 },
            // { date: 645465800, '1': 31 },
            // { date: 125000, '2': 0 },
            // { date: 8798000, '2': 10 },
            // { date: 6454658000, '2': 30 }
            // { date: '2020-01-23', '2': 1 },
            // { date: '2020-01-24', '3': 2 },
            // { date: '2020-01-26', '2': 4 },
            // { date: '2020-01-27', '3': 5 },
          ]
        },
        chartSettings: {
          xAxisType: "time"
          // dimension:['1','2'],
          // metrics:['date']
          // itemStyle: {
          //   color: '#409EFF'
          // },
          // labelMap: {
          //   totalScore: 'Score'
          // },
          // legendName: {
          //   totalScore: 'Score'
          // }
        }
      };
    },
    created() {
      this.$bus.emit("changeHeader", "3");
      // this.practiseListLoading=true
      this.show = false;
    },
    async mounted() {
      this.show = true;
      try {
        const {
          data: res0
        } = await this.$http.post("/contest/getACMTop10", {
          id: Number(this.$route.query.id)
        });
        if (res0.error) {
          this.$message.error(res0.error);
          return;
        }
        this.top10 = res0.data;
        const {
          data: res1
        } = await this.$http.post(
          "/contest/getAllProblemName", {
            id: Number(this.$route.query.id)
          }
        );
        if (res1.error) {
          this.$message.error(res1.error);
          return;
        }
        this.problemList = res1.data;
        const {
          data: res2
        } = await this.$http.post("/contest/getTime", {
          id: Number(this.$route.query.id)
        });
        if (res2.error) {
          this.$message.error(res2.error);
          return;
        }
        this.startTime = res2.data.startTime;
        this.startTime = new Date(this.startTime.replace(/-/g, "/"));
        this.prepareChart();
        const {
          data: res3
        } = await this.$http.post("/contest/getACMRankCount", {
          id: Number(this.$route.query.id)
        });
        if (res3.error) {
          this.$message.error(res3.error);
          return;
        }
        this.count = Number(res3.data);
        this.loading = false;
        this.getRank();
      } catch (err) {
        console.log(err);
        alert(err);
      }
    },
    methods: {
      cellStyle(tb) {
        // console.log(tb)
        if (tb.columnIndex > 3 ) {
          // console.log(tb.row.ACMDetail[tb.columnIndex-4])
          if(!tb.row.ACMDetail[tb.columnIndex-4]){
            return 'height:57px'
          }
          if(tb.row.ACMDetail[tb.columnIndex-4].firstAC){
            return 'background-color:#67C23A;color:#fff;height:57px'
          }
          if(tb.row.ACMDetail[tb.columnIndex-4].ac){
            return 'background-color:rgb(225, 243, 216);height:57px'
          }
          if(!tb.row.ACMDetail[tb.columnIndex-4].ac && tb.row.ACMDetail[tb.columnIndex-4].total){
            return 'background-color:rgb(253, 226, 226);height:57px'
          }
          // tb.row.ACMDetail[tb.columnIndex-4].ac
          // tb.row.ACMDetail[tb.columnIndex-4].total
          // return "font-weight:bold";
        } 
      },
      prepareChart() {
        if(!this.top10){
          return
        }
        for (let i = 0; i < this.top10.length; i++) {
          this.top10[i].ACMDetail = this.top10[i].ACMDetail.sort(function (a, b) {
            return a.lastSubmitTime - b.lastSubmitTime;
          });
          let name = this.top10[i].username;
          this.chartExtend.legend.data.push(name);
          let ac = 0;
          let obj = {
            name: name,
            type: "line",
            smooth: true,
            showSymbol: true,
            data: [
              [this.startTime.getTime(), 0]
            ]
          };
          for (let j = 0; j < this.top10[i].ACMDetail.length; j++) {
            if (this.top10[i].ACMDetail[j].ac) {
              ac++;
              obj.data.push([
                this.countTime(this.top10[i].ACMDetail[j].lastSubmitTime),
                ac
              ]);
            }
          }
          this.chartExtend.series.push(obj);
        }
      },
      prepareTable(data){
        if(!data){
          return
        }
        this.tableData=data
        for(let i=0;i<this.tableData.length;i++){
          for (let j=0;j<this.problemList.length;j++){
            if (this.tableData[i].ACMDetail[j].pid!==this.problemList[j].id){
                  this.tableData[i].ACMDetail.splice(j, 0, {
                    ac: false,
                    cid: 0,
                    firstAC: false,
                    id: 0,
                    lastSubmitTime: -1,
                    pid: 0,
                    total: 0,
                    uid: 0,
                  })
            }
          }
        }
        console.log(this.tableData)
      },
      async getRank() {
        try {
          this.rankLoading = true;
          this.params_init();
          const {
            data: res
          } = await this.$http.post("/contest/getACMRank", {
            cid: Number(this.$route.query.id),
            page: this.page
          });
          if (res.error) {
            this.$message.error(res.error);
            return;
          }
          this.prepareTable(res.data)
          this.rankLoading = false;
        } catch (err) {
          console.log(err);
          alert(err);
        }
      },
      goProblem(val) {
        this.$router.push({
          path: "/contest/answer",
          query: {
            cid: this.$route.query.id,
            pid: val
          }
        });
      },
      countDuration(val) {
        if (val===-1){
          return ""
        }
        var d=val
        var hour = Math.floor(d / 3600);
        d = d % 3600;
        if (hour < 10) {
          hour = "0" + hour;
        }
        var minute = Math.floor(d / 60);
        if (minute < 10) {
          minute = "0" + minute;
        }
        d = d % 60;
        var second = d;
        if (second < 10) {
          second = "0" + second;
        }
        return hour + ":" + minute + ":" + second;
      },
      countTime(val) {
        return this.startTime.getTime() + val * 1000;
      },
      handlePageChange(val) {
        let obj = this.params_query();
        obj.page = Number(val);
        this.fresh(obj);
      },
      fresh(obj) {
        this.$router.push({
          path: "/contest/detail",
          query: obj
        });
      },
      params_query() {
        let obj = {};
        if (this.$route.query.page) {
          obj.page = Number(this.$route.query.page);
        }
        if (this.$route.query.c) {
          obj.c = this.$route.query.c;
        }
        if (this.$route.query.id) {
          obj.id = Number(this.$route.query.id);
        }
        return obj;
      },
      params_init() {
        if (this.$route.query.page) {
          this.page = Number(this.$route.query.page);
        } else {
          this.page = 1;
        }
      },
      indexMethod(index) {
        return 1 + index + (this.page - 1) * 10;
      }
    },
    components: {
      VeLine
    },
    watch: {
      $route() {
        this.getRank();
      }
    }
  };
</script>

<style scoped>
  .box {
    width: 100%;
    background-color: #ffffff;
    border-radius: 10px;
    /* margin-bottom: 20px; */
    box-shadow: 0 2px 12px 0 rgba(0, 0, 0, 0.1);
  }
</style>