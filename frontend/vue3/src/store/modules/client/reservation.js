import Constant from "../../../Constant";
import ReservationService from "../../../services/client/ReservationService";
// import { createToaster } from "@meforma/vue-toaster"; // Debería estar en el componente

export const reservationClient = {
    namespaced: true,

    state: {

    },
    mutations: {
        [Constant.GET_ALL_RESERVATIONS]: (state, payload) => {
            state.reservationslist = payload;
        },
        [Constant.DELETE_ONE_RESERVATION]: (state, payload) => {
            let index = state.reservationslist.findIndex(
                (item) => item.id === payload
            );
            state.reservationslist.splice(index, 1);
        },
    },
    actions: {
        [Constant.GET_ALL_RESERVATIONS]: (store) => {
            ReservationService.getClientReservations().then(data => {
                store.commit(Constant.GET_ALL_RESERVATIONS, data.data);
            }).catch(function () { });

        },
        [Constant.DELETE_ONE_RESERVATION]: (store, payload) => {
            ReservationService.deleteReservationById(payload.id).then(
                store.commit(Constant.DELETE_ONE_RESERVATION, payload.id)
            );
        },
    },
    getters: {
        getReservation(state) {
            return state.reservationslist;
        }
    }
}
