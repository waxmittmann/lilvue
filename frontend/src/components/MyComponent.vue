<template>
    <div>
        <div>Message: {{message}}</div>
        <div>Input Message: {{inputMessage}}</div>
        <button @click="refreshMessage">Refresh</button>
    </div>
</template>

<script>

export default {
  name: 'MyComponent',
  props: {
    inputMessage: String
  },
  data() {
      return {
          message: 'standard message'
      }
  },
  methods: {
    refreshMessage: function() { 
      this.$http.get('/articles/cheese/').then(function(response) {      
        console.log("Response: " + response.data);
        this.message = response.data;
        console.log("Got message " + response.data.message);
      });
    }
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
h3 {
  margin: 40px 0 0;
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
