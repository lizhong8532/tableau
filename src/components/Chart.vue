<template>
  <div style="padding: 10px;">
    <div class="ui grid">
      <div class="eight wide column">
        <h2 class="ui header">Chart</h2>
      </div>
      <div class="eight wide column i-toolbar">
        <i :class="getClassName(item.icon)" class="chart icon" v-on:click="changeChart(item)" v-for="item of charts"></i>
      </div>
    </div>

    <div class="i-chart" ondrop="drop(event)" ondragover="event.preventDefault()">
      <div class="i-list">
        <div class="ui list">
          <div class="item" v-for="(item, i) in list">
            <div class="right floated content">
              <i class="stop icon" :style="getColor(i)"></i>
              <i class="remove icon" v-on:click="removeList(i)" title="Delete the item"></i>
            </div>
            {{ item.sheet }} <span style="font-weight: bold;">{{ item.title }} </span>
          </div>
        </div>
        <div class="ui warning message" v-show="!list.length">
          <p>Please drag data to here from source</p>
        </div>
      </div>

      <div id="chart" class="i-chart-main"></div>
    </div>
  </div>

</template>

<script>
  import c3 from 'c3'
  export default {
    name: 'Chart',
    data () {
      return {
        msg: 'Welcome to Your Vue.js App',
        list: [],
        patterns: ['#1f77b4', '#aec7e8', '#ff7f0e', '#ffbb78', '#2ca02c', '#98df8a', '#d62728', '#ff9896', '#9467bd', '#c5b0d5', '#8c564b', '#c49c94', '#e377c2', '#f7b6d2', '#7f7f7f', '#c7c7c7', '#bcbd22', '#dbdb8d', '#17becf', '#9edae5'],
        currentChart: 'line',
        charts: [
          { icon: 'bar' },
          { icon: 'line' },
          { icon: 'pie' },
          { icon: 'area' }
        ]
      }
    },
    methods: {
      custom () {
        return []
      },
      clearList () {
        this.list.splice(0, 100000)
      },
      pushList (obj) {
        this.list.push(obj)
      },
      removeList (i) {
        let color = this.patterns.splice(i, 1)
        this.patterns.push(color)
        this.list.splice(i, 1)
        window.list.splice(i, 1)
      },
      getClassName (name) {
        if (!this.list.length) {
          name += ' disabled'
        }
        if (name === this.currentChart && this.list.length) {
          name += ' blue'
        }
        return name
      },
      changeChart (item) {
        if (this.currentChart !== item.icon) {
          this.currentChart = item.icon
        }
      },
      renderChart () {
        document.querySelector('#chart').innerHTML = ''
        if (this.list.length) {
          let columns = []
          let arr
          this.list.forEach((item) => {
            arr = JSON.parse(item.valuesJson)
            arr.unshift(item.title)
            columns.push(arr)
          })
          c3.generate({
            bindto: '#chart',
            data: {
              columns: columns,
              type: this.currentChart
            },
            color: {
              pattern: this.patterns
            },
            padding: {
              top: 100
            },
            axis: {
              x: {
                type: 'category',
                categories: JSON.parse(window.labels[this.list[0].sheetIndex])
              }
            }
          })
        }
      },
      getColor (i) {
        return 'color:' + this.patterns[i]
      }
    },
    watch: {
      list () {
        this.renderChart()
      },
      currentChart () {
        this.renderChart()
      }
    },
    mounted () {
      window.chartComponent = this
    }
  }
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style lang="scss" scoped>
  .i-chart-main {
    width: 100%;
    height: 100%;
  }
  .i-toolbar {
    text-align: right;
    i {
      margin-left: 1em;
    }
  }
  .i-chart {
    margin-top: 20px;
    height: 500px;
    border: 1px solid #CCC;
    position: relative;
  }
  .i-list {
    padding: 10px;
    position: absolute;
    top: 10px;
    right: 10px;
    border: 1px solid #DDD;
    width: 200px;
    z-index: 1;
  }

  i.icon:hover {
    cursor: pointer;
  }
</style>
