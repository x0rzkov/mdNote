// The Vue build version to load with the `import` command
// (runtime-only or standalone) has been set in webpack.base.conf with an alias.
import Vue from 'vue'
import App from './App'
import router from './router'
import store from './store'
import VueAxios from 'vue-axios'
import VueAuthenticate from 'vue-authenticate'
import { http } from '@/common'

import 'highlight.js/styles/github.css'
import 'toastr/build/toastr.min.css'
import 'toastr/build/toastr.min.js'
import 'github-markdown-css'

Vue.config.productionTip = false

Vue.use(VueAxios, http)
Vue.use(VueAuthenticate, {
  baseUrl: 'https://mdn0te.herokuapp.com',
  providers: {
    github: {
      clientId: '3ba8b2cde15d9f23ffe3',
      redirectUri: 'https://mdn0te.herokuapp.com/auth/callback/github',
      scope: ['read:user', 'user:email']
    }
  }
})

/* eslint-disable no-new */
new Vue({
  el: '#app',
  router,
  store,
  components: { App },
  template: '<App/>'
})
