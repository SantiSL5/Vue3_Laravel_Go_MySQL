import Constant from "../../../Constant";
import TableService from "../../../services/client/TableService";

export const tableClient = {
    namespaced: true,

    state: {

    },
    mutations: {
        [Constant.GET_ALL_TABLES]: (state, payload) => {
            state.tableslist = payload;
        },
        [Constant.GET_ONE_TABLE]: (state) => {
            state.allUsers;
        },
    },
    actions: {
        [Constant.GET_ALL_TABLES]: (store, payload) => {
            TableService.getAllTables(payload).then(data => {
                store.commit(Constant.GET_ALL_TABLES, data.data);
            });
        },
        [Constant.GET_ONE_TABLE]: (store) => {
            TableService.getTableById().then(data => {
                store.commit(Constant.GET_ONE_TABLE, data.data);
            });
        },
    },
    getters: {
        getTable(state) {
            return state.tableslist;
        }
    }
}
