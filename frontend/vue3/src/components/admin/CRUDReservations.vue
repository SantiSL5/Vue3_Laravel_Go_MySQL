<template>
    <h2 class="mt-2">Reservations</h2>
    <form>
        <div class="mb-3">
            <label for="user" class="form-label">User ID</label>
            <input type="number" class="form-control" id="user" v-model="state.reservation.user" required>
        </div>
        <div class="mb-3">
            <label for="table" class="form-label">Table ID</label>
            <input type="number" class="form-control" id="table" v-model="state.reservation.table" required>
        </div>
        <div class="mb-3">
            <label for="date" class="form-label">Date</label>
            <input type="date" class="form-control" id="date" v-model="state.reservation.dateForm" required
                :min="state.tomorrow">
        </div>
        <div class=" mb-3">
            <label for="turn" class="form-label">Turn</label>
            <select name="turn" id="turn" class="form-control col-5" v-model="state.reservation.turn">
                <option value="" selected="true" disabled="disabled" hidden>Turn</option>
                <option value="lunch">Lunch</option>
                <option value="dinner">Dinner</option>
            </select>
        </div>
        <div class="mb-3">
            <label for="confirmed" class="form-check-label pe-2">Confirmed?</label>
            <input type="checkbox" class="form-check-input" name="confirmed" id="" v-model="state.reservation.confirmed"
                required>
        </div>
        <button type="button" class="btn btn-success mb-3" @click="createReservation()"
            v-if="state.formType == 'create'"
            :disabled="state.reservation.user == '' || state.reservation.table == '' || state.reservation.dateForm == '' || state.reservation.turn == ''">Create</button>
    </form>

    <table class="table">
        <thead>
            <tr>
                <th>ID</th>
                <th>User</th>
                <th>Table</th>
                <th>Date</th>
                <th>Turn</th>
                <th>Confirmed?</th>
                <th>Operations</th>
            </tr>
        </thead>
        <tbody>
            <tr v-for="reservation in reservations" :key="reservation.id">
                <th scope="row">{{ reservation.id }}</th>
                <td>{{ reservation.user.email }}</td>
                <td>{{ reservation.table.id }} ({{ reservation.table.code }})</td>
                <td>{{ reservation.date }}</td>
                <td>{{ reservation.turn }}</td>
                <td>
                    <input type="checkbox" name="confirmed" id="" :checked="reservation.confirmed == 1"
                        @click="updateReservation(reservation)">
                </td>
                <td>
                    <button type="button" @click.stop="deleteReservation(reservation.id)"
                        class="btn btn-danger me-3">Delete</button>
                </td>
            </tr>
        </tbody>
    </table>
</template>

<script>
import Constant from '../../Constant';
import { useStore } from 'vuex'
import { reactive } from "vue";
import { useReservationsToggle } from "../../composables/useReservations";
import { createToaster } from "@meforma/vue-toaster";

export default {
    props: {
        reservations: Array,
    },
    setup() {
        const toaster = createToaster({ position: "top" });
        const store = useStore();
        const state = reactive({
            reservation: { user: "", table: "", dateForm: "", date: "", turn: "", confirmed: true },
            formType: "create",
            onUpdateReservation: null,
            tomorrow: new Date(),
        });

        state.tomorrow.setDate(state.tomorrow.getDate() + 1);
        state.tomorrow = state.tomorrow.toISOString().split('T')[0];

        const changeForm = (reservation = null) => {
            if (reservation) {
                state.formType = "update";
                state.onUpdateReservation = reservation.id
                state.reservation.code = reservation.code;
                state.reservation.category = reservation.category.id;
                state.reservation.capacity = reservation.capacity;
                state.reservation.enabled = reservation.enabled;
            } else {
                state.formType = "create";
                state.onUpdateReservation = "";
                state.reservation.code = "";
                state.reservation.category = "";
                state.reservation.capacity = "";
                state.reservation.enabled = false;
            }
        }

        const createReservation = () => {
            state.reservation.date = state.reservation.dateForm.split("-")
            state.reservation.date = state.reservation.date[2] + "/" + state.reservation.date[1] + "/" + state.reservation.date[0]

            store.dispatch("reservationAdmin/" + Constant.CREATE_ONE_RESERVATION, { reservation: state.reservation })
        }
        const deleteReservation = (id) => {
            store.dispatch("reservationAdmin/" + Constant.DELETE_ONE_RESERVATION, { id });
        }
        const updateReservation = (reservation = null) => {
            let confirmed = reservation.confirmed;
            if (confirmed == 1) {
                confirmed = false;
            } else {
                confirmed = true;
            }
            useReservationsToggle({ confirmed: confirmed }, reservation.id);
            toaster.success(`Reservation updated successfully`);
        }

        return { state, changeForm, createReservation, deleteReservation, updateReservation }
    },
};

</script>