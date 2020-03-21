<template>
      <div class="center-box" v-if="show">
        <el-row>Rank</el-row>
        <ve-histogram
          :data="chartData"
          style="width:100%"
          :settings="chartSettings"
        ></ve-histogram>
        <div style="width:100%">
          <el-table :data="tableData" style="width: 100%">
            <el-table-column type="index" label="#" width="180">
            </el-table-column>
            <el-table-column
              prop="username"
              label="Username"
              min-width="180"
            >
            </el-table-column>
            <el-table-column
              prop="totalScore"
              label="Total Score"
              align="center"
              min-width="180"
            >
            </el-table-column>
            <el-table-column
              :label="item.title"
              min-width="180"
              v-bind:key="i"
              v-for="(item, i) in problemList"
              align="center"
            >
              <template slot-scope="scope">
                <span style="margin-left: 10px" v-if="scope.row.OIDetail[i]">{{
                  scope.row.OIDetail[i].maxScore
                }}</span>
              </template>
            </el-table-column>
          </el-table>
        </div>
      </div>
</template>
<script>
import VeHistogram from 'v-charts/lib/histogram.common';

export default {
  data() {
    return {
      show: false,
      res: {},
      res1: {},
      tableData: [],
      problemList: [],
      chartData: {
        columns: ['username', 'totalScore'],
        rows: [
          { username: 'sads', totalScore: 213 },
          { username: 'asd', totalScore: 123 }
        ]
      },
      chartSettings: {
        labelMap: {
          PV: '访问用户',
          Order: '下单用户'
        },
        legendName: {
          totalScore: 'Score'
        }
      }
    };
  },
  created() {
    this.$bus.emit('changeHeader', '3');
    // this.practiseListLoading=true
    this.show = false;
  },
  async mounted() {
    this.show = true;
    try {
      const { data: res } = await this.$http.post('/contest/getOIRank', {
        id: 1
      });
      if (res.error) {
        this.$message.error(res.error);
        return;
      }
      this.chartData.rows = res.data;
      this.tableData = res.data;

      const { data: res1 } = await this.$http.post(
        '/contest/getAllProblemName',
        {
          id: 1
        }
      );
      if (res1.error) {
        this.$message.error(res1.error);
        return;
      }
      this.problemList = res1.data;
    } catch (err) {
      console.log(err);
      alert(err);
    }
  },
  methods: {},
  components: {
    VeHistogram
  }
};
</script>

<style scoped>
.center-box {
  min-width: 600px;
  margin-top: 20px !important;
  margin: 0 auto;
  width: 100%;
  background-color: rgb(244, 244, 245);
  border-radius: 10px;
}

</style>
