<template>
  <div class="container-fluid">
    <div class="row justify-content-around" v-if="dishtypeslist.length > 0">
      <div class="col-6 justify-content-center" v-for="dishtype in dishtypeslist" :key="dishtype.id">
          <img :src="require(`@/assets/${dishtype.photo}`)" class="img-fluid col-6" />
          <h2>{{ dishtype.name }}</h2>
      </div>
    </div>
    <div v-else>
      <span>Something went wrong</span>
    </div>

  </div>

  <!-- <button @click="redirectReserve(dishtype.id)">Reserve table</button> -->


</template>

<script>
// import { useRouter } from "vue-router";
export default {
  props: {
    dishtypeslist: Array,
  },
  setup() {
    // const router = useRouter();
    // const redirectReserve = (id) => {
    //   let search = { id_dishtype: id };
    //   router.push({
    //     name: "client_table",
    //     params: { filter: btoa(JSON.stringify(search)) },
    //   });
    // };
    // return { redirectReserve };
  },
};
</script>
                        
<style>

</style>