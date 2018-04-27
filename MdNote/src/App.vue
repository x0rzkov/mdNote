<template>
  <div id="app">
    <router-view />
  </div>
</template>

<script>
export default {
  name: 'App'
}
</script>

<style>
@import url('https://fonts.googleapis.com/css?family=Nanum+Gothic');

html, body {
  width: 100%;
  height: 100%;
  margin: 0;
  padding: 0;
}

* {
  margin: 0;
  padding: 0;
  border: 0;
  outline: none;
  box-sizing: border-box;
  font-family: 'Nanum Gothic', sans-serif;
}

#app {
  width: 100%;
  height: 100%;
}

*::-webkit-scrollbar {
  width: 10px;
}

*::-webkit-scrollbar-track {
  background: #f8f8f8;
  border-radius: 5px;
}

*::-webkit-scrollbar-thumb {
  background: #888;
  border-radius: 5px;
}

*::-webkit-scrollbar-thumb:hover {
  background: #777;
}
</style>
