<template>
  <div id="app">
    <el-container>
      <el-aside width="230px">
        <app-header></app-header>
      </el-aside>
        <!-- <el-scrollbar style="height:100%"> -->

      <el-container>

        <el-header style="padding:0;background-color:#ffffff" height="50px" >
          <i class="el-icon-refresh-right"></i>
        </el-header>

          <el-main style="min-width:1050px">
            <router-view>
            </router-view>
          </el-main>
        <!-- <el-footer> -->
        <!-- </el-footer> -->
      </el-container>
        <!-- </el-scrollbar> -->

    </el-container>
  </div>
</template>
<script>
  import Header from '@/components/Header.vue';
  import ScorllBar from "element-ui/lib/scrollbar"


  export default {
    name: 'app',
    components: {
      appHeader: Header,
      elScrollbar: ScorllBar
    },
    data() {
      return {};
    },

  };
</script>
<style scope>
  html,
  body {
    height: 100%;
    width: 100%;
    padding: 0;
    margin: 0;
    border: 0;
    background-color: rgb(244, 244, 245);
    font-family:
      "Helvetica Neue",
      Helvetica,
      "PingFang SC",
      "Hiragino Sans GB",
      "Microsoft YaHei",
      "微软雅黑",
      Arial,
      sans-serif;
  }

  #app {
    width: 100%;
    height: 100%;
    display: flex;
  }

  .el-aside {
    height: 100%;
    background-color: #ffffff;
  }

.el-scrollbar__wrap.default-scrollbar__wrap {
  overflow-x: auto;
}
</style>