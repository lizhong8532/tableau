// The Vue build version to load with the `import` command
// (runtime-only or standalone) has been set in webpack.base.conf with an alias.
import Vue from 'vue'
import App from './App'
import router from './router'

import '../semantic/dist/semantic.min.css'
import '../node_modules/c3/c3.min.css'

import '../node_modules/jquery/dist/jquery.min'
import '../semantic/dist/semantic.min'

window.list = []
window.labels = []

Vue.config.productionTip = false

/* eslint-disable no-new */
new Vue({
  el: '#app',
  router,
  template: '<App/>',
  components: { App }
})

window.drop = (ev) => {
  ev.preventDefault()
  let data = ev.dataTransfer.getData('text')
  if (data.length > 10) {
    if (list.indexOf(data) === -1) {
      list.push(data)
      chartComponent.pushList(transfer(data))
    } else {
      messageComponents.show('This one has been over there, Make sure can\'t duplicate drag the same one.', 'Warning')
    }
  }
}

window.drag = (ev) => {
  ev.dataTransfer.setData('text', ev.target.getAttribute('data'))
}

window.transfer = (text) => {
  let o = {}
  let arr = text.split('###')
  try {
    o.values = JSON.parse(arr[1]).values
    o.valuesJson = JSON.stringify(o.values)
  } catch (e) {
    o.values = []
    o.valuesJson = '[]'
  }
  let subArr = arr[0].split('-')
  o.id = arr[0]
  o.sheet = subArr[2]
  o.title = subArr[3]
  o.sheetIndex = Number(subArr[0])
  return o
}
