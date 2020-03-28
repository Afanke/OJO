<template>
  <div>
    <el-row class="tac">
      <el-menu
        ref="MenuRef"
        :default-active="activeIndex"
        class="el-menu-demo"
        mode="horizontal"
        text-color="#545c64"
        active-text-color="#409EFF"
        @select="handleSelect"
      >
        <!-- <el-menu-item index="0" id="ojName">OJO</el-menu-item> -->
        <el-menu-item id="ojName">OJO</el-menu-item>
        <el-menu-item index="1" class="emi emiw"
          ><p class="el-icon-office-building"></p>
          Home</el-menu-item
        >
        <el-menu-item index="2" class="emi emiw"
          ><p class="el-icon-s-grid"></p>
          Practice</el-menu-item
        >
        <el-menu-item index="3" class="emi emiw"
          ><p class="el-icon-trophy"></p>
          Contests</el-menu-item
        >
        <el-menu-item index="4" class="emi emiw"
          ><p class="el-icon-data-line"></p>
          Status</el-menu-item
        >
        <el-menu-item index="5" class="emi emiw"
          ><p class="el-icon-s-data"></p>
          Rank</el-menu-item
        >
        <el-menu-item index="6" class="emi emiw"
          ><p class="el-icon-info"></p>
          About</el-menu-item
        >

        <!-- <el-menu-item index="7" class="emib"><el-button round>Register</el-button></el-menu-item> -->
        <el-menu-item v-if="!isLogined" style="float:right"
          ><el-button round @click="registerDrawer = true"
            >Register</el-button
          ></el-menu-item
        >
        <!-- <el-menu-item index="8" class="emib"><el-button round>Login</el-button></el-menu-item> -->
        <el-menu-item v-if="!isLogined" class="emib"
          ><el-button round @click="loginDrawer = true"
            >Login</el-button
          ></el-menu-item
        >
        <el-menu-item
          v-if="isLogined"
          style="float:right"
          @click="userDrawer = true"
        >
          <div class="block">
            <el-tag type="info">{{ username }}</el-tag>
          </div></el-menu-item
        >
        <el-menu-item v-if="isLogined" class="emib" @click="userDrawer = true">
          <div class="block">
            <el-avatar
              shape="square"
              :size="40"
              :src="squareUrl"
            ></el-avatar></div
        ></el-menu-item>
      </el-menu>
    </el-row>

    <el-drawer
      title="Login"
      :visible.sync="loginDrawer"
      :before-close="handleClose"
      :with-header="false"
      size="450px"
    >
      <el-form
        label-width="20%"
        :model="loginForm"
        v-loading="loginLoading"
        ref="loginFormRef"
        :rules="loginRules"
        status-icon
      >
        <div style="text-align:center;margin-top:20%">
          <span style="width:100%;font-size:50px">Welcom to OJO</span>
        </div>
        <el-divider></el-divider>
        <el-form-item prop="username">
          <el-input
            v-model="loginForm.username"
            prefix-icon="el-icon-user"
            style="width:80%"
            clearable
            placeholder="Username"
          ></el-input>
        </el-form-item>
        <el-form-item prop="password">
          <el-input
            v-model="loginForm.password"
            prefix-icon="el-icon-lock"
            style="width:80%"
            show-password
            clearable
            placeholder="Password"
          ></el-input>
        </el-form-item>
        <el-form-item prop="captcha">
          <el-col :span="11">
            <!-- <el-input v-model="loginForm.captcha" prefix-icon="el-icon-view" style="width:45%"></el-input> -->
            <el-input
              v-model="loginForm.captcha"
              prefix-icon="el-icon-view"
              clearable
              placeholder="Captcha"
            ></el-input>
          </el-col>
          <el-col class="line" :span="7">
            <el-image
              style="width: 100%; height: 42px;margin-left:19px"
              :src="captchaUrl"
              :fit="fit"
              @click="changeCaptcha"
            ></el-image>
          </el-col>
        </el-form-item>
        <el-form-item label="">
          <el-button type="primary" style="width:80%" @click="login"
            >Login</el-button
          >
        </el-form-item>
        <el-form-item label="" >
          <div style="width:80%">
                 <el-link type="primary" style="float:left;margin-top:-20px" @click="registerDrawer=true">Go to register!</el-link>
            <el-link type="primary" style="float:right;margin-top:-20px" @click="$message('under construction')">Forget password?</el-link>
          </div>
         
        </el-form-item>
      </el-form>
    </el-drawer>

    <el-drawer
      title="Register"
      :visible.sync="registerDrawer"
      :before-close="handleClose"
      :with-header="false"
      size="450px"
    >
      <el-form
        label-width="20%"
        :model="registerForm"
        v-loading="registerLoading"
        ref="registerFormRef"
        :rules="registerRule"
        status-icon
      >
        <div style="text-align:center;margin-top:20%">
          <span style="width:100%;font-size:50px">Welcom to OJO</span>
        </div>
        <el-divider></el-divider>
        <el-form-item prop="username">
          <el-input
            v-model="registerForm.username"
            prefix-icon="el-icon-user"
            style="width:80%"
            clearable
            placeholder="Username"
          ></el-input>
        </el-form-item>
        <el-form-item label="" prop="email">
          <el-input
            v-model="registerForm.email"
            prefix-icon="el-icon-message"
            style="width:80%"
            clearable
            placeholder="Email"
          ></el-input>
        </el-form-item>
        <el-form-item label="" prop="password">
          <el-input
            v-model="registerForm.password"
            prefix-icon="el-icon-lock"
            style="width:80%"
            show-password
            clearable
            placeholder="Password"
          ></el-input>
        </el-form-item>
        <el-form-item label="" prop="passwordAgain">
          <el-input
            v-model="registerForm.passwordAgain"
            prefix-icon="el-icon-lock"
            style="width:80%"
            show-password
            clearable
            placeholder="Password Again"
          ></el-input>
        </el-form-item>
        <el-form-item label="" prop="captcha">
          <el-col :span="11">
            <el-input
              v-model="registerForm.captcha"
              prefix-icon="el-icon-view"
              clearable
              placeholder="Captcha"
            ></el-input>
          </el-col>
          <el-col class="line" :span="7">
            <el-image
              style="width: 100%; height: 42px;margin-left:19px"
              :src="captchaUrl"
              :fit="fit"
              @click="changeCaptcha"
            ></el-image>
          </el-col>
        </el-form-item>
        <el-form-item label="">
          <el-button type="primary" style="width:80%" @click="register"
            >Register</el-button
          >
        </el-form-item>
      </el-form>
    </el-drawer>

    <el-drawer
      title="User"
      :visible.sync="userDrawer"
      :with-header="false"
      size="450px"
    >
      <div>
        <div style="text-align:center;margin-top:20%">
          <div class="block">
            <el-avatar :size="200" :src="circleUrl"></el-avatar>
          </div>
          <p style="font-size:30px">{{ username }}</p>
          <div>
            <el-divider></el-divider>
            <el-row style="width:100%">
              <el-col :span="12" :offset="6"
                ><el-button type="danger" style="width:100%" @click="logout"
                  >Logout</el-button
                ></el-col
              >
            </el-row>
          </div>
        </div>
      </div>
    </el-drawer>
  </div>
