<template>
  <div id="app">
    <header-nav />
    <router-view />
  </div>
</template>

<script>
import HeaderNav from '@/components/HeaderNav/HeaderNav'

export default {
  name: 'App',
  components: {
    HeaderNav
  },
  data () {
    return {
      navOpen: false
    }
  }
}
</script>

<style>
@import url('https://fonts.googleapis.com/css?family=Nanum+Gothic+Coding');

html, body {
  width: 100%;
  height: 100%;
  margin: 0;
  padding: 0;
  overflow-x: hidden;
}

* {
  margin: 0;
  padding: 0;
  border: 0;
  outline: none;
  box-sizing: border-box;
  font-family: 'Nanum Gothic Coding', monospace;
}

#app {
  width: 100%;
  height: 100%;
}
</style>
