<template>

  <div>
    <transition name="slide-fade">
      <div class="center-box" v-if="show">
        <el-row style="height:60px;line-height:60px">
          <span style="font-size:20px;margin-left:20px">Announcement</span>
          <el-button style="float:right;margin-top:15px;margin-right:20px;" class="el-icon-refresh" type="primary"
            size="small" @click="reset">&nbsp;Reset</el-button>
          <el-input style="float:right;width:200px;margin-top:15px;margin-right:20px" placeholder="keywords"
            v-model="keywords" size="small">
            <el-button slot="append" icon="el-icon-search"
              style="color:#ffffff;background-color:#409EFF;border-top-left-radius:0;border-bottom-left-radius:0;margin-right:-21px;margin-top:-7px"
              size="small" @click="handleKeywordsChange"></el-button>
          </el-input>
          <div style="float:right;margin-top:16px;margin-right:80px">
            <my-switch v-model="mine" @toggle="toggle"></my-switch>
          </div>
        </el-row>
        <el-row style="height:1px;float:top;border-top:1px solid #EBEEF5;"></el-row>
        <el-table :data="tableData" style="width: 100%" v-loading="loading" size="small">
          <el-table-column label="ID" prop="id" align="center" min-width="30">
          </el-table-column>
          <el-table-column label="Title" prop="title" align="center" min-width="90">
          </el-table-column>
          <el-table-column label="Creator" prop="creatorName" align="center" min-width="30">
          </el-table-column>
          <el-table-column label="Create Time" prop="createTime" align="center" min-width="60">
          </el-table-column>
          <el-table-column label="Last Update Time" prop="lastUpdateTime" align="center" min-width="60">
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
                  <el-button size="mini" class="el-icon-edit-outline" @click="edit(scope.row.id)"></el-button>
                </el-tooltip>
                <el-tooltip content="Delete" placement="top">
                  <el-button size="mini" class="el-icon-delete" style="color:red" @click="delAnno(scope.row.id)">
                  </el-button>
                </el-tooltip>
              </el-row>
            </template>
          </el-table-column>
        </el-table>
        <el-row style="margin-top:20px;">
          <el-button type="primary" @click="goCreate" size="small" class="el-icon-plus"
            style="margin-left:30px;float:left">Create</el-button>
          <el-pagination @current-change="handlePageChange" :page-size="10" style="float:right;margin-right:30px"
            layout="prev, pager, next" :total="count" :current-page.sync="page">
          </el-pagination>
        </el-row>
        <el-row style="margin-top:20px;">
        </el-row>
      </div>
    </transition>
    <el-dialog title="Create Announcement" :visible.sync="createVisible" width="800px"
      :before-close="handleCreateClose">
      <el-row style="height:40px;line-height:40px;font-size:15px">
        <span style="color:red">*</span>
        <span>&nbsp;Title</span>
      </el-row>
      <el-row>
        <el-input v-model="title" placeholder="Title"></el-input>
      </el-row>
      <el-row style="height:40px;line-height:40px;font-size:15px">
        <span style="color:red">*</span>
        <span>&nbsp;Content</span>
      </el-row>
      <el-row>
        <editor v-model="content"></editor>
      </el-row>
      <el-row style="height:60px;line-height:60px;font-size:15px">
        <span style="color:red">*</span>
        <span style="margin-right:20px">&nbsp;Visible</span>
        <el-switch v-model="visible" active-color="#13ce66" inactive-color="#ff4949">
        </el-switch>
      </el-row>
      <el-row>
        <el-button style="float:right" type="primary" @click="saveCreate">Save</el-button>
        <el-button style="float:right;margin-right:10px" @click="handleCreateClose">Cancel</el-button>
      </el-row>
    </el-dialog>
    <el-dialog title="Edit Announcement" :visible.sync="editVisible" width="800px" :before-close="handleEditClose">
      <el-row style="height:40px;line-height:40px;font-size:15px">
        <span style="color:red">*</span>
        <span>&nbsp;Title</span>
      </el-row>
      <el-row>
        <el-input v-model="Edit.title" placeholder="Title"></el-input>
      </el-row>
      <el-row style="height:40px;line-height:40px;font-size:15px">
        <span style="color:red">*</span>
        <span>&nbsp;Content</span>
      </el-row>
      <el-row>
        <editor v-model="Edit.content"></editor>
      </el-row>
      <el-row style="height:60px;line-height:60px;font-size:15px">
        <span style="color:red">*</span>
        <span style="margin-right:20px">&nbsp;Visible</span>
        <el-switch v-model="Edit.visible" active-color="#13ce66" inactive-color="#ff4949">
        </el-switch>
      </el-row>
      <el-row>
        <el-button style="float:right" type="primary" @click="saveEdit">Save</el-button>
        <el-button style="float:right;margin-right:10px" @click="handleEditClose">Cancel</el-button>
      </el-row>
    </el-dialog>
  </div>

