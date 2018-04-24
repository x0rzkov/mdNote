<template>
  <div id="header-menu-wrapper" :class="{'menu-open': open}">
    <div id="menu-buttons">
      <div class="button" v-for="button in menuButtons" :key="button.id">
        <div class="button-img">
          <img :src="button.img">
        </div>
        <div class="button-text">
          {{ button.text }}
        </div>
      </div>
    </div>
    <div id="category-buttons">
      <div id="category-open-button" @click="categoryButtons.open = !categoryButtons.open">
        <img id="category-open-arrow" :src="require('@/assets/HeaderNav/HeaderMenu/arrow.svg')" :class="{'category-open': categoryButtons.open}"> Category
      </div>
      <div id="category-button-wrapper" v-show="categoryButtons.open">
        <div class="category-button" v-for="button in categoryButtons.buttons" :key="button">
          <img class="category-img" :src="require('@/assets/HeaderNav/HeaderMenu/folder.svg')"> {{ button }}
        </div>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: 'HeaderMenu',
  props: {
    open: {type: Boolean}
  },
  data () {
    return {
      menuButtons: [
        {
          id: 1,
          text: 'New Note',
          img: require('@/assets/HeaderNav/HeaderMenu/add.svg')
        },
        {
          id: 2,
          text: 'All Notes',
          img: require('@/assets/HeaderNav/HeaderMenu/notebook.svg')
        },
        {
          id: 3,
          text: 'Starred',
          img: require('@/assets/HeaderNav/HeaderMenu/star.svg')
        },
        {
          id: 4,
          text: 'Trash',
          img: require('@/assets/HeaderNav/HeaderMenu/garbage.svg')
        }
      ],
      categoryButtons: {
        open: false,
        buttons: [
          'Go',
          'Lol',
          'Docs'
        ]
      }
    }
  }
}
</script>

<style scoped>
#header-menu-wrapper {
  width: 300px;
  height: calc(100vh - 65px);
  position: absolute;
  background-color: #FFECB3;
  top: 65px;
  left: -300px;
  transition: transform 0.3s;
  will-change: transform;
}

.menu-open {
  transform: translateX(300px);
}

#menu-buttons {
  width: 100%;
  display: flex;
  justify-content: center;
  align-items: center;
  flex-direction: column;
  margin-bottom: 30px;
}

.button {
  width: 100%;
  height: 55px;
  font-size: 18px;
  padding: 7px;
  display: flex;
  align-items: center;
  cursor: pointer;
}

.button:hover {
  background-color: rgba(0, 0, 0, 0.1);
}

.button > div {
  display: flex;
  align-items: center;
  justify-content: center;
}

.button-img {
  width: 80px;
  height: 40px;
}

.button-img > img {
  width: 20px;
  height: 20px;
}

.button-text {
  text-align: left;
}

#category-buttons {
  width: 100%;
  display: flex;
  justify-content: center;
  align-items: center;
  flex-direction: column;
}

#category-open-button {
  width: 100%;
  display: flex;
  justify-content: flex-start;
  align-items: center;
  font-size: 18px;
  height: 55px;
  padding: 10px;
  line-height: 30px;
  cursor: pointer;
}

#category-open-button:hover {
  background-color: rgba(0, 0, 0, 0.1);
}

#category-open-arrow {
  width: 15px;
  height: 15px;
  margin: 10px;
  transition: 0.1s;
  will-change: transform;
}

.category-open {
  transform: rotate(90deg);
}

#category-button-wrapper {
  width: 100%;
  display: flex;
  flex-direction: column;
  align-items: flex-end;
}

.category-button {
  width: 90%;
  font-size: 17px;
  height: 33px;
  display: flex;
  align-items: center;
  cursor: pointer;
}

.category-button:hover {
  background-color: rgba(0, 0, 0, 0.1);
}

.category-img {
  width: 17px;
  height: 17px;
  margin: 10px;
}
</style>
