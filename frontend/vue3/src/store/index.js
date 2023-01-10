import { createStore } from "vuex";
import { categoryAdmin } from './modules/admin/category'
import { tableAdmin } from './modules/admin/table'
import { categoryClient } from './modules/client/category'
import { tableClient } from './modules/client/table'
import { dishClient } from './modules/client/dish'


export default createStore({
    state: {
    },
    mutations: {
    },
    actions: {
    },
    modules: {
        categoryAdmin: categoryAdmin,
        tableAdmin: tableAdmin,

        categoryClient: categoryClient,
        tableClient: tableClient,
        dishClient: dishClient,
    }
});