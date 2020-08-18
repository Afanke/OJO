<template>
  <div>
    <transition name="slide-fade">
      <div class="center-box" v-if="show">
        <el-row :gutter="20">
          <el-col :span="18">
            <div class="left-box">
              <el-row>
                  <div style="text-align:center;font-size:20px;line-height:20px;float:left;margin-left:20px">
                    <p>Problem List</p>
                  </div>
                <el-button style="float:right;margin-top:15px;margin-right:20px;" class="el-icon-refresh" type="primary"
                  size="small" @click="reset">&nbsp;Reset</el-button>
                <el-input style="float:right;width:200px;margin-top:15px;margin-right:20px" placeholder="keywords"
                  v-model="keywords" size="small">
                  <el-button slot="append" icon="el-icon-search"
                    style="color:#ffffff;background-color:#409EFF;border-top-left-radius:0;border-bottom-left-radius:0;margin-right:-21px"
                    size="small" @click="handleKeywordsChange"></el-button>
                </el-input>
                <el-select v-model="difficulty" placeholder="Difficulty"
                  style="margin-top:15px;margin-right:20px;float:right;width:140px" size="small"
                  @change="handleDifficultChange">
                  <el-option v-for="item in difficultyOptions" :key="item.value" :label="item.label"
                    :value="item.value">
                  </el-option>
                </el-select>
                <div style="margin-top:18px;margin-right:20px;float:right">
                  <span style="color:grey;font-size:13px">Tags:&nbsp;</span>
                  <el-switch @change="switchTags"  v-model="showTags" >
                </el-switch>
                </div>
              </el-row>
              <el-row>
                <el-table :highlight-current-row="true" size="small" :data="problemList"
                  style="width: 100%;margin-top:-20px;border-radius:6px" v-loading="practiseListLoading">
                  <el-table-column v-if="hasFlag" align="center" label="" width="50">
                    <template slot-scope="scope">
                      <i class="el-icon-check" style="color:#67C23A;font-size:20px;margin-top:2px" v-if="scope.row.flag === 'AC'"></i>
                      <i class="el-icon-minus" style="color:#F56C6C;font-size:20px;margin-top:2px" v-if="scope.row.flag !== 'AC' && scope.row.flag !== ''"></i>
                    </template>
                  </el-table-column>
                  <el-table-column prop="ref" label="#" min-width="80">
                  </el-table-column>
                  <el-table-column prop="title" label="Title" min-width="150">
                    <template slot-scope="scope">
                      <el-link :underline="false" style="font-size:18px" @click="goto(scope.row.id)">
                        {{ scope.row.title }}
                      </el-link>
                    </template>
                  </el-table-column>
                  <el-table-column prop="difficulty" align="center" label="Level" min-width="80">
                    <template slot-scope="scope">
                      <el-button size="mini" type="info" v-if="scope.row.difficulty === 'Casual'">Casual</el-button>
                      <el-button size="mini" type="success" v-if="scope.row.difficulty === 'Easy'">Easy</el-button>
                      <el-button size="mini" type="primary" v-if="scope.row.difficulty === 'Normal'">Normal</el-button>
                      <el-button size="mini" type="warning" v-if="scope.row.difficulty === 'Hard'">Hard</el-button>
                      <el-button size="mini" type="danger" v-if="scope.row.difficulty === 'Crazy'">Crazy</el-button>
                    </template>
                  </el-table-column>
                  <el-table-column prop="statistic.total" label="Total" min-width="80" align="center">
                  </el-table-column>
                  <el-table-column prop="ac_rate" label="AC Rate" min-width="80" align="center">
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
                </el-table>
              </el-row>
            </div>
          </el-col>
          <el-col :span="6">
            <div class="right-box" style="min-height:100px">
              <el-row>
                <el-col :span="6">
                  <div style="text-align:center;font-size:20px;line-height:20px">
                    <p style="padding-left:20px">Tags</p>
                  </div>
                </el-col>
              </el-row>
              <el-row style="width:96%;margin-left:4%;margin-top:-20px;min-height:30px" v-loading="tagsLoading">
                <el-button round size="small" class="tag-btn" v-bind:key="index" v-for="(tag, index) in tags"
                  @click="handleTagChange(tag.id)">{{ tag.name }}</el-button>
              </el-row>
            </div>
          </el-col>
          <el-col :span="18">
            <el-row style="">
              <div class="block">
                <el-pagination style="float:right;" background layout="prev, pager, next" :page-size="20"
                  @current-change="handlePageChange" :current-page="page" :total="count">
                </el-pagination>
              </div>
            </el-row>
          </el-col>
        </el-row>
      </div>
    </transition>
  </div>
