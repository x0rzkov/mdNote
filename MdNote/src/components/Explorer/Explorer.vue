<template>
  <div id="explorer-wrapper">
    <div id="search-box-wrapper">
      <img :src="require('@/assets/Explorer/search.svg')" id="search-img">
      <input type="text" id="search-box" v-model="searchBox" placeholder="Search">
    </div>
    <div class="bar"></div>
    <div id="notes-wrapper">
      <div v-for="note in notes" :key="note.id" class="note-wrapper" @click="$store.dispatch('getNote', note.id)">
        <div class="note">
          <div class="note-title">
            <img :src="require('@/assets/HeaderNav/HeaderMenu/notebook.svg')" >
            <div class="note-title-text">
              {{ note.title }}
            </div>
          </div>
          <div class="note-date">
            {{ note.created_at }}
          </div>
          <div class="note-category">
            {{ note.category }}
          </div>
        </div>
        <div class="bar"></div>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: 'Explorer',
  data () {
    return {
      searchBox: ''
    }
  },
  computed: {
    notes () {
      if (this.searchBox === '') {
        return this.$store.getters.notes
      } else {
        return this.$store.getters.notes.filter(note => {
          return note.title.includes(this.searchBox)
        })
      }
    }
  },
  beforeMount () {
    this.$store.dispatch('getNoteList')
  }
}
</script>

<style scoped>
#explorer-wrapper {
  position: absolute;
  top: 65px;
  width: 300px;
  height: calc(100% - 65px);
  background-color: #FFF8E1;
  display: flex;
  flex-direction: column;
  align-items: center;
}

#search-box-wrapper {
  width: 80%;
  display: flex;
  align-items: center;
}

#search-box-wrapper:hover > #search-box::placeholder {
  color:rgba(0, 0, 0, 0.5);
}

#search-img {
  width: 20px;
  height: 20px;
  margin-left: 15px;
  margin-right: 15px;
}

#search-box {
  width: calc(100% - 20px);
  padding: 7px;
  height: 60px;
  font-size: 18px;
  background-color: rgba(0, 0, 0, 0);
}

#search-box::placeholder {
  font-size: 18px;
  color: rgba(0, 0, 0, 0.2);
  transition: 0.1s;
}

#search-box:focus::placeholder {
  color: rgba(0, 0, 0, 0.5)
}

#notes-wrapper {
  width: 100%;
  height: calc(100% - 60px);
  overflow-y: auto;
}

.note-wrapper {
  width: 100%;
  height: 100px;
  display: flex;
  flex-direction: column;
  align-items: center;
}

.bar {
  width: 80%;
  border-top: 1px solid rgba(0, 0, 0, 0.1);
}

.note {
  width: 80%;
  height: 130px;
  color: rgba(0, 0, 0, 0.3);
  display: flex;
  flex-direction: column;
  align-items: center;
  padding: 7px;
  transition: background-color 0.1s;
  cursor: pointer;
}

.note:hover {
  background-color: rgba(0, 0, 0, 0.1);
}

.note:hover > .note-title {
  color: rgba(0, 0, 0, 0.6);
}

.note-title {
  width: 100%;
  display: flex;
  align-items: center;
  font-size: 18px;
  font-weight: bold;
  height: 30px;
}

.note-title-text {
  text-overflow: ellipsis;
  white-space: nowrap;
  overflow: hidden;
}

.note-title > img {
  width: 16px;
  height: 16px;
  margin-left: 15px;
  margin-right: 15px;
}

.note-date {
  font-size: 13px;
  width: 90%;
  margin-top: 5px;
  margin-bottom: 5px;
}

.note-category {
  font-size: 15px;
  width: 90%;
  text-align: right;
  font-weight: bold;
}
</style>
