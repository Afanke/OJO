<template>
  <div>
    <transition name="slide-fade">
      <div class="center-box" v-if="show">
        <el-row style="height:60px;line-height:60px">
          <span style="font-size:20px;margin-left:20px">Contest List</span>
          <el-button style="float:right;margin-top:15px;margin-right:20px;" class="el-icon-refresh" type="primary"
            size="small" @click="reset">&nbsp;Reset</el-button>
          <el-input style="float:right;width:200px;margin-top:15px;margin-right:20px" placeholder="keywords"
            v-model="keywords" size="small">
            <el-button slot="append" icon="el-icon-search"
              style="color:#ffffff;background-color:#409EFF;border-top-left-radius:0;border-bottom-left-radius:0;margin-right:-21px;margin-top:-7px"
              size="small" @click="handleKeywordsChange"></el-button>
          </el-input>
          <el-select v-model="status" placeholder="Status"
            style="margin-top:1.2px;margin-right:20px;float:right;width:140px" size="small"
            @change="handleStatusChange">
            <el-option v-for="item in StatusOptions" :key="item.value" :label="item.label" :value="item.value">
            </el-option>
          </el-select>
          <el-select v-model="rule" placeholder="Rule"
            style="margin-top:1.2px;margin-right:20px;float:right;width:140px" size="small" @change="handleRuleChange">
            <el-option v-for="item in ruleOptions" :key="item.value" :label="item.label" :value="item.value">
            </el-option>
          </el-select>
          <div style="float:right;margin-top:16px;margin-right:80px">
            <my-switch v-model="mine" @toggle="toggle"></my-switch>
          </div>

        </el-row>
        <el-row style="height:1px;float:top;border-top:1px solid #EBEEF5;"></el-row>
        <el-table :data="tableData" style="width: 100%" v-loading="loading" size="small">
          <el-table-column type="expand">
            <template slot-scope="props">
              <el-form label-position="left" inline class="demo-table-expand">
                <el-form-item label="Start Time">
                  <span>{{ props.row.startTime }}</span>
                </el-form-item>
                <el-form-item label="Create Time">
                  <span>{{ props.row.createTime }}</span>
                </el-form-item>
                <el-form-item label="End Time">
                  <span>{{ props.row.endTime }}</span>
                </el-form-item>
                <el-form-item label="Last Update Time">
                  <span>{{ props.row.lastUpdateTime }}</span>
                </el-form-item>
                <el-form-item label="Punish">
                  <span v-if="props.row.rule==='OI'">{{ props.row.punish }}&nbsp;score</span>
                  <span v-if="props.row.rule==='ACM'">{{ props.row.punish }}&nbsp;second</span>
                </el-form-item>
                <el-form-item label="Submit Limit">
                  <span v-if="props.row.submitLimit===0">--</span>
                  <span v-if="props.row.submitLimit>0">{{ props.row.submitLimit }}</span>
                </el-form-item>
              </el-form>
            </template>
          </el-table-column>
          <el-table-column label="ID" prop="id" align="center" min-width="30">
          </el-table-column>
          <el-table-column label="Title" prop="title" min-width="90" align="center">
          </el-table-column>
          <el-table-column label="Creator" prop="creatorName" align="center" min-width="50">
          </el-table-column>
          <el-table-column align="center" label="Rule" min-width="40">
            <template slot-scope="scope">
              <el-button size="mini" type="primary" plain>{{scope.row.rule}}</el-button>
            </template>
          </el-table-column>
          <el-table-column label="Duration" prop="timeDiff" align="center" min-width="60">
          </el-table-column>
          <el-table-column align="center" label="Status" min-width="60">
            <template slot-scope="scope">
              <el-button size="mini" type="primary" v-if="scope.row.status === 'Not Started'" plain="">Not Started
              </el-button>
              <el-button size="mini" type="success" v-if="scope.row.status === 'Under Way'" plain="">Under Way
              </el-button>
              <el-button size="mini" type="danger" v-if="scope.row.status === 'Ended'" plain="">Ended</el-button>
            </template>
          </el-table-column>
          <el-table-column label="Visible" min-width="80" align="center">
            <template slot-scope="scope">
              <div @click="switchVisible(scope.row)">
                <el-switch v-model="scope.row.visible" active-color="#409eff" inactive-color="#dcdfe6">
                </el-switch>
              </div>
            </template>
          </el-table-column>
          <el-table-column label="Option" min-width="120" align="center">
            <template slot-scope="scope">
              <el-row>
                <el-tooltip content="Edit" placement="top">
                  <el-button size="mini" class="el-icon-edit-outline" @click="editContest(scope.row.id)"></el-button>
                </el-tooltip>
                <el-tooltip content="Problem List" placement="top">
                  <el-button size="mini" class="el-icon-notebook-2" @click="editProblem(scope.row.id)"></el-button>
                </el-tooltip>
                <el-tooltip content="Delete" placement="top">
                  <el-button size="mini" class="el-icon-delete" style="color:red" @click="deleteContest(scope.row.id)">
                  </el-button>
                </el-tooltip>
              </el-row>
            </template>
          </el-table-column>
        </el-table>
        <el-row style="margin-top:20px;">
          <el-button type="primary" @click="goCreate" size="small" class="el-icon-plus"
            style="margin-left:30px;float:left"> Create</el-button>
          <el-pagination @current-change="handlePageChange" :page-size="10" style="float:right;margin-right:30px"
            layout="prev, pager, next" :total="count" :current-page.sync="page">
          </el-pagination>
        </el-row>
        <el-row style="margin-top:20px;">
        </el-row>
      </div>
    </transition>
  </div>

