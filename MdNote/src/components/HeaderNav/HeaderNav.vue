<template>
  <div id="header-nav-wrapper">
    <div id="menu-button" @click="headerMenu = !headerMenu" :class="{'open': headerMenu}">
      <img :src="require('@/assets/HeaderNav/menu.svg')">
    </div>
    <div id="logo-wrapper">
      <img :src="require('@/assets/HeaderNav/Logo.png')">
    </div>
    <div id="sign-in" @click="authenticate('github')">
      Sign In
    </div>
    <header-menu :open="headerMenu" />
  </div>
</template>

<script>
import HeaderMenu from '@/components/HeaderNav/HeaderMenu'

export default {
  name: 'HeaderNav',
  components: {
    HeaderMenu
  },
  computed: {
    headerMenu: {
      get () {
        return this.$store.getters.headerMenuOpen
      },
      set (val) {
        this.$store.dispatch('toggleHeaderMenu')
      }
    }
  },
  methods: {
    authenticate (provider) {
      if (provider === 'github') {
        window.location.href = `https://github.com/login/oauth/authorize?cliend_id=3ba8b2cde15d9f23ffe3&redirect_uri=https://mdn0te.herokuapp.com/auth/callback/github`
      }
    }
  }
}
</script>

<style scoped>
#header-nav-wrapper {
  position: fixed;
  z-index: 100;
  top: 0;
  width: 100vw;
  background-color: #FFD54F;
  height: 65px;
  display: flex;
  align-items: center;
  padding: 5px 20px 5px 20px;
}

#menu-button {
  width: 50px;
  height: 50px;
  display: flex;
  align-items: center;
  justify-content: center;
  border-radius: 50%;
  cursor: pointer;
}

#menu-button:hover {
  background-color: rgba(0, 0, 0, 0.1);
}

.open {
  background-color: rgba(0, 0, 0, 0.07);
}

#menu-button img {
  width: 20px;
  height: 20px;
  filter: opacity(0.5);
}

#logo-wrapper {
  height: 40px;
  margin-left: 20px;
  margin-right: 20px;
}

#logo-wrapper img {
  height: 40px;
}

#sign-in {
  position: absolute;
  right: 15px;
  cursor: pointer;
}
</style>
