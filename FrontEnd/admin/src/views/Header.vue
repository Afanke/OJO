<template>
  <div class="container">
    <el-row style="height:100%;bottom:0;float:left;">
      <el-menu  :unique-opened="true" :default-active="activeIndex" style="height:100%" @select="handleSelect">
        <el-row >
          <img style="width:50%;margin-left:25%" src="../assets/rocket.png" alt="">
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
          <el-menu-item index="2-2">Announcement</el-menu-item>
          <el-menu-item index="2-3">Judge Server</el-menu-item>
          <el-menu-item index="2-4">System Config</el-menu-item>
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
      changeHeader(val) {
        this.activeIndex = val
      },
      handleSelect(key) {
        console.log(key);
        switch (key) {
          case "1":
            this.$router.push("/dashboard")
            break
          case "2-1":
            this.$router.push("/general/user")
            break
          case "2-2":
            this.$router.push("/general/announcement")
            break
          case "2-3":
            this.$router.push("/general/judgeServer")
            break
          case "2-4":
            this.$router.push("/general/system")
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
.container >>> .el-menu{
   border-right:0px!important;
}
</style>