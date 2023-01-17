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

                    <!-- <div class="col mt-2" v-if="state.auth.type != ''">
                        <h6 class="col">{{ state.auth.username }}</h6>
                    </div> -->

                    <div class="col" v-if="state.auth.type != ''">
                        <ul class="navbar-nav">
                            <li class="nav-item" data-bs-toggle="modal" data-bs-target="#modalReservations">
                                <a class="nav-link active" aria-current="page" href="#">{{ state.auth.username }}</a>
                            </li>
                        </ul>

                    </div>

                    <div class="col" v-if="state.auth.type != ''">
                        <button type="button" class="btn btn-info col" @click="logout">Logout</button>
                    </div>

                </span>
            </div>
        </div>
    </nav>
    <!-- Modal -->
    <div class="modal fade" id="modalReservations" tabindex="-1" aria-labelledby="ModalLabel" aria-hidden="true">
        <div class="modal-dialog">
            <div class="modal-content">
                <div class="modal-header">
                    <h1 class="modal-title fs-5" id="ModalLabel">Reservations</h1>
                    <button type="button" class="btn-close" data-bs-dismiss="modal" aria-label="Close"></button>
                </div>
                <div class="modal-body">
                    <table class="table table-striped">
                        <thead>
                            <tr>
                                <th scope="col">Date</th>
                                <th scope="col">Turn</th>
                                <th scope="col">Table</th>
                                <th scope="col">Confirmed</th>
                                <th scope="col">Cancel</th>
                            </tr>
                        </thead>
                        <tbody v-for="reservation in state.reservationslist" :key="reservation.id">
                            <tr>
                                <td>{{ reservation.date }}</td>
                                <td class="uppercase">{{ reservation.turn }}</td>
                                <td>{{ reservation.tablecontent.code }}</td>
                                <td>
                                    <img class="img-fluid" v-if="reservation.confirmed == true"
                                        src="../assets/icons/check.svg" alt="YES" width="25">
                                    <img class="img-fluid" v-if="reservation.confirmed == false"
                                        src="../assets/icons/x.svg" alt="NO" width="25">
                                </td>
                                <td><button type="button" class="btn-close"></button></td>
                            </tr>
                        </tbody>
                    </table>
                </div>
                <div class="modal-footer">
                    <button type="button" class="btn btn-secondary" data-bs-dismiss="modal">Close</button>
                </div>
            </div>
        </div>
    </div>

</template>

<script>
import { reactive, computed } from "vue";
import { useStore } from "vuex";
import Constant from '../Constant';

export default {
    setup() {
        const store = useStore()
        store.dispatch("userClient/" + Constant.PROFILE_USER)
        store.dispatch("reservationClient/" + Constant.GET_ALL_RESERVATIONS)

        const state = reactive({
            user: computed(() => store.getters["userClient/getUser"]),
            auth: computed(() => store.getters["userClient/getAuth"]),
            reservationslist: computed(() => store.getters["reservationClient/getReservation"]),
        });
        const logout = () => {
            store.dispatch("userClient/" + Constant.LOGOUT)
        }

        return { state, logout };
    },
};
</script>

<style>
.uppercase:first-letter {
    text-transform: uppercase;
}
</style>