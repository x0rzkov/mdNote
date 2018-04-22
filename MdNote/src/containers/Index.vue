<template>
  <div id="index-page-wrapper" :class="{'slim': headerMenu}">
    <explorer />
  </div>
</template>

<script>
import Explorer from '@/components/Explorer/Explorer'

export default {
  name: 'Index',
  components: {
    Explorer
  },
  computed: {
    headerMenu () {
      return this.$store.getters.headerMenuOpen
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
}

.slim {
  width: calc(100% - 300px) !important;
}
</style>
