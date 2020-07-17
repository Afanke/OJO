<template>

  <div>
    <transition name="slide-fade">

      <div class="center-box" v-if="show">
        <el-row style="height:60px;line-height:60px">
          <span style="font-size:20px;margin-left:20px">User</span>
          <el-button style="float:right;margin-top:15px;margin-right:20px;" class="el-icon-refresh" type="primary"
            size="small" @click="reset">&nbsp;Reset</el-button>
          <el-input style="float:right;width:200px;margin-top:15px;margin-right:20px" placeholder="keywords"
            v-model="keywords" size="small">
            <el-button slot="append" icon="el-icon-search"
              style="color:#ffffff;background-color:#409EFF;border-top-left-radius:0;border-bottom-left-radius:0;margin-right:-21px;margin-top:-7px"
              size="small" @click="handleKeywordsChange"></el-button>
          </el-input>
          <el-select v-model="type" placeholder="Type"
            style="margin-top:1.2px;margin-right:20px;float:right;width:140px" size="small" @change="handleTypeChange">
            <el-option v-for="item in TypeOptions" :key="item.value" :label="item.label" :value="item.value">
            </el-option>
          </el-select>
          <div style="margin-right:20px;float:right">
            <span style="color:grey;font-size:13px">Show Enabled:&nbsp;</span>
            <el-switch @change="switchShowEnabled" v-model="showEnabled">
            </el-switch>
          </div>
        </el-row>
        <el-row style="height:1px;float:top;border-top:1px solid rgb(233, 233, 235);"></el-row>
        <el-table :data="tableData" style="width: 100%" v-loading="loading" size="small">
          <el-table-column label="ID" prop="id" align="center" min-width="30">
          </el-table-column>
          <el-table-column label="Username" prop="username" align="center" min-width="50">
          </el-table-column>
          <el-table-column label="Real Name" prop="realName" align="center" min-width="50">
          </el-table-column>
          <el-table-column label="Email" prop="email" align="center" min-width="60">
          </el-table-column>
          <el-table-column label="Type" align="center" min-width="50">
            <template slot-scope="scope">
              {{formatType(scope.row.type)}}
            </template>
          </el-table-column>
          <el-table-column label="Create Time" prop="createTime" align="center" min-width="60">
          </el-table-column>
          <el-table-column label="Last Login Time" prop="lastLoginTime" align="center" min-width="60">
          </el-table-column>
          <el-table-column label="Enabled" width="80" v-if="realShowEnabled" align="center" min-width="40">
            <template slot-scope="scope">
              <div @click="switchEnabled(scope.row)">
                <el-switch v-model="scope.row.enabled" active-color="#409eff" inactive-color="#dcdfe6">
                </el-switch>
              </div>
            </template>
          </el-table-column>
          <el-table-column label="Option" width="120" align="center">
            <template slot-scope="scope">
              <el-row>
                <el-tooltip content="Edit" placement="top">
                  <el-button size="mini" class="el-icon-edit-outline" @click="editUser(scope.row)"></el-button>
                </el-tooltip>
                <!-- <el-tooltip content="Delete" placement="top">
                  <el-button size="mini" class="el-icon-delete" style="color:red" @click="deleteUser(scope.row)">
                  </el-button>
                </el-tooltip> -->
              </el-row>
            </template>
          </el-table-column>
        </el-table>
        <el-row style="margin-top:20px;height:40px">
          <el-pagination @current-change="handlePageChange" :page-size="10" style="float:right;margin-right:30px;"
            layout="prev, pager, next" :total="count" :current-page.sync="page">
          </el-pagination>
        </el-row>
      </div>
    </transition>
    <el-dialog :visible.sync="showUserDetail" width="850px" :before-close="handleClose">
      <div slot="title" style="height:0px">
        User Detail
      </div>
      <el-row :gutter="20">
        <el-col :span="3" style="line-height:40px">
          <span style="color:red">*</span>
          <span>&nbsp;Username</span>
        </el-col>
        <el-col :span="9">
          <el-input v-model="userDetail.username" placeholder="Username"></el-input>
        </el-col>
        <el-col :span="3" style="line-height:40px">
          <span style="color:#fff">*</span>
          <span>&nbsp;Real Name</span>
        </el-col>
        <el-col :span="9">
          <el-input v-model="userDetail.realName" placeholder="Real Name"></el-input>
        </el-col>
      </el-row>
      <el-row :gutter="20" style="margin-top:20px">
        <el-col :span="3" style="line-height:40px">
          <span style="color:red">*</span>
          <span>&nbsp;Email</span>
        </el-col>
        <el-col :span="9">
          <el-input v-model="userDetail.email" placeholder="Email"></el-input>
        </el-col>
        <el-col :span="3" style="line-height:40px">
          <span style="color:#fff">*</span>
          <span>&nbsp;Password</span>
        </el-col>
        <el-col :span="9">
          <el-input show-password clearable v-model="password" placeholder="Password"></el-input>
        </el-col>
      </el-row>
      <el-row :gutter="20" style="margin-top:20px">
        <el-col :span="3" style="line-height:40px">
          <span style="color:#fff">*</span>
          <span>&nbsp;School</span>
        </el-col>
        <el-col :span="9">
          <el-input v-model="userDetail.school" placeholder="School"></el-input>
        </el-col>
        <el-col :span="3" style="line-height:40px">
          <span style="color:#fff">*</span>
          <span>&nbsp;Major</span>
        </el-col>
        <el-col :span="9">
          <el-input v-model="userDetail.major" placeholder="Major"></el-input>
        </el-col>
      </el-row>
      <el-row :gutter="20" style="margin-top:20px">
        <el-col :span="3" style="line-height:40px">
          <span style="color:red">*</span>
          <span>&nbsp;Type</span>
        </el-col>
        <el-col :span="6">
          <el-select v-model="userDetail.type" placeholder="Type" >
            <el-option v-for="item in TypeOptions2" :key="item.value" :label="item.label" :value="item.value">
            </el-option>
          </el-select>
        </el-col>
      </el-row>
      <el-row :gutter="20" style="margin-top:20px">
        <el-col :span="3" style="line-height:40px">
          <span style="color:red">*</span>
          <span>&nbsp;Icon Path</span>
        </el-col>
        <el-col :span="21">
          <el-input v-model="userDetail.iconPath" placeholder="Icon Path"></el-input>
        </el-col>
      </el-row>
      <el-row :gutter="20" style="margin-top:20px">
        <el-col :span="3" style="line-height:40px">
          <span style="color:#fff">*</span>
          <span>&nbsp;Signature</span>
        </el-col>
        <el-col :span="21">
          <el-input v-model="userDetail.signature" placeholder="Signature"></el-input>
        </el-col>
      </el-row>
      <el-row :gutter="20" style="margin-top:20px">
        <el-col :span="3" style="line-height:40px">
          <span style="color:#fff">*</span>
          <span>&nbsp;Github</span>
        </el-col>
        <el-col :span="21">
          <el-input v-model="userDetail.github" placeholder="Github"></el-input>
        </el-col>
      </el-row>
      <el-row :gutter="20" style="margin-top:20px">
        <el-col :span="3" style="line-height:40px">
          <span style="color:#fff">*</span>
          <span>&nbsp;Blog</span>
        </el-col>
        <el-col :span="21">
          <el-input v-model="userDetail.blog" placeholder="Blog"></el-input>
        </el-col>
      </el-row>
      <div slot="footer" style="height:5px">
        <div style="float:right;margin-top:-25px">
          <el-button @click="handleClose2">Cancel</el-button>
          <el-button type="primary" @click="saveUserDetail">Save</el-button>
        </div>
      </div>

    </el-dialog>
  </div>

