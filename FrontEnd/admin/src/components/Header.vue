<template>
  <div>
    <el-row style="height:100%;bottom:0;">
      <el-menu :default-active="activeIndex" style="height:100%" @select="handleSelect">
        <el-row style="text-align:center">
          <h1>Hello OJO</h1>
        </el-row>

        <el-menu-item index="1">
          <i class="el-icon-odometer"></i>
          <span slot="title">Dash Board</span>
        </el-menu-item>
        <el-submenu index="2">
          <template slot="title">
            <i class="el-icon-set-up"></i>
            <span style="font-size:15px">General</span>
          </template>
          <el-menu-item index="2-1">User</el-menu-item>
          <el-menu-item index="2-2">Judge Server</el-menu-item>
        </el-submenu>
        <el-submenu index="3">
          <template slot="title">
            <i class="el-icon-s-grid"></i>
            <span style="font-size:15px">Problem</span>
          </template>
          <el-menu-item index="3-1">Problem List</el-menu-item>
          <el-menu-item index="3-2">Tag List</el-menu-item>
        </el-submenu>
        <el-submenu index="4">
          <template slot="title">
            <i class="el-icon-trophy"></i>
            <span style="font-size:15px">Contest</span>
          </template>
          <el-menu-item index="4-1">Contest List</el-menu-item>
          <el-menu-item index="4-2">Create Contest</el-menu-item>
        </el-submenu>
      </el-menu>
      <el-button @click="login1">root</el-button>
      <el-button @click="login2">admin1</el-button>
      <el-button @click="login3">admin2</el-button>
    </el-row>
  </div>
</template>

<script>
  export default {
    data() {
      return {
        activeIndex: '3-1',
      }
    },
    created() {
      this.$bus.on("changeHeader", this.changeHeader)
    },
    methods: {
      async login1() {
        try {
          const {
            data: res
          } = await this.$http.post('/user/login1', {
            username: 'root',
            password: '11111111'
          });
          if (res.error) {
            this.$message.error(res.error)
            return
          }
          this.$message({
            message: res.data,
            type: 'success'
          });
        } catch (err) {
          console.log(err);
          alert(err)
        }
      },
      async login2() {
        try {
          const {
            data: res
          } = await this.$http.post('/user/login1', {
            username: 'admin1',
            password: '11111111'
          });
          if (res.error) {
            this.$message.error(res.error)
            return
          }
          this.$message({
            message: res.data,
            type: 'success'
          });
        } catch (err) {
          console.log(err);
          alert(err)
        }
      },
      async login3() {
        try {
          const {
            data: res
          } = await this.$http.post('/user/login1', {
            username: 'admin2',
            password: '11111111'
          });
          if (res.error) {
            this.$message.error(res.error)
            return
          }
          this.$message({
            message: res.data,
            type: 'success'
          });
        } catch (err) {
          console.log(err);
          alert(err)
        }
      },
      changeHeader(val) {
        this.activeIndex = val
      },
      handleSelect(key) {
        console.log(key);
        switch (key) {
          case "1":
            break
          case "2-1":
            this.$router.push("/general/user")
            break
          case "2-2":
            this.$router.push("/general/judgeServer")
            break
          case "3-1":
            this.$router.push("/problem")
            break
          case "3-2":
            this.$router.push("/problem/tag")
            break
          case "4-1":
            this.$router.push("/contest")
            break
          case "4-2":
            this.$router.push("/contest/create")
            break
        }


      },

    },
    components: {

    }
  }
</script>

<style scoped>

</style>