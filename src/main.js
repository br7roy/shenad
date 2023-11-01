import Vue from 'vue'

import 'normalize.css/normalize.css' // A modern alternative to CSS resets
import ElementUI from 'element-ui'
import 'element-ui/lib/theme-chalk/index.css'
import '@/styles/index.scss' // global css
import App from './App'
import store from './store'
import router from './router'

import '@/icons' // icon
import '@/permission' // permission control
import {
  delCookie,
  get,
  getCookie,
  post,
  setCookie,
  timestamp
} from './utils/axios'
import bus from '@/utils/bus'

Vue.prototype.$bus = bus
Vue.prototype.$get = get
Vue.prototype.$post = post
Vue.prototype.setCookie = setCookie //  设置cookie
Vue.prototype.getCookie = getCookie //  读取cookie
Vue.prototype.delCookie = delCookie //  删除cookie
Vue.prototype.timestamp = timestamp
/**
 * If you don't want to use mock-server
 * you want to use MockJs for mock api
 * you can execute: mockXHR()
 *
 * Currently MockJs will be used in the production environment,
 * please remove it before going online ! ! !
 */
if (process.env.NODE_ENV === 'development') {
  const {mockXHR} = require('../mock')
  mockXHR()
}

// set ElementUI lang to EN
// Vue.use(ElementUI, { locale })
// 如果想要中文版 element-ui，按如下方式声明
Vue.use(ElementUI)

Vue.config.productionTip = false

new Vue({
  el: '#app',
  router,
  store,
  render: h => h(App)
})
