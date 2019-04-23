import Vue from 'vue'
import App from './App.vue'

Vue.config.productionTip = false

console.log("THIS IS CALLED");

new Vue({
 render: h => h(App),
  data: {
    message: 'No cheese yet'
  },
  methods: {
    refreshMessage: function() { 
      console.log("Called refresh message");
    // refreshMessage(resource) {
      this.$http.get('/categories/cheese').then(function(response) {
        console.log("Refreshing");
        this.message = response.data.message;
      });
    },
    doaction: function() { 
      console.log("Hi!")
    }
  }
}).$mount('#app')
