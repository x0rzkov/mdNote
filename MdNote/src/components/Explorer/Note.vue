<template>
  <div class="note" @click="$store.dispatch('getNote', noteId)">
    <div class="note-title">
      <img :src="require('@/assets/HeaderNav/HeaderMenu/notebook.svg')" >
      <div class="note-title-text">
      {{ title }}
      </div>
    </div>
    <div class="note-date">
      {{ createdAt }}
    </div>
    <div class="note-category">
      {{ category }}
    </div>
  </div>
</template>

<script>
export default {
  name: 'Note',
  props: {
    noteId: {type: String},
    title: {type: String},
    createdAt: {type: String},
    category: {type: String},
    isDeleted: {type: Boolean},
    isStarred: {type: Boolean}
  },
  data () {
    return {
      menuOpen: false
    }
  }
}
</script>

<style scoped>
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
  position: relative;
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
