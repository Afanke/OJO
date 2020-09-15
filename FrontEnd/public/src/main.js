import Vue from 'vue'
import App from './App.vue'
import router from './router'
import ElementUI from './element.config'
import 'element-ui/lib/theme-chalk/index.css'
import lang from 'element-ui/lib/locale/lang/en'
import locale from 'element-ui/lib/locale'
import axios from 'axios'
import VueBus from 'vue-bus'
import NProgress from 'nprogress'
import 'nprogress/nprogress.css'

Vue.use(VueBus);
Vue.use(ElementUI);
locale.use(lang)


// axios.defaults.baseURL = 'http://127.0.0.1'
axios.defaults.baseURL = ''
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

Vue.prototype.$http = axios
Vue.prototype.$url=axios.defaults.baseURL
Vue.config.productionTip = true
new Vue({
  router,
  render: h => h(App)
}).$mount('#app')

