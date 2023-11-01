import request from '@/utils/request'

export function login(data) {
  return request({
    // url: '/vue-admin-template/user/login',
    url: '/login',
    method: 'post',
    data
  })
}

export function getInfo(token) {
  return request({
    // url: '/vue-admin-template/user/info',
    url: '/usrinfo',
    method: 'get',
    params: { token }
  })
}

export function logout(token) {
  return request({
    // url: '/vue-admin-template/user/logout',
    url: '/logout',
    method: 'get',
    params: { token }
  })
}
