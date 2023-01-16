import Constant from "../../../Constant";
import ReservationService from "../../../services/admin/ReservationService";
import { createToaster } from "@meforma/vue-toaster"; // Debería estar en el componente

export const reservationAdmin = {
    namespaced: true,

    state: {

    },
    mutations: {
        [Constant.CREATE_ONE_RESERVATION]: (state, payload) => {
            state.reservationslist.push({ ...payload });
        },
        [Constant.GET_ONE_RESERVATION]: (state) => {
            state.allUsers;
        },
        [Constant.DELETE_ONE_RESERVATION]: (state, payload) => {
            let index = state.reservationslist.findIndex(
                (item) => item.id === payload
            );
            state.reservationslist.splice(index, 1);
        },
    },
    actions: {
        [Constant.CREATE_ONE_RESERVATION]: (store, payload) => {
            ReservationService.createReservation(payload.reservation).then(data => {
                const toaster = createToaster({ position: "top" });
                if (data.statusText == "Created") {
                    toaster.success(`Reservation created successfully`);
                    store.commit(Constant.CREATE_ONE_RESERVATION, data.data);
                } else {
                    toaster.error(data.data.Message);
                }
            });
        },
        [Constant.GET_ONE_RESERVATION]: (store) => {
            ReservationService.getReservationById().then(data => {
                store.commit(Constant.GET_ONE_RESERVATION, data.data);
            });
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
