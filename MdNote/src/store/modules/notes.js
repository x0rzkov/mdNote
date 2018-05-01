/* eslint-disable */
import * as types from '../mutation-types'
import {http, getCookie} from '../../common'
import toastr from 'toastr'

toastr.options.positionClass = 'toast-bottom-right'

const state = {
  notes: [],
  currentNote: {
    content: '',
    title: '',
    id: '',
    user_id: '',
    category: '',
    created_at: '',
    deleted_at: '',
    starred: false
  },
  categories: [],
  selectedCategory: ''
}

const getters = {
  notes: state => state.notes,
  currentNote: state => state.currentNote,
  categories: state => state.categories,
  selectedCategory: state => state.selectedCategory
}

const actions = {
  getNote ({commit, dispatch}, payload) {
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
      dispatch('setIsLogin', true)      
    }).catch(err => {
      if (err.response.status === 400) {
        toastr.error('Loading Failed')
      } else if (err.response.status === 401) {
        commit(types.SET_USER_NAME, '')
        toastr.error('Please Sign In')
        dispatch('setIsLogin', false)
      }
      console.log(err)
    })
  },
  getNoteList ({commit, dispatch, getters}) {
    http.get('/note/list', {
      headers: {
        'Authorization': 'JWT ' + getCookie('JWT')
      },
      params: {
        category: getters.selectedCategory
      }
    }).then(response => {
      commit(types.SET_NOTES, response.data.notes)
      commit(types.SET_CATEGORIES, response.data.categories)
      dispatch('setIsLogin', true)      
    }).catch(err => {
      if (err.response.status === 401) {
        commit(types.SET_USER_NAME, '')
        toastr.error('Please Sign In')
        dispatch('setIsLogin', false)      
      }
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
        toastr.success('Saved: ' + response.data.title)      
        dispatch('getNoteList')
        dispatch('setIsLogin', true)              
    }).catch(err => {
      if (err.response.status === 400) {
        toastr.error('Saving Failed')
        dispatch('setIsLogin', false)      
      } else if (err.response.status === 401) {
        commit(types.SET_USER_NAME, '')
        toastr.error('Please Sign In')
      }
    })
  },
  deletedNote({dispatch}, payload) {
    http.delete('/note', null, {
      headers: {
        'Authorization': 'JWT ' + getCookie('JWT')
      }
    }).then(response => {
      toastr.success('Deleting Complete')
      dispatch('newNote')
    }).catch(err => {
      if (err.response.status === 404) {
        toastr.error('Deleting Failed')
      } else if (err.response.status === 401) {
        commit(types.SET_USER_NAME, '')
        toastr.error('Please Sign In')
      }
    })
  },
  newNote({commit}) {
    commit(types.SET_CURRENT_NOTE, {
      content: '',
      title: '',
      id: '',
      user_id: '',
      category: '',
      created_at: '',
      deleted_at: '',
      starred: false
    })
  }
}

const mutations = {
  [types.SET_CURRENT_NOTE] (state, payload) {
    state.currentNote = payload
  },
  [types.SET_NOTES] (state, payload) {
    state.notes = payload
  },
  [types.SET_CATEGORIES] (state, payload) {
    state.categories = payload
  },
  [types.SET_CATEGORY] (state, payload) {
    state.category = payload
  }
}

export default {
  state,
  getters,
  actions,
  mutations
}
