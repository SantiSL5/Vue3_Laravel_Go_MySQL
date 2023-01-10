import Constant from "../../../Constant";
import TableService from "../../../services/admin/TableService";
import { createToaster } from "@meforma/vue-toaster"; // Debería estar en el componente

export const tableAdmin = {
    namespaced: true,

    state: {

    },
    mutations: {
        [Constant.CREATE_ONE_TABLE]: (state, payload) => {
            state.tableslist.push({ ...payload });
        },
        [Constant.GET_ALL_TABLES]: (state, payload) => {
            state.tableslist = payload;
        },
        [Constant.GET_ONE_TABLE]: (state) => {
            state.allUsers;
        },
        [Constant.UPDATE_ONE_TABLE]: (state, payload) => {
            let index = state.tableslist.findIndex(
                item => item.id == payload.id
            );
            state.tableslist[index] = payload;
        },
        [Constant.DELETE_ONE_TABLE]: (state, payload) => {
            let index = state.tableslist.findIndex(
                (item) => item.id === payload
            );
            state.tableslist.splice(index, 1);
        },
    },
    actions: {
        [Constant.CREATE_ONE_TABLE]: (store, payload) => {
            TableService.createTable(payload.table).then(data => {
                const toaster = createToaster({
                    position: "top"
                });
                if (data.statusText == "Created") {
                    toaster.success(`Table created successfully`);
                    store.commit(Constant.CREATE_ONE_TABLE, data.data);
                } else {
                    toaster.error(data.data.Message);
                }
            });
        },
        [Constant.UPDATE_ONE_TABLE]: (store, payload) => {
            TableService.updateTable(payload.table, payload.id).then(data => {
                const toaster = createToaster({
                    position: "top"
                });
                if (data.data.category) {
                    toaster.success(`Table updated successfully`);
                    store.commit(Constant.UPDATE_ONE_TABLE, data.data)
                } else {
                    toaster.error(`Table already exists`);
                }
            });
        },
        [Constant.GET_ALL_TABLES]: (store) => {
            TableService.getAllTables().then(data => {
                store.commit(Constant.GET_ALL_TABLES, data.data);
            });
        },
        [Constant.GET_ONE_TABLE]: (store) => {
            TableService.getTableById().then(data => {
                store.commit(Constant.GET_ONE_TABLE, data.data);
            });
        },
        [Constant.DELETE_ONE_TABLE]: (store, payload) => {
            TableService.deleteTableById(payload.id).then(
                store.commit(Constant.DELETE_ONE_TABLE, payload.id)
            );
        },
    },
    getters: {
        getTable(state) {
            return state.tableslist;
        }
    }
}
