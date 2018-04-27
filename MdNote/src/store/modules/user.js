/* eslint-disable */
import * as types from '../mutation-types'
import {http, getCookie, deleteCookie} from '../../common'
import toastr from 'toastr'

toastr.options.positionClass = 'toast-bottom-right'

const state = {
  userName: ''
}

const getters = {
  userName: state => state.userName,
  isLogin: state => state.userName !== ''
}

const actions = {
  setIsLogin ({commit}, payload) {
    if (payload) {
      let token = getCookie('JWT')
      var base64Url = token.split('.')[1];
      var base64 = base64Url.replace('-', '+').replace('_', '/');
      commit(types.SET_USER_NAME, JSON.parse(window.atob(base64)).name);
    } else {
      commit(types.SET_USER_NAME, '')
    }
  }
}

const mutations = {
  [types.SET_USER_NAME] (state, payload) {
    if (payload === '') {
      deleteCookie('JWT')
    }
    state.userName = payload
  }
}

export default {
  state,
  getters,
  actions,
  mutations
}
