<template>
  <div id="index-page-wrapper" :class="{'slim': headerMenu}" @click="categoryTyping = false">
    <header-nav />
    <explorer />
    <div id="editor" :class="{'close': fullScreen}">
      <div id="text-info">
        <div id="text-category">
          <input type="text" v-model="tempNote.category" @focus="categoryTyping = true" @click.stop="categoryTyping = true" @blur="save; categoryTyping = false">
          <div id="category-list" v-show="categoryTyping">
            <div class="category-option" v-for="category in categories" :key="category" @click="tempNote.category = category; categoryTyping = false">{{ category }} </div>
          </div>
        </div>
        <input type="text" id="text-title" v-model="tempNote.title" @blur="save">
      </div>
      <textarea name="" id="text" v-model="tempNote.content" @blur="save"></textarea>
    </div>
    <div id="viewer" :class="{'full': fullScreen}">
      <div id="tools">
        <img :src="fullScreenImg" @click="fullScreen = !fullScreen">
        <img :src="garbage" @click="tempNote.deleted_at ? $store.dispatch('restoreNote', tempNote.id) : $store.dispatch('deleteNote', tempNote.id)">
        <img :src="starred">
      </div>
      <div id="markdown-title">{{ tempNote.title }}</div>
      <div id="markdown" v-html="htmlSource" class="markdown-body">
      </div>
    </div>
  </div>
</template>

<script>
import Explorer from '@/components/Explorer/Explorer'
import HeaderNav from '@/components/HeaderNav/HeaderNav'
import toastr from 'toastr'
import hljs from 'highlight.js'
import MarkdownIt from 'markdown-it'

toastr.options.closeButton = true

var md = MarkdownIt({
  highlight: function (str, lang) {
    if (lang && hljs.getLanguage(lang)) {
      try {
        return hljs.highlight(lang, str).value
      } catch (__) {}
    }
    return ''
  },
  breaks: true
})

export default {
  name: 'Index',
  components: {
    Explorer,
    HeaderNav
  },
  computed: {
    headerMenu () {
      return this.$store.getters.headerMenuOpen
    },
    fullScreenImg () {
      return this.fullScreen ? require('@/assets/Index/full-screen-exit.svg') : require('@/assets/Index/full-screen.svg')
    },
    garbage () {
      return this.tempNote.deleted_at ? require('@/assets/Index/restore.svg') : require('@/assets/Index/garbage.svg')
    },
    starred () {
      return this.tempNote.starred ? require('@/assets/Index/filled-star.svg') : require('@/assets/Index/star.svg')
    },
    htmlSource () {
      return md.render(this.tempNote.content)
    },
    categories () {
      return this.$store.getters.categories
    }
  },
  data () {
    return {
      tempNote: {
        title: this.$store.getters.currentNote.title,
        content: this.$store.getters.currentNote.content,
        category: this.$store.getters.currentNote.category,
        created_at: this.$store.getters.currentNote.created_at,
        id: this.$store.getters.currentNote.id,
        user_id: this.$store.getters.currentNote.user_id,
        deleted_at: this.$store.getters.currentNote.deleted_at,
        starred: this.$store.getters.currentNote.starred
      },
      fullScreen: false,
      categoryTyping: false
    }
  },
  methods: {
    save () {
      if (this.tempNote.title === '') {
        toastr.warning('Please write title')
      } else if (this.tempNote.content === '') {
        toastr.warning('Please write content')
      } else {
        if (this.isChange(this.tempNote, this.$store.getters.currentNote)) {
          this.$store.dispatch('saveNote', this.tempNote)
        }
      }
    },
    isChange (tempNote, currentNote) {
      for (let prop in tempNote) {
        if (tempNote[prop] !== currentNote[prop]) {
          return true
        }
      }
      return false
    }
  },
  watch: {
    '$store.getters.currentNote' (val) {
      this.tempNote = {
        title: val.title,
        content: val.content,
        category: val.category,
        created_at: val.created_at,
        deleted_at: val.deleted_at,
        starred: val.starred,
        id: val.id,
        user_id: val.user_id
      }
    }
  }
}
</script>

<style scoped>
#index-page-wrapper {
  width: 100%;
  height: 100%;
  float: right;
  transition: 0.3s;
  will-change: width;
  position: relative;
}

.slim {
  width: calc(100% - 300px) !important;
}

#editor {
  position: absolute;
  top: 65px;
  left: 300px;
  width: calc((100% - 300px) / 2);
  height: calc(100% - 65px);
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  overflow-x: hidden;
}

#text-info {
  display: flex;
  align-items: center;
  justify-content: center;
  width: 100%;
}

#text-category {
  width: 20%;
  position: relative;
  display: flex;
  justify-content: center;
}

#text-category > input {
  width: 100%;
  background-color: rgb(240, 240, 240);
  font-size: 30px;
  font-weight: bold;
  line-height: 60px;
  height: 60px;
  border-radius: 15px;
  text-indent: 15px;
}

#category-list {
  width: 80%;
  background-color: rgba(250, 250, 250);
  border: 1px solid rgb(240, 240, 240);
  position: absolute;
  top: 60px;
  max-height: 160px;
  overflow: auto;
}

.category-option {
  width: 100%;
  height: 40px;
  font-weight: bold;
  line-height: 40px;
  text-indent: 10px;
  font-size: 15px;
  overflow: hidden;
  text-overflow: ellipsis;
  cursor: pointer;
}

#text-title {
  width: calc(95% - 20% - 5px);
  background-color: rgb(240, 240, 240);
  font-size: 30px;
  font-weight: bold;
  line-height: 60px;
  height: 60px;
  margin: 5px;
  border-radius: 15px;
  text-indent: 15px;
}

#text {
  width: 95%;
  height: calc(95% - 70px);
  padding: 15px;
  background-color: rgb(240, 240, 240);
  border-radius: 15px;
  resize: none;
  font-size: 18px;
}

#viewer {
  position: absolute;
  top: 65px;
  width: calc((100% - 300px) / 2);
  height: calc(100% - 65px);
  right: 0;
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  overflow-x: hidden;
  overflow-x: hidden;
}

#tools {
  float: right;
  height: 30px;
  width: 100px;
  position: absolute;
  top: 15px;
  right: 15px;
  background-color: rgb(240, 240, 240);
  display: flex;
  align-items: center;
  justify-content: space-around;
}

#tools > img {
  width: 20px;
  height: 20px;
  cursor: pointer;
}

#markdown-title {
  width: 95%;
  font-size: 50px;
  font-weight: bold;
  line-height: 80px;
  height: 80px;
  margin: 5px;
  border-radius: 15px;
  text-indent: 15px;
}

#markdown {
  width: 95%;
  height: calc(95% - 90px);
  padding: 15px;
  font-size: 130%;
  word-break: break-all;
  overflow-y: auto;
}

.full {
  width: calc(100% - 300px) !important;
}

.full > #markdown {
  font-size: 160%;
}

.close {
  width: 0 !important;
}
</style>
