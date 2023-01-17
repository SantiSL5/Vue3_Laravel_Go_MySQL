import { createStore } from "vuex";
import { categoryAdmin } from './modules/admin/category'
import { tableAdmin } from './modules/admin/table'
import { dishTypeAdmin } from './modules/admin/dishtype'
import { dishAdmin } from './modules/admin/dish'
import { userAdmin } from './modules/admin/user'
import { reservationAdmin } from './modules/admin/reservation'

import { categoryClient } from './modules/client/category'
import { tableClient } from './modules/client/table'
import { dishClient } from './modules/client/dish'
import { userClient } from './modules/client/user'
import { reservationClient } from './modules/client/reservation'



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
        userAdmin: userAdmin,
        reservationAdmin: reservationAdmin,

        categoryClient: categoryClient,
        tableClient: tableClient,
        dishClient: dishClient,
        userClient: userClient,
        reservationClient: reservationClient,
    }
});