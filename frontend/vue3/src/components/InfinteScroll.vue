<template>
  <div class="container">
    <div v-for="dishtype in dishtypeslist" :key="dishtype.id">
      {{ dishtype }}

    </div>
  </div>
</template>
   
<script>



export default {
  props: {
    dishtypeslist: Array,
  },
  


}



</script>
    
<style>

</style>