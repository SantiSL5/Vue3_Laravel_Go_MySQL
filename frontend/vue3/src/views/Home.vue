<template>
    <Carousel :categorieslist="state.categorieslist"></Carousel>
    <InfiniteScroll></InfiniteScroll>
</template>
  
<script>
import Constant from "../Constant";
import Carousel from "../components/Carousel.vue";
import InfiniteScroll from "../components/InfinteScroll.vue";
import { reactive, computed } from "vue";
import { useStore } from "vuex";
// import { useRouter } from "vue-router";

export default {
    components: { Carousel, InfiniteScroll },
    setup() {
        const store = useStore();
        const state = reactive({
            categorieslist: computed(() => store.getters["categoryClient/getCategory"]),
        });
        store.dispatch("categoryClient/" + Constant.GET_ALL_CATEGORIES);
        return { state };
    },
};
</script>
  
<style>

</style>
