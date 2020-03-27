<template>
  <div id="app">
    <div class="container">
 
      <el-scrollbar style="height:100%;">
           <div class="header"><app-header></app-header></div>
      <div class="main" style="height:100%">
        <router-view> </router-view>
      </div>
      <!-- <div class="footer"></div> -->
         </el-scrollbar>
   </div>
  </div>
</template>
<script>
import NProgress from 'nprogress';
import 'nprogress/nprogress.css';
import Header from './components/Header.vue';


import ScorllBar from "element-ui/lib/scrollbar"


NProgress.start();
// 给页面设置请求加载进度条效果
window.onload = function() {
  NProgress.done();
};
export default {
  name: 'app',
  components: {
    appHeader: Header,
    elScrollbar:ScorllBar
  },
  data() {
    return {};
  }
};
</script>
<style scope>

#nprogress .bar {
  background: #409eff !important;
  height: 2px;
}
body {
  margin: 0;
  padding: 0;
  border: 0;
  height: 100%;
  background-color: rgb(244, 244, 245);
}
#app {
  width: 100%;
  height: 100%;
}
.container {
  height: 100%;
}
.header {
  width: 100%;
  min-width: 1200px;
  background-color: #ffffff;
  /* position: absolute;
  z-index:10; */
}
.main {
  width: 100%;
  min-width: 1200px;
  height: auto;
  bottom: 0 !important;
  z-index: 0;
  background-color: rgb(244, 244, 245);
  position: relative;
}
.footer{
  width: 100%;
  height:40px;
  /* background-color: red; */
}

.el-scrollbar__wrap {
   overflow-x: hidden!important;
}

</style>
