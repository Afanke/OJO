<template>
  <div id="app">

    <el-container>
  <el-aside width="230px"><app-header></app-header></el-aside>

  <el-container>
    <el-main style="min-width:1050px">
      <router-view> </router-view>
      </el-main>
  </el-container>

</el-container>

  </div>
</template>
<script>
  import Header from '@/components/Header.vue';
  // import ScorllBar from "element-ui/lib/scrollbar"


  export default {
    name: 'app',
    components: {
      appHeader: Header,
      // elScrollbar: ScorllBar
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
  .el-aside{
    height: 100%;
    background-color: #ffffff;
  }

  .el-scrollbar__wrap {
    overflow-x: hidden !important;
  }

  .el-select-dropdown__list {
    margin: 0 0px 15px 0px !important;
  }
</style>