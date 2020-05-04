<template>
  <div>
    <!-- <button @click="show = !show">
    Toggle render
  </button> -->
    <transition name="slide-fade">
      <div v-if="show">
        <div class="center-box" >
          <el-row style="height:60px">
            <span style="float:left;font-size:20px;margin-left:20px;margin-top:15px">Status</span>
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
            <el-table stripe size="small" :data="status" style="width:100%;border-radius:10px" v-loading="loading">
              <el-table-column align="center" label="Submit Time" width="180">
                <template slot-scope="scope">
                  <i class="el-icon-time"></i>
                  <span style="margin-left: 10px">{{
                    scope.row.submitTime | formatDateTime
                  }}</span>
                </template>
              </el-table-column>
              <el-table-column align="center" label="Id" width="180">
                <template slot-scope="scope">
                  <el-link
                    type="primary"
                    :underline="false"
                    @click="gotoResult(scope.row.id)"
                    >{{ scope.row.id }}</el-link
                  >
                </template>
              </el-table-column>
              <el-table-column align="center" label="Status" width="240">
                <template slot-scope="scope">
                  <el-button
                    size="mini"
                    :type="scope.row.status | formatType "
                    >{{scope.row.status |formatFlags}}</el-button
                  >
                 
                </template>
              </el-table-column>
              <el-table-column align="center" label="Problem">
                <template slot-scope="scope">
                  <el-link
                    :underline="false"
                    @click="gotoAnswer(scope.row.pid)"
                    >{{ scope.row.problemName }}</el-link
                  >
                </template>
              </el-table-column>
              <el-table-column align="center" prop="language" label="Language">
              </el-table-column>
              <el-table-column align="center" prop="totalScore" label="Score">
              </el-table-column>
            </el-table>
          </el-row>
        </div>
        <el-row style="width:80%;margin:20px auto 0px">
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
    </transition>
  </div>
</template>
<script>
export default {
  created() {
    this.$bus.emit('changeHeader', '4');
    // this.practiseListLoading=true
  },
  async mounted() {
    this.show = false;
    this.show = true;
    this.loading=true
    this.queryList()
    
  },
  data() {
    return {
      loading:false,
      count: 0,
      page: 1,
      show: false,
      status: [],
    };
  },
  methods: {
    freshBtnCss(){

    },
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
      return obj;
    },
    fresh(obj) {
      this.$router.push({
        path: '/status',
        query: obj
      })
    },
    reset() {
    this.$router.push({
      path: '/status'
    });
    },
    gotoAnswer(val) {
      this.$router.push({
        path: '/practice/answer',
        query: { id: Number(val) }
      });
    },
    gotoResult(val) {
      this.$router.push({
        path: '/practice/result',
        query: { id: Number(val) }
      });
    },
    handlePageChange(val) {
      let obj = this.params_query();
      obj.page = Number(val);
      this.fresh(obj);
    },
   async  queryList(){
      this.loading=true
      this.params_init();
      try {
      const {data:res} = await this.$http.post('/practice/getAllStatus', {
        page: this.page
      });
      // console.log(res);
      if (res.error) {
        this.$message.error(res.error);
      } else {
        this.status = res.data;
        this.loading=false
      }
    } catch (err) {
      console.log(err);
      // alert(err)
    }
    try {
      const {data:res1} = await this.$http.post('/practice/getAllStatusCount');
      // console.log(res1);
      if (res1.error) {
        // this.$message.error(res1.error);
        return
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
    formatFlags:function(value){
        switch(value)
      {
        case "RE":
                    return "Runtime Error"
          break
        case "CE":
                    return "Compile Error"
          break
        case "WA":
                    return "Wrong Answer"
          break;
          case "ISE":
                    return "Internal Server Error"
          break;
        case "TLE":
                    return "Time Limit Exceeded"
          break
        case "MLE":
                    return "Memory Limit Exceeded"
          break
        case "OLE":
                    return "Output Limit Exceeded"
          break
        case "PA":
                    return "Partial Accepted"
          break
        case "Judging":
                    return "Judging"
          break
        case "Pending":
              return "Pending"
          break
        case "AC":    
                  return "Accepted"
          break
          default:
            return "Internal Server Error"
      }
    },
    formatType:function(value){
        switch(value)
      {
        case "RE":
        case "WA":
        case "ISE":
          return "danger"

          break;
        case "TLE":
        case "MLE":
        case "OLE":
        case "CE":
          return "warning"
          break
        case "PA":
        case "Judging":
        case "Pending":
          return "primary"
          break
        case "AC":
          return "success"
          break
          default:
            return "danger"
      }
    }
  }
};
</script>

<style scoped>
.center-box {
  min-width: 600px;
  margin-top: 20px !important;
  margin: 0 auto;
  width: 80%;
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
img {
  /* margin: 0 auto; */
  margin-top: 100px;
  width: 300px;
  height: 300px;
  /* margin: 0 auto; */
  /* border-radius: 50%;  */
  /* left: 50%; */
  /* transform: translate(-5%, 0); */
  background-color: #ffffff;
  /* display: flex;
  justify-content: center;
  align-items: center; */
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
