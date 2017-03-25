<template>
  <div id="app">
    <upload-components></upload-components>
    <message-components></message-components>
    <div class="ui" style="padding: 10px;">
      <div class="i-header-right">
        <button class="ui primary button" onclick="uploadComponents.show()">
          <i class="upload icon"></i>
          Import data source</button>
      </div>
      <a class="ui yellow label">
        <h1 class="ui header">i-Chart</h1>
      </a>
    </div>
    <div class="ui divider"></div>

    <div class="ui grid">
      <div class="four wide column">
        <menu-components></menu-components>
      </div>
      <div class="twelve wide column">
        <chart-components></chart-components>
      </div>
    </div>
  </div>
</template>

<script>

  import MenuComponents from './components/Menu'
  import ChartComponents from './components/Chart'
  import MessageComponents from './components/Message'
  import UploadComponents from './components/Upload'
  export default {
    name: 'app',

    components: {
      MenuComponents,
      ChartComponents,
      MessageComponents,
      UploadComponents
    }
  }
</script>

<style lang="scss">
.i-header-right {
  position: fixed;
  right: 10px;
  top: 10px;
}
</style>
