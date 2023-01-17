<template>
    <nav class="navbar navbar-expand-lg navbar-dark bg-dark">
        <div class="container-fluid">
            <a class="navbar-brand" href="#">Namazu</a>
            <button class="navbar-toggler" type="button" data-bs-toggle="collapse" data-bs-target="#navbarText"
                aria-controls="navbarText" aria-expanded="false" aria-label="Toggle navigation">
                <span class="navbar-toggler-icon"></span>
            </button>
            <div class="collapse navbar-collapse" id="navbarText">
                <ul class="navbar-nav me-auto mb-2 mb-lg-0">
                    <li class="nav-item">
                        <router-link class="nav-link active" aria-current="page" href="#" to="/home">Home</router-link>
                    </li>
                    <li class="nav-item">
                        <router-link class="nav-link" href="#" to="/reservations">Reservations</router-link>
                    </li>
                    <li class="nav-item">
                        <router-link class="nav-link" href="#" to="/menu">Menu</router-link>
                    </li>
                </ul>
                <span class="navbar-text row">
                    {{ state.auth }}
                    <!-- ////////// -->
                    <div class="col" v-if="state.auth.type == 'admin'">
                        <router-link class="nav-link" href="#" to="/admin">
                            <button type="button" class="btn btn-info">PanelAdmin</button>
                        </router-link>
                    </div>
                    <!-- ////////// -->

                    <div class="col" v-if="state.auth.type == ''">
                        <router-link class="nav-link" href="#" to="/login">
                            <button type="button" class="btn btn-info">Login</button>
                        </router-link>
                    </div>

                    <div class="col mt-2" v-if="state.auth.type != ''">
                        <h6 class="col">{{ state.auth.username }}</h6>
                    </div>

                    <div class="col" v-if="state.auth.type != ''">
                        <button type="button" class="btn btn-info col" @click="logout">Logout</button>
                    </div>

                </span>
            </div>
        </div>
    </nav>
</template>

<script>
import { reactive, computed } from "vue";
import { useStore } from "vuex";
import Constant from '../Constant';

export default {
    setup() {
        const store = useStore()
        store.dispatch("userClient/" + Constant.PROFILE_USER)

        const state = reactive({
            user: computed(() => store.getters["userClient/getUser"]),
            auth: computed(() => store.getters["userClient/getAuth"]),
        });
        const logout = () => {
            store.dispatch("userClient/" + Constant.LOGOUT)
        }
        return { state, logout };
    },
};
</script>

<style>

</style>