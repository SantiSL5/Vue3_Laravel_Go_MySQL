<template>
  <div class="container">
    <div class="row mt-5 mb-5">
      <div class="col-sm-12">
        <div id="carouselExampleIndicators" class="carousel slide" data-bs-ride="carousel">
          <div class="carousel-indicators">
            <button v-for="(category,index) in categorieslist" :key="category.id" type="button" data-bs-target="#carouselExampleIndicators" :data-bs-slide-to="index" :class="{ 'active': index === 0 }" aria-current="true" aria-label="Slide"></button>

            <!-- <button type="button" data-bs-target="#carouselExampleIndicators" data-bs-slide-to="1"
              aria-label="Slide 2"></button>
            <button type="button" data-bs-target="#carouselExampleIndicators" data-bs-slide-to="2"
              aria-label="Slide 3"></button> -->
          </div>
          <div class="carousel-inner">
            <div v-for="(category,index) in categorieslist" :key="category.id" class="carousel-item" :class="{ 'active': index === 1 }">
              <img :src="require(`@/assets/${category.photo}`)" class="d-block w-100" alt="...">
            </div>
            <!-- <div class="carousel-item active">
              <img :src="require(`@/assets/logo.png`)" class="d-block w-100" alt="...">
            </div>
            <div class="carousel-item">
              <img :src="require(`@/assets/logo.png`)" class="d-block w-100" alt="...">
            </div>
            <div class="carousel-item">
              <img :src="require(`@/assets/logo.png`)" class="d-block w-100" alt="...">
            </div> -->
          </div>
          <button class="carousel-control-prev" type="button" data-bs-target="#carouselExampleIndicators"
            data-bs-slide="prev">
            <span class="carousel-control-prev-icon" aria-hidden="true"></span>
            <span class="visually-hidden">Previous</span>
          </button>
          <button class="carousel-control-next" type="button" data-bs-target="#carouselExampleIndicators"
            data-bs-slide="next">
            <span class="carousel-control-next-icon" aria-hidden="true"></span>
            <span class="visually-hidden">Next</span>
          </button>
        </div>
      </div>
    </div>
  </div>
</template>
   
<script>
//importing bootstrap 5
// import "bootstrap/dist/css/bootstrap.min.css";
// import "bootstrap/dist/js/bootstrap.min.js";
// import * as Bootstrap from 'bootstrap'
//import DemoComponent from "./components/DemoComponent.vue"
export default {

  props: {
    categorieslist: Array,
  }


}
</script>
    
<style>

</style>