<template>
  <div>
    <transition name="slide-fade">
      <div class="center-box" v-if="show">
        <el-row style="height:62px">
          <span style="float:left;font-size:23px;margin-left:30px;margin-top:15px">Contests</span>
          <el-button style="float:right;margin-top:15px;margin-right:20px;" class="el-icon-refresh" type="primary"
            size="small" @click="reset">&nbsp;Reset</el-button>
          <el-input style="float:right;width:200px;margin-top:15px;margin-right:20px" placeholder="keywords"
            v-model="keywords" size="small">
            <el-button slot="append" icon="el-icon-search"
              style="color:#ffffff;background-color:#409EFF;border-top-left-radius:0;border-bottom-left-radius:0;margin-right:-21px"
              size="small" @click="handleKeywordsChange"></el-button>
          </el-input>
          <el-select v-model="status" placeholder="Status"
            style="margin-top:15px;margin-right:20px;float:right;width:140px" size="small" @change="handleStatusChange">
            <el-option v-for="item in StatusOptions" :key="item.value" :label="item.label" :value="item.value">
            </el-option>
          </el-select>
          <el-select v-model="rule" placeholder="Rule" style="margin-top:15px;margin-right:20px;float:right;width:140px"
            size="small" @change="handleRuleChange">
            <el-option v-for="item in ruleOptions" :key="item.value" :label="item.label" :value="item.value">
            </el-option>
          </el-select>
        </el-row>
        <div v-loading="contestLoading" style="min-height:80px">
          <div v-if="contest.length===0" style="float:left;font-size:20px;width:100%;text-align:center;margin-top:25px">
            No Data</div>
          <el-row v-bind:key="index" v-for="(item, index) in contest" style="height:75px;">
            <img style="width: 45px; height: 45px;float:left;margin-left:30px;margin-top:10px"
              src="@/assets/images/contest.png" />
            <div style="float:left;margin-top:8px;margin-left:10px">
              <el-link style="font-size:20px;" @click="goDetail(item.id)"> {{ item.title }}</el-link>
              <div style="font-size:13px;margin-top:3px">
                <i class="el-icon-unlock" style="color:#409EFF"></i>
                <span>{{ item.startTime }}</span>
                <i class="el-icon-lock" style="color:#409EFF;margin-left:10px"></i>
                <span>{{ item.endTime }}</span>
                <i class="el-icon-timer" style="margin-left:10px;color:#409EFF"></i>
                <span>{{ item.timeDiff }}</span>
                <i class="el-icon-s-operation" style="margin-left:10px;color:#409EFF"></i>
                <span>{{ item.rule }}</span>
              </div>
            </div>
            <div style="float:right;margin-right:50px;margin-top:16px">
              <el-button plain size="small">
                <div :style="item.style">
                  &nbsp;
                </div>
                <span style="margin-left:5px;margin-right:-8px">{{
                  item.status
                }}</span>
              </el-button>
            </div>
            <el-row v-if="index!==0"
              style="height:1px;float:top;margin-top:-5px;border-top:1px solid rgb(233, 233, 235);">
            </el-row>
          </el-row>
        </div>
        <el-pagination style="float:right;margin-top:20px;" background layout="prev, pager, next" :page-size="5"
          @current-change="handlePageChange" :current-page="page" :total="count">
        </el-pagination>
      </div>
    </transition>
  </div>
