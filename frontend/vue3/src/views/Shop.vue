<template>
    <div class="container-lg m-5">
        <div class="table-wrapper">

            <div class="table-title">
                <div class="row">
                    <div class="col-sm-8">
                        <h2><b>AAA</b></h2>
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
</template>
  
<script>
import Constant from "../Constant";
import ListTables from '../components/ListTables.vue';
import { reactive, computed } from "vue";
import { useStore } from "vuex";
// import { useRouter } from "vue-router";

export default {
    components: { ListTables },
    setup() {
        const store = useStore();
        const state = reactive({
            tableslist: computed(() => store.getters["tableClient/getTable"]),
        });
        store.dispatch("tableClient/" + Constant.GET_ALL_TABLES);
        return { state };
    },
};
</script>
  
<style>

</style>
  