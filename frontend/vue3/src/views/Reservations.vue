<template>
    <div class="row container-fluid-lg m-5 justify-content-center">
        <div class="col-11">
            <div class="table-wrapper">
                <div class="table-title">
                    <div class="row">
                        <div class="col-sm-8 row">
                            <input type="date" class="form-control col-5" v-model="state.form.date">
                            <select name="" id="" class="form-control col-5">
                                <option value="lunch" selected="true" disabled="disabled" hidden>Turn</option>
                                <option value="lunch">Lunch</option>
                                <option value="dinner">Dinner</option>
                            </select>
                        </div>
                        <div class="col-sm-4">
                        </div>
                    </div>
                </div>
                <div class="row justify-content-around">
                    <ListTables v-for="tableitem in state.tableslist" :key="tableitem.id" :tableitem="tableitem" />
                </div>
            </div>
        </div>
    </div>
</template>

<script>
import Constant from "../Constant";
import ListTables from '../components/ListTables.vue';
import { reactive, computed } from "vue";
import { useStore } from "vuex";



export default {
    components: { ListTables },
    setup() {
        const store = useStore();
        const state = reactive({
            tableslist: computed(() => store.getters["tableClient/getTable"]),
            form: {
                date: null,
                turn: null,

            }
        });
        store.dispatch("tableClient/" + Constant.GET_ALL_TABLES);
        return { state };
    },
};
</script>

<style>

</style>
