<template>
  <div style="padding: 10px;">
    <h2 class="ui header">Source</h2>
    <div class="ui warning message">
      <p v-show="list.length">Current data as from local files: <span class="blue" style="color: green;font-weight: bold">"{{ fileName }}"</span></p>
      <p v-show="!list.length">First import data source, please!</p>
    </div>

    <div class="ui vertical accordion menu" v-for="(item, sIndex) in list">

      <div class="item">
        <a class="active title blue" v-on:click="toggle(item)"><i class="caret icon" :class="getClass(item)"></i> {{ item.sheet }} </a>
        <div class="active content" v-show="!item.isHide">
          <div class="ui form">
            <div class="grouped fields">
              <div class="ui list">
                <div class="i-item item" v-for="(row, rowIndex) in item.columns" draggable="true" :data="stringify(sIndex, rowIndex, item.sheet, row.title, row, item.labels)" ondragstart="drag(event)">{{ row.title }}({{ sum(row.values) }})</div>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>

</template>

<script>
  export default {
    name: 'Menu',
    data () {
      return {
        list: [],
        fileName: ''
      }
    },

    methods: {
      sum (arr) {
        let count = 0
        arr.forEach((item) => {
          count += item
        })
        return count
      },

      stringify (str1, str2, str3, str4, somethings, labels) {
        return [[str1, str2, str3, str4].join('-'), JSON.stringify(somethings)].join('###')
      },

      check (json) {
        let className = ''
        if (window.list.indexOf(json) >= 0) {
          className = 'disabled'
        }
        return className
      },
      getClass (item) {
        return item.isHide ? 'right' : 'down'
      },
      toggle (item) {
        item.isHide = !item.isHide
      },
      setList (arr) {
        this.list.splice(0, 1000)
        arr.forEach((item) => {
          window.labels.push(JSON.stringify(item.labels))
          item.isHide = false
          this.list.push(item)
        })
      },
      setFileName (name) {
        this.fileName = name
      }
    },

    components: {
    },
    mounted () {
      window.menuComponent = this
    }
  }
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
.i-item {
  cursor: pointer;
}
</style>
