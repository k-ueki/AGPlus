<template>
  <div id="app">
    <Index></Index>
  </div>
</template>

<script>
import Index from './components/Index.vue'
// import ClassList from './components/ClassList.vue'

export default {
  name: 'app',
  components: {
    Index,
    // ClassList
  }
}
</script>

<style>
#app {
  font-family: 'Avenir', Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  margin-top: 60px;
}
</style>
