import request from '@/utils/request'

export function listType(params) {
  return request({
    url: '/app/hxtp',
    method: 'get',
    params
  })
}
