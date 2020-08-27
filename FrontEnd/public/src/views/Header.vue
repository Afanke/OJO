<template>
  <div>
    <el-row class="tac">
      <el-menu ref="MenuRef" :default-active="activeIndex" class="el-menu-demo" mode="horizontal" text-color="#545c64"
        active-text-color="#409EFF" @select="handleSelect">
        <el-menu-item id="ojName">{{OJName}}</el-menu-item>
        <el-menu-item index="1" class="emi emiw">
          <p class="el-icon-office-building"></p>
          Home
        </el-menu-item>
        <el-menu-item index="2" class="emi emiw">
          <p class="el-icon-s-grid"></p>
          Practice
        </el-menu-item>
        <el-menu-item index="3" class="emi emiw">
          <p class="el-icon-trophy"></p>
          Contests
        </el-menu-item>
        <el-submenu index="4">
          <template slot="title">
            <p class="el-icon-data-line"></p>
            Status
          </template>
          <el-menu-item index="4-1" class="emi">Practice Status</el-menu-item>
          <el-menu-item index="4-2" class="emi">Contest Status</el-menu-item>
        </el-submenu>
        <el-submenu index="5">
          <template slot="title">
            <p class="el-icon-s-data"></p>
            Rank
          </template>
          <el-menu-item index="5-1" class="emi">Practice Rank</el-menu-item>
        </el-submenu>
<!--        <el-menu-item index="6" class="emi emiw">-->
<!--          <p class="el-icon-info"></p>-->
<!--          About-->
<!--        </el-menu-item>-->
        <el-menu-item v-if="!isLogin" style="float:right">
          <el-button round @click="registerDrawer = true">Register</el-button>
        </el-menu-item>
        <el-menu-item v-if="!isLogin" class="emib">
          <el-button round @click="loginDrawer = true">Login</el-button>
        </el-menu-item>
        <el-menu-item v-if="isLogin" style="float:right;margin-top:-3px">
          <el-dropdown @command="handleCommand">
            <el-button size="mini" style="font-size:16px">
              {{ username }}
              <i class="el-icon-arrow-down el-icon--right" style="width:9px"></i>
            </el-button>
            <el-dropdown-menu slot="dropdown">
              <el-dropdown-item command="home">Home</el-dropdown-item>
              <el-dropdown-item command="settings">Settings</el-dropdown-item>
              <el-dropdown-item v-if="isAdmin" command="management">Management</el-dropdown-item>
              <el-dropdown-item command="logout" divided>Logout</el-dropdown-item>
            </el-dropdown-menu>
          </el-dropdown>
        </el-menu-item>
        <el-menu-item v-if="isLogin" class="emib" style="margin-top:-1px">
          <div class="block">
            <el-avatar shape="square" :size="40" :src="squareUrl"></el-avatar>
          </div>
        </el-menu-item>
      </el-menu>
    </el-row>

    <el-drawer title="Login" :visible.sync="loginDrawer" :before-close="handleClose" :with-header="false" size="450px">
      <el-form label-width="20%" :model="loginForm" v-loading="loginLoading" ref="loginFormRef" :rules="loginRules"
        status-icon>
        <div style="text-align:center;margin-top:20%">
          <span style="width:100%;font-size:50px">Welcome to OJO</span>
        </div>
        <el-divider></el-divider>
        <el-form-item prop="username">
          <el-input v-model="loginForm.username" prefix-icon="el-icon-user" style="width:80%" clearable
            placeholder="Username"></el-input>
        </el-form-item>
        <el-form-item prop="password">
          <el-input v-model="loginForm.password" prefix-icon="el-icon-lock" style="width:80%" show-password clearable
            placeholder="Password"></el-input>
        </el-form-item>
        <el-form-item prop="captcha">
          <el-col :span="11">
            <el-input v-model="loginForm.captcha" prefix-icon="el-icon-view" clearable placeholder="Captcha"></el-input>
          </el-col>
          <el-col class="line" :span="7">
            <el-image style="width: 100%; height: 42px;margin-left:19px" :src="captchaUrl" :fit="fit"
              @click="changeCaptcha"></el-image>
          </el-col>
        </el-form-item>
        <el-form-item label="">
          <el-button type="primary" style="width:80%" @click="login">Login</el-button>
        </el-form-item>
        <el-form-item label="">
          <div style="width:80%">
            <el-link type="primary" style="float:left;margin-top:-20px" @click="registerDrawer=true">Go to register!
            </el-link>
            <el-link type="primary" style="float:right;margin-top:-20px" @click="goResetPassword">
              Forget password?
            </el-link>
          </div>

        </el-form-item>
      </el-form>
    </el-drawer>

    <el-drawer title="Register" :visible.sync="registerDrawer" :before-close="handleClose" :with-header="false"
      size="450px">
      <el-form label-width="20%" :model="registerForm" v-loading="registerLoading" ref="registerFormRef"
        :rules="registerRule" status-icon>
        <div style="text-align:center;margin-top:20%">
          <span style="width:100%;font-size:50px">Welcome to OJO</span>
        </div>
        <el-divider></el-divider>
        <el-form-item prop="username">
          <el-input v-model="registerForm.username" prefix-icon="el-icon-user" style="width:80%" clearable
            placeholder="Username"></el-input>
        </el-form-item>
        <el-form-item label="" prop="email">
          <el-input v-model="registerForm.email" prefix-icon="el-icon-message" style="width:80%" clearable
            placeholder="Email"></el-input>
        </el-form-item>
        <el-form-item label="" prop="password">
          <el-input v-model="registerForm.password" prefix-icon="el-icon-lock" style="width:80%" show-password clearable
            placeholder="Password"></el-input>
        </el-form-item>
        <el-form-item label="" prop="passwordAgain">
          <el-input v-model="registerForm.passwordAgain" prefix-icon="el-icon-lock" style="width:80%" show-password
            clearable placeholder="Password Again"></el-input>
        </el-form-item>
        <el-form-item label="" prop="captcha">
          <el-col :span="11">
            <el-input v-model="registerForm.captcha" prefix-icon="el-icon-view" clearable placeholder="Captcha">
            </el-input>
          </el-col>
          <el-col class="line" :span="7">
            <el-image style="width: 100%; height: 42px;margin-left:19px" :src="captchaUrl" :fit="fit"
              @click="changeCaptcha"></el-image>
          </el-col>
        </el-form-item>
        <el-form-item label="">
          <el-button type="primary" style="width:80%" @click="register">Register</el-button>
        </el-form-item>
      </el-form>
    </el-drawer>
  </div>
