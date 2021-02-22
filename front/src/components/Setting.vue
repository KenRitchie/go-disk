<template>
  <el-menu default-active="1" class="el-menu-vertical-demo">
    <el-submenu index="1">
      <template slot="title">
        <img src="">
        <span>1</span>
      </template>
      <el-menu-item-group>
        <template slot="title"></template>
        <el-menu-item index="1-1">1</el-menu-item>
        <el-menu-item index="1-2">1</el-menu-item>
        <el-menu-item index="1-3">1</el-menu-item>
      </el-menu-item-group>
    </el-submenu>
  </el-menu>
</template>

<script>
export default {
  name: 'ElementUi',
  data () {
    return {
      msg: 'Welcome to Your Vue.js App'
    }
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
h1, h2 {
  font-weight: normal;
}
ul {
  list-style-type: none;
  padding: 0;
}
li {
  display: inline-block;
  margin: 0 10px;
}
a {
  color: #42b983;
}
</style>
