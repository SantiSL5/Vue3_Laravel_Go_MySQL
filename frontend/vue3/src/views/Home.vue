<template>
    <Carousel></Carousel>
</template>
  
<script>
    import Constant from "../Constant";
    import Carousel from "../components/Carousel.vue";
    import { reactive, computed } from "vue";
    import { useStore } from "vuex";
    // import { useRouter } from "vue-router";

    export default {
        components: { Carousel },
        setup() {
            const store = useStore();
            const state = reactive({
                categorieslist: computed(() => store.getters["categoryClient/getCategory"]),
            });
            store.dispatch("categoryClient/" + Constant.GET_ALL_CATEGORIES);
            console.log(state);
            return { state };
        },
    };
</script>
  
<style>

</style>
  