</template>
<script>
  export default {
    data() {
      return {
        tags: [],
        hasFlag:false,
        showTags:false,
        realShowTags:false,
        tagsLoading: false,
        practiseListLoading: false,
        page: 1,
        show: false,
        count: 0,
        problemList: [],
        loading: false,
        difficulty: '',
        tid: 0,
        keywords: '',
        difficultyOptions: [{
            value: null,
            label: 'All'
          },
          {
            value: 'Casual',
            label: 'Casual'
          },
          {
            value: 'Easy',
            label: 'Easy'
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
      };
    },
    created() {
      this.$bus.emit('changeHeader', '2');
      this.show = false;
    },
    async mounted() {
      this.show = true;
      this.tagsLoading = true;
      await this.queryList();
      try {
        const {
          data: res
        } = await this.$http.post('/practice/getTags');
        console.log(res);
        if (res.error) {
          this.$message.error(res.error)
        } else {
          this.tags = res.data;
          this.tagsLoading = false;
        }
      } catch (err) {
        console.log(err);
        // alert(err)
      }
    },
    methods: {
      switchTags(val){
        this.practiseListLoading=true
        setTimeout(()=>{
                          this.realShowTags=val
        },200)
        setTimeout(()=>{
          this.practiseListLoading=false
        },500)
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
        if (this.$route.query.tid) {
          this.tid = Number(this.$route.query.tid);
        } else {
          this.tid = 0;
        }
        if (this.$route.query.keywords) {
          this.keywords = this.$route.query.keywords;
        } else {
          this.keywords = '';
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
        return obj;
      },
      goto(id) {
        this.$router.push({
          path: '/practice/answer',
          query: {
            id: id
          }
        });
      },
      async queryList() {
        this.practiseListLoading = true;
        this.paramsInit();
        try {
          const {
            data: res1
          } = await this.$http.post('/practice/getCount', {
            page: this.page,
            difficulty: this.difficulty,
            tid: this.tid,
            keywords: this.keywords
          });
          if (res1.error) {
            this.$message.error(res1.error);
            return;
          }
          const {
            data: res
          } = await this.$http.post('/practice/getAll', {
            page: this.page,
            difficulty: this.difficulty,
            tid: this.tid,
            keywords: this.keywords
          });
          if (res.error) {
            this.$message.error(res.error);
            return;
          }
          this.problemList = res.data;
          this.count = res1.data;

          this.hasFlag=false
          for (let i = 0; i < this.problemList.length; i++) {
            let rate =
              this.problemList[i].statistic.ac /
              this.problemList[i].statistic.total;
            this.problemList[i].ac_rate = '0';
            if (isNaN(rate)) {
              this.problemList[i].ac_rate = '--';
            } else {
              this.problemList[i].ac_rate = (rate * 100).toFixed(2) + '%';
            }
            if (this.problemList[i].flag!==""){
              this.hasFlag=true
            }
          }
          this.practiseListLoading = false;
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
      handleTagChange(val) {
        let obj = this.paramsQuery();
        obj.tid = Number(val);
        obj.page = 1;
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
          path: '/practice',
          query: obj
        });
      },
      reset() {
        this.$router.push({
          path: '/practice'
        });
      }
    },
    watch: {
      $route() {
        this.queryList();
      }
    },
    components: {}
  };
</script>

<style scoped>
  .center-box {
    margin: 20px auto 0;
    width: 95%;
    min-width: 1200px;
  }

  .left-box {
    background-color: #ffffff;
    border-radius: 6px;
    box-shadow: 0 2px 12px 0 rgba(0, 0, 0, 0.1);
  }

  .right-box {
    background-color: #ffffff;
    border-radius: 6px;
    box-shadow: 0 2px 12px 0 rgba(0, 0, 0, 0.1);
  }

  .left-box-head {
    height: 60px;
  }

  .tag-btn {
    margin-left: 0;
    margin-right: 5px;
    margin-bottom: 7px;
  }

  .el-row {
    margin-bottom: 20px;
  }

  .el-col {
    border-radius: 4px;
  }

  .grid-content {
    border-radius: 4px;
    min-height: 36px;
  }

  .row-bg {
    padding: 10px 0;
    background-color: #f9fafc;
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