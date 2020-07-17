<template>
  <div>
    <transition name="slide-fade">
      <div class="center-box" v-if="show">
        <el-page-header style="height:60px;line-height:60px;margin-left:20px" title="Back" @back="goBack"
          content="Contest Problem List">
        </el-page-header>
        <!-- <el-row style="height:60px;line-height:60px">
          <span style="font-size:20px;margin-left:20px">Contest Problem List</span>
        </el-row> -->
        <el-row style="height:1px;float:top;border-top:1px solid #EBEEF5;"></el-row>
        <el-table :data="tableData" style="width: 100%" v-loading="loading" size="small">
          <el-table-column label="ID" prop="id" align="center" min-width="30">
          </el-table-column>
          <el-table-column label="Display Id" prop="ref" align="center" min-width="30">
          </el-table-column>
          <el-table-column label="Title" prop="title" align="center" min-width="90">
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
          <el-table-column label="Option" width="120" align="center">
            <template slot-scope="scope">
              <el-row>
                <el-tooltip content="Edit" placement="top">
                  <el-button size="mini" class="el-icon-edit-outline" @click="editProblem(scope.row.id)"></el-button>
                </el-tooltip>
                <el-tooltip content="Delete From Contest" placement="top">
                  <el-button size="mini" class="el-icon-delete" style="color:red" @click="deleteProblem(scope.row.id)">
                  </el-button>
                </el-tooltip>
              </el-row>
            </template>
          </el-table-column>
        </el-table>
        <el-row style="margin-top:20px;">
          <el-button type="primary" @click="goImport" size="small" class="el-icon-plus"
            style="margin-left:30px;float:left">&nbsp;Import</el-button>
        </el-row>
        <el-row style="margin-top:20px;">
        </el-row>
        <el-dialog title="Import From Available Problem" :visible.sync="dialogVisible" class="dialog" width="1000px">
          <el-table :data="tableData2" style="margin-top:-25px" v-loading="loading2">
            <el-table-column label="Id" min-width="30" prop="id" align="center"></el-table-column>
            <el-table-column label="Display Id" min-width="30" prop="ref" align="center"></el-table-column>
            <el-table-column label="title" min-width="30" prop="title" align="center"></el-table-column>
            <el-table-column label="Creator" min-width="30" prop="creatorName" align="center"></el-table-column>
            <el-table-column prop="difficulty" align="center" label="Level" min-width="60">
              <template slot-scope="scope">
                <el-button size="mini" type="info" v-if="scope.row.difficulty === 'Casual'">Casual</el-button>
                <el-button size="mini" type="success" v-if="scope.row.difficulty === 'Eazy'">Eazy</el-button>
                <el-button size="mini" type="primary" v-if="scope.row.difficulty === 'Normal'">Normal</el-button>
                <el-button size="mini" type="warning" v-if="scope.row.difficulty === 'Hard'">Hard</el-button>
                <el-button size="mini" type="danger" v-if="scope.row.difficulty === 'Crazy'">Crazy</el-button>
              </template>
            </el-table-column>
            <el-table-column label="Visible" min-width="30" prop="visible" align="center">
              <template slot-scope="scope">
                <el-tag v-if="scope.row.visible" type="success" size="small">True
                </el-tag>
                <el-tag v-if="!scope.row.visible" type="danger" size="small">False
                </el-tag>
              </template>
            </el-table-column>
            <el-table-column label="Option" width="120" align="center">
              <template slot-scope="scope">
                <el-row>
                  <el-tooltip content="Add" placement="top">
                    <el-button size="mini" class="el-icon-plus" @click="addProblem(scope.row.id)" :disabled="isInList(scope.row.id)"></el-button>
                  </el-tooltip>
                </el-row>
              </template>
            </el-table-column>
          </el-table>
          <el-row style="margin-top:20px">
              <el-pagination @current-change="handlePageChange" :page-size="10" style="float:right;margin-right:30px;"
            layout="prev, pager, next" :total="count" :current-page.sync="page">
          </el-pagination>
          </el-row>
        
        </el-dialog>
      </div>
    </transition>
  </div>

</template>
<script>
  export default {
    data() {
      return {
        show: false,
        loading: true,
        tableData: [],
        tableData2: [],
        dialogVisible: false,
        page: 1,
        count: 0,
        loading2: false,
        id:0
      }
    },
    created() {
      this.$bus.emit("changeHeader", "4-1")
      this.show = false
    },
    async mounted() {
      this.show = true
      this.id=Number(this.$route.query.id)
      this.queryCtsPb()
    },
    methods: {
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
      isInList(id){
        if(!this.tableData){
          return
        }
        for(let i=0;i<this.tableData.length;i++){
          if(this.tableData[i].id===id){
            return true
          }
        }
        return false
      },
      deleteProblem() {},
      async queryCtsPb() {
        this.loading = true
        try {
          const {
            data: res
          } = await this.$http.post('/admin/contest/getCtsProblem', {
            id: this.id
          });
          if (res.error) {
            this.$message.error(res.error)
            return
          }
          this.tableData = res.data
          this.loading = false
        } catch (err) {
          console.log(err);
          alert(err)
        }
      },
      async queryAvaPb() {
        this.loading2 = true
        try {
          const {
            data: res
          } = await this.$http.post('/admin/problem/getAllShared', {
            page: this.page
          });
          if (res.error) {
            this.$message.error(res.error)
            return
          }
          this.tableData2 = res.data
          this.loading2 = false
        } catch (err) {
          console.log(err);
          alert(err)
        }
      },
      async addProblem(pid){
          try {
          const {
            data: res
          } = await this.$http.post('/admin/contest/addProblem', {
            cid: this.id,
            pid:pid
          });
          if (res.error) {
            this.$message.error(res.error)
            return
          }
          this.queryCtsPb()
          this.queryAvaPb()
          this.$message.success(res.data)
        } catch (err) {
          console.log(err);
          alert(err)
        }
      },
      async deleteProblem(pid){
          try {
          const {
            data: res
          } = await this.$http.post('/admin/contest/deleteProblem', {
            cid: this.id,
            pid: pid
          });
          if (res.error) {
            this.$message.error(res.error)
            return
          }
          this.queryCtsPb()
          this.$message.success(res.data)
        } catch (err) {
          console.log(err);
          alert(err)
        }
      },
      handlePageChange(val) {
        this.page=val
      },
      goBack() {
        this.$router.go(-1)
      },
      async goImport() {
        try {
          const {
            data: res
          } = await this.$http.post('/admin/problem/getSharedCount', {});
          if (res.error) {
            this.$message.error(res.error)
            return
          }
          this.count = res.data
        } catch (err) {
          console.log(err);
          alert(err)
        }
        this.dialogVisible = true
        this.queryAvaPb()
      },
    },
    watch: {
      page() {
        this.queryAvaPb();
      }
    },
    components: {}
  };
</script>

<style scoped>
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