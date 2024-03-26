


<template>
    <form @click.prevent="Submit">
      <label>Amount:</label>
      <input type="text" v-model="amount">
      <button type="submit">Pay</button>
    </form>
  </template>
  
  <script>
  define(function(){
    return {
      props: ['flareView'],
      template: template,
      data: function() {
        return {amount: 0};
      },
      methods: {
        submit: function() {
          var self = this;
          var formData = {amount: self.amount};
  
          $flare.http.post('<% .Helpers.UrlForRoute "payment.received" %>', formData)
            .then(function(response){
              console.log(response);
            })
            .catch(function(error){
              console.log(error);
            });
        }
      }
    };
  });
  </script>