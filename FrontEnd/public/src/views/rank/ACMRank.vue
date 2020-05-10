<template>
  <transition name="slide-fade">
    <div class="center-box" v-if="show">
      <div v-if="show" v-loading="loading">
        <div class="content">
          <el-row style="height:60px">
            <span style="float:left;font-size:20px;margin-left:20px;margin-top:15px">ACM Rank</span>
          </el-row>
          <div style="width:90%;float:left;margin-left:5%">
            <ve-histogram :data="chartData" style="width:100%" :settings="chartSettings" :extend="chartExtend"
              :mark-point="markPoint"></ve-histogram>
          </div>
          <el-table :data="tableData" style="width: 100%" v-loading="rankLoading">
            <el-table-column type="index" label="#" min-width="10" align="center" :index="indexMethod">
            </el-table-column>
            <el-table-column prop="username" label="User" align="center" min-width="10">
            </el-table-column>
            <el-table-column prop="signature" label="Signature" align="center" min-width="10">
            </el-table-column>
            <el-table-column prop="ac" label="AC" align="center" min-width="10">
            </el-table-column>
            <el-table-column prop="total" label="Total" align="center" min-width="10">
            </el-table-column>
            <el-table-column prop="rate" label="Rate" align="center" min-width="10">
            </el-table-column>
          </el-table>
        </div>
        <el-pagination style="float:right;margin-top:20px" background layout="prev, pager, next" :page-size="10"
          @current-change="handlePageChange" :current-page="page" :total="count">
        </el-pagination>
      </div>
    </div>
  </transition>
</template>
<script>
  import VeHistogram from 'v-charts/lib/histogram.common.js';
  import 'echarts/lib/component/markPoint';
  export default {
    data() {
      return {
        show: false,
        markPoint: {
          data: [{
            name: 'max',
            type: 'max'
          }]
        },
        tableData: [],
        page: 1,
        count: 5,
        loading: true,
        rankLoading: true,
        chartData: {
          columns: ['username', 'ac', 'total'],
          rows: []
        },
        chartSettings: {
          labelMap: {
            username: 'Username',
            ac: "AC",
            total: "Total"
          },
          legendName: {
            username: 'Username',
            ac: "AC",
            total: "Total"
          }
        },
        chartExtend: {
          color: ["#67C23A", "#E6A23C"],
          xAxis: {
            show: true,
            value: "value",
            axisLine: {
              show: true
            },
            axisTick: {
              show: true
            }
          },

        }
      };
    },
    created() {
      this.$bus.emit("changeHeader", "5-1");
      this.show = false;
    },
    async mounted() {
      this.show = true;
      try {
        const {
          data: res0
        } = await this.$http.post("/rank/getACMTop10", {
          id: Number(this.$route.query.id)
        });
        if (res0.error) {
          this.$message.error(res0.error);
          return;
        }
        this.chartData.rows = res0.data;
        const {
          data: res1
        } = await this.$http.post("/rank/getACMRankCount", {
          id: Number(this.$route.query.id)
        });
        if (res1.error) {
          this.$message.error(res1.error);
          return;
        }
        this.count = Number(res1.data);
        this.loading = false;
        this.getRank();
      } catch (err) {
        console.log(err);
        alert(err);
      }
    },
    methods: {
      async getRank() {
        try {
          this.rankLoading = true;
          this.paramsInit();
          const {
            data: res
          } = await this.$http.post("/rank/getACMRank", {
            page: this.page
          });
          if (res.error) {
            this.$message.error(res.error);
            return;
          }
          this.tableData = res.data
          this.tableData.forEach(e => {
            e.rate = this.getRate(e.ac, e.total)
          })
          this.rankLoading = false;
        } catch (err) {
          console.log(err);
          alert(err);
        }
      },
      handlePageChange(val) {
        let obj = this.paramsQuery();
        obj.page = Number(val);
        this.fresh(obj);
      },
      fresh(obj) {
        this.$router.push({
          path: "/rank/ACMRank",
          query: obj
        });
      },
      paramsQuery() {
        let obj = {};
        if (this.$route.query.page) {
          obj.page = Number(this.$route.query.page);
        }
        return obj;
      },
      paramsInit() {
        if (this.$route.query.page) {
          this.page = Number(this.$route.query.page);
        } else {
          this.page = 1;
        }
      },
      indexMethod(index) {
        return 1 + index + (this.page - 1) * 10;
      },
      getRate(ac, total) {
        let rate = ac / total;
        if (isNaN(rate)) {
          return '--';
        } else {
          return (rate * 100).toFixed(2) + '%';
        }
      }
    },
    components: {
      VeHistogram
    },
    watch: {
      $route() {
        this.getRank();
      }
    }
  };
</script>

<style scoped>
  .center-box {
    min-width: 600px;
    margin-top: 20px !important;
    margin: 0 auto;
    width: 90%;
  }

  .content {
    background-color: #ffffff;
    border-radius: 10px;
    box-shadow: 0 2px 12px 0 rgba(0, 0, 0, 0.1);
    overflow: hidden!important;
  }


  .slide-fade-enter-active {
    transition: all 0.8s ease;
  }

  /* .slide-fade-leave-active {
    transition: all 0.8s cubic-bezier(1, 0.5, 0.8, 1);
  } */

  .slide-fade-enter,
  .slide-fade-leave-to {
    transform: translateY(40px);
    opacity: 0;
  }
</style>