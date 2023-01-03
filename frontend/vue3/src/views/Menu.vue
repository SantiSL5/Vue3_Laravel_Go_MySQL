<template>
    <ListMenu :disheslist="[state.disheslist, state.dishtypeslist]"></ListMenu>
</template>
  
<script>
import Constant from "../Constant";
import ListMenu from "../components/ListMenu.vue";
import { reactive, computed } from "vue";
import { useDishTypeInfinite } from "../composables/useDishTypes";
import { useStore } from "vuex";

export default {
    components: { ListMenu },
    setup() {
        const store = useStore();
        const state = reactive({
            disheslist: computed(() => store.getters["dishClient/getDish"]),
            dishtypeslist: useDishTypeInfinite(),
        });
        store.dispatch("dishClient/" + Constant.GET_ALL_DISHES);
        return { state };
    },
};
</script>
  
<style>

</style>
