<template>
  <div id="app">
    <header-nav />
    <router-view/>
  </div>
</template>

<script>
import HeaderNav from '@/components/HeaderNav/HeaderNav'

export default {
  name: 'App',
  components: {
    HeaderNav
  }
}
</script>

<style>
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
  font-family: 'Noto Sans KR', sans-serif;
}
</style>
