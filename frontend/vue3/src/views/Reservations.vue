<template>
    <div class="row container-fluid-lg m-5 justify-content-center">
        <div class="col-11">
            <div class="table-wrapper">
                <div class="table-title">
                    <div class="row">
                        <div class="col-sm-8 row">
                            <input type="date" class="form-control col-5" v-model="state.form.dateForm"
                                :min="state.today">
                            <select name="" id="" class="form-control col-5" v-model="state.form.turn">
                                <!-- <option value="lunch" selected="true" disabled="disabled" hidden>Turn</option> -->
                                <option value="lunch">Lunch</option>
                                <option value="dinner">Dinner</option>
                            </select>
                        </div>
                        <div class="col-sm-4">
                        </div>
                    </div>
                </div>
                <div class="row justify-content-around">
                    <ListTables v-for="tableitem in state.tableslist" :key="tableitem.id" :tableitem="tableitem"
                        @reservetable="reserve" />
                </div>
                <Pagination :totalpages="state.totalpages" @changepage="loadnewtables" />

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
        store.dispatch("tableClient/" + Constant.GET_ALL_TABLES);


        const state = reactive({
            tableslist: computed(() => store.getters["tableClient/getTable"]),
            form: {
                date: null,
                dateForm: null,
                turn: null,
            },
            today: new Date(),
            totalpages: 0,
        });

        state.today.setDate(state.today.getDate() + 1);
        state.today = state.today.toISOString().split('T')[0];

        setTimeout(() => {
            // console.log(state.tableslist.length);
            state.totalpages = computed(() => Math.ceil(store.getters["tableClient/getTable"].length / 6));
            // console.log(state.totalpages);
        }, 100);


        const loadnewtables = (page) => {
            console.log(page);
            // state.show_tablelist = computed(() =>
            //     state.save_tablelist.slice(page * 6, page * 6 + 6)
            // );
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

        return { state, loadnewtables, reserve };
    },
};
</script>

<style>

</style>
