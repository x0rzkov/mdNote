import Vue from 'vue'
import Vuex from 'vuex'

import HeaderNav from '@/store/modules/header-nav'

Vue.use(Vuex)

export default new Vuex.Store({
  modules: {
    HeaderNav
  }
})
