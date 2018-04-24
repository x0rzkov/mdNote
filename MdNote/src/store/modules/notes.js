/* eslint-disable */
import * as types from '../mutation-types'
import {http, getCookie} from '../../common'

const state = {
  notes: [
    {
      id: 1,
      title: 'Hello world',
      date: '2017',
      category: 'Go'
    },
    {
      id: 2,
      title: 'H2ello dssddssddddworld',
      date: '2017',
      category: 'Go'
    },
    {
      id: 3,
      title: 'Heello world',
      date: '2017',
      category: 'Go'
    },
    {
      id: 4,
      title: 'Hedllo world',
      date: '2017',
      category: 'Go'
    },
    {
      id: 5,
      title: 'Hello worl123d',
      date: '2017',
      category: 'Go'
    },
    {
      id: 6,
      title: 'Hello world',
      date: '2017',
      category: 'Go'
    },
    {
      id: 7,
      title: 'Hello1 worl1d',
      date: '2017',
      category: 'Go'
    },
    {
      id: 8,
      title: 'Hello world6',
      date: '2017',
      category: 'Go'
    },
    {
      id: 9,
      title: 'Hello world32',
      date: '2017',
      category: 'Go'
    },
    {
      id: 10,
      title: 'Hello world',
      date: '2017',
      category: 'Go'
    }
  ],
  currentNote: {}
}

const getters = {
  notes: state => state.notes,
  currentNote: state => state.currentNote
}

const actions = {
  getNote ({commit, payload}) {
    http.get('/note', {
      headers: {
        'Authorization': 'JWT ' + getCookie('JWT')
      },
      params: {
        id: payload.id
      }
    }).then(response => {
      commit(types.SET_CURRENT_NOTE, response.data)
    }).catch(err => {
      console.log(err)
    })
  },
  getNoteList ({commit, payload}) {
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
  saveNote({commit, payload}) {
    http.put('/note', JSON.stringify(payload), {
      headers: {
        'Authorization': 'JWT ' + getCookie('JWT')
      }
    }).then(response => {
       commit(types.SET_CURRENT_NOTE, response.data)
    }).catch(err => {
      console.log(err)
    })
  }
}

const mutations = {
  [types.SET_CURRENT_NOTE] (state, payload) {
    state.headerMenuOpen = payload
  },
  [types.SET_NOTES] (state, payload) {
    state.notes = JSON.parse(payload)
  }
}

export default {
  state,
  getters,
  actions,
  mutations
}