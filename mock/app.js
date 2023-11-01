const Mock = require('mockjs')

const d = Mock.mock({
  'code': 0,
  'iterm': [{
    Name: '画像1',
    Type: 0
  }, {
    Name: '画像2',
    Type: 1
  }, {
    Name: '画像3',
    Type: 2
  }]
})

module.exports = [
  {
    url: '/app/hxtp',
    type: 'get',
    response: config => {
      const item = d.iterm
      return {
        code: 0,
        data: item
      }
    }
  }
]
