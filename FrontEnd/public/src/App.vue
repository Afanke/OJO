<template>
  <div id="app">
    <div class="container">

      <el-scrollbar style="height:100%;">
        <div class="header">
          <app-header></app-header>
        </div>
        <div class="main">
          <router-view> </router-view>
        </div>
        <div class="footer"></div>
      </el-scrollbar>
    </div>
  </div>
</template>
<script>
  import NProgress from 'nprogress';
  import 'nprogress/nprogress.css';
  import AppHeader from '@/components/Header.vue';


  import ScorllBar from "element-ui/lib/scrollbar"


  NProgress.start();
  // 给页面设置请求加载进度条效果
  window.onload = function () {
    NProgress.done();
  };
  export default {
    name: 'app',
    components: {
      appHeader: AppHeader,
      elScrollbar: ScorllBar
    },
    data() {
      return {
      };
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
      "微软雅黑",
      "Microsoft YaHei",

      "PingFang SC",
      "Hiragino Sans GB",
      "Helvetica Neue",
      Helvetica,
      Arial,
      sans-serif;
  }

  #nprogress .bar {
    background: #409eff !important;
    height: 2px;
  }

  #app {
    width: 100%;
    height: 100%;
  }

  .container {
    height: 100%;
  }

  /* .header { */
  /* width: 100%;
  min-width: 1200px;
  background-color: #ffffff; */
  /* position: absolute;
  z-index:10; */
  /* } */
  .main {
    margin-top:20px;
    min-width:1000px;
    /* float:right;
    width:auto;
    right:20px;
    left:250px; */
    background-color: rgb(244, 244, 245);
    /* position: relative; */
  }

  .footer {
    width: 100%;
    height: 40px;
    /* background-color: red; */
  }

  .el-scrollbar__wrap {
    overflow-x: hidden !important;
  }

  .el-select-dropdown__list {
    margin: 0 0px 15px 0px !important;
  }
</style>