</template>
<script>
  import SwitchButton from "@/components/Switch.vue"
    import Editor from '@/components/Editor'
  export default {
    data() {
      return {
        show: false,
        loading: true,
        keywords: '',
        count: 0,
        page: 1,
        mine: false,
        tableData: [],
        createVisible: false,
        title: "",
        content: "",
        visible: false,
        editVisible: false,
        Edit: {
          id: 0,
          title: "",
          content: "",
          visible: false,
        }
      }
    },
    created() {
      this.$bus.emit("changeHeader", "2-2")
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
      createCheck() {
        if (this.content === "") {
          this.$message.error("content is required")
          return false
        }
        if (this.title === "") {
          this.$message.error("title is required")
          return false
        }
        return true
      },
      editCheck() {
        if (this.Edit.content === "") {
          this.$message.error("content is required")
          return false
        }
        if (this.Edit.title === "") {
          this.$message.error("title is required")
          return false
        }
        return true
      },
      async edit(id) {
        try {
          const {
            data: res
          } = await this.$http.post('/admin/announcement/getDetail', {
            id: id
          });
          if (res.error) {
            this.$message.error(res.error)
            return
          }
          this.Edit.id = id
          this.Edit.title = res.data.title
          this.Edit.content = res.data.content
          this.Edit.visible = res.data.visible
          this.editVisible = true
        } catch (err) {
          console.log(err);
          alert(err)
        }
      },
      async saveEdit() {
        if (!this.editCheck()) {
          return
        }
        try {
          const {
            data: res
          } = await this.$http.post('/admin/announcement/update', {
            id: this.Edit.id,
            title: this.Edit.title,
            content: this.Edit.content,
            visible: this.Edit.visible
          });
          if (res.error) {
            this.$message.error(res.error)
            return
          }
          this.$message({
            message: res.data,
            type: 'success'
          });
          this.handleEditClose()
          this.queryList()
        } catch (err) {
          console.log(err);
          alert(err)
        }
      },
      async saveCreate() {
        if (!this.createCheck()) {
          return
        }
        try {
          const {
            data: res
          } = await this.$http.post('/admin/announcement/add', {
            title: this.title,
            content: this.content,
            visible: this.visible
          });
          if (res.error) {
            this.$message.error(res.error)
            return
          }
          this.$message({
            message: res.data,
            type: 'success'
          });
          this.handleCreateClose()
          this.queryList()
        } catch (err) {
          console.log(err);
          alert(err)
        }
      },
      async delAnno(id) {
          try {
          const {
            data: res
          } = await this.$http.post('/admin/announcement/delete', {
            id:id
          });
          if (res.error) {
            this.$message.error(res.error)
            return
          }
          this.$message({
            message: res.data,
            type: 'success'
          });
          this.queryList()
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
        if (this.$route.query.keywords) {
          this.keywords = this.$route.query.keywords;
        } else {
          this.keywords = '';
        }
        if (this.$route.query.mine) {
          if (typeof (this.$route.query.mine) === typeof true) {
            this.mine = this.$route.query.mine
          } else {
            this.mine = this.$route.query.mine === "true";
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
        if (this.$route.query.keywords) {
          obj.keywords = this.$route.query.keywords;
        }
        obj.mine = !!this.$route.query.mine;
        return obj;
      },
      async queryList() {
        this.loading = true
        this.paramsInit()
        try {
          const {
            data: res
          } = await this.$http.post('/admin/announcement/getAll', {
            page: this.page,
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
          } = await this.$http.post('/admin/announcement/getCount', {
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
            } = await this.$http.post('/admin/announcement/setVisibleTrue', {
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
            } = await this.$http.post('/admin/announcement/setVisibleFalse', {
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
      handleCreateClose(done) {
        this.title = ""
        this.content = ""
        this.visible = false
        this.createVisible = false
        if (done) {
          done()
        }
      },
      handleEditClose(done) {
        this.Edit.id = 0
        this.Edit.title = ""
        this.Edit.content = ""
        this.Edit.visible = false
        this.editVisible = false
        if (done) {
          done()
        }
      },
      handlePageChange(val) {
        let obj = this.paramsQuery();
        obj.page = Number(val);
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
          path: '/general/announcement',
          query: obj
        });
      },
      reset() {
        this.$router.push({
          path: '/general/announcement'
        });
      },

      goCreate() {
        this.createVisible = true
      },

    },
    watch: {
      $route() {
        this.queryList();
      }
    },
    components: {
      mySwitch: SwitchButton,
      editor: Editor
    }
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