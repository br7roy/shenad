import axios from 'axios'
import {mapState} from 'vuex'
import store from '@/store'

axios.defaults.timeout = 10000

axios.defaults.baseURL = ''
axios.defaults.headers.post['Content-Type'] = 'application/x-www-form-urlencoded;charset=UTF-8'

const param = {
  $store: store,
  ...mapState(['app.form.prdEnv', 'app.form.schUserName'])
}

const headers = {
  'Content-Type': 'application/json',
  'Accept': 'application/json',
  'token': window.localStorage.getItem['token']
}
if (process.env.NODE_ENV === 'production') {
  console.log('shenad run  product')
} else {
  // 为开发环境修改配置...
  console.log('shenad run dev')
}

export function get(url, params) {
  return new Promise((resolve, reject) => {
    axios.get(url, params, {
      headers: headers
    })
      .then(res => {
        resolve(res.data)
      })
      .catch(err => {
        reject(err.data)
      })
  })
}

export function post(url, params) {
  return new Promise((resolve, reject) => {
    axios.post(url, params, {
      headers: headers
    })
      .then(res => {
        resolve(res.data)
      })
      .catch(err => {
        reject(err.data)
      })
  })
}

// 设置cookie
export function setCookie(c_name, value, expiredays) {
  var exdate = new Date()
  exdate.setDate(exdate.getDate() + (expiredays * 24 * 60 * 60 * 1000))
  document.cookie = c_name + '=' + value + ((expiredays == null) ? '' : ';expires=' + exdate.toGMTString())
}
// 获取cookie
export function getCookie(cname) {
  var name = cname + '='
  var ca = document.cookie.split(';')
  for (var i = 0; i < ca.length; i++) {
    var c = ca[i].trim()
    if (c.indexOf(name) == 0) { return c.substring(name.length, c.length) }
  }
  return ''
}
// 删除cookie
export function delCookie(name) {
  var exp = new Date()
  exp.setTime(exp.getTime() - 1)
  var cval = getCookie(name)
  if (cval) {
    document.cookie = name + '=' + cval + ';expires=' + exp.toGMTString()
  }
}

// 时间戳转换
export function timestamp(inputTime) {
  var date = new Date(inputTime)
  var y = date.getFullYear()
  var m = date.getMonth() + 1
  m = m < 10 ? ('0' + m) : m
  var d = date.getDate()
  d = d < 10 ? ('0' + d) : d
  return y + m + d
}

export function dateFilter(time) {
  if (!time) {
    // 当时间是null或者无效格式时我们返回空
    return ''
  } else {
    const date = new Date(time) // 时间戳为10位需*1000，时间戳为13位的话不需乘1000
    const dateNumFun = (num) => num < 10 ? `0${num}` : num // 使用箭头函数和三目运算以及es6字符串的简单操作。因为只有一个操作不需要{} ，目的就是数字小于10，例如9那么就加上一个0，变成09，否则就返回本身。        // 这是es6的解构赋值。
    const [Y, M, D, h, m, s] = [
      date.getFullYear(),
      dateNumFun(date.getMonth() + 1),
      dateNumFun(date.getDate()),
      dateNumFun(date.getHours()),
      dateNumFun(date.getMinutes()),
      dateNumFun(date.getSeconds())
    ]
    return `${Y}${M}${D}` // 一定要注意是反引号，否则无效。
  }
}

// 设置请求拦截器
axios.interceptors.request.use((request) => {
  // config.headers.Authorization = 本地的token
  request.data = holeStageParamInterceptor(request.data)
  return request
})

function holeStageParamInterceptor(data) {
  if (data) {
    if (data.hasOwnProperty('startTime')) {
      data.startTime = timestamp(data.startTime)
    }
    if (data.hasOwnProperty('endTime')) {
      data.endTime = timestamp(data.endTime)
    }
    // data.setProperty('prdEnv', param.$store.getters.prdEnv)
    data.prdEnv = param.$store.getters.prdEnv
    data.schUserName = param.$store.getters.schUserName
    console.log('axios.')
    console.log(param.$store.getters.schUserName)
    console.log(data.schUserName)
  }
  console.log('requestData:\n')
  console.log(data)
  return data
}

// 设置响应拦截器
axios.interceptors.response.use(res => {
  return res
}, err => {
  // 出错的时候设置跳转到哪里
  return Promise.reject(err)
})

export default axios
