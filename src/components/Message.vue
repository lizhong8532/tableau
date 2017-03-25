<template>
  <div class="ui basic modal" id="message">
    <i class="close icon"></i>
    <div class="header">{{ title }}</div>
    <div class="image content">
      <div class="image">
        <i class="alarm outline icon"></i>
      </div>
      <div class="description">
        <p>{{ msg }}</p>
      </div>
    </div>
    <div class="actions">
      <div class="two fluid ui inverted buttons">
        <div class="ui ok green basic inverted button">
          <i class="checkmark icon"></i>
          Yes
        </div>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: 'Message',
  data () {
    return {
      title: '',
      msg: ''
    }
  },
  methods: {
    show (msg, title) {
      this.msg = msg
      this.title = title
      $('#message').modal('show')
    }
  },
  mounted () {
    window.messageComponents = this
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>

</style>
