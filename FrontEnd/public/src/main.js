import Vue from 'vue'
import App from './App.vue'
import router from './router'
import ElementUI from './element.config'
import 'element-ui/lib/theme-chalk/index.css'
import lang from 'element-ui/lib/locale/lang/en'
import locale from 'element-ui/lib/locale'
import axios from 'axios'
import VueBus from 'vue-bus'
import NProgress from 'nprogress' // 引入nprogress插件
import 'nprogress/nprogress.css'  // 这个nprogress样式必须引入

Vue.use(VueBus);
Vue.use(ElementUI);
locale.use(lang)


// axios.defaults.baseURL = 'http://49.234.91.99'
axios.defaults.baseURL = 'http://127.0.0.1'
axios.defaults.withCredentials=true
NProgress.configure({showSpinner: false});
NProgress.configure({minimum:0.1});
NProgress.configure({ease:'cubic-bezier',speed:500});
router.beforeEach((to, from, next) => {
  NProgress.start()
  next()
})
router.afterEach(() => {
  NProgress.done()
})

String.prototype.replaceAll = function(s1,s2){
  return this.replace(new RegExp(s1,"gm"),s2);
}

// axios.interceptors.request.use(
//   config => {
//     NProgress.start() // 设置加载进度条(开始..)
//     return config
//   },
//   error => {
//     return Promise.reject(error)
//   }
// )
// // axios响应拦截器
// axios.interceptors.response.use(
//   function(response) {
//     NProgress.done() // 设置加载进度条(结束..)
//     return response
//   },
//   function(error) {
//     return Promise.reject(error)
//   }
// )

Vue.prototype.$http = axios
Vue.prototype.$url=axios.defaults.baseURL
Vue.config.productionTip = true
new Vue({
  router,
  render: h => h(App)
}).$mount('#app')

