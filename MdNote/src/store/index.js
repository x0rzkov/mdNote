import Vue from 'vue'
import Vuex from 'vuex'

import HeaderNav from '@/store/modules/header-nav'
import Notes from '@/store/modules/notes'

Vue.use(Vuex)

export default new Vuex.Store({
  modules: {
    HeaderNav,
    Notes
  }
})
