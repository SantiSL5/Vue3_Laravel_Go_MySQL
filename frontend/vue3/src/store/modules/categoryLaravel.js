import Constant from "../../Constant";
import CategoryService from "../../services/CategoryServiceLaravel";

export const categoryLaravel = {
    namespaced: true,

    state: {

    },
    mutations: {
        [Constant.GET_ALL_CATEGORIES]: (state, payload) => {
            state.categorieslist = payload;
            console.log(state);
        },
        [Constant.GET_ONE_CATEGORY]: (state) => {
            state.allUsers;
        },
    },
    actions: {
        [Constant.GET_ALL_CATEGORIES]: (store) => {
            CategoryService.getAllCategories().then(data => {
                console.log(data.data);
                store.commit(Constant.GET_ALL_CATEGORIES, data.data);
            });
        },
        [Constant.GET_ONE_CATEGORY]: (store) => {
            CategoryService.getCategoryById().then(data => {
                store.commit(Constant.GET_ONE_CATEGORY, data.data);
            });
        },
    },
    getters: {
        getCategory(state) {
            return state.categorieslist;
        }
    }
}
