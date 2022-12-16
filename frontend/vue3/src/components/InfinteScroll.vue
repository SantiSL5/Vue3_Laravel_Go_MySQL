<template>
  <div class="container">
    <div class="row mt-5 mb-5">
      <div class="col-sm-12">
        <div id="carouselExampleIndicators" class="carousel slide" data-bs-ride="carousel" data-bs-interval="3000">
          <div class="carousel-indicators">
            <button v-for="(category,index) in categorieslist" :key="category.id" type="button" data-bs-target="#carouselExampleIndicators" :data-bs-slide-to="index" :class="{ 'active': index === 0 }" aria-current="true" aria-label="Slide"></button>
          </div>
          <div class="carousel-inner">
            <div v-for="(category,index) in categorieslist" :key="category.id" class="carousel-item" :class="{ 'active': index === 0 }">
              <img :src="require(`@/assets/${category.photo}`)" class="img-fluid d-block w-100 photo-carousel" alt="...">
            </div>
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

export default {

  props: {
    categorieslist: Array,
  }


}
</script>
    
<style>
  .photo-carousel {
    height: 600px;
  }
  .carousel-control-next, .carousel-control-prev{
    background-color: rgba(117, 116, 116, 0.445);
  }
</style>