<template>
    <Carousel :categorieslist="state.categorieslist"></Carousel>
    <div class="text-center mb-3">
        <h1>Menu</h1>
        <InfiniteScroll :dishtypeslist="state.dishtypeslist"></InfiniteScroll>
    </div>
</template>
  
<script>
import Constant from "../Constant";
import Carousel from "../components/Carousel.vue";
import InfiniteScroll from "../components/InfinteScroll.vue";
import { useDishTypeInfinite } from "../composables/useDishTypes";
import { reactive, computed } from "vue";
import { useStore } from "vuex";

export default {
    components: { Carousel, InfiniteScroll },
    setup() {
        const store = useStore();
        const state = reactive({
            categorieslist: computed(() => store.getters["categoryClient/getCategory"]),
            dishtypeslist: useDishTypeInfinite(0, 4),
            scroll: {
                offset: 2,
                active: "no",
                continue: true
            },
        });
        store.dispatch("categoryClient/" + Constant.GET_ALL_CATEGORIES);

        const makeScroll = () => {
            let scrollTop = document.documentElement.scrollTop;
            let scrollHeight = document.documentElement.scrollHeight;
            let clientHeight = document.documentElement.clientHeight;
            if (state.scroll.continue) {
                if (scrollTop + clientHeight >= scrollHeight) {
                    state.scroll.offset += 2;

                    state.listTmp = useDishTypeInfinite(state.scroll.offset, 2);

                    setTimeout(() => {
                        state.listTmp.map((item) => {
                            state.dishtypeslist.push(item);
                        });
                        if (state.dishtypeslist.length < state.scroll.offset) {
                            state.scroll.continue = false;
                        }
                    }, 100);
                }
            }
        };

        return { state, makeScroll };
    },
    mounted() {
        window.addEventListener("scroll", this.makeScroll);
    },
};
</script>
  
<style>

</style>
