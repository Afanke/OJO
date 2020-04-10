// The Vue build version to load with the `import` command
// (runtime-only or standalone) has been set in webpack.base.conf with an alias.
import Vue from 'vue'

import router from './router'
import '@vant/touch-emulator';
import axios from 'axios'
import App from './App'


Vue.prototype.axios=axios;


Vue.config.productionTip = false




import { NavBar } from 'vant';
import 'vant/lib/nav-bar/style';
Vue.use(NavBar);
import { Col, Row } from 'vant';
import 'vant/lib/col/style';
import 'vant/lib/row/style';
Vue.use(Col);
Vue.use(Row);
import { Popup } from 'vant';
import 'vant/lib/popup/style';
Vue.use(Popup);
import { Field } from 'vant';
import 'vant/lib/field/style';
Vue.use(Field);
import { Cell, CellGroup } from 'vant';
import 'vant/lib/cell/style';
import 'vant/lib/cell-group/style';
Vue.use(Cell);
Vue.use(CellGroup);
import { Button } from 'vant';
import 'vant/lib/button/style';
Vue.use(Button);
/* eslint-disable no-new */


new Vue({
  router,
  render: h => h(App)
}).$mount('#app')
