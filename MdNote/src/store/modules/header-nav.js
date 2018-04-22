/* eslint-disable */
import * as types from '../mutation-types'

const state = {
  headerMenuOpen: false
}

const getters = {
  headerMenuOpen: state => state.headerMenuOpen
}

const actions = {
  toggleHeaderMenu ({commit, getters}) {
    commit(types.SET_HEADER_MENU_OPEN, !getters.headerMenuOpen)
  }
}

const mutations = {
  [types.SET_HEADER_MENU_OPEN] (state, payload) {
    state.headerMenuOpen = payload
  }
}

export default {
  state,
  getters,
  actions,
  mutations
}