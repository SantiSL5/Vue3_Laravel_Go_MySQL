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
                        <a class="nav-link pointer" @click="state.crud = 'categories'">Categories</a>
                        <a class="nav-link pointer" @click="state.crud = 'tables'">Tables</a>
                        <a class="nav-link pointer" @click="state.crud = 'dishtypes'">Dish Types</a>
                        <a class="nav-link pointer" @click="state.crud = 'dishes'">Dishes</a>
                        <a class="nav-link pointer" @click="state.crud = 'users'">Users</a>
                        <a class="nav-link pointer" @click="state.crud = 'reservations'">Reservations</a>
                    </div>
                </div>
            </div>
        </nav>

        <Categories v-if="state.crud == 'categories'" :categories="state.categories"></Categories>
        <Tables v-if="state.crud == 'tables'" :tables="state.tables"></Tables>
        <DishTypes v-if="state.crud == 'dishtypes'" :dishtypes="state.dishtypes"></DishTypes>
        <Dishes v-if="state.crud == 'dishes'" :dishes="state.dishes"></Dishes>
        <Users v-if="state.crud == 'users'" :users="state.users"></Users>
        <Reservations v-if="state.crud == 'reservations'" :reservations="state.reservations"></Reservations>

    </div>
</template>
  
<script>
import Constant from "../Constant";
import { reactive, computed } from "vue";
import { useStore } from "vuex";
import Categories from "../components/admin/CRUDCategories.vue";
import Tables from "../components/admin/CRUDTables.vue";
import DishTypes from "../components/admin/CRUDDishTypes.vue";
import Dishes from "../components/admin/CRUDDishes.vue";
import Users from "../components/admin/CRUDUsers.vue";
import Reservations from "../components/admin/CRUDReservations.vue";


export default {
    components: { Categories, Tables, DishTypes, Dishes, Users, Reservations },
    setup() {
        const store = useStore();
        const state = reactive({
            categories: computed(() => store.getters["categoryAdmin/getCategory"]),
            tables: computed(() => store.getters["tableAdmin/getTable"]),
            dishtypes: computed(() => store.getters["dishTypeAdmin/getDishType"]),
            dishes: computed(() => store.getters["dishAdmin/getDish"]),
            users: computed(() => store.getters["userAdmin/getUser"]),
            reservations: computed(() => store.getters["reservationAdmin/getReservation"]),
            crud: "categories"
        });
        store.dispatch("categoryAdmin/" + Constant.GET_ALL_CATEGORIES);
        store.dispatch("tableAdmin/" + Constant.GET_ALL_TABLES);
        store.dispatch("dishTypeAdmin/" + Constant.GET_ALL_DISHTYPES);
        store.dispatch("dishAdmin/" + Constant.GET_ALL_DISHES);
        store.dispatch("userAdmin/" + Constant.GET_ALL_USERS);
        store.dispatch("reservationAdmin/" + Constant.GET_ALL_RESERVATIONS);
        return { state };
    },
};

</script>
  
<style>
.pointer {
    cursor: pointer;
}
</style>
