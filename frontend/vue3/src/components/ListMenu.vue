<template>
  <div class="container-fluid row">
    <div class="col-2">
      <div class="sticky-top">
        <div class="row justify-content-around mt-5">

          <div id="navbar-dishtype" class="list-group col-10">
            <a v-for="dishtype in disheslist[1]" :key="dishtype.id" class="list-group-item list-group-item-action" :href="'#' + dishtype.name">{{ dishtype.name }}</a>
          </div>

          
          
          
          
          
          
          <!-- <nav id="navbar-example3" class="navbar navbar-light bg-light flex-column align-items-stretch p-3">
            <nav class="nav nav-pills flex-column" v-for="dishtype in disheslist[1]" :key="dishtype.id">
              <a class="nav-link" :href="'#' + dishtype.name">{{ dishtype.name }}</a>
            </nav>
          </nav> -->



        </div>
      </div>
    </div>
    <div class="col-9 justify-content-center mt-3" data-bs-spy="scroll" data-bs-target="#navbar-dishtype" data-bs-offset="0"
      tabindex="0">
      <div v-for="dishtype in disheslist[1]" :key="dishtype.id" :id="dishtype.name">
        <h1  class="row">{{ dishtype.name }}</h1>
        <div class="row" v-for="dish in disheslist[0]" :key="dish.id">
          <div v-if="dish.typecontent.name == dishtype.name" class="row">

            <div class="col-10">
              <h4>{{ dish.dish }}</h4>
              <h6>{{ dish.desc }}</h6>
            </div>
            <div class="col-2 text-end">
              <h6>{{ dish.price }} €</h6>
            </div>
          </div>
          </div>
        </div>
      </div>

  </div>
</template>

<script>
export default {
  props: {
    disheslist: Array,
  },
};
// for (let index = 0; index < array.length; index++) {
//   tipos (composable)
//     h1
  
//   for (let index = 0; index < array.length; index++) {


//     platos
//   }

// }
</script>
                        
<style>

</style>