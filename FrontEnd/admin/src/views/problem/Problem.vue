<template>

  <div>
    <transition name="slide-fade">

      <div class="center-box" v-if="show">
        <el-row style="height:60px;line-height:60px">
          <span style="font-size:20px;margin-left:20px">Problem List</span>
          <el-button style="float:right;margin-top:15px;margin-right:20px;" class="el-icon-refresh" type="primary"
            size="small" @click="reset">&nbsp;Reset</el-button>
          <el-input style="float:right;width:200px;margin-top:15px;margin-right:20px" placeholder="keywords"
            v-model="keywords" size="small">
            <el-button slot="append" icon="el-icon-search"
              style="color:#ffffff;background-color:#409EFF;border-top-left-radius:0;border-bottom-left-radius:0;margin-right:-21px;margin-top:-7px"
              size="small" @click="handleKeywordsChange"></el-button>
          </el-input>
          <el-select v-model="difficulty" placeholder="Difficulty"
            style="margin-top:1.2px;margin-right:20px;float:right;width:140px" size="small"
            @change="handleDifficultChange">
            <el-option v-for="item in difficultyOptions" :key="item.value" :label="item.label" :value="item.value">
            </el-option>
          </el-select>
          <div style="float:right;margin-top:16px;margin-right:80px">
            <my-switch v-model="mine" @toggle="toggle"></my-switch>
          </div>
          <div style="margin-right:20px;float:right">
            <span style="color:grey;font-size:13px">Tags:&nbsp;</span>
            <el-switch @change="switchTags" v-model="showTags">
            </el-switch>
          </div>
        </el-row>
        <el-row style="height:1px;float:top;border-top:1px solid #EBEEF5;"></el-row>
        <el-table :data="tableData" style="width: 100%" v-loading="loading" size="small">
          <el-table-column label="ID" prop="id" align="center" min-width="30">
          </el-table-column>
          <el-table-column label="Display Id" prop="ref" align="center" min-width="30">
          </el-table-column>
          <el-table-column label="Title" prop="title" min-width="90">
          </el-table-column>
          <el-table-column label="Creator" prop="creatorName" align="center" min-width="30">
          </el-table-column>
          <el-table-column label="Create Time" prop="createTime" align="center" min-width="60">
          </el-table-column>
          <el-table-column label="Last Update Time" prop="lastUpdateTime" align="center" min-width="60">
          </el-table-column>
          <el-table-column prop="difficulty" align="center" label="Level" min-width="60">
            <template slot-scope="scope">
              <el-button size="mini" type="info" v-if="scope.row.difficulty === 'Casual'">Casual</el-button>
              <el-button size="mini" type="success" v-if="scope.row.difficulty === 'Eazy'">Eazy</el-button>
              <el-button size="mini" type="primary" v-if="scope.row.difficulty === 'Normal'">Normal</el-button>
              <el-button size="mini" type="warning" v-if="scope.row.difficulty === 'Hard'">Hard</el-button>
              <el-button size="mini" type="danger" v-if="scope.row.difficulty === 'Crazy'">Crazy</el-button>
            </template>
          </el-table-column>
          <el-table-column v-if="realShowTags" prop="tags" align="center" label="Tags" min-width="120">
            <template slot-scope="scope">
              <div v-bind:key="i" v-for="(tag, i) in scope.row.tags">
                <div style="height:35px;">
                  <el-button size="mini" type="primary" style="margin-top:3px" plain>{{ tag.name }}
                  </el-button>
                </div>
              </div>
            </template>
          </el-table-column>
          <el-table-column label="Visible" width="80" align="center">
            <template slot-scope="scope">
              <!-- <el-switch v-model="scope.row.visible" active-color="#13ce66" inactive-color="#ff4949">
              </el-switch> -->
              <div @click="switchVisible(scope.row)">
                <el-switch v-model="scope.row.visible" active-color="#409eff" inactive-color="#dcdfe6">
                </el-switch>
              </div>
            </template>
          </el-table-column>
          <el-table-column label="Option" width="120" align="center">
            <template slot-scope="scope">
              <el-row>
                <el-tooltip content="Edit" placement="top">
                  <el-button size="mini" class="el-icon-edit-outline" @click="editProblem(scope.row.id)"></el-button>
                </el-tooltip>
                <el-tooltip content="Delete" placement="top">
                  <el-button size="mini" class="el-icon-delete" style="color:red" @click="deleteProblem(scope.row)">
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
        difficulty: '',
        keywords: '',
        realShowTags: false,
        showTags: false,
        count: 0,
        page: 1,
        mine: false,
        tableData: [{
            id: '12987122',
            name: '好滋好味鸡蛋仔',
            category: '江浙小吃、小吃零食',
            desc: '荷兰优质淡奶，奶香浓而不腻',
            address: '上海市普陀区真北路',
            shop: '王小虎夫妻店',
            shopId: '10333'
          },
          {
            id: '12987123',
            name: '好滋好味鸡蛋仔',
            category: '江浙小吃、小吃零食',
            desc: '荷兰优质淡奶，奶香浓而不腻',
            address: '上海市普陀区真北路',
            shop: '王小虎夫妻店',
            shopId: '10333'
          }, {
            id: '12987125',
            name: '好滋好味鸡蛋仔',
            category: '江浙小吃、小吃零食',
            desc: '荷兰优质淡奶，奶香浓而不腻',
            address: '上海市普陀区真北路',
            shop: '王小虎夫妻店',
            shopId: '10333'
          }, {
            id: '12987126',
            name: '好滋好味鸡蛋仔',
            category: '江浙小吃、小吃零食',
            desc: '荷兰优质淡奶，奶香浓而不腻',
            address: '上海市普陀区真北路',
            shop: '王小虎夫妻店',
            shopId: '10333'
          }
        ],
        difficultyOptions: [{
            value: null,
            label: 'All'
          },
          {
            value: 'Casual',
            label: 'Casual'
          },
          {
            value: 'Eazy',
            label: 'Eazy'
          },
          {
            value: 'Normal',
            label: 'Normal'
          },
          {
            value: 'Hard',
            label: 'Hard'
          },
          {
            value: 'Crazy',
            label: 'Crazy'
          }
        ]
      }
    },
    created() {
      this.$bus.emit("changeHeader", "3-1")
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
      async editProblem(id) {
        try {
          const {
            data: res
          } = await this.$http.post('/admin/problem/tryEdit', {
            id: id
          });
          if (res.error) {
            this.$message.error(res.error)
            return
          }
          this.$router.push({
            path: "/problem/edit",
            query: {
              id: id
            }
          })
        } catch (err) {
          console.log(err);
          alert(err)
        }
      },
      deleteProblem() {

      },
      paramsInit() {
        if (this.$route.query.page) {
          this.page = Number(this.$route.query.page);
        } else {
          this.page = 1;
        }
        if (this.$route.query.difficulty) {
          this.difficulty = this.$route.query.difficulty;
        } else {
          this.difficulty = '';
        }
        if (this.$route.query.keywords) {
          this.keywords = this.$route.query.keywords;
        } else {
          this.keywords = '';
        }
        if (this.$route.query.mine) {
          if (typeof (this.$route.query.mine) === typeof(true)) {
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
        if (this.$route.query.difficulty) {
          obj.difficulty = this.$route.query.difficulty;
        }
        if (this.$route.query.tid) {
          obj.tid = Number(this.$route.query.tid);
        }
        if (this.$route.query.keywords) {
          obj.keywords = this.$route.query.keywords;
        }
        if (this.$route.query.mine) {
          obj.mine = true
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
          } = await this.$http.post('/admin/problem/getAll', {
            page: this.page,
            difficulty: this.difficulty,
            keywords: this.keywords,
            mine: this.mine,
          });
          if (res.error) {
            this.$message.error(res.error)
            return
          }
          this.tableData = res.data
          const {
            data: res1
          } = await this.$http.post('/admin/problem/getCount', {
            difficulty: this.difficulty,
            keywords: this.keywords,
            mine: this.mine,
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
            } = await this.$http.post('/admin/problem/setVisibleTrue', {
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
            } = await this.$http.post('/admin/problem/setVisibleFalse', {
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
      handlePageChange(val) {
        let obj = this.paramsQuery();
        obj.page = Number(val);
        this.fresh(obj);
      },
      handleDifficultChange(val) {
        let obj = this.paramsQuery();
        obj.difficulty = val;
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
          path: '/problem',
          query: obj
        });
      },
      reset() {
        this.$router.push({
          path: '/problem'
        });
      },
      goCreate() {
        this.$router.push({
          path: '/problem/create'
        });
      },
      switchTags(val) {
        this.loading = true
        setTimeout(() => {
          this.realShowTags = val
        }, 200)
        setTimeout(() => {
          this.loading = false
        }, 500)
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