</template>
<script>
  import SwitchButton from "@/components/Switch.vue"
  export default {
    data() {
      return {
        show: false,
        loading: true,
        rule: '',
        keywords: '',
        count: 0,
        page: 1,
        mine: false,
        tableData: [],
        status: 0,
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
            value: 'OI',
            label: 'OI'
          },
          {
            value: 'ACM',
            label: 'ACM'
          }
        ]
      }
    },
    created() {
      this.$bus.emit("changeHeader", "4-1")
      this.show = false
    },
    async mounted() {
      this.show = true
      this.queryList()
    },
    methods: {
      toggle(checked) {
        this.mine = checked
        let obj = this.paramsQuery();
        obj.mine = this.mine;
        obj.page = 1;
        this.fresh(obj);
      },
      async editContest(id) {
        try {
          const {
            data: res
          } = await this.$http.post('/admin/contest/tryEdit', {
            id: id
          });
          if (res.error) {
            this.$message.error(res.error)
            return
          }
          this.$router.push({
            path: "/contest/edit",
            query: {
              id: id
            }
          })
        } catch (err) {
          console.log(err);
          alert(err)
        }
      },
      async deleteContest(id) {
        try {
          const {
            data: res
          } = await this.$http.post('/admin/contest/deleteContest', {
            id: id
          });
          if (res.error) {
            this.$message.error(res.error)
            return
          }
          this.$message.success(res.data)
          this.queryList()
        } catch (err) {
          console.log(err);
          alert(err)
        }
      },
      async editProblem(id) {
        try {
          const {
            data: res
          } = await this.$http.post('/admin/contest/tryEdit', {
            id: id
          });
          if (res.error) {
            this.$message.error(res.error)
            return
          }
          this.$router.push({
            path: "/contest/problem",
            query: {
              id: id
            }
          })
        } catch (err) {
          console.log(err);
          alert(err)
        }
      },
      paramsInit() {
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
        if (this.$route.query.mine) {
          if (typeof (this.$route.query.mine) === typeof (true)) {
            this.mine = this.$route.query.mine
          } else {
            if (this.$route.query.mine === "true") {
              this.mine = true
            } else {
              this.mine = false
            }
          }
        } else {
          this.mine = false
        }
      },
      paramsQuery() {
        let obj = {};
        if (this.$route.query.page) {
          obj.page = Number(this.$route.query.page);
        }
        if (this.$route.query.rule) {
          obj.rule = this.$route.query.rule;
        }
        if (this.$route.query.keywords) {
          obj.keywords = this.$route.query.keywords;
        }
        if (this.$route.query.status) {
          obj.status = Number(this.$route.query.status);
        }
        if (this.$route.query.mine) {
          if (typeof (this.$route.query.mine) === typeof (true)) {
            obj.mine = this.$route.query.mine
          } else {
            if (this.$route.query.mine === "true") {
              obj.mine = true
            } else {
              obj.mine = false
            }
          }
        } else {
          obj.mine = false
        }
        return obj;
      },
      async queryList() {
        this.loading = true
        this.paramsInit()
        try {
          const {
            data: res
          } = await this.$http.post('/admin/contest/getAll', {
            page: this.page,
            rule: this.rule,
            keywords: this.keywords,
            mine: this.mine,
            status: this.status
          });
          if (res.error) {
            this.$message.error(res.error)
            return
          }
          this.tableData = res.data
          this.handleStatus()
          const {
            data: res1
          } = await this.$http.post('/admin/contest/getCount', {
            difficulty: this.difficulty,
            keywords: this.keywords,
            mine: this.mine,
            status: this.status,
            rule: this.rule,
          });
          if (res1.error) {
            this.$message.error(res1.error)
            return
          }
          this.count = res1.data
          this.loading = false
        } catch (err) {
          console.log(err);
          alert(err)
        }
      },
      async switchVisible(obj) {
        console.log(obj)
        if (obj.visible) {
          try {
            const {
              data: res
            } = await this.$http.post('/admin/contest/setVisibleTrue', {
              id: obj.id
            });
            if (res.error) {
              this.$message.error(res.error)
              obj.visible = false
              return
            }
            obj.visible = true
          } catch (err) {
            console.log(err);
            // alert(err)
          }
        } else {
          try {
            const {
              data: res
            } = await this.$http.post('/admin/contest/setVisibleFalse', {
              id: obj.id
            });
            if (res.error) {
              this.$message.error(res.error)
              obj.visible = true
              return
            }
            obj.visible = false
          } catch (err) {
            console.log(err);
            // alert(err)
          }
        }
      },
      handleStatus() {
        for (let i = 0; i < this.tableData.length; i++) {
          let startTime = new Date(this.tableData[i].startTime.replace(/-/g, '/'));
          let now = new Date(this.tableData[i].now.replace(/-/g, '/'));
          let endTime = new Date(this.tableData[i].endTime.replace(/-/g, '/'));
          let timeDiff = endTime - startTime;
          if (timeDiff < 3600000) {
            this.tableData[i].timeDiff =
              this.toDecimal(timeDiff / 60000) + ' minutes';
          } else if (3600000 <= timeDiff && timeDiff < 86400000) {
            this.tableData[i].timeDiff =
              this.toDecimal(timeDiff / 3600000) + ' hours';
          } else if (86400000 <= timeDiff && timeDiff < 2592000000) {
            this.tableData[i].timeDiff =
              this.toDecimal(timeDiff / 86400000) + ' days';
          } else if (2592000000 <= timeDiff && timeDiff < 31104000000) {
            this.tableData[i].timeDiff =
              this.toDecimal(timeDiff / 2592000000) + ' months';
          } else {
               this.tableData[i].timeDiff =
              this.toDecimal(timeDiff / 31104000000) + ' years';
          }
          if (now.getTime() < startTime.getTime()) {
            this.tableData[i].status = 'Not Started';
          } else if (now.getTime() > endTime.getTime()) {
            this.tableData[i].status = 'Ended';
          } else {
            this.tableData[i].status = 'Under Way';
          }
        }
      },
      toDecimal(x) {
        let f = parseFloat(x);
        if (isNaN(f)) {
          return;
        }
        f = Math.round(x * 100) / 100;
        return f;
      },
      handleStatusChange(val) {
        let obj = this.paramsQuery();
        obj.status = Number(val);
        obj.page = 1;
        this.fresh(obj);
      },
      handlePageChange(val) {
        let obj = this.paramsQuery();
        obj.page = Number(val);
        this.fresh(obj);
      },
      handleRuleChange(val) {
        let obj = this.paramsQuery();
        obj.rule = val;
        obj.page = 1;
        this.fresh(obj);
      },
      handleKeywordsChange() {
        let obj = this.paramsQuery();
        obj.keywords = this.keywords;
        obj.page = 1;
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
      },
      goCreate() {
        this.$router.push({
          path: '/contest/create'
        });
      },

    },
    watch: {
      $route() {
        this.queryList();
      }
    },
    components: {
      mySwitch: SwitchButton
    }
  };
</script>

<style scoped>
  .center-box>>>.el-form-item__label {
    text-align: left;
    vertical-align: middle;
    float: left;
    font-size: 14px;
    /* color: #606266; */
    min-width: 90px;
    color: #99a9bf;
    line-height: 40px;
    padding: 0 12px 0 0;
    box-sizing: border-box;
  }


  .demo-table-expand .el-form-item {
    margin-right: 0;
    margin-bottom: 0;
    width: 50%;
  }

  .center-box {
    height: auto;
    background-color: #ffffff;
    border-radius: 10px;
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