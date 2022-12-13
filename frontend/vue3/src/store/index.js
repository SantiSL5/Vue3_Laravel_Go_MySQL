import { createStore } from "vuex";
import { categoryAdmin } from './modules/admin/category'
import { categoryClient } from './modules/client/category'

export default createStore({
    state: {
    },
    mutations: {
    },
    actions: {
    },
    modules: {
        categoryAdmin: categoryAdmin,
        categoryClient: categoryClient
    }
});