import { createStore } from "vuex";
import { categoryLaravel } from './modules/categoryLaravel'
import { categoryGo } from './modules/categoryGo'

export default createStore({
    state: {
    },
    mutations: {
    },
    actions: {
    },
    modules: {
        categoryLaravel: categoryLaravel,
        categoryGo: categoryGo
    }
});