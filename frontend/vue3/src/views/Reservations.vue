<template>
    <div class="row container-fluid-lg m-5 justify-content-center">
        <div class="col-11">
            <div class="table-wrapper">
                <div class="table-title">
                    <div class="row">
                        <div class="col-sm-4 row">
                            <div class="mb-2">
                                <label for="date" class="form-label">Date:</label>
                                <input type="date" name="date" id="date" class="form-control col-5"
                                    v-model="state.form.dateForm" :min="state.today">
                            </div>
                            <div class="mb-3">
                                <label for="turn" class="form-label">Turn:</label>
                                <select name="turn" id="turn" class="form-control col-5" v-model="state.form.turn">
                                    <!-- <option value="lunch" selected="true" disabled="disabled" hidden>Turn</option> -->
                                    <option value="lunch">Lunch</option>
                                    <option value="dinner">Dinner</option>
                                </select>
                            </div>
                        </div>

                        <div class="col-sm-4 mt-2">
                            Minimal Capacity:
                            <input type="number" class="form-control col-5" v-model="state.capacity"
                                :min="state.capacity" @change="applyFilter">

                        </div>
                    </div>
                </div>
                <div class="row justify-content-around">
                    <ListTables v-for="tableitem in state.tableslist.tables" :key="tableitem.id" :tableitem="tableitem"
                        @reservetable="reserve" />
                    <h1 v-if="state.tableslist.tables.length == 0">No tables available</h1>
                </div>
                <div v-if="state.totalpages != 0">
                    <Pagination :totalpages="state.totalpages" @changepage="loadnewtables" />
                </div>

            </div>
        </div>
    </div>
</template>

<script>
import Constant from "../Constant";
import ListTables from '../components/ListTables.vue';
import Pagination from '../components/Pagination.vue';
import { reactive, computed } from "vue";
import { useStore } from "vuex";
import { createToaster } from "@meforma/vue-toaster";
import { useReservationsReserve } from "../composables/useReservations";

export default {
    components: { ListTables, Pagination },
    setup() {
        const toaster = createToaster({ position: "top" });
        const store = useStore();
        store.dispatch("tableClient/" + Constant.GET_ALL_TABLES, { offset: 0, limit: 6, capacity: 4 });


        const state = reactive({
            tableslist: computed(() => store.getters["tableClient/getTable"]),
            form: {
                date: null,
                dateForm: null,
                turn: null,
            },
            today: new Date(),
            totalpages: 0,
            capacity: 4
        });

        state.today.setDate(state.today.getDate() + 1);
        state.today = state.today.toISOString().split('T')[0];

        setTimeout(() => {
            state.totalpages = Math.ceil(state.tableslist.count / 6);
        }, 200);

        const applyFilter = () => {
            store.dispatch("tableClient/" + Constant.GET_ALL_TABLES, { offset: 0, limit: 6, capacity: state.capacity });
            setTimeout(() => {
                state.totalpages = Math.ceil(state.tableslist.count / 6);
            }, 200);
        }

        const loadnewtables = (page) => {
            console.log(page);
            store.dispatch("tableClient/" + Constant.GET_ALL_TABLES, { offset: 6 * page, limit: 6, capacity: state.capacity });
        }

        const reserve = (id) => {
            if (store._state.data.userClient.user.type != "client") {
                toaster.error(`Being logged is required`);
            } else {
                if (state.form.dateForm != null && state.form.turn != null) {
                    state.form.date = state.form.dateForm.split("-")
                    state.form.date = state.form.date[2] + "/" + state.form.date[1] + "/" + state.form.date[0]

                    useReservationsReserve({ table: id, date: state.form.date, turn: state.form.turn });
                    toaster.success(`Reserved successfully`);
                } else {
                    toaster.error(`Date and turn required`);
                }
            }
        }

        return { state, loadnewtables, reserve, applyFilter };
    },
};
</script>

<style>

</style>
