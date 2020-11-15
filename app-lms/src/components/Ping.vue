<template>
  <div>
    <button type="button" class="btn btn-primary">{{ msg }}</button>
  </div>
</template>
<script>
import axios from 'axios';

export default {
  name: 'Ping',
  data: function() {
    return {
      msg: "ping"
    }
  },   
  methods: {
    getMessage: function() {
      const path = 'http://localhost:8080/v1/ping';
      axios.get(path)
        .then((res) => {
          this.msg = res.data;
        })
        .catch((error) => {
          // eslint-disable-next-line
          console.error(error);
        });
    },
  },
  created: function() {
    this.getMessage();
  }, 
}
</script>