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
        <div><input style="outline: none" type="file" @change="selectFile" id="file" name="first-name" placeholder=""></div>
        <div style="margin-top: 10px;" id="i_file_info">
          <p>Size: <span id="i_file_size"></span></p>
          <p>Type: <span id="i_file_type"></span></p>
        </div>

        <div id="i_progress_box" style="margin-top: 20px">
          <div class="ui progress" data-value="0" data-total="100" id="i_progress">
            <div class="bar"></div>
            <div class="label">Uploading Files <span id="i_progress_label"></span>%</div>
          </div>
        </div>

        <div id="i_message" style="display: none; margin-top: 40px;">
          <div class="ui negative message">
            <p id="i_message_p"></p>
          </div>
        </div>

        <div id="i_file_upload_btn" style="margin-top:40px;">
          <button class="ui primary button" v-on:click="uploadFile()">Start upload</button>
        </div>
      </div>
    </div>
    <div class="actions">
      <div class="ui black deny button">
        Cancel
      </div>
      <div class="ui positive right labeled icon button" style="display: none">
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
      title: '',
      timer: null,
      isProcess: false
    }
  },
  methods: {
    show (msg, title) {
      $('#i_file_upload_btn').find('button').removeClass('disabled')
      $('#i_progress').progress({
        percent: 0
      })
      document.querySelector('#file').value = ''
      document.querySelector('#i_file_info').style.display = 'none'
      document.querySelector('#i_progress_box').style.display = 'none'
      document.querySelector('#i_message').style.display = 'none'
      document.querySelector('#i_file_upload_btn').style.display = 'none'
      document.querySelector('#i_file_size').innerHTML = ''
      document.querySelector('#i_file_type').innerHTML = ''
      document.querySelector('#i_progress_label').innerHTML = '0'
      this.isProcess = false
      this.fileInfo.size = 0
      this.fileInfo.type = ''
      this.msg = msg
      this.title = title
      $('#upload').modal('show')
    },
    hide () {
      $('#upload').modal('hide')
      this.isProcess = false
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

        document.querySelector('#i_file_info').style.display = 'block'
        document.querySelector('#i_file_upload_btn').style.display = 'block'
        document.querySelector('#i_file_size').innerHTML = this.fileInfo.size
        document.querySelector('#i_file_type').innerHTML = this.fileInfo.type
//        this.uploadFile()
      }
    },
    message (message) {
      clearTimeout(this.timer)
      this.timer = null
      document.querySelector('#i_message_p').innerHTML = message
      document.querySelector('#i_message').style.display = 'block'
      this.timer = setTimeout(() => {
        document.querySelector('#i_message').style.display = 'none'
      }, 5000)
    },
    uploadFile () {
      if (!this.isProcess) {
        this.isProcess = true
        document.querySelector('#i_progress_box').style.display = 'block'
        let btn = document.querySelector('#i_file_upload_btn')
        $(btn).find('button').addClass('disabled')
        window.setTimeout(() => {
          let formData = new FormData()
          formData.append('file', this.preFile)
          let myXhr
          $.ajax({
            url: '/api/import',
            type: 'POST',
            xhr: () => {  // custom xhr
              myXhr = $.ajaxSettings.xhr()
              if (myXhr.upload) { // if upload property exists
                myXhr.upload.addEventListener('progress', (e) => {
                  let value = Number((100 * (e.loaded / e.total)).toFixed())
                  document.querySelector('#i_progress_label').innerHTML = value
                  $('#i_progress').progress({
                    percent: value
                  })
                }, false)
              }
              return myXhr
            },
            // Ajax events
            success: (data) => {
              if (Array.isArray(data)) {
                menuComponent.setList(data)
                chartComponent.clearList()
                menuComponent.setFileName(this.preFile.name)
                this.hide()
              } else {
                this.message('No result, Make check your file has data and format correct!')
                this.isProcess = false
                document.querySelector('#i_progress_label').innerHTML = '0'
                $(btn).find('button').removeClass('disabled')
                $('#i_progress').progress({
                  percent: 0
                })
              }
            },
            error: (e) => {
              $(btn).removeClass('disabled')
              this.message(e.responseText)
              this.isProcess = false
            },
            // Form data
            data: formData,
            // Options to tell jQuery not to process data or worry about the content-type
            cache: false,
            contentType: false,
            processData: false
          }, 'json')
        }, 1000)
      }
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
