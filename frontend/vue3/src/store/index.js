import { createStore } from "vuex";
import { categoryAdmin } from './modules/admin/category'
import { tableAdmin } from './modules/admin/table'
import { dishTypeAdmin } from './modules/admin/dishtype'
import { dishAdmin } from './modules/admin/dish'
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
        dishTypeAdmin: dishTypeAdmin,
        dishAdmin: dishAdmin,
        
        categoryClient: categoryClient,
        tableClient: tableClient,
        dishClient: dishClient,
    }
});