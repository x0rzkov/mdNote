/* eslint-disable */
import * as types from '../mutation-types'
import {http, getCookie} from '../../common'
import toastr from 'toastr'

toastr.options.positionClass = 'toast-bottom-right'
toastr.options.preventDuplicates = true;

const state = {
  notes: [],
  currentNote: {
    content: '',
    title: '',
    category: ''
  },
  categories: [],
  selectedCategory: '',
  selectedDirectory: ''
}

const getters = {
  notes: state => state.notes,
  currentNote: state => state.currentNote,
  categories: state => state.categories,
  selectedCategory: state => state.selectedCategory,
  selectedDirectory: state => state.selectedDirectory
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
        toastr.error('Please Sign In')
        dispatch('setIsLogin', false)
      }
      console.log(err)
    })
  },
  getNoteList ({commit, dispatch, getters}) {
    if (getters.selectedDirectory == 'All Notes') {
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
          toastr.error('Please Sign In')
          dispatch('setIsLogin', false)  
        }
        console.log(err)
      })
    } else if (getters.selectedDirectory == 'Starred') {
      http.get('/note/list/starred', {
        headers: {
          'Authorization': 'JWT ' + getCookie('JWT')
        }
      }).then(response => {
        commit(types.SET_NOTES, response.data)
        dispatch('setIsLogin', true)      
      }).catch(err => {
        if (err.response.status === 401) {
          toastr.error('Please Sign In')
          dispatch('setIsLogin', false)  
        }
        console.log(err)
      })
    }
  },
  saveNote ({commit, dispatch}, payload) {
    http.put('/note', payload, {
      headers: {
        'Authorization': 'JWT ' + getCookie('JWT')
      }
    }).then(response => {
        commit(types.SET_CURRENT_NOTE, response.data)
        toastr.success('Saved: ' + response.data.title)      
        dispatch('setIsLogin', true)
        dispatch('getNoteList')
    }).catch(err => {
      if (err.response.status === 400) {
        toastr.error('Saving Failed')
      } else if (err.response.status === 401) {
        dispatch('setIsLogin', false)      
        toastr.error('Please Sign In')
      }
    })
  },
  deleteNote ({dispatch}, payload) {
    http.delete('/note', {
      headers: {
        'Authorization': 'JWT ' + getCookie('JWT')
      },
      params: {
        id: payload
      }
    }).then(response => {
      toastr.success('Deleting Complete')
      dispatch('setIsLogin', true)      
      dispatch('newNote')
      dispatch('getNoteList')
    }).catch(err => {
      if (err.response.status === 404) {
        toastr.error('Deleting Failed')
      } else if (err.response.status === 401) {
        dispatch('setIsLogin', false)      
        toastr.error('Please Sign In')
      }
    })
  },
  newNote ({commit}) {
    commit(types.SET_CURRENT_NOTE, {
      content: '',
      title: '',
      category: '',
    })
  },
  getDeletedNoteList ({commit, dispatch, getters}) {
    http.get('/note/list/deleted', {
      headers: {
        'Authorization': 'JWT ' + getCookie('JWT')
      }
    }).then(response => {
      commit(types.SET_NOTES, response.data)
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
  restoreNote ({dispatch, getters}, payload) {
    http.post('/note/restore', {id: payload}, {
      headers: {
        'Authorization': 'JWT ' + getCookie('JWT')
      }
    }).then(response => {
        toastr.success('Restoring Complete')      
        dispatch('getDeletedNoteList')
        dispatch('newNote')
        dispatch('setIsLogin', true)
    }).catch(err => {
      if (err.response.status === 400) {
        toastr.error('Restoring Failed')
      } else if (err.response.status === 401) {
        dispatch('setIsLogin', false)      
        toastr.error('Please Sign In')
      }
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
    state.selectedCategory = payload
  },
  [types.SET_DIRECTORY] (state, payload) {
    state.selectedDirectory = payload
  }
}

export default {
  state,
  getters,
  actions,
  mutations
}
