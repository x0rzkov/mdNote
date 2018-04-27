import Vue from 'vue'
import Vuex from 'vuex'

import HeaderNav from '@/store/modules/header-nav'
import Notes from '@/store/modules/notes'
import User from '@/store/modules/user'

Vue.use(Vuex)

export default new Vuex.Store({
  modules: {
    HeaderNav,
    Notes,
    User
  }
})
