<template>
  <div id="index-page-wrapper" :class="{'slim': headerMenu}">
    <explorer />
    <div id="editor" :class="{'close': fullScreen}">
      <textarea name="" id="text" v-model="source" @blur="save"></textarea>
    </div>
    <div id="viewer" :class="{'full': fullScreen}">
      <div id="tools">
        <img :src="fullScreenImg" id="full-screen-img" @click="fullScreen = !fullScreen">
      </div>
      <div id="markdown" v-html="htmlSource">
      </div>
    </div>
  </div>
</template>

<script>
import Explorer from '@/components/Explorer/Explorer'
var hljs = require('highlight.js')

var md = require('markdown-it')({
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
    Explorer
  },
  computed: {
    headerMenu () {
      return this.$store.getters.headerMenuOpen
    },
    fullScreenImg () {
      return this.fullScreen ? require('@/assets/Index/full-screen-exit.svg') : require('@/assets/Index/full-screen.svg')
    },
    htmlSource () {
      return md.render(this.source)
    }
  },
  data () {
    return {
      source: '',
      fullScreen: false
    }
  },
  methods: {
    save () {
      console.log('saving')
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
  align-items: center;
  justify-content: center;
  overflow-x: hidden;
}

#text {
  width: 95%;
  height: 95%;
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
  align-items: center;
  justify-content: center;
  overflow-x: hidden;
  overflow-x: hidden;
}

#tools {
  float: right;
  width: 30px;
  height: 30px;
  position: absolute;
  top: 15px;
  right: 15px;
  background-color: rgb(240, 240, 240);
  display: flex;
  align-items: center;
  justify-content: center;
}

#full-screen-img {
  width: 15px;
  height: 15px;
  cursor: pointer;
}

#markdown {
  width: 95%;
  height: 95%;
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
