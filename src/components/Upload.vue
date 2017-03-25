<template>
  <div class="ui modal" id="upload">
    <i class="close icon"></i>
    <div class="header">
      Import the data sources
    </div>
    <div class="image content">
      <div class="ui medium image">
        <img src="../assets/excel.jpg">
      </div>
      <div class="description">
        <div class="ui header">You can upload your local file by here</div>
        <p>Please pick your file in the box</p>
        <div><input type="file" @change="selectFile" id="file" name="first-name" placeholder=""></div>
        <p v-show="fileInfo.type">
          Size: {{ fileInfo.size }} , Type: {{ fileInfo.type }}, <button class="ui primary button" v-on:click="uploadFile()">Start upload</button>
        </p>
        <p>Is it okay to use this photo?</p>
      </div>
    </div>
    <div class="actions">
      <div class="ui black deny button">
        Cancel
      </div>
      <div class="ui positive right labeled icon button" v-show="fileInfo.type">
        Yep, Go on
        <i class="checkmark icon"></i>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: 'Upload',
  data () {
    return {
      fileInfo: {},
      title: ''
    }
  },
  methods: {
    show (msg, title) {
      document.querySelector('#file').value = ''
      this.fileInfo.size = 0
      this.fileInfo.type = ''
      this.msg = msg
      this.title = title
      $('#upload').modal('show')
    },
    hide () {
      $('#upload').modal('hide')
    },
    selectFile (e) {
      let files = e.target.files || e.dataTransfer.files
      if (files.length) {
        this.preUploadFile(files[0])
      }
    },
    preUploadFile (file) {
      let reader = new window.FileReader()
      reader.readAsDataURL(file)
      reader.onload = (e) => {
        this.fileInfo.type = file.type
        if (file.size > 1024 * 1024) {
          this.fileInfo.size = (Math.round(file.size * 100 / (1024 * 1024)) / 100).toString() + 'MB'
        } else {
          this.fileInfo.size = (Math.round(file.size * 100 / 1024) / 100).toString() + 'KB'
        }
        this.preFile = file
        this.uploadFile()
      }
    },
    uploadFile () {
      let formData = new FormData()
      let myXhr

      formData.append('file', this.preFile)
      $.ajax({
        url: '/api/import',
        type: 'POST',
        xhr: () => {  // custom xhr
          myXhr = $.ajaxSettings.xhr()
          if (myXhr.upload) { // if upload property exists
//            myXhr.upload.addEventListener('progress', progressHandlingFunction, false)
          }
          return myXhr
        },
        // Ajax events
        success: (data) => {
          menuComponent.setList(data)
          chartComponent.clearList()
          menuComponent.setFileName(this.preFile.name)
          this.hide()
        },
        error: (e) => {
          messageComponents.show(e.responseText)
        },
        // Form data
        data: formData,
        // Options to tell jQuery not to process data or worry about the content-type
        cache: false,
        contentType: false,
        processData: false
      }, 'json')
    }
  },
  mounted () {
    window.uploadComponents = this
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>

</style>
