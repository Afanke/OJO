<template>
  <div>
        <div class="center-box">
          <el-row style="height:60px">
            <span
              style="float:left;font-size:20px;margin-left:20px;margin-top:15px"
              >Status</span
            >
            <el-button
              style="float:right;margin-top:15px;margin-right:20px;"
              class="el-icon-refresh"
              type="primary"
              size="small"
              @click="reset"
              >&nbsp;Reset</el-button
            >
          </el-row>
          <el-row>
            <el-table
              stripe
              size="small"
              :data="status"
              style="width:100%;border-radius:10px"
              v-loading="loading"
            >
              <el-table-column align="center" label="Submit Time" width="180">
                <template slot-scope="scope">
                  <i class="el-icon-time"></i>
                  <span style="margin-left: 10px">{{
                    scope.row.submitTime | formatDateTime
                  }}</span>
                </template>
              </el-table-column>
              <el-table-column align="center" label="Id" min-width="80">
                <template slot-scope="scope">
                  <el-link
                    type="primary"
                    :underline="false"
                    @click="gotoResult(scope.row.id)"
                    >{{ scope.row.id }}</el-link
                  >
                </template>
              </el-table-column>
              <el-table-column align="center" label="Status" min-width="120">
                <template slot-scope="scope">
                  <el-button
                    size="mini"
                    :type="scope.row.status | formatType"
                    >{{ scope.row.status | formatFlags }}</el-button
                  >
                </template>
              </el-table-column>
              <el-table-column align="center" label="Problem" min-width="80" >
                <template slot-scope="scope">
                  <el-link
                    :underline="false"
                    @click="gotoAnswer(scope.row.pid)"
                    >{{ scope.row.problemName }}</el-link
                  >
                </template>
              </el-table-column>
              <el-table-column align="center" prop="language" label="Language" min-width="80">
              </el-table-column>
              <el-table-column align="center" prop="totalScore" label="Score" min-width="80">
              </el-table-column>
            </el-table>
          </el-row>
        </div>
        <el-row style="margin:20px auto 0px">
          <el-pagination
            style="float:right;"
            hide-on-single-page
            background=""
            layout="prev, pager, next"
            :page-size="10"
            @current-change="handlePageChange"
            :current-page="page"
            :total="count"
          >
          </el-pagination>
        </el-row>
  </div>
</template>
<script>
export default {
  created() {
    this.$bus.emit('changeHeader', '3');
    this.show = false;
    // this.practiseListLoading=true
  },
  async mounted() {
    this.show = true;
    this.loading = true;
    this.queryList();
  },
  data() {
    return {
      loading: false,
      count: 0,
      page: 1,
      show: false,
      status: [],
    };
  },
  methods: {
    freshBtnCss() {},
    params_init() {
      if (this.$route.query.page) {
        this.page = Number(this.$route.query.page);
      } else {
        this.page = 1;
      }
    },
    params_query() {
      let obj = {};
      if (this.$route.query.page) {
        obj.page = Number(this.$route.query.page);
      }
      if (this.$route.query.c) {
        obj.c = Number(this.$route.query.c);
      }
      if (this.$route.query.id) {
        obj.id = Number(this.$route.query.id);
      }
      return obj;
    },
    fresh(obj) {
      this.$router.push({
        path: '/contest/detail',
        query: obj
      });
    },
    reset() {
      var obj=this.params_query()
      obj.page=1
      this.$router.push({
        path: '/contest/detail',
        query: obj
      });
    },
    gotoAnswer(val) {
      this.$router.push({
        path: '/contest/answer',
        query: { cid: this.$route.query.id ,pid:Number(val) }
      });
    },
    gotoResult(val) {
      this.$router.push({
        path: '/contest/result',
        query: { id: Number(val) }
      });
    },
    handlePageChange(val) {
      let obj = this.params_query();
      obj.page = Number(val);
      this.fresh(obj);
    },
    async queryList() {
      this.loading = true;
      this.params_init();
      try {
        const { data: res } = await this.$http.post('/contest/getAllStatus', {
          page: this.page,
          cid:Number(this.$route.query.id) 
        });
        // console.log(res);
        if (res.error) {
          this.$message.error(res.error);
          return
        } else {
          this.status = res.data;
          this.loading = false;
        }
      } catch (err) {
        console.log(err);
        // alert(err)
      }
      try {
        const { data: res1 } = await this.$http.post('/contest/getAllStatusCount', {
          cid:Number(this.$route.query.id) 
        });
        if (res1.error) {
          this.$message.error(res1.error);
          return;
        } else {
          this.count = res1.data;
        }
      } catch (err) {
        console.log(err);
        // alert(err)
      }
    }
  },
  watch: {
    $route() {
      this.queryList();
    }
  },
  components: {},
  filters: {
    formatDateTime: function(value) {
      let d = new Date(value);
      let a =
        d.getFullYear() +
        '-' +
        (d.getMonth() + 1 < 10 ? '0' + (d.getMonth() + 1) : d.getMonth() + 1) +
        '-' +
        (d.getDate() < 10 ? '0' + d.getDate() : d.getDate()) +
        ' ' +
        (d.getHours() < 10 ? '0' + d.getDate() : d.getDate()) +
        ':' +
        (d.getMinutes() < 10 ? '0' + d.getMinutes() : d.getMinutes()) +
        ':' +
        (d.getSeconds() < 10 ? '0' + d.getSeconds() : d.getSeconds());
      return a;
    },
    formatFlags: function(value) {
      switch (value) {
        case 'RE':
          return 'Runtime Error';
          break;
        case 'CE':
          return 'Compile Error';
          break;
        case 'WA':
          return 'Wrong Answer';
          break;
        case 'ISE':
          return 'Internal Server Error';
          break;
        case 'TLE':
          return 'Time Limit Exceeded';
          break;
        case 'MLE':
          return 'Memory Limit Exceeded';
          break;
        case 'OLE':
          return 'Output Limit Exceeded';
          break;
        case 'PA':
          return 'Partial Accepted';
          break;
        case 'Judging':
          return 'Judging';
          break;
        case 'Pending':
          return 'Pending';
          break;
        case 'AC':
          return 'Accepted';
          break;
      }
    },
    formatType: function(value) {
      switch (value) {
        case 'RE':
        case 'WA':
        case 'ISE':
          return 'danger';

          break;
        case 'TLE':
        case 'MLE':
        case 'OLE':
        case 'CE':
          return 'warning';
          break;
        case 'PA':
        case 'Judging':
        case 'Pending':
          return 'primary';
          break;
        case 'AC':
          return 'success';
          break;
      }
    }
  }
};
</script>

<style scoped>
.center-box {
  margin: 0 auto;
  width: 100%;
  background-color: #ffffff;
  border-radius: 10px;
  box-shadow: 0 2px 12px 0 rgba(0, 0, 0, 0.1);
}

.el-col {
  text-align: center;
  display: flex;
  justify-content: center;
  align-items: center;
}
.slide-fade-enter-active {
  transition: all 0.8s ease;
}
.slide-fade-leave-active {
  transition: all 0.8s cubic-bezier(1, 0.5, 0.8, 1);
}
.slide-fade-enter, .slide-fade-leave-to
/* .slide-fade-leave-active for below version 2.1.8 */ {
  transform: translateY(40px);
  opacity: 0;
}
</style>
