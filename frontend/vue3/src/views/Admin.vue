<template>
    <div class="container">
        <nav class="navbar navbar-expand-lg bg-light">
            <div class="container-fluid">
                <a class="navbar-brand">Tables</a>
                <button class="navbar-toggler" type="button" data-bs-toggle="collapse"
                    data-bs-target="#navbarNavAltMarkup" aria-controls="navbarNavAltMarkup" aria-expanded="false"
                    aria-label="Toggle navigation">
                    <span class="navbar-toggler-icon"></span>
                </button>
                <div class="collapse navbar-collapse" id="navbarNavAltMarkup">
                    <div class="navbar-nav">
                        <a class="nav-link" @click="state.crud = 'categories'">Categories</a>
                        <a class="nav-link" @click="state.crud = 'tables'">Tables</a>
                        <a class="nav-link" @click="state.crud = 'dishtypes'">Dish Types</a>
                        <!-- <a class="nav-link">Features</a>
                        <a class="nav-link">Pricing</a> -->
                    </div>
                </div>
            </div>
        </nav>

        <Categories v-if="state.crud == 'categories'" :categories="state.categories"></Categories>
        <Tables v-if="state.crud == 'tables'" :tables="state.tables"></Tables>
        <DishTypes v-if="state.crud == 'dishtypes'" :dishtypes="state.dishtypes"></DishTypes>

    </div>
</template>
  
<script>
import Constant from "../Constant";
import { reactive, computed } from "vue";
import { useStore } from "vuex";
import Categories from "../components/admin/CRUDCategories.vue";
import Tables from "../components/admin/CRUDTables.vue";
import DishTypes from "../components/admin/CRUDDishTypes.vue";


export default {
    components: { Categories, Tables, DishTypes },
    setup() {
        const store = useStore();
        const state = reactive({
            categories: computed(() => store.getters["categoryAdmin/getCategory"]),
            tables: computed(() => store.getters["tableAdmin/getTable"]),
            dishtypes: computed(() => store.getters["dishTypeAdmin/getDishType"]),
            crud: "categories"
        });
        store.dispatch("categoryAdmin/" + Constant.GET_ALL_CATEGORIES);
        store.dispatch("tableAdmin/" + Constant.GET_ALL_TABLES);
        store.dispatch("dishTypeAdmin/" + Constant.GET_ALL_DISHTYPES);
        return { state };
    },
};

</script>
  
<style>

</style>
