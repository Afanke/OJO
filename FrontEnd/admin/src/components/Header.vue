<template>
  <div>
    <el-row style="height:100%;bottom:0;">
      <el-menu :default-active="activeIndex" style="height:100%" @select="handleSelect">
        <el-row style="text-align:center">
          <h1>Hello OJO</h1>
        </el-row>

        <el-submenu index="1">
          <template slot="title">
            <i class="el-icon-location"></i>
            <span>导航一</span>
          </template>
          <el-menu-item-group>
            <template slot="title">分组一</template>
            <el-menu-item index="1-1">选项1</el-menu-item>
            <el-menu-item index="1-2">选项2</el-menu-item>
          </el-menu-item-group>
          <el-menu-item-group title="分组2">
            <el-menu-item index="1-3">选项3</el-menu-item>
          </el-menu-item-group>
          <el-submenu index="1-4">
            <template slot="title">选项4</template>
            <el-menu-item index="1-4-1">选项1</el-menu-item>
          </el-submenu>
        </el-submenu>
        <el-menu-item index="2">
          <i class="el-icon-menu"></i>
          <span slot="title">导航二</span>
        </el-menu-item>
        <el-submenu index="3">
          <template slot="title">
            <i class="el-icon-s-grid"></i>
            <span style="font-size:15px">Problem</span>
          </template>
          <el-menu-item index="3-1">Problem List</el-menu-item>
          <el-menu-item index="3-2">Create Problem</el-menu-item>
        </el-submenu>
        <el-submenu index="4">
          <template slot="title">
            <i class="el-icon-location"></i>
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
    created(){
      this.$bus.on("changeHeader",this.changeHeader)
    },
    methods: {
      changeHeader(val){
        this.activeIndex=val
      },
      handleSelect(key) {
        console.log(key);
        switch (key) {
          case "1":
            break
          case "2":
            break
          case "3-1":
            this.$router.push("/problem")
            break
          case "3-2":
            this.$router.push("/problem/create")
            console.log('asdad')
            break
            break
          case "4":
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