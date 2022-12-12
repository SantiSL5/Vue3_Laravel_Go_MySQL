import Constant from "../../Constant";
import CategoryService from "../../services/CategoryServiceGo";

export const categoryGo = {
    namespaced: true,

    state: {

    },
    mutations: {
        [Constant.GET_ALL_CATEGORIES_GO]: (state) => {
            state.allUsers;
        },
        [Constant.GET_ONE_CATEGORY]: (state) => {
            state.allUsers;
        },
    },
    actions: {
        [Constant.GET_ALL_CATEGORIES_GO]: (store) => {
            CategoryService.getAllCategories().then(data => {
                store.commit(Constant.GET_ALL_CATEGORIES_GO, data.data);
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
            return state.category;
        }
    }
}
