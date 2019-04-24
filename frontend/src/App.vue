<template>
  <div id="app">
    <div style='padding-bottom: 20px'>
      <div>Message is: {{message}}</div>
    </div>
    <button v-on:click="refreshMessage">Refresh Message</button>
    <!-- <HelloWorld msg="Welcome to Your Vue.js App"/> -->
    <MyComponent inputMessage="Welcome to Your Vue.js App"/>
    <button v-on:click="doaction">Do action</button>
  </div>
</template>

<script>
import HelloWorld from './components/HelloWorld.vue'
import MyComponent from './components/MyComponent.vue'


import Vue from 'vue';
import VueResource from 'vue-resource';
Vue.use(VueResource);

export default {
  name: 'app',
  components: {
    MyComponent
  },
  data() {
    return {
      message: 'hek hek'
    }
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
