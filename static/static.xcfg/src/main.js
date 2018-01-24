import Vue from 'vue'
import 'xiaomimi'

import router from './router/router'

import App from './App.vue'
document.title='xcfg -- 远程配置文件管理系统'
new Vue({
  router,
  el: '#app',
  render: h => h(App)
})
