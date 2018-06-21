<template>
  <div id="viewer-wrapper">
    <div id="markdown-title">
      {{ title }}
      <div id="markdown-tag">
        {{ category }}
      </div>
    </div>
    <div id="markdown" v-html="htmlSource" class="markdown-body" />
  </div>
</template>

<script>
export default {
  name: 'Viewer',
  props: {
    htmlSource: {
      type: String,
      required: true
    },
    title: {
      type: String,
      required: true
    },
    category: {
      type: String,
      required: true
    }
  }
}
</script>

<style scoped>
#viewer-wrapper {
  width: 100%;
  height: 100%;
}

#markdown-title {
  width: 100%;
  height: 70px;
  word-break: break-all;
  font-size: 50px;
  font-weight: bold;
  line-height: 40px;
  padding: 15px;
  border-radius: 15px;
  text-align: center;
  color: #FFD54F;
  position: relative;
}

#markdown-tag {
  position: absolute;
  right: 0;
  bottom: 0;
  height: 20px;
  font-size: 15px;
  color: red;
  display: none;
  line-height: 20px;
}

#markdown {
  width: 100%;
  height: calc(100% - 70px);
  padding: 15px;
  font-size: 130%;
  word-break: break-all;
  overflow-y: auto;
}

@media print {
  #viewer-wrapper {
    display: block !important;
    height: auto;
    overflow: visible;
  }

  #markdown-title {
    text-align: center;
    text-overflow: ellipsis;
    color: #FFD54F;
  }

  #markdown-tag {
    display: block;
  }

  #markdown {
    overflow: visible;
  }

  #markdown * {
    page-break-inside: avoid;
  }
}
</style>
