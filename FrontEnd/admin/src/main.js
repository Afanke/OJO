import Vue from 'vue'
import App from './App.vue'
import router from './router'
import ElementUI from './element.config'
import 'element-ui/lib/theme-chalk/index.css'
import axios from 'axios'
import VueBus from 'vue-bus'


Vue.use(VueBus);

Vue.use(ElementUI)

// axios.defaults.baseURL = 'http://49.234.91.99'
// axios.defaults.baseURL = 'http://127.0.0.1'
axios.defaults.baseURL = ''
axios.defaults.withCredentials=true

Vue.prototype.$http = axios
Vue.prototype.$url=axios.defaults.baseURL
Vue.config.productionTip = true

String.prototype.replaceAll = function(s1,s2){
  return this.replace(new RegExp(s1,"gm"),s2);
}

new Vue({
  router,
  render: h => h(App)
}).$mount('#app')

