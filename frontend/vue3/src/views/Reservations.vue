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
                                    v-model="state.form.dateForm" :min="state.today" @change="applyFilter">
                            </div>
                            <div class="mb-3">
                                <label for="turn" class="form-label">Turn:</label>
                                <select name="turn" id="turn" class="form-control col-5" v-model="state.form.turn"
                                    @change="applyFilter">
                                    <!-- <option value="lunch" selected="true" disabled="disabled" hidden>Turn</option> -->
                                    <option value="lunch">Lunch</option>
                                    <option value="dinner">Dinner</option>
                                </select>
                            </div>
                        </div>

                        <div class="col-sm-4 mt-2">
                            Minimal Capacity:
                            <input type="number" class="form-control col-5" v-model="state.form.capacity"
                                @change="applyFilter">

                            <div class="mt-3">
                                Category:
                                <select name="turn" id="turn" class="form-control col-5" v-model="state.form.category"
                                    @change="applyFilter">
                                    <option value="">All</option>
                                    <option v-for="category in state.categories" :key="category.id"
                                        :value="category.name">
                                        {{ category.name }}
                                    </option>
                                </select>

                            </div>
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
        store.dispatch("categoryClient/" + Constant.GET_ALL_CATEGORIES);

        const state = reactive({
            tableslist: computed(() => store.getters["tableClient/getTable"]),
            form: {
                date: null,
                dateForm: null,
                turn: "lunch",
                capacity: 4,
                category: "",
            },
            today: new Date(),
            dayFormatted: null,
            totalpages: 0,
            categories: computed(() => store.getters["categoryClient/getCategory"]),
        });

        state.today.setDate(state.today.getDate() + 1);
        state.today = state.today.toISOString().split('T')[0];

        state.form.dateForm = state.today;

        state.dayFormatted = state.today.split("-");
        state.dayFormatted = state.dayFormatted[2] + "/" + state.dayFormatted[1] + "/" + state.dayFormatted[0];

        store.dispatch("tableClient/" + Constant.GET_ALL_TABLES, { offset: 0, limit: 6, capacity: 4, date: state.dayFormatted, turn: state.form.turn });

        setTimeout(() => {
            state.totalpages = Math.ceil(state.tableslist.count / 6);
        }, 200);

        const applyFilter = () => {
            state.form.date = state.form.dateForm.split("-");
            state.form.date = state.form.date[2] + "/" + state.form.date[1] + "/" + state.form.date[0];

            store.dispatch("tableClient/" + Constant.GET_ALL_TABLES, { offset: 0, limit: 6, capacity: state.form.capacity, date: state.form.date, turn: state.form.turn, category: state.form.category });

            state.totalpages = 0;
            
            setTimeout(() => {
                state.totalpages = Math.ceil(state.tableslist.count / 6);
            }, 200);
        }

        const loadnewtables = (page) => {
            state.form.date = state.form.dateForm.split("-");
            state.form.date = state.form.date[2] + "/" + state.form.date[1] + "/" + state.form.date[0];

            console.log({ offset: 6 * page, limit: 6, capacity: state.form.capacity, date: state.form.date, turn: state.form.turn, category: state.form.category });
            store.dispatch("tableClient/" + Constant.GET_ALL_TABLES, { offset: 6 * page, limit: 6, capacity: state.form.capacity, date: state.form.date, turn: state.form.turn, category: state.form.category });
        }

        const reserve = (id) => {
            if (store._state.data.userClient.user.type != "client") {
                toaster.error(`Being logged is required`);
            } else {
                if (state.form.dateForm != null && state.form.turn != null) {
                    state.form.date = state.form.dateForm.split("-")
                    state.form.date = state.form.date[2] + "/" + state.form.date[1] + "/" + state.form.date[0]

                    state.tableslist.tables.map(function (value, index) {
                        if (value.id == id) {
                            state.tableslist.tables[index].enabled = false;
                        }
                    })

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