</template>

<script>
  export default {
    data() {
      let checkPasswordAgain = (rule, value, callback) => {
        if (value !== this.registerForm.password) {
          callback(new Error("Password does not match"));
        } else {
          callback();
        }
      };
      let checkEmail = (rule, value, callback) => {
        let reg = /^([a-zA-Z]|[0-9])(\w|-)+@[a-zA-Z0-9]+\.([a-zA-Z]{2,4})$/;
        if (reg.test(value)) {
          callback();
        } else {
          callback(new Error("Email format is not correct"));
        }
      };
      return {
        OJName:"",
        squareUrl: "https://cube.elemecdn.com/9/c2/f0ee8a3c7c9638a54940382568c9dpng.png",
        fit: "contain",
        captchaUrl: this.$url + "/user/captcha",
        isLogin: false,
        activeIndex: "0",
        activeIndex2: "1",
        username: "",
        isAdmin:false,
        loginDrawer: false,
        registerDrawer: false,
        loginLoading: false,
        registerLoading: false,
        loginForm: {
          username: "",
          password: "",
          captcha: ""
        },
        registerForm: {
          username: "",
          email: "",
          password: "",
          passwordAgain: "",
          captcha: ""
        },
        loginRules: {
          username: [{
              required: true,
              message: "Username is required",
              trigger: "change"
            },
            {
              min: 2,
              max: 24,
              message: "The length of username must between 2 and 24 ",
              trigger: "change"
            }
          ],
          password: [{
              required: true,
              message: "password is required",
              trigger: "change"
            },
            {
              min: 8,
              max: 32,
              message: "The length of password must between 8 and 32 ",
              trigger: "change"
            }
          ],
          captcha: [{
            required: true,
            message: "captcha is required",
            trigger: "change"
          }, ]
        },
        registerRule: {
          username: [{
              required: true,
              message: "Username is required",
              trigger: "change"
            },
            {
              min: 2,
              max: 24,
              message: "The length of username must between 2 and 24 ",
              trigger: "change"
            }
          ],
          email: [{
              required: true,
              message: "Email is required",
              trigger: "change"
            },
            {
              min: 4,
              max: 32,
              message: "The length of email must between 4 and 32 ",
              trigger: "change"
            },
            {
              validator: checkEmail,
              trigger: "change"
            }
          ],
          password: [{
              required: true,
              message: "password is required",
              trigger: "change"
            },
            {
              min: 8,
              max: 32,
              message: "The length of email must between 8 and 32 ",
              trigger: "change"
            }
          ],
          passwordAgain: [{
              required: true,
              message: "password is required",
              trigger: "change"
            },
            {
              min: 8,
              max: 32,
              message: "The length of email must between 8 and 32 ",
              trigger: "change"
            },
            {
              validator: checkPasswordAgain,
              trigger: "change"
            }
          ],
          captcha: [{
            required: true,
            message: "captcha is required",
            trigger: "change"
          }, ]
        }
      };
    },
    created() {
      this.$bus.on("changeHeader", this.changeHeader);
      this.$bus.on("changeUserIcon", this.getDetail)
      this.$bus.on("OJName", this.changeOJName)
    },
    mounted() {
      this.getDetail()
    },
    methods: {
      changeOJName(name){
        this.OJName=name
      },
      async getDetail() {
        try {
          const {data: res} = await this.$http.post("/user/getDetail", {
            id: this.userId
          });
          if (res.error) {
            return
          }
          this.isAdmin=res.data.type>1
          this.isLogin = true;
          this.username = res.data.username;
          this.userId = res.data.id
          this.squareUrl = this.$http.defaults.baseURL + res.data.iconPath
        } catch (err) {
          console.log(err);
          alert(err);
        }
      },
      changeHeader(key) {
        this.activeIndex = key;
      },
      changeCaptcha() {
        this.captchaUrl = this.$url + "/user/captcha?k=" + Math.random();
      },
      handleSelect(key, keyPath) {
        console.log(key)
        switch (key) {
          case "1":
            this.$router.push("/home");
            break;
          case "2":
            this.$router.push("/practice");
            break;
          case "3":
            this.$router.push("/contest");
            break;
          case "4-1":
            this.$router.push("/status/practice");
            break;
          case "4-2":
            this.$router.push("/status/contest");
            break;
          case "5-1":
            this.$router.push("/rank/PracticeRank");
            break;
          case "6":
            this.$router.push("/about");
            break;
        }
      },
      login() {
        this.$refs.loginFormRef.validate(async valid => {
          if (!valid) return;
          try {
            this.loginLoading = true;
            const {
              data: res
            } = await this.$http.post("/user/login", this.loginForm);
            this.loginLoading = false;
            if (res.error) {
              this.$message.error(res.error);
              return;
            }
            this.loginDrawer = false;
            this.username = res.data.username;
            this.userId=res.data.id;
            this.isLogin = true;
            await this.getDetail()
            this.$message.success("Welcome " + this.username + " !");
          } catch (err) {
            console.log(err);
            alert(err);
            this.loginLoading = false;
          } finally {
            this.loginForm.captcha = "";
            this.loginForm.password = "";
            this.changeCaptcha();
          }
        });
      },
      register() {
        this.$refs.registerFormRef.validate(async valid => {
          if (!valid) return;
          try {
            this.registerLoading = true;
            const {
              data: res
            } = await this.$http.post(
              "/user/register",
              this.registerForm
            );
            this.registerLoading = false;
            if (res.error) {
              this.$message.error(res.error);
              return;
            }
            this.registerDrawer = false;
            this.$message.success("register successfully!");
          } catch (err) {
            console.log(err);
            alert(err);
            this.registerLoading = false;
          } finally {
            this.registerForm.captcha = "";
            this.registerForm.password = "";
            this.changeCaptcha();
          }
        });
      },
      async logout() {
        try {
          const {
            data: res
          } = await this.$http.post("/user/logout", this.registerForm);
          if (res.error) {
            this.$message.error(res.error);
            return;
          }
          this.username = "";
          this.isLogin = false;
          await this.$router.push("/home")
        } catch (err) {
          console.log(err);
          alert(err);
        }
      },
      handleClose(done) {
        done();
      },
      goHome() {
        this.$router.push({
          path: "/user/home",
          query: {
            id: this.userId
          }
        });
      },
      goSettings() {
        this.$router.push({
          path: "/user/settings",
        });
      },
      goManagement() {
        window.open("/admin")
      },
      goResetPassword(){
        this.loginDrawer=false
        this.$router.push({
          path: "/user/resetPassword",
        });
      },
      handleCommand(command) {
        switch (command) {
          case "logout":
            this.logout()
            break
          case "home":
            this.goHome()
            break
          case "settings":
            this.goSettings()
            break
          case "management":
            this.goManagement()
            break
        }

      }
    }
  };
</script>
<style scoped>
  .tac {
    width: 100%;
    min-width: 1250px;
    background-color: #ffffff;
  }

  #ojName {
    font-size: 20px;
  }

  .emi:hover {
    color: #ffffff !important;
    background-color: rgb(121, 187, 255) !important;
  }

  .emiw {
    width: 130px;
    text-align: center;
    /* color:transparent; */
  }

  .emib {
    float: right !important;
    margin-right: -20px !important;
    /* margin-left: -10px!important; */
    /* color:transparent; */

  }

  .el-menu-demo {
    box-shadow: 0 1px 8px 0 rgba(0, 0, 0, 0.1)
  }
</style>