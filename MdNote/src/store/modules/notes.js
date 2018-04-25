/* eslint-disable */
import * as types from '../mutation-types'
import {http, getCookie} from '../../common'
import toastr from 'toastr'

const state = {
  notes: [
  ],
  currentNote: {
    content: '',
    title: '',
    id: '',
    user_id: '',
    category: ''
  }
}

const getters = {
  notes: state => state.notes,
  currentNote: state => state.currentNote
}

const actions = {
  getNote ({commit}, payload) {
    toastr.info('Loading')
    http.get('/note', {
      headers: {
        'Authorization': 'JWT ' + getCookie('JWT')
      },
      params: {
        id: payload
      }
    }).then(response => {
      toastr.success('Loading Complete')
      commit(types.SET_CURRENT_NOTE, response.data)
    }).catch(err => {
      toastr.error('Loading Failed')      
      console.log(err)
    })
  },
  getNoteList ({commit}, payload = {}) {
    http.get('/note/list', {
      headers: {
        'Authorization': 'JWT ' + getCookie('JWT')
      },
      params: {
        category: payload.category || ''
      }
    }).then(response => {
      commit(types.SET_NOTES, response.data)
    }).catch(err => {
      console.log(err)
    })
  },
  saveNote({commit, dispatch}, payload) {
    http.put('/note', payload, {
      headers: {
        'Authorization': 'JWT ' + getCookie('JWT')
      }
    }).then(response => {
       commit(types.SET_CURRENT_NOTE, response.data)
       dispatch('getNoteList')
    }).catch(err => {
      console.log(err)
    })
  },
  newNote({commit}) {
    commit(types.SET_CURRENT_NOTE, {
      content: '',
      title: '',
      id: '',
      user_id: '',
      category: ''
    })
  }
}

const mutations = {
  [types.SET_CURRENT_NOTE] (state, payload) {
    state.currentNote = payload
  },
  [types.SET_NOTES] (state, payload) {
    state.notes = payload
  }
}

export default {
  state,
  getters,
  actions,
  mutations
}