</template>
<script>
  export default {
    data() {
      return {
        password: '',
        show: false,
        type: 0,
        count: 0,
        showEnabled: false,
        realShowEnabled: false,
        loading: true,
        page: 1,
        showUserDetail: false,
        keywords: '',
        userDetail: {},
        tableData: [],
        TypeOptions: [{
            value: 0,
            label: 'All',
          },
          {
            value: 1,
            label: 'Normal User',
          },
          {
            value: 2,
            label: 'Admin',
          },
          {
            value: 3,
            label: 'Super Admin',
          },
        ],
        TypeOptions2: [{
            value: 1,
            label: 'Normal User',
          },
          {
            value: 2,
            label: 'Admin',
          },
          {
            value: 3,
            label: 'Super Admin',
          },
        ]
      }
    },
    created() {
      this.$bus.emit("changeHeader", "2-1")
    },
    async mounted() {
      this.show = true
      this.queryList()
    },
    methods: {
      handleClose2() {
        this.showUserDetail = false
        this.password = ''
      },
      handleClose(done) {
        this.password = ''
        done()
      },
      async editUser(val) {
        console.log(val)
        try {
          const {
            data: res
          } = await this.$http.post('/admin/user/getDetail', {
            id: val.id
          });
          if (res.error) {
            this.$message.error(res.error)
            return
          }
          this.userDetail = res.data
          this.showUserDetail = true
          console.log(this.userDetail)
        } catch (err) {
          console.log(err);
          alert(err)
        }
      },
      async saveUserDetail() {
        try {
          const {
            data: res
          } = await this.$http.post('/admin/user/updateDetail', {
            id: this.userDetail.id,
            username: this.userDetail.username,
            password: this.password,
            realName: this.userDetail.realName,
            type: this.userDetail.type,
            email: this.userDetail.email,
            signature: this.userDetail.signature,
            school: this.userDetail.school,
            blog: this.userDetail.blog,
            major: this.userDetail.major,
            github: this.userDetail.github,
            iconPath: this.userDetail.iconPath,
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
          this.showUserDetail = false
          this.password = ''
        } catch (err) {
          console.log(err);
          alert(err)
        }
      },
      async switchEnabled(obj){
        console.log(obj)
        if (obj.enabled) {
          try {
            const {
              data: res
            } = await this.$http.post('/admin/user/enable', {
              id:obj.id
            });
            if (res.error) {
              this.$message.error(res.error)
              obj.enabled=false
              return
            }
              obj.enabled=true
          } catch (err) {
            console.log(err);
          }
        } else {
          try {
            const {
              data: res
            } = await this.$http.post('/admin/user/disable', {
              id:obj.id
            });
            if (res.error) {
              this.$message.error(res.error)
              obj.enabled=true
              return
            }
              obj.enabled=false
          } catch (err) {
            console.log(err);
            // alert(err)
          }
        }
      },
      deleteUser(val) {
        this.$message("under construction")
      },
      initParams() {
        if (this.$route.query.page) {
          this.page = Number(this.$route.query.page);
        } else {
          this.page = 1;
        }
        if (this.$route.query.type) {
          this.type = Number(this.$route.query.type);
        } else {
          this.type = 0;
        }
        if (this.$route.query.keywords) {
          this.keywords = this.$route.query.keywords;
        } else {
          this.keywords = '';
        }
      },
      queryParams() {
        let obj = {};
        if (this.$route.query.page) {
          obj.page = Number(this.$route.query.page);
        }
        if (this.$route.query.type) {
          obj.type = Number(this.$route.query.type);
        }
        if (this.$route.query.keywords) {
          obj.keywords = this.$route.query.keywords;
        }
        return obj;
      },
      switchShowEnabled(val) {
        this.loading = true
        setTimeout(() => {
          this.realShowEnabled = val
        }, 200)
        setTimeout(() => {
          this.loading = false
        }, 500)
      },
      async queryList() {
        this.loading = true
        this.initParams()
        try {
          const {
            data: res
          } = await this.$http.post('/admin/user/getAll', {
            page: this.page,
            type: this.type,
            keywords: this.keywords
          });
          if (res.error) {
            this.$message.error(res.error)
            return
          }
          this.tableData = res.data
          const {
            data: res1
          } = await this.$http.post('/admin/user/getCount', {
            type: this.type,
            keywords: this.keywords
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
      handleTypeChange(val) {
        let obj = this.queryParams();
        obj.type = Number(val);
        obj.page = 1;
        this.fresh(obj);
      },
      handlePageChange(val) {
        let obj = this.queryParams();
        obj.page = Number(val);
        this.fresh(obj);
      },
      handleKeywordsChange() {
        let obj = this.queryParams();
        obj.keywords = this.keywords;
        obj.page = 1;
        this.fresh(obj);
      },
      fresh(obj) {
        this.$router.push({
          path: '/general/user',
          query: obj
        });
      },
      reset() {
        this.$router.push({
          path: '/general/user'
        });
      },
      formatType(val) {
        switch (val) {
          case 1:
            return "Normal User"
          case 2:
            return "Admin"
          case 3:
            return "Super Admin"
        }
      },
    },
    components: {},
    watch: {
      $route() {
        this.queryList();
      }
    },
  };
</script>

<style scoped>
  .center-box {
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