</template>

<script>
export default {
  // props: {
  //   index: "1"
  // },
  data() {
    var checkPasswordAgain = (rule, value, callback) => {
      if (value != this.registerForm.password) {
        callback(new Error("Password does not match"));
      } else {
        callback();
      }
    };
    var checkEmail = (rule, value, callback) => {
      var reg = /^([a-zA-Z]|[0-9])(\w|\-)+@[a-zA-Z0-9]+\.([a-zA-Z]{2,4})$/;
      if (reg.test(value)) {
        callback();
      } else {
        callback(new Error("Email format is not correct"));
      }
    };
    return {
      circleUrl:
        "https://cube.elemecdn.com/3/7c/3ea6beec64369c2642b92c6726f1epng.png",
      squareUrl:
        "https://cube.elemecdn.com/9/c2/f0ee8a3c7c9638a54940382568c9dpng.png",
      userDrawer: false,
      fit: "contain",
      captchaUrl: this.$url + "/user/captcha",
      isLogined: false,
      activeIndex: "0",
      activeIndex2: "1",
      loginDrawer: false,
      username: "",
      registerDrawer: false,
      // direction: 'rtl',
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
        username: [
          {
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
        password: [
          {
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
        captcha: [
          { required: true, message: "captcha is required", trigger: "change" },
        ]
      },
      registerRule: {
        username: [
          {
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
        email: [
          { required: true, message: "Email is required", trigger: "change" },
          {
            min: 4,
            max: 32,
            message: "The length of email must between 4 and 32 ",
            trigger: "change"
          },
          { validator: checkEmail, trigger: "change" }
        ],
        password: [
          {
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
        passwordAgain: [
          {
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
          { validator: checkPasswordAgain, trigger: "change" }
        ],
        captcha: [
          { required: true, message: "captcha is required", trigger: "change" },
        ]
      }
    };
  },
  created() {
    this.$bus.on("changeHeader", this.changeHeader);
  },
  async mounted() {
    try {
      const {data:res} = await this.$http.post("/user/getInfo");
      console.log(res)
      if (res.error) {
        // this.$message.error(res.error)
      } else {
        this.username = res.data.username;
        this.isLogined = true;
        this.$message.success("Welcome " + res.data.username + " !");
      }
    } catch (err) {
      console.log(err);
      alert(err);
    }
  },
  methods: {
    changeHeader(key) {
      this.activeIndex = key;
    },
    changeCaptcha() {
      this.captchaUrl = this.$url + "/user/captcha?k=" + Math.random();
    },
    handleSelect(key, keyPath) {
      switch (keyPath[0]) {
        case "1":
          this.$router.push("/home");
          break;
        case "2":
          this.$router.push("/practice");
          break;
        case "3":
          this.$router.push("/contest");
          break;
        case "4":
          this.$router.push("/status");
          break;
        case "5":
          this.$router.push("/rank");
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
          const {data:res} = await this.$http.post("/user/login", this.loginForm);
          // console.log(res)
          this.loginLoading = false;
          if (res.error) {
            this.$message.error(res.error);
            return;
          }
          this.loginDrawer = false;
          this.username = res.data.username;
          this.isLogined = true;
          this.$message.success("Welcome "+this.username+" !");
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
        // eslint-disable-next-line no-unused-vars
        // const { data: res } = await this.$http.post('login', this.loginForm)
        try {
          this.registerLoading = true;
          const {data:res} = await this.$http.post(
            "/user/register",
            this.registerForm
          );
          // console.log(res)
          this.registerLoading = false;
          if (res.error) {
            this.$message.error(res.error);
            return;
          }
          this.registerDrawer = false;
          this.$message.success("Register successfully,now you can log in");
        } catch (err) {
          console.log(err);
          alert(err);
          this.registerLoading = false;
        } finally {
          this.registerForm.captcha = "";
          this.registerForm.password = "";
          this.changeCaptcha();
        }
        // this.$router.push('/home')
      });
    },
    async logout() {
      try {
        const {data:res} = await this.$http.post("/user/logout", this.registerForm);
        // console.log(res)
        if (res.error) {
          this.$message.error(res.error);
          return;
        }
        this.userDrawer = false;
        this.username = "";
        this.isLogined = false;
        this.$message.success("log out successfully");
      } catch (err) {
        console.log(err);
        alert(err);
      }
    },
    handleClose(done) {
      done();
      // this.$confirm('确认关闭？')
      //   .then(_ => {
      //     done();
      //   })
      //   .catch(_ => {});
    }
  }
};
</script>
<style scoped>
.tac{
  width: 100%;
  min-width: 1200px;
  background-color: #ffffff;
}
#ojName {
  font-size: 30px;
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
/* .el-menu-demo{
  opacity: 0.1;/
} */
</style>