</template>
<script>
  export default {
    data() {
      return {
        show: false,
        keywords: '',
        status: 0,
        rule: '',
        count: 5,
        page: 1,
        contestLoading: true,
        contest: [],
        StatusOptions: [{
            value: 0,
            label: 'All'
          },
          {
            value: 1,
            label: 'Not Started'
          },
          {
            value: 2,
            label: 'Under Way'
          },
          {
            value: 3,
            label: 'Ended'
          }
        ],
        ruleOptions: [{
            value: null,
            label: 'All'
          },
          {
            value: 'ACM',
            label: 'ACM'
          },
          {
            value: 'OI',
            label: 'OI'
          }
        ],
      };
    },
    created() {
      this.$bus.emit('changeHeader', '3');
      // this.practiseListLoading=true
      this.show = false;
    },
    mounted() {
      this.show = true;
      this.queryList();
    },
    methods: {
      goDetail(val) {
        this.$router.push({
          path: "/contest/detail",
          query: {
            id: val
          }
        })
      },
      params_init() {
        if (this.$route.query.page) {
          this.page = Number(this.$route.query.page);
        } else {
          this.page = 1;
        }
        if (this.$route.query.rule) {
          this.rule = this.$route.query.rule;
        } else {
          this.rule = '';
        }
        if (this.$route.query.status) {
          this.status = Number(this.$route.query.status);
        } else {
          this.status = null;
        }
        if (this.$route.query.keywords) {
          this.keywords = this.$route.query.keywords;
        } else {
          this.keywords = '';
        }
      },
      params_query() {
        let obj = {};
        if (this.$route.query.page) {
          obj.page = Number(this.$route.query.page);
        }
        if (this.$route.query.rule) {
          obj.rule = this.$route.query.rule;
        }
        if (this.$route.query.status) {
          obj.status = Number(this.$route.query.status);
        }
        if (this.$route.query.keywords) {
          obj.keywords = this.$route.query.keywords;
        }
        return obj;
      },
      async queryList() {
        try {
          this.params_init();
          this.contestLoading = true;
          const {
            data: res
          } = await this.$http.post('/contest/getAll', {
            page: this.page,
            rule: this.rule,
            status: this.status,
            keywords: this.keywords
          });
          const {
            data: res1
          } = await this.$http.post('/contest/getCount', {
            rule: this.rule,
            status: this.status,
            keywords: this.keywords
          });
          console.log(res);
          if (res.error) {
            this.$message.error(res.error);
          } else {
            this.contestLoading = false;
            this.contest = res.data;
            this.count = res1.data
            this.handleStatus();
          }
        } catch (err) {
          console.log(err);
          alert(err);
        }
      },
      handleStatus() {
        for (var i = 0; i < this.contest.length; i++) {
          var startTime = new Date(this.contest[i].startTime.replace(/-/g, '/'));
          var now = new Date(this.contest[i].now.replace(/-/g, '/'));
          var endTime = new Date(this.contest[i].endTime.replace(/-/g, '/'));
          var timeDiff = endTime - startTime;
          if (timeDiff < 3600000) {
            this.contest[i].timeDiff =
              this.toDecimal(timeDiff / 60000) + ' minutes';
          } else if (3600000 < timeDiff && timeDiff < 86400000) {
            this.contest[i].timeDiff =
              this.toDecimal(timeDiff / 3600000) + ' hours';
          } else if (86400000 < timeDiff && timeDiff < 2592000000) {
            this.contest[i].timeDiff =
              this.toDecimal(timeDiff / 86400000) + ' days';
          } else if (2592000000 < timeDiff && timeDiff < 31104000000) {
            this.contest[i].timeDiff =
              this.toDecimal(timeDiff / 2592000000) + ' months';
          }
          if (now.getTime() < startTime.getTime()) {
            this.contest[i].status = 'Not Start';
            this.contest[i].style =
              'float:left;margin-left:-10px;width:12px;height:12px;border-radius:6px;background:#409EFF';
          } else if (now.getTime() > endTime.getTime()) {
            this.contest[i].status = 'Ended';
            this.contest[i].style =
              'float:left;margin-left:-10px;width:12px;height:12px;border-radius:6px;background:#F56C6C';
          } else {
            this.contest[i].status = 'Under Way';
            this.contest[i].style =
              'float:left;margin-left:-10px;width:12px;height:12px;border-radius:6px;background:#67C23A';
          }
        }
      },
      toDecimal(x) {
        var f = parseFloat(x);
        if (isNaN(f)) {
          return;
        }
        f = Math.round(x * 100) / 100;
        return f;
      },
      handleRuleChange(val) {
        let obj = this.params_query();
        obj.rule = val;
        obj.page = 1;
        this.fresh(obj);
      },
      handleStatusChange(val) {
        let obj = this.params_query();
        obj.status = Number(val);
        obj.page = 1;
        this.fresh(obj);
      },
      handleKeywordsChange(val) {
        let obj = this.params_query();
        obj.keywords = this.keywords;
        obj.page = 1;
        this.fresh(obj);
      },
      handlePageChange(val) {
        let obj = this.params_query();
        obj.page = Number(val);
        this.fresh(obj);
      },
      fresh(obj) {
        this.$router.push({
          path: '/contest',
          query: obj
        });
      },
      reset() {
        this.$router.push({
          path: '/contest'
        });
      }
    },
    watch: {
      $route() {
        this.queryList();
      }
    },
  };
</script>

<style scoped>
  .center-box {
    min-width: 600px;
    margin-top: 20px !important;
    margin: 0 auto;
    width: 95%;
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

  .slide-fade-enter,
  .slide-fade-leave-to

  /* .slide-fade-leave-active for below version 2.1.8 */
    {
    transform: translateY(40px);
    opacity: 0;
  }
</style>